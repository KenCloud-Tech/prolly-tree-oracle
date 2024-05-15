package config

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/model"
)

const (
	URL             = "ws://127.0.0.1:8545"
	OracleOwner     = "0xCb5f3F9E5b36E6e4070aF11FF9E2b01D74D079E3"
	ContractAddress = "0x292320687B85839EFc977003Ac544AfA2F485757"
)

var OracleContract *Oracle.Oracle

var Dbs map[string]*model.OracleModel
