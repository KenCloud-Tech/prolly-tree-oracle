package api

import (
	"context"
	"fmt"
	"path"
	"strings"
	"time"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ipfs/go-cid"
	"go.uber.org/zap"
)

func CreatEventListener(ctx context.Context, logger *zap.SugaredLogger) {
	for {
		// Create channels for logs
		Logs := make(chan *Oracle.OracleCreate)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}

		// start Listening...
		logger.Info("CreateEvent Listening ...")
		eventSub, err := config.OracleContract.WatchCreate(opts, Logs)
		if err != nil {
			logger.Error("Failed to subscribe to Get events:", err)
			time.Sleep(5 * time.Second)
			close(Logs)
			continue
		}
	LOOP:
		for {
			select {
			case err := <-eventSub.Err():
				logger.Errorf("[Error in Event CREATE]:", err)
				break LOOP
			case event := <-Logs:
				logger.Info("Received create event ", event.ReqID)
				create(ctx, event, logger)
			}
		}
	}
}

// create memory collection
func create(ctx context.Context, event *Oracle.OracleCreate, logger *zap.SugaredLogger) {
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	dbName := event.DbName
	logger.Info("Create collection dbName: ", event.DbName)
	logger.Info("Create collection ColName: ", event.ColName)
	logger.Info("Create collection PrimaryKey: ", event.PrimaryKey)
	if db, ok := config.Dbs[dbName]; ok {
		logger.Info("existed dbName: ", event.DbName)
		colName := event.ColName
		pK := strings.Split(event.PrimaryKey, ",")
		_, err := db.Collection(ctx, colName, pK...)
		if err != nil {
			logger.Error("Create collection ERROR: ", err)
			info := fmt.Sprintf("Create collection ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.CreatRsp(tps, event.ReqID, statement, dbName, colName, event.Owner, info)
			return
		}
		statement = true
		info := fmt.Sprintf("Create collection Success: %v", dbName)
		//response to oracle
		_, err = config.OracleContract.CreatRsp(tps, event.ReqID, statement, dbName, colName, event.Owner, info)
		if err != nil {
			logger.Error("Req function get an Error : ", err)
			db.DeleteCol(colName)
		} else {
			logger.Info("[", colName, "]", "Create collection success")
		}
	} else {
		logger.Info("non existed dbName: ", event.DbName)
		dbPath, emptyRoot, err := createEmptyDbFile(ctx, config.SaveDataPath, event.DbName)
		if err != nil {
			logger.Error("New Database ERROR: ", err)
			info := fmt.Sprintf("New Database ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.CreatRsp(tps, event.ReqID, statement, dbName, event.ColName, event.Owner, info)
			return
		}

		db, err := indexer.ImportFromFile(dbPath, emptyRoot)
		if err != nil {
			logger.Error("New Database ERROR: ", err)
			info := fmt.Sprintf("New Database ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.CreatRsp(tps, event.ReqID, statement, dbName, event.ColName, event.Owner, info)
			return
		}
		config.Dbs[dbName] = db
		colName := event.ColName
		pK := strings.Split(event.PrimaryKey, ",")
		_, err = db.Collection(ctx, colName, pK...)
		if err != nil {
			logger.Error("Create collection ERROR: ", err)
			info := fmt.Sprintf("Create collection ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.CreatRsp(tps, event.ReqID, statement, dbName, colName, event.Owner, info)
			return
		}
		logger.Info("non existed dbName created success: ", event.DbName)
		statement = true
		info := fmt.Sprintf("Create non existed collection Success: %v", dbName)
		//response to oracle
		_, err = config.OracleContract.CreatRsp(tps, event.ReqID, statement, dbName, colName, event.Owner, info)
		if err != nil {
			logger.Error("Req function get an Error : ", err)
			db.DeleteCol(colName)
		} else {
			logger.Info("[", colName, "]", "Create db success")
		}
	}

}

func createEmptyDbFile(ctx context.Context, savepath string, dbName string) (string, cid.Cid, error) {
	db, err := indexer.NewMemoryDatabase(ctx)
	if err != nil {
		return "", cid.Undef, err
	}
	fileName := dbName + ".car"
	dbFile := path.Join(savepath, fileName)
	err = db.ExportToFile(ctx, dbFile)
	if err != nil {
		return "", cid.Undef, err
	}
	defer db.Close()
	return dbFile, db.RootCid(), nil
}
