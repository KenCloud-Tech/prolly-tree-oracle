package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var gasPerByteByUrl *big.Int
var gasPerByteByUrlOnce sync.Once

func ImportEventListener(ctx context.Context) {
	for {
		restartflag := false
		// Import channels for logs
		Logs := make(chan *Oracle.OracleImportFromUrl)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}
		eventSub, err := config.OracleContract.WatchImportFromUrl(opts, Logs)
		if err != nil {
			log.Fatal("Failed to subscribe to Import events:", err)
			time.Sleep(5 * time.Second)
			close(Logs)
			continue
		}
		// start Listening...
		log.Println("ImportEvent Listening ...")
		for {
			select {
			case err := <-eventSub.Err():
				log.Println("[Error in Event IMPORT:", err)
				restartflag = true
				break
			case event := <-Logs:
				log.Println("Received put event ", event.ReqID)
				importByUrl(ctx, event)
			}
			if restartflag {
				log.Println("[restart ImportEventListener for loop]:", err)
				time.Sleep(5 * time.Second)
				close(Logs)
				break
			}
		}
	}
}

// Import Data to memory db
func importByUrl(ctx context.Context, event *Oracle.OracleImportFromUrl) {
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	bigInt := big.Int{}
	size := bigInt.SetInt64(0)
	colName := event.ColName
	dbName := event.DbName
	db := config.Dbs[dbName]
	dbC, err := db.Collection(ctx, colName, "")
	if err != nil {
		log.Println("Get collection ERROR: ", err)
		info := fmt.Sprintf("Get collection ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
		return
	}
	resp, err := http.Get(event.Url)
	defer resp.Body.Close()
	ContentLength := resp.ContentLength
	size = bigInt.SetInt64(ContentLength)

	if err != nil {
		log.Println("Get datas ERROR: ", err)
		info := fmt.Sprintf("Get datas ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
		return
	}
	if resp.StatusCode != 200 {
		log.Println("Get datas Fail, StatusCode = ", resp.StatusCode)
		info := fmt.Sprintf("Get datas Fail, StatusCode = %v", resp.StatusCode)
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
		return
	}
	if ContentLength == 0 {
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
		log.Println("Get gas per bytes ERROR: ", err)
		info := fmt.Sprintf("Get gas per bytes ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
		return
	}

	if event.Value.Int64() > gasPerByteByUrl.Int64()*ContentLength {
		info := fmt.Sprintf("The data is too large and the gas paid is insufficient.")
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
		return
	}

	switch event.Format {
	case "csv":
		body, err := io.ReadAll(resp.Body)
		reader := strings.NewReader(string(body))
		err = IngestCSV(ctx, reader, dbC)
		err = db.ApplyChanges(ctx)
		if err != nil {
			log.Println("Unmarshal csv ERROR: ", err)
			info := fmt.Sprintf("Unmarshal csv ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, size, event.Sender, info)
			return
		} else {
			statement = true
		}
	case "ndjson":
		body, _ := io.ReadAll(resp.Body)
		reader := strings.NewReader(string(body))
		// insert Data
		err = dbC.IndexNDJSON(ctx, reader)
		if err != nil {
			log.Println("Import Data ERROR: ", err)
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
		log.Println("Req function get an Error : ", err)
	} else {
		log.Println("[", event.ColName, "]", "Import Data success")
	}
}
