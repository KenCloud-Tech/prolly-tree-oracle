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
	OracleAddress       = config.ContractAddress
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
	_, err = toc.Get(tps, dbName, recordID, "catchData(bytes)")
	if err != nil {
		log.Println("Failed to send transaction: ", err)
	}

	time.Sleep(time.Second * 2)

	//search data
	val1 := OracleTest.OracleInterfaceSearchController{K: "name", Str: "Bob", ComDataType: "string"}
	_, err = toc.Search(tps, dbName, val1, "equal", "catchData(bytes)")
	if err != nil {
		log.Println("Failed to send transaction: ", err)
	}
	val2 := OracleTest.OracleInterfaceSearchController{K: "age", Integer: 20, ComDataType: "int"}
	_, err = toc.Search(tps, dbName, val2, "equal", "catchData(bytes)")
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

	// 设置一个用于接收信号的通道
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// 等待信号
	sig := <-sigs
	log.Printf("Received signal %s, exiting...\n", sig)
	// 这里可以添加清理工作，比如关闭打开的文件、数据库连接等
	// 然后退出程序
	os.Exit(0)

}

func InitClient() {
	var err error
	Client, err = ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	PrivateKey, err = crypto.HexToECDSA(UserPrimaryKey[2:]) // 移除前缀"0x"
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	// 将字符串地址转换为`common.Address`类型
	contractAddr := common.HexToAddress(TestContractAddress)
	// 使用合约地址和client创建一个新的实例
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

	// 将字符串地址转换为`common.Address`类型
	contractAddr := common.HexToAddress(OracleAddress)
	// 使用合约地址和client创建一个新的实例
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
	Logs := make(chan *OracleTest.OracleTestSendData)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	eventSub, err := toc.WatchSendData(opts, Logs)
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
	// 从私钥生成TransactOpts
	auth, err := bind.NewKeyedTransactorWithChainID(PrivateKey, big.NewInt(ChainID))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// 设置Gas价格和Gas限额，这些可以通过客户端查询来更智能地设置
	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = GasLimit
	return auth
}
