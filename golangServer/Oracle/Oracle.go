// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Oracle

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"statement\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"ReqState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"colName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"primaryKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"colName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recordID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"get\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getCol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"colName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getIndex\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getRootCid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"colName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"Key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"index\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"colName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"put\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"colName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"querys\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"search\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AllowWrite\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"statement\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"colName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"CreatRsp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"colName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"primaryKey\",\"type\":\"string\"}],\"name\":\"Create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"colName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"recordID\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"}],\"name\":\"Get\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"}],\"name\":\"GetCol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"statement\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"GetColRsp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"colName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"}],\"name\":\"GetIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"statement\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"GetIndexRsp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ReqID\",\"type\":\"uint256\"}],\"name\":\"GetReqState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"}],\"name\":\"GetRootCid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"statement\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"GetRootCidRsp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"statement\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"GetRsp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"colName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"idx\",\"type\":\"string\"}],\"name\":\"Index\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"statement\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"IndexRsp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"colName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Put\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"statement\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"PutRsp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"colName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"querys\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"}],\"name\":\"Search\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"statement\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"SearchRsp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// GetReqState is a free data retrieval call binding the contract method 0x8812330b.
//
// Solidity: function GetReqState(uint256 ReqID) view returns(bool)
func (_Oracle *OracleCaller) GetReqState(opts *bind.CallOpts, ReqID *big.Int) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "GetReqState", ReqID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetReqState is a free data retrieval call binding the contract method 0x8812330b.
//
// Solidity: function GetReqState(uint256 ReqID) view returns(bool)
func (_Oracle *OracleSession) GetReqState(ReqID *big.Int) (bool, error) {
	return _Oracle.Contract.GetReqState(&_Oracle.CallOpts, ReqID)
}

// GetReqState is a free data retrieval call binding the contract method 0x8812330b.
//
// Solidity: function GetReqState(uint256 ReqID) view returns(bool)
func (_Oracle *OracleCallerSession) GetReqState(ReqID *big.Int) (bool, error) {
	return _Oracle.Contract.GetReqState(&_Oracle.CallOpts, ReqID)
}

// AllowWrite is a paid mutator transaction binding the contract method 0xb89f561f.
//
// Solidity: function AllowWrite(address to) returns()
func (_Oracle *OracleTransactor) AllowWrite(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "AllowWrite", to)
}

// AllowWrite is a paid mutator transaction binding the contract method 0xb89f561f.
//
// Solidity: function AllowWrite(address to) returns()
func (_Oracle *OracleSession) AllowWrite(to common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AllowWrite(&_Oracle.TransactOpts, to)
}

// AllowWrite is a paid mutator transaction binding the contract method 0xb89f561f.
//
// Solidity: function AllowWrite(address to) returns()
func (_Oracle *OracleTransactorSession) AllowWrite(to common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AllowWrite(&_Oracle.TransactOpts, to)
}

// CreatRsp is a paid mutator transaction binding the contract method 0x8276f971.
//
// Solidity: function CreatRsp(uint256 reqID, bool statement, string dbName, string colName, address sender, string info) returns()
func (_Oracle *OracleTransactor) CreatRsp(opts *bind.TransactOpts, reqID *big.Int, statement bool, dbName string, colName string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "CreatRsp", reqID, statement, dbName, colName, sender, info)
}

// CreatRsp is a paid mutator transaction binding the contract method 0x8276f971.
//
// Solidity: function CreatRsp(uint256 reqID, bool statement, string dbName, string colName, address sender, string info) returns()
func (_Oracle *OracleSession) CreatRsp(reqID *big.Int, statement bool, dbName string, colName string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.CreatRsp(&_Oracle.TransactOpts, reqID, statement, dbName, colName, sender, info)
}

// CreatRsp is a paid mutator transaction binding the contract method 0x8276f971.
//
// Solidity: function CreatRsp(uint256 reqID, bool statement, string dbName, string colName, address sender, string info) returns()
func (_Oracle *OracleTransactorSession) CreatRsp(reqID *big.Int, statement bool, dbName string, colName string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.CreatRsp(&_Oracle.TransactOpts, reqID, statement, dbName, colName, sender, info)
}

// Create is a paid mutator transaction binding the contract method 0x2e4bf45d.
//
// Solidity: function Create(string dbName, string colName, string primaryKey) returns()
func (_Oracle *OracleTransactor) Create(opts *bind.TransactOpts, dbName string, colName string, primaryKey string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "Create", dbName, colName, primaryKey)
}

