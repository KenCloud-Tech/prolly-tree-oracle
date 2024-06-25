package api

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"io"
	"log"
	"net/http"
	"strings"
)

func ImportEventListener() {
	// Import channels for logs
	Logs := make(chan *Oracle.OracleImportFromUrl)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	eventSub, err := config.OracleContract.WatchImportFromUrl(opts, Logs)
	if err != nil {
		log.Fatal("Failed to subscribe to Import events:", err)
	}
	// start Listening...
	log.Println("ImportEvent Listening ...")
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event IMPORT:", err)
		case event := <-Logs:
			log.Println("Received put event ", event.ReqID)
			importByUrl(event)
		}
	}
}

// Import data to memory db
func importByUrl(event *Oracle.OracleImportFromUrl) {
	var statement bool
	tps := GenTransactOpts(config.GasLimit)

	colName := event.ColName
	ctx := context.Background()
	dbName := event.DbName
	db := config.Dbs[dbName]
	dbC, err := db.Collection(colName, "")
	if err != nil {
		log.Println("Get collection ERROR: ", err)
		info := fmt.Sprintf("Get collection ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, event.Sender, info)
		return
	}

	resp, err := http.Get(event.Url)
	defer resp.Body.Close()
	if err != nil {
		log.Println("Get datas ERROR: ", err)
		info := fmt.Sprintf("Get datas ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, event.Sender, info)
		return
	}
	if resp.StatusCode != 200 {
		log.Println("Get datas Fail, StatusCode = ", resp.StatusCode)
		info := fmt.Sprintf("Get datas Fail, StatusCode = %v", resp.StatusCode)
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, event.Sender, info)
		return
	}
	if resp.Header.Get("Content-Length") == "0" {
		info := fmt.Sprintf("Empty content, Content-Length = 0")
		statement = false
		//response to oracle
		config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, event.Sender, info)
		return
	}

	switch event.Format {
	case "csv":
		body, err := io.ReadAll(resp.Body)
		reader := strings.NewReader(string(body))
		err = IngestCSV(ctx, reader, dbC)
		err = db.ApplyChanges(ctx)
		if err != nil {
			log.Println("Unmarshal csv ERROR: ", err)
			info := fmt.Sprintf("Unmarshal csv ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, event.Sender, info)
			return
		} else {
			statement = true
		}
	case "ndjson":
		body, _ := io.ReadAll(resp.Body)
		jsonArrayStr := string(body)

		cleanedStr := strings.Trim(jsonArrayStr, "[]")
		objects := strings.Split(cleanedStr, `},{`)
		var resultBuilder bytes.Buffer
		for i, obj := range objects {
			var prefix, suffix string
			if i == 0 {
				prefix = ""
			} else {
				prefix = "{"
			}
			if i == len(objects)-1 {
				suffix = ""
			} else {
				suffix = "}"
			}
			// 将处理后的对象追加到结果字符串中，并在每个对象之间添加换行符
			resultBuilder.WriteString(prefix + obj + suffix + "\n")
		}
		resultStr := resultBuilder.String()

		reader := strings.NewReader(resultStr)
		// insert data
		err = dbC.IndexNDJSON(ctx, reader)
		if err != nil {
			log.Println("Import data ERROR: ", err)
			info := fmt.Sprintf("Import data ERROR: %v", err)
			statement = false
			//response to oracle
			config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, event.Sender, info)
			return
		} else {
			statement = true
		}
	}
	//response to oracle
	_, err = config.OracleContract.ImportFromUrlRsp(tps, event.ReqID, statement, event.Sender, "")
	if err != nil {
		log.Println("Req function get an error : ", err)
	} else {
		log.Println("[", event.ColName, "]", "Import data success")
	}
}
