//package config
//
//import (
//	"Oracle.com/golangServer/Oracle"
//	"crypto/ecdsa"
//	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
//	"github.com/ethereum/go-ethereum/ethclient"
//)
//
//const (
//	URL                   = "wss://eth-sepolia.g.alchemy.com/v2/s6fIWpm5hg-zvYZrcFa2qY7MHiOLeHR0"
//	OracleOwnerPrivateKey = "0xab4a55cab59349201260aec4ae102603cce8a279ead9f83f56ad1f4407e72fa5"
//	ContractAddress       = "0xDCA07264B37e60455F5cA4e572cFCa80a26D691e"
//	ChainID               = 11155111
//	SaveDataPath          = "golangServer/savedDatas/"
//)
//
//var (
//	Client *ethclient.Client
//
//	PrivateKey *ecdsa.PrivateKey
//
//	OracleContract *Oracle.Oracle
//
//	Dbs map[string]*indexer.Database //mapping dbName address to db
//
//	GasLimit uint64 = 300000
//)

// config.go
package config

import (
	"Oracle.com/golangServer/Oracle"
	"crypto/ecdsa"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	SaveDataPath = "golangServer/savedDatas/"
)

var (
	URL                   string
	OracleOwnerPrivateKey string
	ContractAddress       string
	ChainID               int64
	Client                *ethclient.Client
	PrivateKey            *ecdsa.PrivateKey
	OracleContract        *Oracle.Oracle
	Dbs                   map[string]*indexer.Database // mapping dbName address to db
	GasLimit              uint64                       = 300000
)

// SetConfig sets the configuration variables.
func SetConfig(url, privateKey, contractAddress string, chainID int64) {
	URL = url
	OracleOwnerPrivateKey = privateKey
	ContractAddress = contractAddress
	ChainID = chainID
}

// SetClient sets the Ethereum client.
func SetClient(client *ethclient.Client) {
	Client = client
}

// SetPrivateKey sets the private key.
func SetPrivateKey(privateKey *ecdsa.PrivateKey) {
	PrivateKey = privateKey
}

// SetOracleContract sets the Oracle contract.
func SetOracleContract(contract *Oracle.Oracle) {
	OracleContract = contract
}

// SetDatabases sets the database map.
func SetDatabases(dbs map[string]*indexer.Database) {
	Dbs = dbs
}
