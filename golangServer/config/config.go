package config

import (
	"Oracle.com/golangServer/Oracle"
	"crypto/ecdsa"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	URL                   = "ws://127.0.0.1:8545"
	OracleOwnerPrivateKey = "0x0ff4f26be978b617efce2f61859d5298bb676c423ebbab8f281f905333e72928"
	ContractAddress       = "0x7F221E7aa73Bc3BEd7f13ac3F6c74B8412ffe8aC"
	ChainID               = 1337
	SaveDataPath          = "golangServer/savedDatas/"

	MaxDataSizeByUrl = 1024 * 1024 * 128 //128Mb
)

var (
	Client *ethclient.Client

	PrivateKey *ecdsa.PrivateKey

	OracleContract *Oracle.Oracle

	Dbs map[string]*indexer.Database //mapping dbName address to db

	GasLimit uint64 = 300000
)
