package api

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"context"
	"fmt"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/basicnode"
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

	dbName := event.DbName
	ctx, dbC := config.Dbs[dbName].Ctx, config.Dbs[dbName].Db
	method := event.Method

	switch method {
	case "equal":
		query := indexer.Query{
			Equal: map[string]ipld.Node{
				event.Val.K: creatNode(event),
			},
		}
		results, err := dbC.Search(ctx, query)
		if err != nil {
			log.Println("[", event.DbName, "]", "Search data  ERROR: ", err)
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
	//case "compare":
	case "sort":
		query := indexer.Query{
			Sort: event.Val.Str,
		}
		results, err := dbC.Search(ctx, query)
		if err != nil {
			log.Println("[", event.DbName, "]", "Search data  ERROR: ", err)
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
			log.Println("[", event.DbName, "]", "Search data  ERROR: ", err)
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
			log.Println("[", event.DbName, "]", "Search data  ERROR: ", err)
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
	config.OracleContract.SearchRsp(tps, event.ReqID, statement, data, event.CallBack, event.Sender, "")
	log.Println("[", event.DbName, "]", "Get data success")
}

func creatNode(event *Oracle.OracleSearch) (node ipld.Node) {
	switch event.Val.ComDataType {
	case "string":
		node = basicnode.NewString(event.Val.Str)
	case "int":
		node = basicnode.NewInt(event.Val.Integer)
	case "uint":
		node = basicnode.NewUint(event.Val.Uinteger)
	case "bytes":
		node = basicnode.NewBytes(event.Val.Bytess)
	case "bool":
		node = basicnode.NewBool(event.Val.Boolean)
	}
	return
}
