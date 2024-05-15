package api

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
)

func PutEventListener() {
	// Put channels for logs
	Logs := make(chan *Oracle.OraclePut)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	eventSub, err := config.OracleContract.WatchPut(opts, Logs)
	if err != nil {
		log.Fatal("Failed to subscribe to Put events:", err)
	}
	// start Listening...
	log.Println("PutEvent Listening ...")
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event CREAT]:", err)
		case event := <-Logs:
			log.Println("Received put event ", event.ReqID)
			put(event)
		}
	}
}

// Put data to memory db
func put(event *Oracle.OraclePut) {
	var statement bool
	tps := &bind.TransactOpts{
		From: common.BytesToAddress([]byte(config.OracleOwner)),
	}

	dbName := event.DbName
	ctx, dbC := config.Dbs[dbName].Ctx, config.Dbs[dbName].Db

	strData := string(event.Data)
	reader := strings.NewReader(strData)
	// insert data
	err := dbC.IndexNDJSON(ctx, reader)
	if err != nil {
		log.Println("Put data ERROR: ", err)
		info := fmt.Sprintf("Put data ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.PutRsp(tps, event.ReqID, statement, event.Sender, info)
		return
	} else {
		statement = true
	}
	//response to oracle
	config.OracleContract.PutRsp(tps, event.ReqID, statement, event.Sender, "")
	log.Println("[", event.DbName, "]", "Put data success")
}
