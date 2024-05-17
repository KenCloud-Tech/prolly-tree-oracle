package test

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	OracleTest "Oracle.com/golangServer/test/testContract"
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

const (
	URL                 = "ws://127.0.0.1:8545"
	UserPrimaryKey      = "0x9945e953c3b37004a00d238d4b20a561e263cc578fb14fde6872cf76222ff702"
	OracleAddress       = config.ContractAddress
	TestContractAddress = "0x5c2a450cA28C5af7E95682ca262D42D0e8103a1e"
	ChainID             = 1337
)

var (
	Client *ethclient.Client

	PrivateKey *ecdsa.PrivateKey

	toc *OracleTest.OracleTest

	GasLimit uint64 = 300000
)

func TestOracle(t *testing.T) {
	InitClient()
	tps := GenTransactOpts(GasLimit)
	var err error

	_, err = toc.SetOracle(tps, common.HexToAddress(OracleAddress))
	if err != nil {
		log.Println("Failed to send transaction: ", err)
	}

	dbName := "users"
	// create memory db
	_, err = toc.Create(tps, dbName, "name")
	if err != nil {
		log.Println("Failed to send transaction: ", err)
	}
	time.Sleep(time.Second * 2)

	//create index in age field
	_, err = toc.Index(tps, dbName, "age")
	if err != nil {
		log.Println("Failed to send transaction: ", err)
	}

	time.Sleep(time.Second * 2)

	// put data
	data := []byte(`{"name":"Alice", "age": 18}
					{"name":"Bob", "age": 19}
					{"name":"Albert", "age": 20}
					{"name":"Clearance and Steve", "age":18}`)
	_, err = toc.Put(tps, dbName, data)
	if err != nil {
		log.Println("Failed to send transaction: ", err)
	}

	time.Sleep(time.Second * 2)

	//get data
	recordID := []byte{129, 99, 66, 111, 98}
	_, err = toc.Get(tps, dbName, recordID, "CBFunc(bytes)")
	if err != nil {
		log.Println("Failed to send transaction: ", err)
	}

	time.Sleep(time.Second * 2)

	//search data
	val1 := OracleTest.OracleInterfaceSearchController{K: "name", Str: "Bob", DataType: "string"}
	_, err = toc.Search(tps, dbName, val1, "equal", "CBFunc(bytes)")
	if err != nil {
		log.Println("Failed to send transaction: ", err)
	}
	val2 := OracleTest.OracleInterfaceSearchController{K: "age", Integer: 20, DataType: "int"}
	_, err = toc.Search(tps, dbName, val2, "equal", "CBFunc(bytes)")
	if err != nil {
		log.Println("Failed to send transaction: ", err)
	}

	time.Sleep(time.Second * 2)

	//allowWrite
	ADD := "0x00C696904c0CCE30D19704A762035Eb67eC3580C"
	_, err = toc.AllowWrite(tps, dbName, common.HexToAddress(ADD))
	if err != nil {
		log.Println("Failed to send transaction: ", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	log.Printf("Received signal %s, exiting...\n", sig)
	os.Exit(0)

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
	toc, err = OracleTest.NewOracleTest(contractAddr, Client)
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
		log.Fatal("Failed to subscribe to Create events:", err)
	}
	// start Listening...
	log.Println("RspListener is Listening ...")
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event CREAT]:", err)
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
	eventSub, err := toc.WatchCatchData(opts, Logs)
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
			log.Println("Received data: ", string(event.Data))
		}
	}
}

func GenTransactOpts(GasLimit uint64) *bind.TransactOpts {
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
	return auth
}
