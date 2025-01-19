package api

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

func GetCollections(ctx context.Context, logger *zap.SugaredLogger) {

	for {
		// Get channels for logs
		Logs := make(chan *Oracle.OracleGetCol)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}
		eventSub, err := config.OracleContract.WatchGetCol(opts, Logs)
		if err != nil {
			logger.Errorf("Failed to subscribe to GetCollections events %v", err)
			time.Sleep(5 * time.Second)
			close(Logs)
			continue
		}
		// start Listening...
		logger.Info("GetCollectionsEvent Listening ...")
	LOOP:
		for {
			select {
			case err := <-eventSub.Err():
				logger.Error("[Error in Event GetCollections]:", err)
				break LOOP
			case event := <-Logs:
				logger.Infof("Received GetCollections event ", event.ReqID)
				getCollection(ctx, event, logger)
			}
		}
	}
}

func getCollection(ctx context.Context, event *Oracle.OracleGetCol, logger *zap.SugaredLogger) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("stacktrace from panic: " + string(debug.Stack()))
		}
	}()
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	db := config.Dbs[event.DbName]
	cols, err := db.ListCollections(ctx)
	if err != nil {
		logger.Error("[", event.DbName, "]", "List Collections ERROR: ", err)
		info := fmt.Sprintf("List Collections ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetColRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}

	jsonBytes, err := json.Marshal(cols)
	if err != nil {
		logger.Error("[", event.DbName, "]", "Trans to json ERROR: ", err)
		info := fmt.Sprintf("Trans to json ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetColRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	result, err := json.Marshal([][]byte{jsonBytes})
	if err != nil {
		logger.Error("Marshal Results ERROR: ", err)
		info := fmt.Sprintf("Marshal Results ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	statement = true
	//response to oracle
	_, err = config.OracleContract.GetColRsp(tps, event.ReqID, statement, result, event.CallBack, event.Sender, "")
	if err != nil {
		logger.Error("Req function get an Error : ", err)
	} else {
		logger.Info("[Get collections success]")
	}
}

func GetIndexes(ctx context.Context, logger *zap.SugaredLogger) {
	for {
		// Get channels for logs
		Logs := make(chan *Oracle.OracleGetIndex)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}
		eventSub, err := config.OracleContract.WatchGetIndex(opts, Logs)
		if err != nil {
			logger.Errorf("Failed to subscribe to GetIndexes events:", err)
			time.Sleep(5 * time.Second)
			close(Logs)
			continue
		}
		// start Listening...
		logger.Info("GetIndexesEvent Listening ...")
	LOOP:
		for {
			select {
			case err := <-eventSub.Err():
				logger.Error("[Error in Event GetIndexes]:", err)
				break LOOP
			case event := <-Logs:
				logger.Info("Received GetIndexes event ", event.ReqID)
				getIndex(ctx, event, logger)
			}
		}
	}
}

func getIndex(ctx context.Context, event *Oracle.OracleGetIndex, logger *zap.SugaredLogger) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("stacktrace from panic: " + string(debug.Stack()))
		}
	}()

	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	db := config.Dbs[event.DbName]
	colName := event.ColName
	col, err := db.Collection(ctx, colName, "")
	if err != nil {
		logger.Error("Get collection ERROR: ", err)
		info := fmt.Sprintf("Get collection ERROR: %v", err)
		statement = false
		//response to oracle
		err = sendTx(ctx, config.Client, func() (*types.Transaction, error) {
			return config.OracleContract.GetIndexRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		})
		if err != nil {
			logger.Error("response error resp %v", err)
		}
		return
	}
	indx, err := col.Indexes(ctx)
	if err != nil {
		logger.Error("Get indexes ERROR: ", err)
		info := fmt.Sprintf("Get indexes ERROR: %v", err)
		statement = false
		//response to oracle
		err = sendTx(ctx, config.Client, func() (*types.Transaction, error) {
			return config.OracleContract.GetIndexRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		})
		if err != nil {
			logger.Error("response error resp %v", err)
		}
		return
	}
	var indexes []string
	for _, i := range indx {
		indexes = append(indexes, i.Fields()...)
	}
	jsonBytes, err := json.Marshal(indexes)
	if err != nil {
		logger.Error("[", event.DbName, "]", "Trans to json ERROR: ", err)
		info := fmt.Sprintf("Trans to json ERROR: %v", err)
		statement = false
		//response to oracle
		err = sendTx(ctx, config.Client, func() (*types.Transaction, error) {
			return config.OracleContract.GetIndexRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		})
		if err != nil {
			logger.Error("response error resp %v", err)
		}
		return
	}

	result, err := json.Marshal([][]byte{jsonBytes})
	if err != nil {
		logger.Error("Marshal Results ERROR: ", err)
		info := fmt.Sprintf("Marshal Results ERROR: %v", err)
		statement = false
		//response to oracle
		err = sendTx(ctx, config.Client, func() (*types.Transaction, error) {
			return config.OracleContract.GetIndexRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		})
		if err != nil {
			logger.Error("response error resp %v", err)
		}
		return
	}
	statement = true
	//response to oracle
	err = sendTx(ctx, config.Client, func() (*types.Transaction, error) {
		return config.OracleContract.GetIndexRsp(tps, event.ReqID, statement, result, event.CallBack, event.Sender, "")
	})
	if err != nil {
		logger.Error("Req function get an Error : ", err)
	} else {
		logger.Info("[Get indexes success]")
	}
}
