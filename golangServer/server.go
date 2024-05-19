package main

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/api"
	"Oracle.com/golangServer/config"
	"Oracle.com/golangServer/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
	"os/signal"
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
	config.Dbs = make(map[string]*model.OracleModel)
}

func main() {
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
	// Cleaning work can be added here, such as closing open files, database connections, etc.
	os.Exit(0)
}
