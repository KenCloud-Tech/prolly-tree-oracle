package api

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ipld/go-ipld-prime/printer"
	"log"
)

func SearchEventListener() {
	// Get channels for logs
	Logs := make(chan *Oracle.OracleSearch)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	eventSub, err := config.OracleContract.WatchSearch(opts, Logs)
	if err != nil {
		log.Fatal("Failed to subscribe to Search events:", err)
	}
	// start Listening...
	log.Println("SearchEvent Listening ...")
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event SEARCH]:", err)
		case event := <-Logs:
			log.Println("Received search event ", event.ReqID)
			search(event)
		}
	}
}

// Search data from memory db
func search(event *Oracle.OracleSearch) {
	var statement bool
	tps := GenTransactOpts(config.GasLimit)
	var data string

	colName := event.ColName
	db := config.Dbs[event.DbName]
	dbC, err := db.Collection(colName, "")
	if err != nil {
		log.Println("Get collection ERROR: ", err)
		info := fmt.Sprintf("Get collection ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	querys := make([]Queryer, 1)
	err = json.Unmarshal(event.Querys, &querys)
	if err != nil {
		log.Println("Unmarshal ERROR: ", err)
		info := fmt.Sprintf("Unmarshal ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	k := 1
	for _, q := range querys {
		data += "\nQuery:" + string(k) + "\n"
		k++
		query := q.Queryer2indexerQ()
		results, err := dbC.Search(context.Background(), query)
		if err != nil {
			v := "Search fail."
			data += v
			continue
		}
		for record := range results {
			if record.Data == nil {
				v := "Data is not exist"
				data += v
				continue
			}
			data = data + "\n" + printer.Sprint(record.Data)
		}
	}
	result, err := json.Marshal(data)
	if err != nil {
		log.Println("Marshal Results ERROR: ", err)
		info := fmt.Sprintf("Marshal Results ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	statement = true
	//response to oracle
	_, err = config.OracleContract.SearchRsp(tps, event.ReqID, statement, result, event.CallBack, event.Sender, "")
	if err != nil {
		log.Println("Req function get an error : ", err)
	} else {
		log.Println("[", event.ColName, "]", "Search data success")
	}
}
