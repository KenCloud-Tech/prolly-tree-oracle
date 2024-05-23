package api

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"context"
	"fmt"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ipld/go-ipld-prime"
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
	var data []byte

	colName := event.ColName
	ctx := context.Background()
	db := config.Dbs[event.Cid]
	dbC, err := db.Collection(colName, "")
	if err != nil {
		log.Println("Get collection ERROR: ", err)
		info := fmt.Sprintf("Get collection ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	method := event.Val.Method

	switch method {
	case "equal":
		query := indexer.Query{
			Equal: map[string]ipld.Node{
				event.Val.K: creatNode(event),
			},
		}
		results, err := dbC.Search(ctx, query)
		if err != nil {
			log.Println("[", event.ColName, "]", "Search data  ERROR: ", err)
			info := fmt.Sprintf("Search data  ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.SearchRsp(tps, event.ReqID, statement, data, event.CallBack, event.Sender, info)
			return
		}
		record := <-results
		node := record.Data
		if node == nil {
			statement = false
			//response to oracle
			config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, "Data is not exist")
			return
		}
		data = nodeTobyte(node)
		statement = true
	case "compare":
		query := indexer.Query{
			Compare: &indexer.CompareCondition{Cmp: indexer.Op(event.Val.ComOp), IndexName: event.Val.K, IndexVal: creatNode(event)},
		}
		results, err := dbC.Search(ctx, query)
		if err != nil {
			log.Println("[", event.ColName, "]", "Search data  ERROR: ", err)
			info := fmt.Sprintf("Search data  ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.SearchRsp(tps, event.ReqID, statement, data, event.CallBack, event.Sender, info)
			return
		}
		record := <-results
		node := record.Data
		if node == nil {
			statement = false
			//response to oracle
			config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, "Data is not exist")
			return
		}
		data = nodeTobyte(node)
		statement = true
	case "sort":
		query := indexer.Query{
			Sort: event.Val.Str,
		}
		results, err := dbC.Search(ctx, query)
		if err != nil {
			log.Println("[", event.ColName, "]", "Search data  ERROR: ", err)
			info := fmt.Sprintf("Search data  ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.SearchRsp(tps, event.ReqID, statement, data, event.CallBack, event.Sender, info)
			return
		}
		record := <-results
		node := record.Data
		if node == nil {
			statement = false
			//response to oracle
			config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, "Data is not exist")
			return
		}
		data = nodeTobyte(node)
		statement = true
	case "limit":
		query := indexer.Query{
			Limit: int(event.Val.Integer),
		}
		results, err := dbC.Search(ctx, query)
		if err != nil {
			log.Println("[", event.ColName, "]", "Search data  ERROR: ", err)
			info := fmt.Sprintf("Search data  ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.SearchRsp(tps, event.ReqID, statement, data, event.CallBack, event.Sender, info)
			return
		}
		record := <-results
		node := record.Data
		if node == nil {
			statement = false
			//response to oracle
			config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, "Data is not exist")
			return
		}
		data = nodeTobyte(node)
		statement = true
	case "skip":
		query := indexer.Query{
			Skip: int(event.Val.Integer),
		}
		results, err := dbC.Search(ctx, query)
		if err != nil {
			log.Println("[", event.ColName, "]", "Search data  ERROR: ", err)
			info := fmt.Sprintf("Search data  ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.SearchRsp(tps, event.ReqID, statement, data, event.CallBack, event.Sender, info)
			return
		}
		record := <-results
		node := record.Data
		if node == nil {
			statement = false
			//response to oracle
			config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, "Data is not exist")
			return
		}
		data = nodeTobyte(node)
		statement = true
	}
	//response to oracle
	_, err = config.OracleContract.SearchRsp(tps, event.ReqID, statement, data, event.CallBack, event.Sender, "")
	if err != nil {
		log.Println("Req function get an error : ", err)
	} else {
		log.Println("[", event.ColName, "]", "Get data success")
	}
}
