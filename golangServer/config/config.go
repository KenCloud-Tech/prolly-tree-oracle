package config

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/model"
)

const (
	URL             = "ws://127.0.0.1:8545"
	OracleOwner     = "0xAB40890Ff19775832AC05cb184533Bc5C7Cdc32E"
	ContractAddress = "0x569229C06ED415E32523FCa73744d43A689BC66E"
)

var OracleContract *Oracle.Oracle

var Dbs map[string]*model.OracleModel
