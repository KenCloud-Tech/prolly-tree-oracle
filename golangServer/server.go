// main.go
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/api"
	"Oracle.com/golangServer/config"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

	// Connect to the Ethereum client
	var err error
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
	dbs, err := loadDb(ctx, config.SaveDataPath)
	if err != nil {
		mainLogger.Errorf("Init error %v", err)
		return
	}
	logger.Info("Loading local db successfully. ")

	config.SetDatabases(dbs)
	// Meta info listener
	go api.GetCollections(ctx, sugar.Named("get_collection"))
	go api.GetIndexes(ctx, sugar.Named("get_index"))
	go api.GetRootCid(ctx, sugar.Named("get_rootcid"))
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

	err = saveDB(ctx, config.SaveDataPath, config.Dbs) // Save db
	if err != nil {
		mainLogger.Errorf("save db file error %w", err)
		return
	}
	logger.Info("All data saved successfully. ")
}

func loadDb(ctx context.Context, savePath string) (map[string]*indexer.Database, error) {
	// Load saved databases if they exist
	dbs := make(map[string]*indexer.Database)
	file, err := os.Open(path.Join(savePath, "paths"))
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("error opening file: %w", err)
		}
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			filename := scanner.Text()

			parts := strings.Split(filename, "/")
			lastPart := parts[len(parts)-1]
			fileNameParts := strings.Split(lastPart, ".")
			result := fileNameParts[0]

			db, err := indexer.ImportFromFile(path.Join(savePath, filename))
			if err != nil {
				return nil, fmt.Errorf("an error occurred while scanning path: %w", err)
			}
			dbs[result] = db
		}
		// Check whether errors are encountered during the Scan process
		if scanner.Err() != nil {
			return nil, fmt.Errorf("an error occurred while scanning the file %w", err)
		}
	}
	return dbs, nil
}

func saveDB(ctx context.Context, savePath string, dbs map[string]*indexer.Database) error {
	tmpDir, err := os.MkdirTemp(os.TempDir(), "*")
	if err != nil {
		return fmt.Errorf("create tmp dir %w", err)
	}

	metaPath := path.Join(tmpDir, "paths")
	file, err := os.OpenFile(metaPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return fmt.Errorf("create paths file fail %w", err)
	}

	defer file.Close()

	for cid, db := range dbs {
		if err != nil {
			return fmt.Errorf("unable to create or open file: %w", err)
		}
		// Path: tmpDir/{RootCid}.car
		fileName := cid + ".car"
		// Use fmt.Fprintln to write a string to a file
		if _, err := fmt.Fprintln(file, fileName); err != nil {
			return fmt.Errorf("unable to write to file: %w", err)

		}
		// Save db as db.car
		err = db.ExportToFile(ctx, path.Join(tmpDir, fileName))
		if err != nil {
			return fmt.Errorf("export to file fail %w", err)
		}
	}

	err = os.RemoveAll(savePath)
	if err != nil {
		return fmt.Errorf("remove  savepath %w", err)
	}

	err = os.Rename(tmpDir, savePath)
	if err != nil {
		return fmt.Errorf("move temp path %w", err)
	}
	return nil
}