// Create is a paid mutator transaction binding the contract method 0x2e4bf45d.
//
// Solidity: function Create(string dbName, string colName, string primaryKey) returns()
func (_Oracle *OracleSession) Create(dbName string, colName string, primaryKey string) (*types.Transaction, error) {
	return _Oracle.Contract.Create(&_Oracle.TransactOpts, dbName, colName, primaryKey)
}

// Create is a paid mutator transaction binding the contract method 0x2e4bf45d.
//
// Solidity: function Create(string dbName, string colName, string primaryKey) returns()
func (_Oracle *OracleTransactorSession) Create(dbName string, colName string, primaryKey string) (*types.Transaction, error) {
	return _Oracle.Contract.Create(&_Oracle.TransactOpts, dbName, colName, primaryKey)
}

// Get is a paid mutator transaction binding the contract method 0x0405473e.
//
// Solidity: function Get(string dbName, string colName, bytes recordID, string callBack) returns()
func (_Oracle *OracleTransactor) Get(opts *bind.TransactOpts, dbName string, colName string, recordID []byte, callBack string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "Get", dbName, colName, recordID, callBack)
}

// Get is a paid mutator transaction binding the contract method 0x0405473e.
//
// Solidity: function Get(string dbName, string colName, bytes recordID, string callBack) returns()
func (_Oracle *OracleSession) Get(dbName string, colName string, recordID []byte, callBack string) (*types.Transaction, error) {
	return _Oracle.Contract.Get(&_Oracle.TransactOpts, dbName, colName, recordID, callBack)
}

// Get is a paid mutator transaction binding the contract method 0x0405473e.
//
// Solidity: function Get(string dbName, string colName, bytes recordID, string callBack) returns()
func (_Oracle *OracleTransactorSession) Get(dbName string, colName string, recordID []byte, callBack string) (*types.Transaction, error) {
	return _Oracle.Contract.Get(&_Oracle.TransactOpts, dbName, colName, recordID, callBack)
}

// GetCol is a paid mutator transaction binding the contract method 0xe722d92e.
//
// Solidity: function GetCol(string dbName, string callBack) returns()
func (_Oracle *OracleTransactor) GetCol(opts *bind.TransactOpts, dbName string, callBack string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "GetCol", dbName, callBack)
}

// GetCol is a paid mutator transaction binding the contract method 0xe722d92e.
//
// Solidity: function GetCol(string dbName, string callBack) returns()
func (_Oracle *OracleSession) GetCol(dbName string, callBack string) (*types.Transaction, error) {
	return _Oracle.Contract.GetCol(&_Oracle.TransactOpts, dbName, callBack)
}

// GetCol is a paid mutator transaction binding the contract method 0xe722d92e.
//
// Solidity: function GetCol(string dbName, string callBack) returns()
func (_Oracle *OracleTransactorSession) GetCol(dbName string, callBack string) (*types.Transaction, error) {
	return _Oracle.Contract.GetCol(&_Oracle.TransactOpts, dbName, callBack)
}

// GetColRsp is a paid mutator transaction binding the contract method 0x604cd7d7.
//
// Solidity: function GetColRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleTransactor) GetColRsp(opts *bind.TransactOpts, reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "GetColRsp", reqID, statement, data, callBack, sender, info)
}

// GetColRsp is a paid mutator transaction binding the contract method 0x604cd7d7.
//
// Solidity: function GetColRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleSession) GetColRsp(reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.GetColRsp(&_Oracle.TransactOpts, reqID, statement, data, callBack, sender, info)
}

// GetColRsp is a paid mutator transaction binding the contract method 0x604cd7d7.
//
// Solidity: function GetColRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleTransactorSession) GetColRsp(reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.GetColRsp(&_Oracle.TransactOpts, reqID, statement, data, callBack, sender, info)
}

// GetIndex is a paid mutator transaction binding the contract method 0x0a362255.
//
// Solidity: function GetIndex(string dbName, string colName, string callBack) returns()
func (_Oracle *OracleTransactor) GetIndex(opts *bind.TransactOpts, dbName string, colName string, callBack string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "GetIndex", dbName, colName, callBack)
}

// GetIndex is a paid mutator transaction binding the contract method 0x0a362255.
//
// Solidity: function GetIndex(string dbName, string colName, string callBack) returns()
func (_Oracle *OracleSession) GetIndex(dbName string, colName string, callBack string) (*types.Transaction, error) {
	return _Oracle.Contract.GetIndex(&_Oracle.TransactOpts, dbName, colName, callBack)
}

// GetIndex is a paid mutator transaction binding the contract method 0x0a362255.
//
// Solidity: function GetIndex(string dbName, string colName, string callBack) returns()
func (_Oracle *OracleTransactorSession) GetIndex(dbName string, colName string, callBack string) (*types.Transaction, error) {
	return _Oracle.Contract.GetIndex(&_Oracle.TransactOpts, dbName, colName, callBack)
}

