package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func SearchEventListener(ctx context.Context) {
	for {
		restartflag := false
		// Get channels for logs
		Logs := make(chan *Oracle.OracleSearch)
		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}
		eventSub, err := config.OracleContract.WatchSearch(opts, Logs)
		if err != nil {
			log.Fatal("Failed to subscribe to Search events:", err)
			time.Sleep(5 * time.Second)
			close(Logs)
			continue
		}
		// start Listening...
		log.Println("SearchEvent Listening ...")
		for {
			select {
			case err := <-eventSub.Err():
				log.Println("[Error in Event SEARCH]:", err)
				restartflag = true
				break
			case event := <-Logs:
				log.Println("Received search event ", event.ReqID)
				search(ctx, event)
			}
			if restartflag {
				log.Println("[restart SearchEventListener for loop]:", err)
				time.Sleep(5 * time.Second)
				close(Logs)
				break
			}
		}
	}
}

// Search Data from memory db
func search(ctx context.Context, event *Oracle.OracleSearch) {
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

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
	dbC, err := db.Collection(ctx, colName, "")
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
		resultChan, err := dbC.Search(ctx, query)
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
