package main

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/api"
	"Oracle.com/golangServer/config"
	"Oracle.com/golangServer/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	client, err := ethclient.Dial(config.URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 将字符串地址转换为`common.Address`类型
	contractAddr := common.HexToAddress(config.ContractAddress)
	// 使用合约地址和client创建一个新的实例
	config.OracleContract, err = Oracle.NewOracle(contractAddr, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Oracle contract: %v", err)
	}
	config.Dbs = make(map[string]*model.OracleModel)
}

func main() {
	go api.CreatEventListener()
	go api.PutEventListener()
	go api.GetEventListener()
	go api.IndexEventListener()
	go api.SearchEventListener()

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
