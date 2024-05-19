package api

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"Oracle.com/golangServer/model"
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
			creat(event)
		}
	}
}

// create memory db
func creat(event *Oracle.OracleCreate) {
	var statement bool
	tps := GenTransactOpts(config.GasLimit)

	ctx := context.Background()
	db, err := indexer.NewMemoryDatabase()
	if err != nil {
		log.Println("New Memory Database ERROR: ", err)
		info := fmt.Sprintf("New Memory Database ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.CreatRsp(tps, event.ReqID, statement, event.DbName, event.Owner, info)
		return
	}
	name := event.DbName
	pK := strings.Split(event.PrimaryKey, ",")
	collection, err := db.Collection(name, pK...)
	if err != nil {
		log.Println("Create memory db ERROR: ", err)
		info := fmt.Sprintf("Create memory db ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.CreatRsp(tps, event.ReqID, statement, event.DbName, event.Owner, info)
		return
	}

	config.Dbs[event.DbName] = &model.OracleModel{DbName: event.DbName, Ctx: ctx, Db: collection}
	statement = true
	//response to oracle
	_, err = config.OracleContract.CreatRsp(tps, event.ReqID, statement, event.DbName, event.Owner, "")
	if err != nil {
		log.Println("Req function get an error : ", err)
		delete(config.Dbs, event.DbName)
	} else {
		log.Println("[", event.DbName, "]", "Create memory db success")
	}
}
