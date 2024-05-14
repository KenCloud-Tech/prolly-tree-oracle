package main

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/api"
	"Oracle.com/golangServer/config"
	"Oracle.com/golangServer/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
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
	api.CreatEventListener()

	for true {
	}
}