// GetIndexRsp is a paid mutator transaction binding the contract method 0xe97ec78b.
//
// Solidity: function GetIndexRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleTransactor) GetIndexRsp(opts *bind.TransactOpts, reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "GetIndexRsp", reqID, statement, data, callBack, sender, info)
}

// GetIndexRsp is a paid mutator transaction binding the contract method 0xe97ec78b.
//
// Solidity: function GetIndexRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleSession) GetIndexRsp(reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.GetIndexRsp(&_Oracle.TransactOpts, reqID, statement, data, callBack, sender, info)
}

// GetIndexRsp is a paid mutator transaction binding the contract method 0xe97ec78b.
//
// Solidity: function GetIndexRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleTransactorSession) GetIndexRsp(reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.GetIndexRsp(&_Oracle.TransactOpts, reqID, statement, data, callBack, sender, info)
}

// GetRootCid is a paid mutator transaction binding the contract method 0xf04c04df.
//
// Solidity: function GetRootCid(string dbName, string callBack) returns()
func (_Oracle *OracleTransactor) GetRootCid(opts *bind.TransactOpts, dbName string, callBack string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "GetRootCid", dbName, callBack)
}

// GetRootCid is a paid mutator transaction binding the contract method 0xf04c04df.
//
// Solidity: function GetRootCid(string dbName, string callBack) returns()
func (_Oracle *OracleSession) GetRootCid(dbName string, callBack string) (*types.Transaction, error) {
	return _Oracle.Contract.GetRootCid(&_Oracle.TransactOpts, dbName, callBack)
}

// GetRootCid is a paid mutator transaction binding the contract method 0xf04c04df.
//
// Solidity: function GetRootCid(string dbName, string callBack) returns()
func (_Oracle *OracleTransactorSession) GetRootCid(dbName string, callBack string) (*types.Transaction, error) {
	return _Oracle.Contract.GetRootCid(&_Oracle.TransactOpts, dbName, callBack)
}

// GetRootCidRsp is a paid mutator transaction binding the contract method 0x2f94a9a6.
//
// Solidity: function GetRootCidRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleTransactor) GetRootCidRsp(opts *bind.TransactOpts, reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "GetRootCidRsp", reqID, statement, data, callBack, sender, info)
}

// GetRootCidRsp is a paid mutator transaction binding the contract method 0x2f94a9a6.
//
// Solidity: function GetRootCidRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleSession) GetRootCidRsp(reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.GetRootCidRsp(&_Oracle.TransactOpts, reqID, statement, data, callBack, sender, info)
}

// GetRootCidRsp is a paid mutator transaction binding the contract method 0x2f94a9a6.
//
// Solidity: function GetRootCidRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleTransactorSession) GetRootCidRsp(reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.GetRootCidRsp(&_Oracle.TransactOpts, reqID, statement, data, callBack, sender, info)
}

// GetRsp is a paid mutator transaction binding the contract method 0x7f96591d.
//
// Solidity: function GetRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleTransactor) GetRsp(opts *bind.TransactOpts, reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "GetRsp", reqID, statement, data, callBack, sender, info)
}

// GetRsp is a paid mutator transaction binding the contract method 0x7f96591d.
//
// Solidity: function GetRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleSession) GetRsp(reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.GetRsp(&_Oracle.TransactOpts, reqID, statement, data, callBack, sender, info)
}

// GetRsp is a paid mutator transaction binding the contract method 0x7f96591d.
//
// Solidity: function GetRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleTransactorSession) GetRsp(reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.GetRsp(&_Oracle.TransactOpts, reqID, statement, data, callBack, sender, info)
}

// Index is a paid mutator transaction binding the contract method 0x38e4fc73.
//
// Solidity: function Index(string dbName, string colName, string idx) returns()
func (_Oracle *OracleTransactor) Index(opts *bind.TransactOpts, dbName string, colName string, idx string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "Index", dbName, colName, idx)
}

// Index is a paid mutator transaction binding the contract method 0x38e4fc73.
//
// Solidity: function Index(string dbName, string colName, string idx) returns()
func (_Oracle *OracleSession) Index(dbName string, colName string, idx string) (*types.Transaction, error) {
	return _Oracle.Contract.Index(&_Oracle.TransactOpts, dbName, colName, idx)
}

// Index is a paid mutator transaction binding the contract method 0x38e4fc73.
//
// Solidity: function Index(string dbName, string colName, string idx) returns()
func (_Oracle *OracleTransactorSession) Index(dbName string, colName string, idx string) (*types.Transaction, error) {
	return _Oracle.Contract.Index(&_Oracle.TransactOpts, dbName, colName, idx)
}

