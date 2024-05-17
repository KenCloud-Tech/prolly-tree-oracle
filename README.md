
# Prolly Tree Oracle

## What
* EVM oracle
* Based on [ipld prolly tree indexer](https://github.com/RangerMauve/ipld-prolly-indexer)
* Enable contracts to create/read/update/delete to prolly trees.
* Index over collections to enable faster queries
* charge gas or eth for transactions (e.g. if the field you query isn't indexed, charge more)
* maybe gas costs per lPLD block loaded? That way people deploying oracles can set theiown costs.
* allow querying by existing cid
* advertise prolly tree root on lPFS network to query offchain
* oracle publishes tree under ipns
* can we get read/write access gating based on wallets or multisigs?

## API:

* create(owner wallet, prolly config params)=>id

* allowWrite(wallet)

* put(id,collection,JSON document, primaryKey?) => primaryKey

* index(id, collection, fields)

* get(id, collection, primaryKey)->JSON

* search(id, collection, query {equals, sort, limit, skip})

## Architecture
![](/sources/Architecture.png)


## API Implementation:
* **function Create(string dbName, string primaryKey)**<br>
To initialize a new database, you need to provide the name and primary key of the database. 
<br>There can be multiple primary keys, separated by ",".<br>
**NOTE:** <br>
   1.  The owner of db is obtained by "msg.sender" when calling the Create method and does not need to be used as a parameter.<br>
   2.  The interface for creating a new db by polly-tree-indexer does not require config, so the parameters of the API only require "dbName" and "PrimaryKey"<br>


* **function AllowWrite(string dbName, address to)**<br>
  Grant user "to" read and write permissions on database "dbName".<br>
**NOTE:**<br>
  1. When authorizing, in addition to providing the authorized user, you should also provide the authorized db<br>


* **Put(string dbName, bytes data)**<br>
  Save json data "data" to database "dbName".<br>
  **NOTE:**<br>
  1. When putting data, the collection can be obtained through "dbName" and does not require "PrimaryKey"
  

* **Get(string dbName, bytes recordID, string callBack)**<br>
  To obtain data from database "dbName" based on "recordID", you need to provide a callback function to receive data.<br>
  For example, the callback function I defined in the consumer contract is:`function CallBackFunc(bytes calldata data)`.<br>
  So, the callBack parameter of string type is `"CallBackFunc(bytes)"`.<br>
  **NOTE:**<br>
  1. The get method of polly-tree-indexer obtains data through "recordID"
  2. The collection can be obtained through "dbName"
  3. Need to provide a callback function to receive data


* **Index(string dbName, string index)**<br>
  Add index "index" to database "dName".<br>
  **NOTE:**<br>
  1. The collection can be obtained through "dbName"


* **Search(string dbName, SearchController Val, string Method, string callBack)**<br>
  Obtain data through the "method" method in the database "dName", methods:`{equal, compare, sort, limit, skip}`<br>
  **NOTE:**<br>
  1. The collection can be obtained through "dbName"<br>
  2. Need to provide a callback function to receive data
  3. Values of multiple data types may be used when searching, so "SearchController" is added and "mehod" is used to indicate the search method.
  
  You need to provide a "SearchController", defined as follows:<br>
````
struct SearchController{
        string k; // The key for search
        
        // The value used for search,
        // Select the appropriate member according to the type of value
        string str;
        int64 integer;
        uint64 uinteger;
        bytes bytess;
        bool boolean;

        string DataType; // The type of value, assigned as "string", "int", "uint", "bytes" or "bool",
        string comOp; // When the search method is compare, you need to provide compare optiong: "GreaterThan" or "LessThan"
    }
````
Similarly, a callback function needs to be provided to receive data<br>


# Usage
### [Click here to learn how to start the service](https://www.youtube.com/watch?v=_T4nLVJaQ5k)
## Consumer Contract
```
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Oracle} from "./oracle.sol";

contract OracleTest{
    Oracle public or;
    function setOracle(address add) public{
        or = Oracle(add);
    }
    //allow a address write
    function AllowWrite(string calldata dbName, address to) external{
        or.AllowWrite(dbName,to);
    }

    //creat a new db
    function Create(string calldata dbName, string calldata primaryKey) external{
        or.Create(dbName,primaryKey);
    }

    // put data to db
    function Put(string calldata dbName, bytes calldata data) external{
        or.Put(dbName,data);
    }

    //get data from db
    function Get(string calldata dbName, bytes calldata recordID, string calldata callBack) external{
        or.Get(dbName,recordID,callBack);
    }

    //creat index
    function Index(string calldata dbName, string calldata Key) external{
        or.Index(dbName,Key);
    }
    
    //query by {equals, compare, sort, limit, skip}
    function Search(string calldata dbName, Oracle.SearchController calldata Val, string calldata Method, string calldata callBack) external{
        or.Search(dbName,Val,Method,callBack);
    }

    // catch data and emit event
    event CatchData(bytes data);
    function CBFunc(bytes calldata data) external{
        //You can process the obtained data here.
        emit CatchData(data);
    }

}
```
Compile the **Consumer Contract** to obtain the ABI file, and use the ***[abigen](https://geth.ethereum.org/docs/tools/abigen)*** to generate **Golang package** corresponding to the contract:
`abigen --abi={ Path to abi.json } --pkg={ Package name } --out={ The path to the generated .go file }`
## Golang Client
Use a Golang test program to call the request function in the Consumer Contract for testing the oracle.
```
package test

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"Oracle.com/golangServer/test/testContract"
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
	UserPrimaryKey      = "0xdadd9db9c1c0d5098de6f6138051ecf3a9ff17ac33082c96f2757fbc34529b9a"
	OracleAddress       = "0x825E9Ec368aE314949FD60D098e3C616cCBcD0BF"
	TestContractAddress = "0xf6C221Caa234D3b1D7AB0e2Aa228bE532ce7f710"
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
```
