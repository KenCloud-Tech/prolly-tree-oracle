package api

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"context"
	"fmt"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"strings"
)

func CreatEventListener() {
	// Create channels for logs
	Logs := make(chan *Oracle.OracleCreate)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}
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
			create(event)
		}
	}
}

// create memory collection
func create(event *Oracle.OracleCreate) {
	var statement bool
	tps := GenTransactOpts(config.GasLimit)

	cid := event.Cid
	if db, ok := config.Dbs[cid]; ok {
		colName := event.ColName
		pK := strings.Split(event.PrimaryKey, ",")
		_, err := db.Collection(colName, pK...)
		if err != nil {
			log.Println("Create memory db ERROR: ", err)
			info := fmt.Sprintf("Create memory db ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.CreatRsp(tps, event.ReqID, statement, db.RootCid().String(), cid, event.ColName, event.Owner, info)
			return
		}
		delete(config.Dbs, cid)
		config.Dbs[db.RootCid().String()] = db
		statement = true
		//response to oracle
		_, err = config.OracleContract.CreatRsp(tps, event.ReqID, statement, db.RootCid().String(), cid, event.ColName, event.Owner, "")
		if err != nil {
			log.Println("Req function get an error : ", err)
			delete(config.Dbs, event.ColName)
		} else {
			log.Println("[", event.ColName, "]", "Create memory db success")
		}
	} else {
		db, err := indexer.NewMemoryDatabase()
		if err != nil {
			log.Println("New Memory Database ERROR: ", err)
			info := fmt.Sprintf("New Memory Database ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.CreatRsp(tps, event.ReqID, statement, "", "", event.ColName, event.Owner, info)
			return
		}
		colName := event.ColName
		pK := strings.Split(event.PrimaryKey, ",")
		_, err = db.Collection(colName, pK...)
		if err != nil {
			log.Println("Create collection ERROR: ", err)
			info := fmt.Sprintf("Create collection ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.CreatRsp(tps, event.ReqID, statement, "", "", event.ColName, event.Owner, info)
			return
		}
		config.Dbs[db.RootCid().String()] = db
		statement = true
		//response to oracle
		_, err = config.OracleContract.CreatRsp(tps, event.ReqID, statement, db.RootCid().String(), "", event.ColName, event.Owner, "")
		if err != nil {
			log.Println("Req function get an error : ", err)
			delete(config.Dbs, event.ColName)
		} else {
			log.Println("[", event.ColName, "]", "Create memory db success")
		}
	}

}
