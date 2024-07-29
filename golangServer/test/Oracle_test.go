package test

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	OracleTest "Oracle.com/golangServer/test/testContract"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
	"time"
)

const (
	URL                 = "ws://127.0.0.1:8545"
	UserPrimaryKey      = "0x16f99b7743f83f132f858ed0aa62cd05e3f9afa3cd5838c7cb3138be6e2756cd"
	OracleAddress       = config.ContractAddress
	TestContractAddress = "0x60C3037baD781109830AC6FB48110A5db5BDC4dC"
	ChainID             = 1337
)

var (
	Client *ethclient.Client

	PrivateKey *ecdsa.PrivateKey

	test_contract *OracleTest.OracleTest

	GasLimit uint64 = 300000
	Value           = big.NewInt(1000000)

	dbName    = "MyDb"
	colName   = "users"
	callBack  = "CBFunc(bytes)"
	urlcsv    = "http://127.0.0.1:8080/data.csv"
	urlndjson = "http://127.0.0.1:8080/ndjson"
)

func TestOracle(t *testing.T) {
	InitClient()
	tps := GenTransactOpts(GasLimit, big.NewInt(0))
	var err error

	_, err = test_contract.SetOracle(tps, common.HexToAddress(OracleAddress))

	////create memory db
	//tps = GenTransactOpts(GasLimit, Value)
	//_, err = test_contract.Create(tps, dbName, colName, "name", Value)
	//time.Sleep(time.Second)
	//
	////create index in age field
	//tps = GenTransactOpts(GasLimit, Value)
	//_, err = test_contract.Index(tps, dbName, colName, "age", Value)
	//
	//time.Sleep(time.Second)
	//
	////put data
	//tps = GenTransactOpts(GasLimit, Value)
	//data := []byte(`{"name":"Alice", "age": 18}
	//				{"name":"Bob", "age": 19}
	//				{"name":"Albert", "age": 20}
	//				{"name":"Clearance and Steve", "age":18}`)
	//_, err = test_contract.Put(tps, dbName, colName, data, Value)
	//time.Sleep(time.Second)
	//
	////get data
	//tps = GenTransactOpts(GasLimit, Value)
	//recordID := []byte{129, 99, 66, 111, 98}
	//_, err = test_contract.Get(tps, dbName, colName, recordID, callBack, Value)
	//
	//time.Sleep(time.Second)

	//search data
	tps = GenTransactOpts(GasLimit, Value)
	type eq struct {
		IndexName string
		IndexVal  any
	}
	type cmp struct {
		//	GreaterThan Op = "GreaterThan"
		//	LessThan    Op = "LessThan"
		Op        string
		IndexName string
		IndexVal  any
	}
	type Queryer struct {
		Equal   eq
		Compare cmp
		Sort    string
		Limit   int
		Skip    int
	}
	quserys := []Queryer{
		{
			Equal: eq{IndexName: "name", IndexVal: "Bob"},
		},
		{
			Compare: cmp{
				Op:        "GreaterThan",
				IndexName: "age",
				IndexVal:  "18",
			},
		},
	}
	mq, _ := json.Marshal(quserys)
	_, err = test_contract.Search(tps, dbName, colName, mq, callBack, Value)
	time.Sleep(time.Second)

	//allowWrite
	tps = GenTransactOpts(GasLimit, Value)
	ADD := "0x00C696904c0CCE30D19704A762035Eb67eC3580C"
	_, err = test_contract.AllowWrite(tps, common.HexToAddress(ADD), Value)
	time.Sleep(time.Second)

	// get collections
	tps = GenTransactOpts(GasLimit, Value)
	_, err = test_contract.GetCol(tps, dbName, callBack, Value)
	time.Sleep(time.Second)

	//get indexes
	tps = GenTransactOpts(GasLimit, Value)
	_, err = test_contract.GetIndex(tps, dbName, colName, callBack, Value)
	time.Sleep(time.Second)

	//get rootCid
	tps = GenTransactOpts(GasLimit, Value)
	_, err = test_contract.GetRootCid(tps, dbName, callBack, Value)
	time.Sleep(time.Second)

	//create memory db
	dbName = "demoCsv"
	colName = "demoCol"
	tps = GenTransactOpts(GasLimit, Value)
	_, err = test_contract.Create(tps, dbName, colName, "Series_Title", Value)
	time.Sleep(time.Second)
	//put csv data by url
	tps = GenTransactOpts(GasLimit, Value)
	_, err = test_contract.ImportFromUrl(tps, dbName, colName, urlcsv, "csv", Value)
	time.Sleep(time.Second)

	//create memory db
	dbName = "demoNdjson"
	colName = "demoCol"
	tps = GenTransactOpts(GasLimit, Value)
	_, err = test_contract.Create(tps, dbName, colName, "name", Value)
	time.Sleep(time.Second)
	//put ndjson by url
	tps = GenTransactOpts(GasLimit, Value)
	_, err = test_contract.ImportFromUrl(tps, dbName, colName, urlndjson, "ndjson", Value)
	time.Sleep(time.Second)

	fmt.Println(err)
}

func InitClient() {
	var err error
	Client, err = ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	PrivateKey, err = crypto.HexToECDSA(UserPrimaryKey[2:]) // Remove prefix "0x"
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	// Convert string address to `common.Address` type
	contractAddr := common.HexToAddress(TestContractAddress)
	// Create a new instance using the contract address and client
	test_contract, err = OracleTest.NewOracleTest(contractAddr, Client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Oracle contract: %v", err)
	}
	go RspListener()
	go CatchListener()
}

func RspListener() {
	cli, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractAddr := common.HexToAddress(OracleAddress)
	OOOC, err := Oracle.NewOracle(contractAddr, cli)
	// Create channels for logs
	Logs := make(chan *Oracle.OracleReqState)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	eventSub, err := OOOC.WatchReqState(opts, Logs)
	if err != nil {
		log.Fatal("Failed to subscribe events:", err)
	}
	// start Listening...
	log.Println("RspListener is Listening ...")
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error]:", err)
		case event := <-Logs:
			log.Printf("[Rsp: %v]  state: %v  info: %v", event.ReqID, event.Statement, event.Info)
		}
	}
}

func CatchListener() {
	// Get channels for logs
	Logs := make(chan *OracleTest.OracleTestCatchData)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	eventSub, err := test_contract.WatchCatchData(opts, Logs)
	if err != nil {
		log.Fatal("Failed to subscribe to Get events:", err)
	}
	// start Listening...
	log.Println("CatchDataListener Listening ...")
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event CREAT]:", err)
		case event := <-Logs:
			log.Println("Received data: ")
			data := make([][]byte, 1)
			json.Unmarshal(event.Data, &data)
			var i any
			for k, v := range data {
				json.Unmarshal(v, &i)
				log.Printf("data%v: %v", k, i)
			}
		}
	}
}

func GenTransactOpts(GasLimit uint64, Value *big.Int) *bind.TransactOpts {
	// Generate TransactOpts from private key
	auth, err := bind.NewKeyedTransactorWithChainID(PrivateKey, big.NewInt(ChainID))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Set gas prices and gas limits, these can be set more intelligently through client queries
	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = GasLimit
	auth.Value = Value
	return auth
}
