package api

import (
	"context"
	"fmt"
	"log"
	"strings"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func PutEventListener(ctx context.Context) {
	// Put channels for logs
	Logs := make(chan *Oracle.OraclePut)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: ctx, Start: nil}
	eventSub, err := config.OracleContract.WatchPut(opts, Logs)
	if err != nil {
		log.Fatal("Failed to subscribe to Put events:", err)
	}
	// start Listening...
	log.Println("PutEvent Listening ...")
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event PUT]:", err)
		case event := <-Logs:
			log.Println("Received put event ", event.ReqID)
			put(ctx, event)
		}
	}
}

// Put Data to memory db
func put(ctx context.Context, event *Oracle.OraclePut) {
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	log.Println("CreateEvent DbName ...", event.DbName)
	log.Println("CreateEvent ColName ...", event.ColName)
	colName := event.ColName
	dbName := event.DbName
	db := config.Dbs[dbName]
	dbC, err := db.Collection(ctx, colName, "")
	if err != nil {
		log.Println("Get collection ERROR: ", err)
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
		log.Println("Put Data ERROR: ", err)
		info := fmt.Sprintf("Put Data ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.PutRsp(tps, event.ReqID, statement, event.Sender, info)
		return
	} else {
		statement = true
	}
	//response to oracle
	_, err = config.OracleContract.PutRsp(tps, event.ReqID, statement, event.Sender, "")
	if err != nil {
		log.Println("Req function get an Error : ", err)
	} else {
		log.Println("[", event.ColName, "]", "Put Data success")
	}
}
