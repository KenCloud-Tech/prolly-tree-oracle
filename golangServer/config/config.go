package config

import (
	"Oracle.com/golangServer/Oracle"
	"crypto/ecdsa"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	URL                   = "ws://127.0.0.1:8545"
	OracleOwnerPrivateKey = "0xe0c960bf16551bd3d69c54a8cfeaab6d7fe4d4ca8e404449b704945e27a06687"
	ContractAddress       = "0x0E51070026990dedF1f63be6214f43ea7dd408c0"
	ChainID               = 1337
	SaveDataPath          = "golangServer/savedDatas/"
)

var (
	Client *ethclient.Client

	PrivateKey *ecdsa.PrivateKey

	OracleContract *Oracle.Oracle

	Dbs map[string]*indexer.Database //mapping dbName address to db

	GasLimit uint64 = 300000
)
