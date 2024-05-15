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

func IndexEventListener() {
	// Create channels for logs
	Logs := make(chan *Oracle.OracleIndex)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	eventSub, err := config.OracleContract.WatchIndex(opts, Logs)
	if err != nil {
		log.Fatal("Failed to subscribe to Index events:", err)
	}
	// start Listening...
	log.Println("IndexEvent Listening ...")
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event INDEX]:", err)
		case event := <-Logs:
			log.Println("Received index event ", event.ReqID)
			index(event)
		}
	}
}

// Create index
func index(event *Oracle.OracleIndex) {
	var statement bool
	tps := &bind.TransactOpts{
		From: common.BytesToAddress([]byte(config.OracleOwner)),
	}
	dbName := event.DbName
	ctx, dbC := config.Dbs[dbName].Ctx, config.Dbs[dbName].Db
	pK := strings.Split(event.Key, ",")
	_, err := dbC.CreateIndex(ctx, pK...)
	if err != nil {
		log.Println("Create Index ERROR: ", err)
		info := fmt.Sprintf("Create Index ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.IndexRsp(tps, event.ReqID, statement, event.Sender, info)
		return
	}

	statement = true
	config.OracleContract.IndexRsp(tps, event.ReqID, statement, event.Sender, "")
	log.Println("[", event.DbName, "]", "Create index success")
}
