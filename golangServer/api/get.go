package api

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
)

func GetEventListener() {
	// Get channels for logs
	Logs := make(chan *Oracle.OracleGet)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	eventSub, err := config.OracleContract.WatchGet(opts, Logs)
	if err != nil {
		log.Fatal("Failed to subscribe to Get events:", err)
	}
	// start Listening...
	log.Println("GetEvent Listening ...")
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event CREAT]:", err)
		case event := <-Logs:
			log.Println("Received get event ", event.ReqID)
			get(event)
		}
	}
}

// Get data from memory db
func get(event *Oracle.OracleGet) {
	var statement bool
	tps := GenTransactOpts(config.GasLimit)

	colName := event.ColName
	ctx := context.Background()
	db := config.Dbs[event.Cid]
	col, err := db.Collection(colName, "")
	if err != nil {
		log.Println("Get collection ERROR: ", err)
		info := fmt.Sprintf("Get collection ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	node, err := col.Get(ctx, event.RecordID)
	if err != nil {
		log.Println("[", event.ColName, "]", "Get data ERROR: ", err)
		info := fmt.Sprintf("Get data ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	if node == nil {
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, "Data is not exist")
		return
	}
	data := nodeTobyte(node)
	statement = true
	//response to oracle
	_, err = config.OracleContract.GetRsp(tps, event.ReqID, statement, data, event.CallBack, event.Sender, "")
	if err != nil {
		log.Println("Req function get an error : ", err)
	} else {
		log.Println("[", event.ColName, "]", "Get data success")
	}
}
