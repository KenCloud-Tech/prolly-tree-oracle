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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"primaryKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"Key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"Val\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"get\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"put\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AllowWrite\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"statement\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreatRsp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"primaryKey\",\"type\":\"string\"}],\"name\":\"Create\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Val\",\"type\":\"string\"}],\"name\":\"Get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"statement\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"GetRsp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Put\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reqID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"statement\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"PutRsp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// AllowWrite is a paid mutator transaction binding the contract method 0x4f0f6f07.
//
// Solidity: function AllowWrite(string dbName, address to) returns(uint256 reqID)
func (_Oracle *OracleTransactor) AllowWrite(opts *bind.TransactOpts, dbName string, to common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "AllowWrite", dbName, to)
}

// AllowWrite is a paid mutator transaction binding the contract method 0x4f0f6f07.
//
// Solidity: function AllowWrite(string dbName, address to) returns(uint256 reqID)
func (_Oracle *OracleSession) AllowWrite(dbName string, to common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AllowWrite(&_Oracle.TransactOpts, dbName, to)
}

// AllowWrite is a paid mutator transaction binding the contract method 0x4f0f6f07.
//
// Solidity: function AllowWrite(string dbName, address to) returns(uint256 reqID)
func (_Oracle *OracleTransactorSession) AllowWrite(dbName string, to common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AllowWrite(&_Oracle.TransactOpts, dbName, to)
}

// CreatRsp is a paid mutator transaction binding the contract method 0x4009c0aa.
//
// Solidity: function CreatRsp(uint256 reqID, bool statement, string dbName, address owner) returns()
func (_Oracle *OracleTransactor) CreatRsp(opts *bind.TransactOpts, reqID *big.Int, statement bool, dbName string, owner common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "CreatRsp", reqID, statement, dbName, owner)
}

// CreatRsp is a paid mutator transaction binding the contract method 0x4009c0aa.
//
// Solidity: function CreatRsp(uint256 reqID, bool statement, string dbName, address owner) returns()
func (_Oracle *OracleSession) CreatRsp(reqID *big.Int, statement bool, dbName string, owner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.CreatRsp(&_Oracle.TransactOpts, reqID, statement, dbName, owner)
}

// CreatRsp is a paid mutator transaction binding the contract method 0x4009c0aa.
//
// Solidity: function CreatRsp(uint256 reqID, bool statement, string dbName, address owner) returns()
func (_Oracle *OracleTransactorSession) CreatRsp(reqID *big.Int, statement bool, dbName string, owner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.CreatRsp(&_Oracle.TransactOpts, reqID, statement, dbName, owner)
}

// Create is a paid mutator transaction binding the contract method 0x38a96ba3.
//
// Solidity: function Create(string dbName, string primaryKey) returns(uint256 reqID)
func (_Oracle *OracleTransactor) Create(opts *bind.TransactOpts, dbName string, primaryKey string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "Create", dbName, primaryKey)
}

// Create is a paid mutator transaction binding the contract method 0x38a96ba3.
//
// Solidity: function Create(string dbName, string primaryKey) returns(uint256 reqID)
func (_Oracle *OracleSession) Create(dbName string, primaryKey string) (*types.Transaction, error) {
	return _Oracle.Contract.Create(&_Oracle.TransactOpts, dbName, primaryKey)
}

// Create is a paid mutator transaction binding the contract method 0x38a96ba3.
//
// Solidity: function Create(string dbName, string primaryKey) returns(uint256 reqID)
func (_Oracle *OracleTransactorSession) Create(dbName string, primaryKey string) (*types.Transaction, error) {
	return _Oracle.Contract.Create(&_Oracle.TransactOpts, dbName, primaryKey)
}

// Get is a paid mutator transaction binding the contract method 0xb04b7826.
//
// Solidity: function Get(string dbName, string Key, string Val) returns(uint256 reqID)
func (_Oracle *OracleTransactor) Get(opts *bind.TransactOpts, dbName string, Key string, Val string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "Get", dbName, Key, Val)
}

// Get is a paid mutator transaction binding the contract method 0xb04b7826.
//
// Solidity: function Get(string dbName, string Key, string Val) returns(uint256 reqID)
func (_Oracle *OracleSession) Get(dbName string, Key string, Val string) (*types.Transaction, error) {
	return _Oracle.Contract.Get(&_Oracle.TransactOpts, dbName, Key, Val)
}

// Get is a paid mutator transaction binding the contract method 0xb04b7826.
//
// Solidity: function Get(string dbName, string Key, string Val) returns(uint256 reqID)
func (_Oracle *OracleTransactorSession) Get(dbName string, Key string, Val string) (*types.Transaction, error) {
	return _Oracle.Contract.Get(&_Oracle.TransactOpts, dbName, Key, Val)
}

// GetRsp is a paid mutator transaction binding the contract method 0xad6e3c9f.
//
// Solidity: function GetRsp(uint256 reqID, bool statement, bytes data, address sender) returns()
func (_Oracle *OracleTransactor) GetRsp(opts *bind.TransactOpts, reqID *big.Int, statement bool, data []byte, sender common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "GetRsp", reqID, statement, data, sender)
}

// GetRsp is a paid mutator transaction binding the contract method 0xad6e3c9f.
//
// Solidity: function GetRsp(uint256 reqID, bool statement, bytes data, address sender) returns()
func (_Oracle *OracleSession) GetRsp(reqID *big.Int, statement bool, data []byte, sender common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.GetRsp(&_Oracle.TransactOpts, reqID, statement, data, sender)
}

// GetRsp is a paid mutator transaction binding the contract method 0xad6e3c9f.
//
// Solidity: function GetRsp(uint256 reqID, bool statement, bytes data, address sender) returns()
func (_Oracle *OracleTransactorSession) GetRsp(reqID *big.Int, statement bool, data []byte, sender common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.GetRsp(&_Oracle.TransactOpts, reqID, statement, data, sender)
}

