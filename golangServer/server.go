// main.go
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/api"
	"Oracle.com/golangServer/config"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func init() {
	// Define command-line flags
	urlFlag := flag.String("url", "", "The URL of the Ethereum node")
	privateKeyFlag := flag.String("privateKey", "", "The private key of the Oracle owner")
	contractAddressFlag := flag.String("contractAddress", "", "The address of the Oracle contract")
	chainIDFlag := flag.Int("chainID", 0, "The chain ID of the Ethereum network")

	flag.Parse()

	// Set configuration values based on flags
	config.SetConfig(*urlFlag, *privateKeyFlag, *contractAddressFlag, int64(*chainIDFlag))

	// Connect to the Ethereum client
	var err error
	client, err := ethclient.Dial(config.URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	config.SetClient(client)

	// Parse the private key
	privateKey, err := crypto.HexToECDSA(config.OracleOwnerPrivateKey[2:]) // Remove prefix "0x"
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}
	config.SetPrivateKey(privateKey)

	// Convert string address to `common.Address` type
	contractAddr := common.HexToAddress(config.ContractAddress)
	// Create a new instance using the contract address and client
	oracleContract, err := Oracle.NewOracle(contractAddr, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Oracle contract: %v", err)
	}
	config.SetOracleContract(oracleContract)

	// Load saved databases if they exist
	dbs := make(map[string]*indexer.Database)
	file, err := os.Open(fmt.Sprint(config.SaveDataPath, "paths"))
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln("Error opening file: ", err)
		}
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			path := scanner.Text()

			parts := strings.Split(path, "/")
			lastPart := parts[len(parts)-1]
			fileNameParts := strings.Split(lastPart, ".")
			result := fileNameParts[0]

			db, err := indexer.ImportFromFile(path)
			if err != nil {
				log.Printf("An error occurred while scanning path: %v", err)
			}
			dbs[result] = db
		}
		// Check whether errors are encountered during the Scan process
		if scanner.Err() != nil {
			log.Fatalf("An error occurred while scanning the file: %v", err)
		}
		log.Println("Loading local db successfully. ")
	}
	config.SetDatabases(dbs)
}

func main() {

	ctx := context.Background()
	// Meta info listener
	go api.GetCollections(ctx)
	go api.GetIndexes(ctx)
	go api.GetRootCid(ctx)
	// Service listener
	go api.CreatEventListener(ctx)
	go api.PutEventListener(ctx)
	go api.IndexEventListener(ctx)
	go api.GetEventListener(ctx)
	go api.SearchEventListener(ctx)
	go api.ImportEventListener(ctx)
	go api.GetRootCidEventListener(ctx)

	// Set up a channel for receiving signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// Wait for signal
	sig := <-sigs
	log.Printf("Received signal %s, exiting...\n", sig)

	saveDB(ctx) // Save db

	os.Exit(0)
}

func saveDB(ctx context.Context) {
	path := fmt.Sprint(config.SaveDataPath, "paths")
	os.Remove(path)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()

	for cid, db := range config.Dbs {
		if err != nil {
			log.Printf("Unable to create or open file: %s\n", err)
			return
		}
		// Path: savePath/{RootCid}.car
		path := fmt.Sprint(config.SaveDataPath, cid, ".car")
		// Use fmt.Fprintln to write a string to a file
		if _, err := fmt.Fprintln(file, path); err != nil {
			log.Printf("Unable to write to file: %s\n", err)
			return
		}
		// Save db as db.car
		db.ExportToFile(ctx, path)

	}
	log.Println("All data saved successfully. ")
}