// IndexRsp is a paid mutator transaction binding the contract method 0xd82ae8b8.
//
// Solidity: function IndexRsp(uint256 reqID, bool statement, address sender, string info) returns()
func (_Oracle *OracleTransactor) IndexRsp(opts *bind.TransactOpts, reqID *big.Int, statement bool, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "IndexRsp", reqID, statement, sender, info)
}

// IndexRsp is a paid mutator transaction binding the contract method 0xd82ae8b8.
//
// Solidity: function IndexRsp(uint256 reqID, bool statement, address sender, string info) returns()
func (_Oracle *OracleSession) IndexRsp(reqID *big.Int, statement bool, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.IndexRsp(&_Oracle.TransactOpts, reqID, statement, sender, info)
}

// IndexRsp is a paid mutator transaction binding the contract method 0xd82ae8b8.
//
// Solidity: function IndexRsp(uint256 reqID, bool statement, address sender, string info) returns()
func (_Oracle *OracleTransactorSession) IndexRsp(reqID *big.Int, statement bool, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.IndexRsp(&_Oracle.TransactOpts, reqID, statement, sender, info)
}

// Put is a paid mutator transaction binding the contract method 0x717a9408.
//
// Solidity: function Put(string dbName, string colName, bytes data) returns()
func (_Oracle *OracleTransactor) Put(opts *bind.TransactOpts, dbName string, colName string, data []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "Put", dbName, colName, data)
}

// Put is a paid mutator transaction binding the contract method 0x717a9408.
//
// Solidity: function Put(string dbName, string colName, bytes data) returns()
func (_Oracle *OracleSession) Put(dbName string, colName string, data []byte) (*types.Transaction, error) {
	return _Oracle.Contract.Put(&_Oracle.TransactOpts, dbName, colName, data)
}

// Put is a paid mutator transaction binding the contract method 0x717a9408.
//
// Solidity: function Put(string dbName, string colName, bytes data) returns()
func (_Oracle *OracleTransactorSession) Put(dbName string, colName string, data []byte) (*types.Transaction, error) {
	return _Oracle.Contract.Put(&_Oracle.TransactOpts, dbName, colName, data)
}

// PutRsp is a paid mutator transaction binding the contract method 0xa244d08d.
//
// Solidity: function PutRsp(uint256 reqID, bool statement, address sender, string info) returns()
func (_Oracle *OracleTransactor) PutRsp(opts *bind.TransactOpts, reqID *big.Int, statement bool, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "PutRsp", reqID, statement, sender, info)
}

// PutRsp is a paid mutator transaction binding the contract method 0xa244d08d.
//
// Solidity: function PutRsp(uint256 reqID, bool statement, address sender, string info) returns()
func (_Oracle *OracleSession) PutRsp(reqID *big.Int, statement bool, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.PutRsp(&_Oracle.TransactOpts, reqID, statement, sender, info)
}

// PutRsp is a paid mutator transaction binding the contract method 0xa244d08d.
//
// Solidity: function PutRsp(uint256 reqID, bool statement, address sender, string info) returns()
func (_Oracle *OracleTransactorSession) PutRsp(reqID *big.Int, statement bool, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.PutRsp(&_Oracle.TransactOpts, reqID, statement, sender, info)
}

// Search is a paid mutator transaction binding the contract method 0x38ae4ea1.
//
// Solidity: function Search(string dbName, string colName, bytes querys, string callBack) returns()
func (_Oracle *OracleTransactor) Search(opts *bind.TransactOpts, dbName string, colName string, querys []byte, callBack string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "Search", dbName, colName, querys, callBack)
}

// Search is a paid mutator transaction binding the contract method 0x38ae4ea1.
//
// Solidity: function Search(string dbName, string colName, bytes querys, string callBack) returns()
func (_Oracle *OracleSession) Search(dbName string, colName string, querys []byte, callBack string) (*types.Transaction, error) {
	return _Oracle.Contract.Search(&_Oracle.TransactOpts, dbName, colName, querys, callBack)
}

// Search is a paid mutator transaction binding the contract method 0x38ae4ea1.
//
// Solidity: function Search(string dbName, string colName, bytes querys, string callBack) returns()
func (_Oracle *OracleTransactorSession) Search(dbName string, colName string, querys []byte, callBack string) (*types.Transaction, error) {
	return _Oracle.Contract.Search(&_Oracle.TransactOpts, dbName, colName, querys, callBack)
}

