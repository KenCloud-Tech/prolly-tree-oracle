package api

import (
	"context"
	"fmt"
	"log"
	"strings"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func CreatEventListener(ctx context.Context) {
	// Create channels for logs
	Logs := make(chan *Oracle.OracleCreate)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: ctx, Start: nil}
	eventSub, err := config.OracleContract.WatchCreate(opts, Logs)
	if err != nil {
		log.Fatal("Failed to subscribe to Create events:", err)
	}
	// start Listening...
	log.Println("CreateEvent Listening ...")
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event CREAT]:", err)
		case event := <-Logs:
			log.Println("Received creat event ", event.ReqID)
			create(ctx, event)
		}
	}
}

// create memory collection
func create(ctx context.Context, event *Oracle.OracleCreate) {
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	dbName := event.DbName
	if db, ok := config.Dbs[dbName]; ok {
		colName := event.ColName
		pK := strings.Split(event.PrimaryKey, ",")
		_, err := db.Collection(ctx, colName, pK...)
		if err != nil {
			log.Println("Create collection ERROR: ", err)
			info := fmt.Sprintf("Create collection ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.CreatRsp(tps, event.ReqID, statement, dbName, colName, event.Owner, info)
			return
		}
		statement = true
		//response to oracle
		_, err = config.OracleContract.CreatRsp(tps, event.ReqID, statement, dbName, colName, event.Owner, "")
		if err != nil {
			log.Println("Req function get an Error : ", err)
			db.DeleteCol(colName)
		} else {
			log.Println("[", colName, "]", "Create collection success")
		}
	} else {
		db, err := indexer.NewMemoryDatabase()
		if err != nil {
			log.Println("New Memory Database ERROR: ", err)
			info := fmt.Sprintf("New Memory Database ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.CreatRsp(tps, event.ReqID, statement, dbName, event.ColName, event.Owner, info)
			return
		}
		config.Dbs[dbName] = db
		colName := event.ColName
		pK := strings.Split(event.PrimaryKey, ",")
		_, err = db.Collection(ctx, colName, pK...)
		if err != nil {
			log.Println("Create collection ERROR: ", err)
			info := fmt.Sprintf("Create collection ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.CreatRsp(tps, event.ReqID, statement, dbName, colName, event.Owner, info)
			return
		}
		statement = true
		//response to oracle
		_, err = config.OracleContract.CreatRsp(tps, event.ReqID, statement, dbName, colName, event.Owner, "")
		if err != nil {
			log.Println("Req function get an Error : ", err)
			db.DeleteCol(colName)
		} else {
			log.Println("[", colName, "]", "Create memory db success")
		}
	}

}
