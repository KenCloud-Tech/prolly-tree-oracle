package api

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
)

func GetCollections() {
	// Get channels for logs
	Logs := make(chan *Oracle.OracleGetCol)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	eventSub, err := config.OracleContract.WatchGetCol(opts, Logs)
	if err != nil {
		log.Fatal("Failed to subscribe to GetCollections events:", err)
	}
	// start Listening...
	log.Println("GetCollectionsEvent Listening ...")
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event GetCollections]:", err)
		case event := <-Logs:
			log.Println("Received GetCollections event ", event.ReqID)
			var statement bool
			tps := GenTransactOpts(config.GasLimit)

			db := config.Dbs[event.Cid]
			cols := db.ListCollections()
			jsonBytes, err := json.Marshal(cols)
			if err != nil {
				log.Println("[", event.Cid, "]", "Trans to json ERROR: ", err)
				info := fmt.Sprintf("Trans to json ERROR: %v", err)
				statement = false
				//response to oracle
				config.OracleContract.GetColRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
			}
			result, err := json.Marshal([][]byte{jsonBytes})
			if err != nil {
				log.Println("Marshal Results ERROR: ", err)
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
				log.Println("Req function get an error : ", err)
			} else {
				log.Println("[Get collections success]")
			}
		}
	}
}

func GetIndexes() {
	// Get channels for logs
	Logs := make(chan *Oracle.OracleGetIndex)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	eventSub, err := config.OracleContract.WatchGetIndex(opts, Logs)
	if err != nil {
		log.Fatal("Failed to subscribe to GetIndexes events:", err)
	}
	// start Listening...
	log.Println("GetIndexesEvent Listening ...")
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event GetIndexes]:", err)
		case event := <-Logs:
			log.Println("Received GetIndexes event ", event.ReqID)
			var statement bool
			ctx := context.Background()
			tps := GenTransactOpts(config.GasLimit)

			db := config.Dbs[event.Cid]
			colName := event.ColName
			col, err := db.Collection(colName, "")
			if err != nil {
				log.Println("Get collection ERROR: ", err)
				info := fmt.Sprintf("Get collection ERROR: %v", err)
				statement = false
				//response to oracle
				config.OracleContract.GetIndexRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
				return
			}
			indx, err := col.Indexes(ctx)
			if err != nil {
				log.Println("Get indexes ERROR: ", err)
				info := fmt.Sprintf("Get indexes ERROR: %v", err)
				statement = false
				//response to oracle
				config.OracleContract.GetIndexRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
				return
			}
			var indexes []string
			for _, i := range indx {
				indexes = append(indexes, i.Fields()...)
			}
			jsonBytes, err := json.Marshal(indexes)
			if err != nil {
				log.Println("[", event.Cid, "]", "Trans to json ERROR: ", err)
				info := fmt.Sprintf("Trans to json ERROR: %v", err)
				statement = false
				//response to oracle
				config.OracleContract.GetIndexRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
			}

			result, err := json.Marshal([][]byte{jsonBytes})
			if err != nil {
				log.Println("Marshal Results ERROR: ", err)
				info := fmt.Sprintf("Marshal Results ERROR: %v", err)
				statement = false
				//response to oracle
				config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
				return
			}
			statement = true
			//response to oracle
			_, err = config.OracleContract.GetIndexRsp(tps, event.ReqID, statement, result, event.CallBack, event.Sender, "")
			if err != nil {
				log.Println("Req function get an error : ", err)
			} else {
				log.Println("[Get indexes success]")
			}
		}
	}
}