// SearchRsp is a paid mutator transaction binding the contract method 0xbcbdd19b.
//
// Solidity: function SearchRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleTransactor) SearchRsp(opts *bind.TransactOpts, reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "SearchRsp", reqID, statement, data, callBack, sender, info)
}

// SearchRsp is a paid mutator transaction binding the contract method 0xbcbdd19b.
//
// Solidity: function SearchRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleSession) SearchRsp(reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.SearchRsp(&_Oracle.TransactOpts, reqID, statement, data, callBack, sender, info)
}

// SearchRsp is a paid mutator transaction binding the contract method 0xbcbdd19b.
//
// Solidity: function SearchRsp(uint256 reqID, bool statement, bytes data, string callBack, address sender, string info) returns()
func (_Oracle *OracleTransactorSession) SearchRsp(reqID *big.Int, statement bool, data []byte, callBack string, sender common.Address, info string) (*types.Transaction, error) {
	return _Oracle.Contract.SearchRsp(&_Oracle.TransactOpts, reqID, statement, data, callBack, sender, info)
}

// OracleReqStateIterator is returned from FilterReqState and is used to iterate over the raw logs and unpacked data for ReqState events raised by the Oracle contract.
type OracleReqStateIterator struct {
	Event *OracleReqState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleReqStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleReqState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleReqState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleReqStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleReqStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleReqState represents a ReqState event raised by the Oracle contract.
type OracleReqState struct {
	ReqID     *big.Int
	User      common.Address
	Statement bool
	Info      string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReqState is a free log retrieval operation binding the contract event 0x23ea00673ffe456fb58cc2d191a5da9f49ac9f21503ddf1b6f10d20ef97bd566.
//
// Solidity: event ReqState(uint256 reqID, address user, bool statement, string info)
func (_Oracle *OracleFilterer) FilterReqState(opts *bind.FilterOpts) (*OracleReqStateIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "ReqState")
	if err != nil {
		return nil, err
	}
	return &OracleReqStateIterator{contract: _Oracle.contract, event: "ReqState", logs: logs, sub: sub}, nil
}

// WatchReqState is a free log subscription operation binding the contract event 0x23ea00673ffe456fb58cc2d191a5da9f49ac9f21503ddf1b6f10d20ef97bd566.
//
// Solidity: event ReqState(uint256 reqID, address user, bool statement, string info)
func (_Oracle *OracleFilterer) WatchReqState(opts *bind.WatchOpts, sink chan<- *OracleReqState) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "ReqState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleReqState)
				if err := _Oracle.contract.UnpackLog(event, "ReqState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReqState is a log parse operation binding the contract event 0x23ea00673ffe456fb58cc2d191a5da9f49ac9f21503ddf1b6f10d20ef97bd566.
//
// Solidity: event ReqState(uint256 reqID, address user, bool statement, string info)
func (_Oracle *OracleFilterer) ParseReqState(log types.Log) (*OracleReqState, error) {
	event := new(OracleReqState)
	if err := _Oracle.contract.UnpackLog(event, "ReqState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the Oracle contract.
type OracleCreateIterator struct {
	Event *OracleCreate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleCreate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleCreate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleCreate represents a Create event raised by the Oracle contract.
type OracleCreate struct {
	ReqID      *big.Int
	DbName     string
	ColName    string
	PrimaryKey string
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0xae273ad1c3c01ee1ad66d33c1c243a8e3a3b065b780eb2aceb3931c27732eb32.
//
// Solidity: event create(uint256 reqID, string dbName, string colName, string primaryKey, address owner)
func (_Oracle *OracleFilterer) FilterCreate(opts *bind.FilterOpts) (*OracleCreateIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "create")
	if err != nil {
		return nil, err
	}
	return &OracleCreateIterator{contract: _Oracle.contract, event: "create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0xae273ad1c3c01ee1ad66d33c1c243a8e3a3b065b780eb2aceb3931c27732eb32.
//
// Solidity: event create(uint256 reqID, string dbName, string colName, string primaryKey, address owner)
func (_Oracle *OracleFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *OracleCreate) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "create")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleCreate)
				if err := _Oracle.contract.UnpackLog(event, "create", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreate is a log parse operation binding the contract event 0xae273ad1c3c01ee1ad66d33c1c243a8e3a3b065b780eb2aceb3931c27732eb32.
//
// Solidity: event create(uint256 reqID, string dbName, string colName, string primaryKey, address owner)
func (_Oracle *OracleFilterer) ParseCreate(log types.Log) (*OracleCreate, error) {
	event := new(OracleCreate)
	if err := _Oracle.contract.UnpackLog(event, "create", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleGetIterator is returned from FilterGet and is used to iterate over the raw logs and unpacked data for Get events raised by the Oracle contract.
type OracleGetIterator struct {
	Event *OracleGet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleGetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleGet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleGet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleGetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleGetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleGet represents a Get event raised by the Oracle contract.
type OracleGet struct {
	ReqID    *big.Int
	DbName   string
	ColName  string
	RecordID []byte
	CallBack string
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGet is a free log retrieval operation binding the contract event 0xf95ce09bbae71e94492156b11ac31168e7b11ec7a264b71db25038ce94877d26.
//
// Solidity: event get(uint256 reqID, string dbName, string colName, bytes recordID, string callBack, address sender)
func (_Oracle *OracleFilterer) FilterGet(opts *bind.FilterOpts) (*OracleGetIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "get")
	if err != nil {
		return nil, err
	}
	return &OracleGetIterator{contract: _Oracle.contract, event: "get", logs: logs, sub: sub}, nil
}

// WatchGet is a free log subscription operation binding the contract event 0xf95ce09bbae71e94492156b11ac31168e7b11ec7a264b71db25038ce94877d26.
//
// Solidity: event get(uint256 reqID, string dbName, string colName, bytes recordID, string callBack, address sender)
func (_Oracle *OracleFilterer) WatchGet(opts *bind.WatchOpts, sink chan<- *OracleGet) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "get")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleGet)
				if err := _Oracle.contract.UnpackLog(event, "get", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGet is a log parse operation binding the contract event 0xf95ce09bbae71e94492156b11ac31168e7b11ec7a264b71db25038ce94877d26.
//
// Solidity: event get(uint256 reqID, string dbName, string colName, bytes recordID, string callBack, address sender)
func (_Oracle *OracleFilterer) ParseGet(log types.Log) (*OracleGet, error) {
	event := new(OracleGet)
	if err := _Oracle.contract.UnpackLog(event, "get", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleGetColIterator is returned from FilterGetCol and is used to iterate over the raw logs and unpacked data for GetCol events raised by the Oracle contract.
type OracleGetColIterator struct {
	Event *OracleGetCol // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleGetColIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleGetCol)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleGetCol)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleGetColIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleGetColIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleGetCol represents a GetCol event raised by the Oracle contract.
type OracleGetCol struct {
	ReqID    *big.Int
	DbName   string
	CallBack string
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGetCol is a free log retrieval operation binding the contract event 0xd583b378be23e9954c91371fcd0b84a56ccd2789fb45f90c191308682c92a5a9.
//
// Solidity: event getCol(uint256 reqID, string dbName, string callBack, address sender)
func (_Oracle *OracleFilterer) FilterGetCol(opts *bind.FilterOpts) (*OracleGetColIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "getCol")
	if err != nil {
		return nil, err
	}
	return &OracleGetColIterator{contract: _Oracle.contract, event: "getCol", logs: logs, sub: sub}, nil
}

// WatchGetCol is a free log subscription operation binding the contract event 0xd583b378be23e9954c91371fcd0b84a56ccd2789fb45f90c191308682c92a5a9.
//
// Solidity: event getCol(uint256 reqID, string dbName, string callBack, address sender)
func (_Oracle *OracleFilterer) WatchGetCol(opts *bind.WatchOpts, sink chan<- *OracleGetCol) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "getCol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleGetCol)
				if err := _Oracle.contract.UnpackLog(event, "getCol", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGetCol is a log parse operation binding the contract event 0xd583b378be23e9954c91371fcd0b84a56ccd2789fb45f90c191308682c92a5a9.
//
// Solidity: event getCol(uint256 reqID, string dbName, string callBack, address sender)
func (_Oracle *OracleFilterer) ParseGetCol(log types.Log) (*OracleGetCol, error) {
	event := new(OracleGetCol)
	if err := _Oracle.contract.UnpackLog(event, "getCol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleGetIndexIterator is returned from FilterGetIndex and is used to iterate over the raw logs and unpacked data for GetIndex events raised by the Oracle contract.
type OracleGetIndexIterator struct {
	Event *OracleGetIndex // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleGetIndexIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleGetIndex)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleGetIndex)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleGetIndexIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleGetIndexIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleGetIndex represents a GetIndex event raised by the Oracle contract.
type OracleGetIndex struct {
	ReqID    *big.Int
	DbName   string
	ColName  string
	CallBack string
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGetIndex is a free log retrieval operation binding the contract event 0x212fbb17e3c8775a57cc5fe569eeaecbbecb483a46a1d1cdf955d6e06f1c00fb.
//
// Solidity: event getIndex(uint256 reqID, string dbName, string colName, string callBack, address sender)
func (_Oracle *OracleFilterer) FilterGetIndex(opts *bind.FilterOpts) (*OracleGetIndexIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "getIndex")
	if err != nil {
		return nil, err
	}
	return &OracleGetIndexIterator{contract: _Oracle.contract, event: "getIndex", logs: logs, sub: sub}, nil
}

// WatchGetIndex is a free log subscription operation binding the contract event 0x212fbb17e3c8775a57cc5fe569eeaecbbecb483a46a1d1cdf955d6e06f1c00fb.
//
// Solidity: event getIndex(uint256 reqID, string dbName, string colName, string callBack, address sender)
func (_Oracle *OracleFilterer) WatchGetIndex(opts *bind.WatchOpts, sink chan<- *OracleGetIndex) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "getIndex")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleGetIndex)
				if err := _Oracle.contract.UnpackLog(event, "getIndex", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGetIndex is a log parse operation binding the contract event 0x212fbb17e3c8775a57cc5fe569eeaecbbecb483a46a1d1cdf955d6e06f1c00fb.
//
// Solidity: event getIndex(uint256 reqID, string dbName, string colName, string callBack, address sender)
func (_Oracle *OracleFilterer) ParseGetIndex(log types.Log) (*OracleGetIndex, error) {
	event := new(OracleGetIndex)
	if err := _Oracle.contract.UnpackLog(event, "getIndex", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleGetRootCidIterator is returned from FilterGetRootCid and is used to iterate over the raw logs and unpacked data for GetRootCid events raised by the Oracle contract.
type OracleGetRootCidIterator struct {
	Event *OracleGetRootCid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleGetRootCidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleGetRootCid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleGetRootCid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleGetRootCidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleGetRootCidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleGetRootCid represents a GetRootCid event raised by the Oracle contract.
type OracleGetRootCid struct {
	ReqID    *big.Int
	DbName   string
	CallBack string
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGetRootCid is a free log retrieval operation binding the contract event 0x70c0ae29feae948daed8cf8bf7eef0ba8b6a16919665052795aa432ecf40421a.
//
// Solidity: event getRootCid(uint256 reqID, string dbName, string callBack, address sender)
func (_Oracle *OracleFilterer) FilterGetRootCid(opts *bind.FilterOpts) (*OracleGetRootCidIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "getRootCid")
	if err != nil {
		return nil, err
	}
	return &OracleGetRootCidIterator{contract: _Oracle.contract, event: "getRootCid", logs: logs, sub: sub}, nil
}

// WatchGetRootCid is a free log subscription operation binding the contract event 0x70c0ae29feae948daed8cf8bf7eef0ba8b6a16919665052795aa432ecf40421a.
//
// Solidity: event getRootCid(uint256 reqID, string dbName, string callBack, address sender)
func (_Oracle *OracleFilterer) WatchGetRootCid(opts *bind.WatchOpts, sink chan<- *OracleGetRootCid) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "getRootCid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleGetRootCid)
				if err := _Oracle.contract.UnpackLog(event, "getRootCid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGetRootCid is a log parse operation binding the contract event 0x70c0ae29feae948daed8cf8bf7eef0ba8b6a16919665052795aa432ecf40421a.
//
// Solidity: event getRootCid(uint256 reqID, string dbName, string callBack, address sender)
func (_Oracle *OracleFilterer) ParseGetRootCid(log types.Log) (*OracleGetRootCid, error) {
	event := new(OracleGetRootCid)
	if err := _Oracle.contract.UnpackLog(event, "getRootCid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleIndexIterator is returned from FilterIndex and is used to iterate over the raw logs and unpacked data for Index events raised by the Oracle contract.
type OracleIndexIterator struct {
	Event *OracleIndex // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleIndexIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleIndex)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleIndex)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleIndexIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleIndexIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleIndex represents a Index event raised by the Oracle contract.
type OracleIndex struct {
	ReqID   *big.Int
	DbName  string
	ColName string
	Key     string
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIndex is a free log retrieval operation binding the contract event 0x5d687c727a6613fa6c66be3985b0d845711a1361f747c3509bf4c758f7187e14.
//
// Solidity: event index(uint256 reqID, string dbName, string colName, string Key, address sender)
func (_Oracle *OracleFilterer) FilterIndex(opts *bind.FilterOpts) (*OracleIndexIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "index")
	if err != nil {
		return nil, err
	}
	return &OracleIndexIterator{contract: _Oracle.contract, event: "index", logs: logs, sub: sub}, nil
}

// WatchIndex is a free log subscription operation binding the contract event 0x5d687c727a6613fa6c66be3985b0d845711a1361f747c3509bf4c758f7187e14.
//
// Solidity: event index(uint256 reqID, string dbName, string colName, string Key, address sender)
func (_Oracle *OracleFilterer) WatchIndex(opts *bind.WatchOpts, sink chan<- *OracleIndex) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "index")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleIndex)
				if err := _Oracle.contract.UnpackLog(event, "index", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIndex is a log parse operation binding the contract event 0x5d687c727a6613fa6c66be3985b0d845711a1361f747c3509bf4c758f7187e14.
//
// Solidity: event index(uint256 reqID, string dbName, string colName, string Key, address sender)
func (_Oracle *OracleFilterer) ParseIndex(log types.Log) (*OracleIndex, error) {
	event := new(OracleIndex)
	if err := _Oracle.contract.UnpackLog(event, "index", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OraclePutIterator is returned from FilterPut and is used to iterate over the raw logs and unpacked data for Put events raised by the Oracle contract.
type OraclePutIterator struct {
	Event *OraclePut // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OraclePutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OraclePut)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OraclePut)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OraclePutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OraclePutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OraclePut represents a Put event raised by the Oracle contract.
type OraclePut struct {
	ReqID   *big.Int
	DbName  string
	ColName string
	Data    []byte
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPut is a free log retrieval operation binding the contract event 0x946f3555fdbe19f5d6d0ed7b13cd52d57a314514c5654f993d96f9fe49fd6651.
//
// Solidity: event put(uint256 reqID, string dbName, string colName, bytes data, address sender)
func (_Oracle *OracleFilterer) FilterPut(opts *bind.FilterOpts) (*OraclePutIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "put")
	if err != nil {
		return nil, err
	}
	return &OraclePutIterator{contract: _Oracle.contract, event: "put", logs: logs, sub: sub}, nil
}

// WatchPut is a free log subscription operation binding the contract event 0x946f3555fdbe19f5d6d0ed7b13cd52d57a314514c5654f993d96f9fe49fd6651.
//
// Solidity: event put(uint256 reqID, string dbName, string colName, bytes data, address sender)
func (_Oracle *OracleFilterer) WatchPut(opts *bind.WatchOpts, sink chan<- *OraclePut) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "put")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OraclePut)
				if err := _Oracle.contract.UnpackLog(event, "put", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePut is a log parse operation binding the contract event 0x946f3555fdbe19f5d6d0ed7b13cd52d57a314514c5654f993d96f9fe49fd6651.
//
// Solidity: event put(uint256 reqID, string dbName, string colName, bytes data, address sender)
func (_Oracle *OracleFilterer) ParsePut(log types.Log) (*OraclePut, error) {
	event := new(OraclePut)
	if err := _Oracle.contract.UnpackLog(event, "put", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleSearchIterator is returned from FilterSearch and is used to iterate over the raw logs and unpacked data for Search events raised by the Oracle contract.
type OracleSearchIterator struct {
	Event *OracleSearch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleSearchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleSearch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleSearch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleSearchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleSearchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleSearch represents a Search event raised by the Oracle contract.
type OracleSearch struct {
	ReqID    *big.Int
	DbName   string
	ColName  string
	Querys   []byte
	CallBack string
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSearch is a free log retrieval operation binding the contract event 0xf9c679450fa2336656a505b21ba8b20cf414bd0e75f0f89cc57be8e08d7a6c5c.
//
// Solidity: event search(uint256 reqID, string dbName, string colName, bytes querys, string callBack, address sender)
func (_Oracle *OracleFilterer) FilterSearch(opts *bind.FilterOpts) (*OracleSearchIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "search")
	if err != nil {
		return nil, err
	}
	return &OracleSearchIterator{contract: _Oracle.contract, event: "search", logs: logs, sub: sub}, nil
}

// WatchSearch is a free log subscription operation binding the contract event 0xf9c679450fa2336656a505b21ba8b20cf414bd0e75f0f89cc57be8e08d7a6c5c.
//
// Solidity: event search(uint256 reqID, string dbName, string colName, bytes querys, string callBack, address sender)
func (_Oracle *OracleFilterer) WatchSearch(opts *bind.WatchOpts, sink chan<- *OracleSearch) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "search")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleSearch)
				if err := _Oracle.contract.UnpackLog(event, "search", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSearch is a log parse operation binding the contract event 0xf9c679450fa2336656a505b21ba8b20cf414bd0e75f0f89cc57be8e08d7a6c5c.
//
// Solidity: event search(uint256 reqID, string dbName, string colName, bytes querys, string callBack, address sender)
func (_Oracle *OracleFilterer) ParseSearch(log types.Log) (*OracleSearch, error) {
	event := new(OracleSearch)
	if err := _Oracle.contract.UnpackLog(event, "search", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
