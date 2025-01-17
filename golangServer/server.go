// main.go
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path"
	"syscall"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/api"
	"Oracle.com/golangServer/config"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-cid"
	"go.uber.org/zap"
)

func Init(logger *zap.SugaredLogger) error {
	// Define command-line flags
	urlFlag := flag.String("url", "", "The URL of the Ethereum node")
	privateKeyFlag := flag.String("privateKey", "", "The private key of the Oracle owner")
	contractAddressFlag := flag.String("contractAddress", "", "The address of the Oracle contract")
	chainIDFlag := flag.Int("chainID", 0, "The chain ID of the Ethereum network")
	dataFlag := flag.String("data", "data", "The storage path of data")
	flag.Parse()

	// Set configuration values based on flags
	config.SetConfig(*dataFlag, *urlFlag, *privateKeyFlag, *contractAddressFlag, int64(*chainIDFlag))

	_, err := os.Stat(config.SaveDataPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(config.SaveDataPath, 0755)
			if err != nil {
				return err
			}
			err = os.WriteFile(path.Join(config.SaveDataPath, "meta.json"), []byte("[]"), 0666)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("inspect path %w", err)
		}
	}
	// Connect to the Ethereum client
	client, err := ethclient.Dial(config.URL)
	if err != nil {
		return fmt.Errorf("failed to connect to the Ethereum client: %w", err)
	}
	config.SetClient(client)

	// Parse the private key
	privateKey, err := crypto.HexToECDSA(config.OracleOwnerPrivateKey[2:]) // Remove prefix "0x"
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}
	config.SetPrivateKey(privateKey)

	// Convert string address to `common.Address` type
	contractAddr := common.HexToAddress(config.ContractAddress)
	// Create a new instance using the contract address and client
	oracleContract, err := Oracle.NewOracle(contractAddr, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Oracle contract: %w", err)
	}
	config.SetOracleContract(oracleContract)

	return nil
}

func main() {
	ctx := context.Background()
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	mainLogger := sugar.Named("main")
	err := Init(sugar.Named("init"))
	if err != nil {
		mainLogger.Errorf("Init error %v", err)
		return
	}
	logger.Info("Config init successfully. ")
	cwd, err := os.Getwd()
	if err != nil {
		mainLogger.Errorf("get cwd directory %v", err)
		return
	}
	mainLogger.Info("attemp to load db from", path.Join(cwd, config.SaveDataPath))
	dbs, err := loadDb(ctx, mainLogger, config.SaveDataPath)
	if err != nil {
		mainLogger.Errorf("Init error %v", err)
		return
	}
	mainLogger.With("dbNum", len(dbs)).Info("Loading local db successfully. ")

	config.SetDatabases(dbs)
	// Meta info listener
	go api.GetCollections(ctx, sugar.Named("get_collection"))
	go api.GetIndexes(ctx, sugar.Named("get_index"))
	// Service listener
	go api.CreatEventListener(ctx, sugar.Named("create_event"))
	go api.PutEventListener(ctx, sugar.Named("put_event"))
	go api.IndexEventListener(ctx, sugar.Named("index_event"))
	go api.GetEventListener(ctx, sugar.Named("get_event"))
	go api.SearchEventListener(ctx, sugar.Named("search_event"))
	go api.ImportEventListener(ctx, sugar.Named("import_event"))
	go api.GetRootCidEventListener(ctx, sugar.Named("rootcid_event"))

	// Set up a channel for receiving signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// Wait for signal
	sig := <-sigs
	mainLogger.Info("Received signal %s, exiting...", sig)

	err = saveDB(ctx, *mainLogger, config.Dbs, config.SaveDataPath) // Save db
	if err != nil {
		mainLogger.Errorf("save db file error %w", err)
		return
	}
	logger.Info("All data saved successfully. ")
}

func loadDb(ctx context.Context, logger *zap.SugaredLogger, savePath string) (map[string]*indexer.Database, error) {
	// Load saved databases if they exist
	dbs := make(map[string]*indexer.Database)
	file, err := os.Open(path.Join(savePath, "meta.json"))
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("error opening file: %w", err)
		}
	} else {
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("read meta %w", err)
		}
		var metas []config.DbMeta
		err = json.Unmarshal(content, &metas)
		if err != nil {
			return nil, fmt.Errorf("unmarshal meta %w", err)
		}
		for _, meta := range metas {
			treeRoot, err := cid.Decode(meta.RootCid)
			if err != nil {
				return nil, fmt.Errorf("parse tree root %w", err)
			}
			db, err := indexer.ImportFromFile(path.Join(savePath, meta.FileName), treeRoot)
			if err != nil {
				return nil, fmt.Errorf("an error occurred while scanning path: %w", err)
			}
			logger.Infof("load db %s rootcid %s", meta.Name, db.RootCid())
			dbs[meta.Name] = db
		}
	}
	return dbs, nil
}

func saveDB(ctx context.Context, logger zap.SugaredLogger, dbs map[string]*indexer.Database, savePath string) error {
	var metas []config.DbMeta
	for dbName, db := range dbs {
		err := db.Close()
		if err != nil {
			logger.Errorf("flush db file %v", err)
		}
		logger.Infof("flush db %s rootcid %s", dbName, db.RootCid().String())
		metas = append(metas, config.DbMeta{
			Name:     dbName,
			FileName: dbName + ".car",
			RootCid:  db.RootCid().String(),
		})
	}

	metaBytes, err := json.Marshal(metas)
	if err != nil {
		return fmt.Errorf("marshal metas %w", err)
	}

	err = os.WriteFile(path.Join(savePath, "meta.json"), metaBytes, 0666)
	if err != nil {
		return fmt.Errorf("write metas file %w", err)
	}

	logger.Infof("flush db successfully")
	return nil
}
