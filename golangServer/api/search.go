package api

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
)

type SearchResults struct {
	Status bool
	Data   string
	Error  string
}
type Results struct {
	Id           int
	QueryResults []SearchResults
	Error        string
}

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

// Search Data from memory db
func search(event *Oracle.OracleSearch) {
	var statement bool
	tps := GenTransactOpts(config.GasLimit)

	colName := event.ColName
	db, ok := config.Dbs[event.DbName]
	if !ok {
		log.Println("Collection does not exist")
		info := fmt.Sprintf("Collection does not exist")
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	dbC, err := db.Collection(colName, "")
	if err != nil {
		log.Println("Get collection ERROR: ", err)
		info := fmt.Sprintf("Get collection ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	var querys []Queryer
	err = json.Unmarshal(event.Querys, &querys)
	if err != nil {
		log.Println("Unmarshal ERROR: ", err)
		info := fmt.Sprintf("Unmarshal ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	var r []Results
	for i, q := range querys {
		query := q.Queryer2indexerQ()
		resultChan, err := dbC.Search(context.Background(), query)
		if err != nil {
			r = append(r, Results{
				Id:    i,
				Error: err.Error(),
			})
			continue
		}
		var records []SearchResults
		for record := range resultChan {
			if record.Data == nil {
				records = append(records, SearchResults{
					Status: false,
					Error:  "record not found",
				})
				continue
			}
			records = append(records, SearchResults{
				Status: true,
				Data:   string(nodeTobyte(record.Data)),
			})
		}
		r = append(r, Results{
			Id:           i,
			QueryResults: records,
		})
	}
	result, err := json.Marshal(r)
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
		log.Println("Req function get an Error : ", err)
	} else {
		log.Println("[", event.ColName, "]", "Search Data success")
	}
}
