package api

import (
	"Oracle.com/golangServer/config"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"
)

func GenTransactOpts(GasLimit uint64) *bind.TransactOpts {
	// Generate TransactOpts from private key
	auth, err := bind.NewKeyedTransactorWithChainID(config.PrivateKey, big.NewInt(config.ChainID))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Set gas prices and gas limits, these can be set more intelligently through client queries
	gasPrice, err := config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = GasLimit
	return auth
}