// Put is a paid mutator transaction binding the contract method 0x36acab85.
//
// Solidity: function Put(string dbName, bytes data) returns(uint256 reqID)
func (_Oracle *OracleTransactor) Put(opts *bind.TransactOpts, dbName string, data []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "Put", dbName, data)
}

// Put is a paid mutator transaction binding the contract method 0x36acab85.
//
// Solidity: function Put(string dbName, bytes data) returns(uint256 reqID)
func (_Oracle *OracleSession) Put(dbName string, data []byte) (*types.Transaction, error) {
	return _Oracle.Contract.Put(&_Oracle.TransactOpts, dbName, data)
}

// Put is a paid mutator transaction binding the contract method 0x36acab85.
//
// Solidity: function Put(string dbName, bytes data) returns(uint256 reqID)
func (_Oracle *OracleTransactorSession) Put(dbName string, data []byte) (*types.Transaction, error) {
	return _Oracle.Contract.Put(&_Oracle.TransactOpts, dbName, data)
}

// PutRsp is a paid mutator transaction binding the contract method 0x24834abb.
//
// Solidity: function PutRsp(uint256 reqID, bool statement, address sender) returns()
func (_Oracle *OracleTransactor) PutRsp(opts *bind.TransactOpts, reqID *big.Int, statement bool, sender common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "PutRsp", reqID, statement, sender)
}

// PutRsp is a paid mutator transaction binding the contract method 0x24834abb.
//
// Solidity: function PutRsp(uint256 reqID, bool statement, address sender) returns()
func (_Oracle *OracleSession) PutRsp(reqID *big.Int, statement bool, sender common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.PutRsp(&_Oracle.TransactOpts, reqID, statement, sender)
}

// PutRsp is a paid mutator transaction binding the contract method 0x24834abb.
//
// Solidity: function PutRsp(uint256 reqID, bool statement, address sender) returns()
func (_Oracle *OracleTransactorSession) PutRsp(reqID *big.Int, statement bool, sender common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.PutRsp(&_Oracle.TransactOpts, reqID, statement, sender)
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
	PrimaryKey string
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x5e08f3373543c9fa1a22411d63cfcba03baf04918709ad2be36cb7904c71f0ab.
//
// Solidity: event create(uint256 reqID, string dbName, string primaryKey, address owner)
func (_Oracle *OracleFilterer) FilterCreate(opts *bind.FilterOpts) (*OracleCreateIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "create")
	if err != nil {
		return nil, err
	}
	return &OracleCreateIterator{contract: _Oracle.contract, event: "create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x5e08f3373543c9fa1a22411d63cfcba03baf04918709ad2be36cb7904c71f0ab.
//
// Solidity: event create(uint256 reqID, string dbName, string primaryKey, address owner)
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

// ParseCreate is a log parse operation binding the contract event 0x5e08f3373543c9fa1a22411d63cfcba03baf04918709ad2be36cb7904c71f0ab.
//
// Solidity: event create(uint256 reqID, string dbName, string primaryKey, address owner)
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
	ReqID  *big.Int
	DbName string
	Key    string
	Val    string
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGet is a free log retrieval operation binding the contract event 0x8ed560550951980644309528009409dc7bc1277cf8741c4e608057a6882b6537.
//
// Solidity: event get(uint256 reqID, string dbName, string Key, string Val, address sender)
func (_Oracle *OracleFilterer) FilterGet(opts *bind.FilterOpts) (*OracleGetIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "get")
	if err != nil {
		return nil, err
	}
	return &OracleGetIterator{contract: _Oracle.contract, event: "get", logs: logs, sub: sub}, nil
}

// WatchGet is a free log subscription operation binding the contract event 0x8ed560550951980644309528009409dc7bc1277cf8741c4e608057a6882b6537.
//
// Solidity: event get(uint256 reqID, string dbName, string Key, string Val, address sender)
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

// ParseGet is a log parse operation binding the contract event 0x8ed560550951980644309528009409dc7bc1277cf8741c4e608057a6882b6537.
//
// Solidity: event get(uint256 reqID, string dbName, string Key, string Val, address sender)
func (_Oracle *OracleFilterer) ParseGet(log types.Log) (*OracleGet, error) {
	event := new(OracleGet)
	if err := _Oracle.contract.UnpackLog(event, "get", log); err != nil {
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
	ReqID  *big.Int
	DbName string
	Data   []byte
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPut is a free log retrieval operation binding the contract event 0xd229ece7f92194d85c5405f897fa56a4f5cbfbee1f872ba130b491006fc58738.
//
// Solidity: event put(uint256 reqID, string dbName, bytes data, address sender)
func (_Oracle *OracleFilterer) FilterPut(opts *bind.FilterOpts) (*OraclePutIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "put")
	if err != nil {
		return nil, err
	}
	return &OraclePutIterator{contract: _Oracle.contract, event: "put", logs: logs, sub: sub}, nil
}

// WatchPut is a free log subscription operation binding the contract event 0xd229ece7f92194d85c5405f897fa56a4f5cbfbee1f872ba130b491006fc58738.
//
// Solidity: event put(uint256 reqID, string dbName, bytes data, address sender)
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

// ParsePut is a log parse operation binding the contract event 0xd229ece7f92194d85c5405f897fa56a4f5cbfbee1f872ba130b491006fc58738.
//
// Solidity: event put(uint256 reqID, string dbName, bytes data, address sender)
func (_Oracle *OracleFilterer) ParsePut(log types.Log) (*OraclePut, error) {
	event := new(OraclePut)
	if err := _Oracle.contract.UnpackLog(event, "put", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
