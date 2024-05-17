package config

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/model"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	URL                   = "ws://127.0.0.1:8545"
	OracleOwnerPrivateKey = "0xa1c02aed7d70e580c62576f066ce6708aa01546a3490339aeb7193e5cede8b6d"
	ContractAddress       = "0xF7Cd6920dc8B1380c2c0751f78a292f160C39FAf"
	ChainID               = 1337
)

var (
	Client *ethclient.Client

	PrivateKey *ecdsa.PrivateKey

	OracleContract *Oracle.Oracle

	Dbs map[string]*model.OracleModel

	GasLimit uint64 = 300000
)
