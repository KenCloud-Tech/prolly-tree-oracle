package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetCollections(ctx context.Context) {

	for {
		restartflag := false
		// Get channels for logs
		Logs := make(chan *Oracle.OracleGetCol)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}
		eventSub, err := config.OracleContract.WatchGetCol(opts, Logs)
		if err != nil {
			log.Fatal("Failed to subscribe to GetCollections events:", err)
			time.Sleep(5 * time.Second)
			close(Logs)
			continue
		}
		// start Listening...
		log.Println("GetCollectionsEvent Listening ...")
		for {
			select {
			case err := <-eventSub.Err():
				log.Println("[Error in Event GetCollections]:", err)
				restartflag = true
				break
			case event := <-Logs:
				log.Println("Received GetCollections event ", event.ReqID)
				var statement bool
				tps := GenTransactOpts(ctx, config.GasLimit)

				db := config.Dbs[event.DbName]
				cols, err := db.ListCollections(ctx)
				if err != nil {
					log.Println("[", event.DbName, "]", "List Collections ERROR: ", err)
					info := fmt.Sprintf("List Collections ERROR: %v", err)
					statement = false
					//response to oracle
					config.OracleContract.GetColRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
				}

				jsonBytes, err := json.Marshal(cols)
				if err != nil {
					log.Println("[", event.DbName, "]", "Trans to json ERROR: ", err)
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
					log.Println("Req function get an Error : ", err)
				} else {
					log.Println("[Get collections success]")
				}
			}
			if restartflag {
				log.Println("[restart GetCollections for loop]:", err)
				time.Sleep(5 * time.Second)
				close(Logs)
				break
			}
		}
	}
}

func GetIndexes(ctx context.Context) {
	for {
		restartflag := false
		// Get channels for logs
		Logs := make(chan *Oracle.OracleGetIndex)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}
		eventSub, err := config.OracleContract.WatchGetIndex(opts, Logs)
		if err != nil {
			log.Fatal("Failed to subscribe to GetIndexes events:", err)
			time.Sleep(5 * time.Second)
			close(Logs)
			continue
		}
		// start Listening...
		log.Println("GetIndexesEvent Listening ...")
		for {
			select {
			case err := <-eventSub.Err():
				log.Println("[Error in Event GetIndexes]:", err)
				restartflag = true
				break
			case event := <-Logs:
				log.Println("Received GetIndexes event ", event.ReqID)
				var statement bool
				ctx := ctx
				tps := GenTransactOpts(ctx, config.GasLimit)

				db := config.Dbs[event.DbName]
				colName := event.ColName
				col, err := db.Collection(ctx, colName, "")
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
					log.Println("[", event.DbName, "]", "Trans to json ERROR: ", err)
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
					log.Println("Req function get an Error : ", err)
				} else {
					log.Println("[Get indexes success]")
				}
			}
			if restartflag {
				log.Println("[restart GetIndexes for loop]:", err)
				time.Sleep(5 * time.Second)
				close(Logs)
				break
			}
		}
	}
}

func GetRootCid(ctx context.Context) {
	for {
		// Get channels for logs
		Logs := make(chan *Oracle.OracleGetRootCid)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}
		eventSub, err := config.OracleContract.WatchGetRootCid(opts, Logs)
		if err != nil {
			log.Fatal("Failed to subscribe to WatchGetRootCid events:", err)
			time.Sleep(5 * time.Second)
			close(Logs)
			continue
		}
		// start Listening...
		log.Println("WatchGetRootCid Listening ...")
		for {
			select {
			case err := <-eventSub.Err():
				log.Println("[Error in Event WatchGetRootCid]:", err)
				break
			case event := <-Logs:
				log.Println("Received WatchGetRootCid event ", event.ReqID)
				var statement bool
				tps := GenTransactOpts(ctx, config.GasLimit)

				db := config.Dbs[event.DbName]
				rootCid := db.RootCid().String()
				marshal, err := json.Marshal(rootCid)
				result, err := json.Marshal([][]byte{marshal})
				if err != nil {
					log.Println("Marshal Results ERROR: ", err)
					info := fmt.Sprintf("Marshal Results ERROR: %v", err)
					statement = false
					//response to oracle
					config.OracleContract.GetRootCidRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
					return
				}
				statement = true
				//response to oracle
				_, err = config.OracleContract.GetRootCidRsp(tps, event.ReqID, statement, result, event.CallBack, event.Sender, "")
				if err != nil {
					log.Println("Req function get an Error : ", err)
				} else {
					log.Println("[Get RootCid success]")
				}
			}
			if err != nil {
				log.Println("[break GetRootCid for loop]:", err)
				time.Sleep(5 * time.Second)
				close(Logs)
				break
			} else {
				log.Println("[continue GetRootCid for loop]:", err)
				time.Sleep(5 * time.Second)
				continue
			}
		}
	}
}
