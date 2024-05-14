package api

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"Oracle.com/golangServer/model"
	"context"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func CreatEventListener() {
	// Create channels for logs
	createLogs := make(chan *Oracle.OracleCreate)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	eventSub, err := config.OracleContract.WatchCreate(opts, createLogs)
	if err != nil {
		log.Fatal("Failed to subscribe to Create events:", err)
	}
	// start Listening...
	log.Println("CreateEvent Listening :", config.ContractAddress)
	// 循环遍历事件订阅和错误通道，等待事件发生或错误发生
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event CREAT]:", err)
		case event := <-createLogs:
			log.Println("Received creat event")
			do(event)
		}
	}
}

func do(event *Oracle.OracleCreate) {
	// create memory db
	var statement bool
	ctx := context.Background()
	db, err := indexer.NewMemoryDatabase()
	if err != nil {
		log.Println("Create memory db ERROR: ", err)
		statement = false
	}
	collection, err := db.Collection("users", "name")
	if err != nil {
		log.Println("Create memory db ERROR: ", err)
		statement = false
	}
	config.Dbs[event.DbName] = &model.OracleModel{DbID: event.DbName, Ctx: ctx, Db: collection}
	statement = true
	//response to oracle
	tps := &bind.TransactOpts{
		From: common.BytesToAddress([]byte(config.OracleOwner)),
	}
	config.OracleContract.CreatRsp(tps, event.ReqID, statement, event.DbName, event.Owner)
	if v, ok := config.Dbs[event.DbName]; ok {
		log.Println("Create memory db Success: ", v)
	}
}
