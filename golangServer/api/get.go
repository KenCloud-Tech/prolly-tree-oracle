package api

import (
	"context"
	"fmt"
	"log"
	"time"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetEventListener(ctx context.Context) {

	for {
		// Get channels for logs
		Logs := make(chan *Oracle.OracleGet)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}
		eventSub, err := config.OracleContract.WatchGet(opts, Logs)
		if err != nil {
			log.Fatal("Failed to subscribe to Get events:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		// start Listening...
		log.Println("GetEvent Listening ...")
		for {
			select {
			case err := <-eventSub.Err():
				log.Println("[Error in Event GET]:", err)
				break
			case event := <-Logs:
				log.Println("Received get event ", event.ReqID)
				get(ctx, event)
			}
			if err != nil {
				log.Println("[break GetEventListener for loop]:", err)
				time.Sleep(5 * time.Second)
				break
			}
		}
	}
}

// Get Data from memory db
func get(ctx context.Context, event *Oracle.OracleGet) {
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	colName := event.ColName
	db := config.Dbs[event.DbName]
	col, err := db.Collection(ctx, colName, "")
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
		log.Println("[", event.ColName, "]", "Get Data ERROR: ", err)
		info := fmt.Sprintf("Get Data ERROR: %v", err)
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
	result := nodeTobyte(node)
	statement = true
	//response to oracle
	_, err = config.OracleContract.GetRsp(tps, event.ReqID, statement, result, event.CallBack, event.Sender, "")
	if err != nil {
		log.Println("Req function get an Error : ", err)
	} else {
		log.Println("[", event.ColName, "]", "Get Data success")
	}
}
