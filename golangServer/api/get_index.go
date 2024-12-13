package api

import (
	"context"
	"log"
	"time"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	cid "github.com/ipfs/go-cid"
)

func GetRootCidEventListener(ctx context.Context) {
	for {
		restartflag := false
		// Create a channel for logs
		logs := make(chan *Oracle.OracleGetRootCid)

		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}

		// Start Listening...
		log.Println("GetRootCidEvent Listening ...")
		eventSub, err := config.OracleContract.WatchGetRootCid(opts, logs)
		if err != nil {
			log.Fatal("Failed to subscribe to Get events:", err)
			time.Sleep(5 * time.Second)
			close(logs)
			continue
		}
		for {
			select {
			case err := <-eventSub.Err():
				log.Println("[Error in Event GETROOTCID]:", err)
				restartflag = true
				break
			case event := <-logs:
				log.Println("Received get root cid event ", event.ReqID)
				getRootCid(ctx, event)
			}
			if restartflag {
				log.Println("[restart GetRootCidEventListener for loop]:", err)
				time.Sleep(5 * time.Second)
				close(logs)
				break
			}
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
