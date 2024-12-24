package api

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	cid "github.com/ipfs/go-cid"
	"log"
)

func GetRootCidEventListener(ctx context.Context) {
	// Create a channel for logs
	logs := make(chan *Oracle.OracleGetRootCid)

	// Subscribe to each event
	opts := &bind.WatchOpts{Context: ctx, Start: nil}

	// Start Listening...
	log.Println("GetRootCidEvent Listening ...")
	eventSub, err := config.OracleContract.WatchGetRootCid(opts, logs)
	if err != nil {
		log.Fatal("Failed to subscribe to Get events:", err)
	}
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event GETROOTCID]:", err)
			break
		case event := <-logs:
			log.Println("Received get root cid event ", event.ReqID)
			getRootCid(ctx, event)
		}
	}
}

// Get root cid from database
func getRootCid(ctx context.Context, event *Oracle.OracleGetRootCid) {
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	dbName := event.DbName
	db, exists := config.Dbs[dbName]
	if !exists || db == nil {
		log.Println("Database not found or is nil for dbName:", dbName)
		statement = false
		// Response to oracle
		config.OracleContract.GetRootCidRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, "Database is not initialized")
		return
	}

	log.Println("Attempting to get RootCid for database:", dbName)
	rootCid := db.RootCid()
	if rootCid == cid.Undef {
		statement = false
		// Response to oracle
		config.OracleContract.GetRootCidRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, "Root CID is undefined")
		return
	}

	statement = true
	data := rootCid.Bytes() // Convert cid to bytes
	// Response to oracle
	_, err := config.OracleContract.GetRootCidRsp(tps, event.ReqID, statement, data, event.CallBack, event.Sender, "")
	if err != nil {
		log.Println("Req function encountered an error: ", err)
	} else {
		log.Println("[", dbName, "]", "Get Root CID success")
	}
}
