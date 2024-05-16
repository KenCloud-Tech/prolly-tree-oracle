package config

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/model"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	URL                   = "ws://127.0.0.1:8545"
	OracleOwnerPrivateKey = "0xfd6bb83200bb7dafb2e247fcfcd88c296d6d65f6542c5ad3cb66856ddac160f2"
	ContractAddress       = "0x825E9Ec368aE314949FD60D098e3C616cCBcD0BF"
	ChainID               = 1337
)

var (
	Client *ethclient.Client

	PrivateKey *ecdsa.PrivateKey

	OracleContract *Oracle.Oracle

	Dbs map[string]*model.OracleModel

	GasLimit uint64 = 300000
)
