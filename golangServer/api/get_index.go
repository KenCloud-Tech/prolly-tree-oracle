package api

import (
	"context"
	"time"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	cid "github.com/ipfs/go-cid"
	"go.uber.org/zap"
)

func GetRootCidEventListener(ctx context.Context, logger *zap.SugaredLogger) {
	for {
		// Create a channel for logs
		logs := make(chan *Oracle.OracleGetRootCid)

		// Subscribe to each event
		opts := &bind.WatchOpts{Context: ctx, Start: nil}

		// Start Listening...
		logger.Info("GetRootCidEvent Listening ...")
		eventSub, err := config.OracleContract.WatchGetRootCid(opts, logs)
		if err != nil {
			logger.Error("Failed to subscribe to Get events:", err)
			time.Sleep(5 * time.Second)
			close(logs)
			continue
		}
	LOOP:
		for {
			select {
			case err := <-eventSub.Err():
				logger.Error("[Error in Event GETROOTCID]:", err)
				break LOOP
			case event := <-logs:
				logger.Info("Received get root cid event ", event.ReqID)
				getRootCid(ctx, event, logger)
			}
		}
	}
}

// Get root cid from database
func getRootCid(ctx context.Context, event *Oracle.OracleGetRootCid, logger *zap.SugaredLogger) {
	var statement bool
	tps := GenTransactOpts(ctx, config.GasLimit)

	dbName := event.DbName
	db, exists := config.Dbs[dbName]
	if !exists || db == nil {
		logger.Error("Database not found or is nil for dbName: %s", dbName)
		statement = false
		// Response to oracle
		config.OracleContract.GetRootCidRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, "Database is not initialized")
		return
	}

	logger.Infof("Attempting to get RootCid for database:", dbName)
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
		logger.Error("Req function encountered an error: ", err)
	} else {
		logger.Info("[", dbName, "]", "Get Root CID success")
	}
}
