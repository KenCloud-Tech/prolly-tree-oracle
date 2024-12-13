package api

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func IndexEventListener(ctx context.Context) {
	for {
		// Create channels for logs
		Logs := make(chan *Oracle.OracleIndex)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}
		eventSub, err := config.OracleContract.WatchIndex(opts, Logs)
		if err != nil {
			log.Fatal("Failed to subscribe to Index events:", err)
			time.Sleep(5 * time.Second)
			close(Logs)
			continue
		}
		// start Listening...
		log.Println("IndexEvent Listening ...")
		for {
			select {
			case err := <-eventSub.Err():
				log.Println("[Error in Event INDEX]:", err)
				break
			case event := <-Logs:
				log.Println("Received index event ", event.ReqID)
				index(ctx, event)
			}
			if err != nil {
				log.Println("[break IndexEventListener for loop]:", err)
				time.Sleep(5 * time.Second)
				close(Logs)
				break
			} else {
				log.Println("[continue IndexEvent for loop]:", err)
				time.Sleep(5 * time.Second)
				continue
			}
		}
	}
}

// Create index
func index(ctx context.Context, event *Oracle.OracleIndex) {
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	colName := event.ColName
	db := config.Dbs[event.DbName]
	dbC, err := db.Collection(ctx, colName, "")
	if err != nil {
		log.Println("Get collection ERROR: ", err)
		info := fmt.Sprintf("Get collection ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.IndexRsp(tps, event.ReqID, statement, event.Sender, info)
		return
	}
	pK := strings.Split(event.Key, ",")
	_, err = dbC.CreateIndex(ctx, pK...)
	if err != nil {
		log.Println("Create Index ERROR: ", err)
		info := fmt.Sprintf("Create Index ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.IndexRsp(tps, event.ReqID, statement, event.Sender, info)
		return
	}

	statement = true
	_, err = config.OracleContract.IndexRsp(tps, event.ReqID, statement, event.Sender, "")
	if err != nil {
		log.Println("Req function get an Error : ", err)
	} else {
		log.Println("[", event.ColName, "]", "Create index success")
	}
}
