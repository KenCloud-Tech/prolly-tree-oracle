package api

import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"go.uber.org/zap"
)

func IndexEventListener(ctx context.Context, logger *zap.SugaredLogger) {
	for {
		// Create channels for logs
		Logs := make(chan *Oracle.OracleIndex)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}
		eventSub, err := config.OracleContract.WatchIndex(opts, Logs)
		if err != nil {
			logger.Errorf("Failed to subscribe to Index events:", err)
			time.Sleep(5 * time.Second)
			close(Logs)
			continue
		}
		// start Listening...
		logger.Info("IndexEvent Listening ...")
	LOOP:
		for {
			select {
			case err := <-eventSub.Err():
				logger.Errorf("[Error in Event INDEX]:", err)
				break LOOP
			case event := <-Logs:
				logger.Info("Received index event ", event.ReqID)
				index(ctx, event, logger)
			}
		}
	}
}

// Create index
func index(ctx context.Context, event *Oracle.OracleIndex, logger *zap.SugaredLogger) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("stacktrace from panic: " + string(debug.Stack()))
		}
	}()
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	colName := event.ColName
	db := config.Dbs[event.DbName]
	dbC, err := db.Collection(ctx, colName)
	if err != nil {
		logger.Errorf("Get collection ERROR: ", err)
		info := fmt.Sprintf("Get collection ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.IndexRsp(tps, event.ReqID, statement, event.Sender, info)
		return
	}
	pK := strings.Split(event.Key, ",")
	_, err = dbC.CreateIndex(ctx, pK...)
	if err != nil {
		logger.Errorf("Create Index ERROR: ", err)
		info := fmt.Sprintf("Create Index ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.IndexRsp(tps, event.ReqID, statement, event.Sender, info)
		return
	}

	statement = true
	_, err = config.OracleContract.IndexRsp(tps, event.ReqID, statement, event.Sender, "")
	if err != nil {
		logger.Errorf("Req function get an Error : ", err)
	} else {
		logger.Info("[", event.ColName, "]", "Create index success")
	}
}
