package api

import (
	"Oracle.com/golangServer/config"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"
)

func GenTransactOpts(GasLimit uint64) *bind.TransactOpts {
	// 从私钥生成TransactOpts
	auth, err := bind.NewKeyedTransactorWithChainID(config.PrivateKey, big.NewInt(config.ChainID))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// 设置Gas价格和Gas限额，这些可以通过客户端查询来更智能地设置
	gasPrice, err := config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = GasLimit
	return auth
}
