package api

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"go.uber.org/zap"
)

func GetEventListener(ctx context.Context, logger *zap.SugaredLogger) {
	for {
		// Get channels for logs
		Logs := make(chan *Oracle.OracleGet)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}
		eventSub, err := config.OracleContract.WatchGet(opts, Logs)
		if err != nil {
			logger.Error("Failed to subscribe to Get events:", err)
			time.Sleep(5 * time.Second)
			close(Logs)
			continue
		}
		// start Listening...
		logger.Info("GetEvent Listening ...")
	LOOP:
		for {
			select {
			case err := <-eventSub.Err():
				logger.Error("[Error in Event GET]:", err)
				break LOOP
			case event := <-Logs:
				logger.Info("Received get event ", event.ReqID)
				get(ctx, event, logger)
			}
		}
	}
}

// Get Data from memory db
func get(ctx context.Context, event *Oracle.OracleGet, logger *zap.SugaredLogger) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("stacktrace from panic: " + string(debug.Stack()))
		}
	}()
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	colName := event.ColName
	db := config.Dbs[event.DbName]
	if db == nil {
		logger.Errorf("Get DB ERROR: ", db)
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, "DB is not exist")
		return
	}
	col, err := db.Collection(ctx, colName)
	if err != nil {
		logger.Errorf("Get collection ERROR: ", err)
		info := fmt.Sprintf("Get collection ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	node, err := col.Get(ctx, event.RecordID)
	if err != nil {
		logger.Errorf("[", event.ColName, "]", "Get Data ERROR: ", err)
		info := fmt.Sprintf("Get Data ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	if node == nil {
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, "Data is not exist")
		return
	}
	result := nodeTobyte(node)
	statement = true
	//response to oracle
	_, err = config.OracleContract.GetRsp(tps, event.ReqID, statement, result, event.CallBack, event.Sender, "")
	if err != nil {
		logger.Errorf("Req function get an Error : ", err)
	} else {
		logger.Info("[", event.ColName, "]", "Get Data success")
	}
}
