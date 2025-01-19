package api

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

var gasPerByteByUrl *big.Int
var gasPerByteByUrlOnce sync.Once

func ImportEventListener(ctx context.Context, logger *zap.SugaredLogger) {
	for {
		// Import channels for logs
		Logs := make(chan *Oracle.OracleImportFromUrl)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}
		eventSub, err := config.OracleContract.WatchImportFromUrl(opts, Logs)
		if err != nil {
			logger.Errorf("Failed to subscribe to Import events:", err)
			time.Sleep(5 * time.Second)
			close(Logs)
			continue
		}
		// start Listening...
		logger.Info("ImportEvent Listening ...")
	LOOP:
		for {
			select {
			case err := <-eventSub.Err():
				logger.Errorf("[Error in Event IMPORT:", err)
				break LOOP
			case event := <-Logs:
				logger.Info("Received import event ", event.ReqID)
				importByUrl(ctx, event, logger)
			}
		}
	}
}

// Import Data to memory db
func importByUrl(ctx context.Context, event *Oracle.OracleImportFromUrl, logger *zap.SugaredLogger) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("stacktrace from panic: " + string(debug.Stack()))
		}
	}()
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	bigInt := big.Int{}
	size := bigInt.SetInt64(0)
	colName := event.ColName
	dbName := event.DbName
	db := config.Dbs[dbName]
	dbC, err := db.Collection(ctx, colName, "")
	if err != nil {
		logger.Errorf("Get collection ERROR: ", err)
		info := fmt.Sprintf("Get collection ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
		return
	}
	resp, err := http.Get(event.Url)
	if err != nil {
		logger.Errorf("Request url %s content ERROR: ", event.Url, err)
		info := fmt.Sprintf("Request url content: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
		return
	}
	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("Read response body ERROR: ", err)
		info := fmt.Sprintf("Read response body ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
		return
	}
	contentLength := len(body)

	if resp.StatusCode != 200 {
		logger.Errorf("Get datas Fail, StatusCode = ", resp.StatusCode)
		info := fmt.Sprintf("Get datas Fail, StatusCode = %v", resp.StatusCode)
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
		return
	}
	if contentLength == 0 {
		info := fmt.Sprintf("Empty content, Content-Length = 0")
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
		return
	}

	gasPerByteByUrlOnce.Do(func() {
		gasPerByteByUrl, err = config.OracleContract.GasPerByteByUrl(&bind.CallOpts{From: common.HexToAddress(config.ContractAddress)})
	})
	if err != nil {
		logger.Errorf("Get gas per bytes ERROR: ", err)
		info := fmt.Sprintf("Get gas per bytes ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
		return
	}

	logger.Infof("import event paid. pay(%d) price(%d), total(%d)",
		event.Value.Int64(), gasPerByteByUrl.Int64(), gasPerByteByUrl.Int64()*int64(contentLength))
	if event.Value.Int64() < gasPerByteByUrl.Int64()*int64(contentLength) {
		info := fmt.Sprintf("The data is too large and the gas paid is insufficient. pay(%d) unit(%d), total(%d)",
			event.Value.Int64(), gasPerByteByUrl.Int64(), gasPerByteByUrl.Int64()*int64(contentLength))
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
		return
	}

	switch event.Format {
	case "csv":
		reader := strings.NewReader(string(body))
		err = IngestCSV(ctx, reader, dbC)
		err = db.ApplyChanges(ctx)
		if err != nil {
			logger.Errorf("Unmarshal csv ERROR: ", err)
			info := fmt.Sprintf("Unmarshal csv ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
			return
		} else {
			statement = true
		}
	case "ndjson":

		reader := strings.NewReader(string(body))
		// insert Data
		err = dbC.IndexNDJSON(ctx, reader)
		if err != nil {
			logger.Errorf("Import Data ERROR: ", err)
			info := fmt.Sprintf("Import Data ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
			return
		} else {
			statement = true
		}
	}
	//response to oracle
	_, err = config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, "")
	if err != nil {
		logger.Errorf("Req function get an Error : ", err)
	} else {
		logger.Info("[", event.ColName, "]", "Import Data success")
	}
}
