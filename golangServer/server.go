package main

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/api"
	"Oracle.com/golangServer/config"
	"bufio"
	"context"
	"fmt"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func init() {
	var err error
	config.Client, err = ethclient.Dial(config.URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	config.PrivateKey, err = crypto.HexToECDSA(config.OracleOwnerPrivateKey[2:]) // Remove prefix "0x"
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	// Convert string address to `common.Address` type
	contractAddr := common.HexToAddress(config.ContractAddress)
	// Create a new instance using the contract address and client
	config.OracleContract, err = Oracle.NewOracle(contractAddr, config.Client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Oracle contract: %v", err)
	}

	config.Dbs = make(map[string]*indexer.Database)
	// load saved dbs if exist
	file, err := os.Open(fmt.Sprint(config.SaveDataPath, "paths"))
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln("Error opening file: ", err)
		}
	} else {
		// import db from file
		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			path := scanner.Text()

			parts := strings.Split(path, "/")
			lastPart := parts[len(parts)-1]
			fileNameParts := strings.Split(lastPart, ".")
			result := fileNameParts[0]

			config.Dbs[result], err = indexer.ImportFromFile(path)
			if err != nil {
				log.Printf("An error occurred while scanning path: %v", err)
			}
		}
		// Check whether errors are encountered during the Scan process
		if scanner.Err() != nil {
			log.Fatalf("An error occurred while scanning the file: %v", err)
		}
		log.Println("Loading local db successfully. ")
	}
}

func main() {
	// M eta info listener
	go api.GetCollections()
	go api.GetIndexes()
	// Service listener
	go api.CreatEventListener()
	go api.PutEventListener()
	go api.IndexEventListener()
	go api.GetEventListener()
	go api.SearchEventListener()

	// Set up a channel for receiving signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// wait for signal
	sig := <-sigs
	log.Printf("Received signal %s, exiting...\n", sig)

	saveDB() //save db

	os.Exit(0)
}

func saveDB() {
	path := fmt.Sprint(config.SaveDataPath, "paths")
	os.Remove(path)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()

	for cid, db := range config.Dbs {
		if err != nil {
			log.Printf("Unable to create or open file: %s\n", err)
			return
		}
		// path : savePath/{RootCid}.car
		path := fmt.Sprint(config.SaveDataPath, cid, ".car")
		// Use fmt.Fprintln to write a string to a file
		if _, err := fmt.Fprintln(file, path); err != nil {
			log.Printf("Unable to write to file: %s\n", err)
			return
		}
		// Save db as db.car
		db.ExportToFile(context.Background(), path)

	}
	log.Println("All data saved successfully. ")
}
