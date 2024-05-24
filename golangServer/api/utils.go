package api

import (
	"Oracle.com/golangServer/config"
	"bytes"
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
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

func nodeTobyte(node ipld.Node) []byte {
	var buf bytes.Buffer
	if err := dagjson.Encode(node, &buf); err != nil {
		log.Fatalf("Trans node to json fail: %v", err)
	}

	var formatted bytes.Buffer
	if err := json.Indent(&formatted, buf.Bytes(), "", "  "); err != nil {
		log.Fatalf("Format json fail: %v", err)
	}
	return formatted.Bytes()
}
