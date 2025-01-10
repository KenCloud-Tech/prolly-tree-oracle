package api

import (
	"context"
	"fmt"
	"strings"
	"time"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"go.uber.org/zap"
)

func PutEventListener(ctx context.Context, logger *zap.SugaredLogger) {
	for {
		// Put channels for logs
		Logs := make(chan *Oracle.OraclePut)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}
		eventSub, err := config.OracleContract.WatchPut(opts, Logs)
		if err != nil {
			logger.Error("Failed to subscribe to Put events:", err)
			time.Sleep(5 * time.Second)
			close(Logs)
			continue
		}
		// start Listening...
		logger.Info("PutEvent Listening ...")
	LOOP:
		for {
			select {
			case err := <-eventSub.Err():
				logger.Error("[Error in Event PUT]:", err)
				break LOOP
			case event := <-Logs:
				logger.Info("Received put event ", event.ReqID)
				put(ctx, event, logger)
			}
		}
	}
}

// Put Data to memory db
func put(ctx context.Context, event *Oracle.OraclePut, logger *zap.SugaredLogger) {
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	colName := event.ColName
	dbName := event.DbName
	logger.Info("put collection dbName: ", event.DbName)
	logger.Info("put collection ColName: ", event.ColName)
	db := config.Dbs[dbName]
	if db == nil {
		logger.Error("Get DB ERROR: ", dbName)
		statement = false
		//response to oracle
		config.OracleContract.PutRsp(tps, event.ReqID, statement, event.Sender, "DB is not exist")
		return
	}
	dbC, err := db.Collection(ctx, colName, "")
	if err != nil {
		logger.Error("Get collection ERROR: ", err)
		info := fmt.Sprintf("Get collection ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.PutRsp(tps, event.ReqID, statement, event.Sender, info)
		return
	}
	strData := string(event.Data)
	reader := strings.NewReader(strData)
	// insert Data
	err = dbC.IndexNDJSON(ctx, reader)
	if err != nil {
		logger.Error("Put Data ERROR: ", err)
		info := fmt.Sprintf("Put Data ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.PutRsp(tps, event.ReqID, statement, event.Sender, info)
		return
	} else {
		logger.Info("Put Data Success")
		statement = true
	}
	//response to oracle
	_, err = config.OracleContract.PutRsp(tps, event.ReqID, statement, event.Sender, "")
	if err != nil {
		logger.Error("Req function get an Error : ", err)
	} else {
		logger.Info("[", event.ColName, "]", "Put Data success")
	}
}
