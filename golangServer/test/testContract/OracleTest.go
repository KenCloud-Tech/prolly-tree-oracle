// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OracleTest

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

// OracleInterfaceSearchController is an auto generated low-level Go binding around an user-defined struct.
type OracleInterfaceSearchController struct {
	K        string
	Str      string
	Integer  int64
	Uinteger uint64
	Bytess   []byte
	Boolean  bool
	DataType string
	ComOp    string
}

// OracleTestMetaData contains all meta data concerning the OracleTest contract.
var OracleTestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"CatchData\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AllowWrite\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"CBFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"primaryKey\",\"type\":\"string\"}],\"name\":\"Create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"recordID\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"}],\"name\":\"Get\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Key\",\"type\":\"string\"}],\"name\":\"Index\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Put\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dbName\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"k\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"integer\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"uinteger\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"bytess\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"boolean\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"DataType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"comOp\",\"type\":\"string\"}],\"internalType\":\"structOracleInterface.SearchController\",\"name\":\"Val\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"Method\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"}],\"name\":\"Search\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"or\",\"outputs\":[{\"internalType\":\"contractOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OracleTestABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleTestMetaData.ABI instead.
var OracleTestABI = OracleTestMetaData.ABI

// OracleTest is an auto generated Go binding around an Ethereum contract.
type OracleTest struct {
	OracleTestCaller     // Read-only binding to the contract
	OracleTestTransactor // Write-only binding to the contract
	OracleTestFilterer   // Log filterer for contract events
}

// OracleTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleTestSession struct {
	Contract     *OracleTest       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleTestCallerSession struct {
	Contract *OracleTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OracleTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTestTransactorSession struct {
	Contract     *OracleTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OracleTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleTestRaw struct {
	Contract *OracleTest // Generic contract binding to access the raw methods on
}

// OracleTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleTestCallerRaw struct {
	Contract *OracleTestCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTestTransactorRaw struct {
	Contract *OracleTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleTest creates a new instance of OracleTest, bound to a specific deployed contract.
func NewOracleTest(address common.Address, backend bind.ContractBackend) (*OracleTest, error) {
	contract, err := bindOracleTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleTest{OracleTestCaller: OracleTestCaller{contract: contract}, OracleTestTransactor: OracleTestTransactor{contract: contract}, OracleTestFilterer: OracleTestFilterer{contract: contract}}, nil
}

// NewOracleTestCaller creates a new read-only instance of OracleTest, bound to a specific deployed contract.
func NewOracleTestCaller(address common.Address, caller bind.ContractCaller) (*OracleTestCaller, error) {
	contract, err := bindOracleTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTestCaller{contract: contract}, nil
}

// NewOracleTestTransactor creates a new write-only instance of OracleTest, bound to a specific deployed contract.
func NewOracleTestTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTestTransactor, error) {
	contract, err := bindOracleTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTestTransactor{contract: contract}, nil
}

// NewOracleTestFilterer creates a new log filterer instance of OracleTest, bound to a specific deployed contract.
func NewOracleTestFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleTestFilterer, error) {
	contract, err := bindOracleTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleTestFilterer{contract: contract}, nil
}

// bindOracleTest binds a generic wrapper to an already deployed contract.
func bindOracleTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleTest *OracleTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleTest.Contract.OracleTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleTest *OracleTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleTest.Contract.OracleTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleTest *OracleTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleTest.Contract.OracleTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleTest *OracleTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleTest *OracleTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleTest *OracleTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleTest.Contract.contract.Transact(opts, method, params...)
}

// Or is a free data retrieval call binding the contract method 0x2d5e74d8.
//
// Solidity: function or() view returns(address)
func (_OracleTest *OracleTestCaller) Or(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleTest.contract.Call(opts, &out, "or")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Or is a free data retrieval call binding the contract method 0x2d5e74d8.
//
// Solidity: function or() view returns(address)
func (_OracleTest *OracleTestSession) Or() (common.Address, error) {
	return _OracleTest.Contract.Or(&_OracleTest.CallOpts)
}

// Or is a free data retrieval call binding the contract method 0x2d5e74d8.
//
// Solidity: function or() view returns(address)
func (_OracleTest *OracleTestCallerSession) Or() (common.Address, error) {
	return _OracleTest.Contract.Or(&_OracleTest.CallOpts)
}

// AllowWrite is a paid mutator transaction binding the contract method 0x4f0f6f07.
//
// Solidity: function AllowWrite(string dbName, address to) returns()
func (_OracleTest *OracleTestTransactor) AllowWrite(opts *bind.TransactOpts, dbName string, to common.Address) (*types.Transaction, error) {
	return _OracleTest.contract.Transact(opts, "AllowWrite", dbName, to)
}

// AllowWrite is a paid mutator transaction binding the contract method 0x4f0f6f07.
//
// Solidity: function AllowWrite(string dbName, address to) returns()
func (_OracleTest *OracleTestSession) AllowWrite(dbName string, to common.Address) (*types.Transaction, error) {
	return _OracleTest.Contract.AllowWrite(&_OracleTest.TransactOpts, dbName, to)
}

// AllowWrite is a paid mutator transaction binding the contract method 0x4f0f6f07.
//
// Solidity: function AllowWrite(string dbName, address to) returns()
func (_OracleTest *OracleTestTransactorSession) AllowWrite(dbName string, to common.Address) (*types.Transaction, error) {
	return _OracleTest.Contract.AllowWrite(&_OracleTest.TransactOpts, dbName, to)
}

// CBFunc is a paid mutator transaction binding the contract method 0xa374f08b.
//
// Solidity: function CBFunc(bytes data) returns()
func (_OracleTest *OracleTestTransactor) CBFunc(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _OracleTest.contract.Transact(opts, "CBFunc", data)
}

// CBFunc is a paid mutator transaction binding the contract method 0xa374f08b.
//
// Solidity: function CBFunc(bytes data) returns()
func (_OracleTest *OracleTestSession) CBFunc(data []byte) (*types.Transaction, error) {
	return _OracleTest.Contract.CBFunc(&_OracleTest.TransactOpts, data)
}

// CBFunc is a paid mutator transaction binding the contract method 0xa374f08b.
//
// Solidity: function CBFunc(bytes data) returns()
func (_OracleTest *OracleTestTransactorSession) CBFunc(data []byte) (*types.Transaction, error) {
	return _OracleTest.Contract.CBFunc(&_OracleTest.TransactOpts, data)
}

// Create is a paid mutator transaction binding the contract method 0x38a96ba3.
//
// Solidity: function Create(string dbName, string primaryKey) returns()
func (_OracleTest *OracleTestTransactor) Create(opts *bind.TransactOpts, dbName string, primaryKey string) (*types.Transaction, error) {
	return _OracleTest.contract.Transact(opts, "Create", dbName, primaryKey)
}

// Create is a paid mutator transaction binding the contract method 0x38a96ba3.
//
// Solidity: function Create(string dbName, string primaryKey) returns()
func (_OracleTest *OracleTestSession) Create(dbName string, primaryKey string) (*types.Transaction, error) {
	return _OracleTest.Contract.Create(&_OracleTest.TransactOpts, dbName, primaryKey)
}

// Create is a paid mutator transaction binding the contract method 0x38a96ba3.
//
// Solidity: function Create(string dbName, string primaryKey) returns()
func (_OracleTest *OracleTestTransactorSession) Create(dbName string, primaryKey string) (*types.Transaction, error) {
	return _OracleTest.Contract.Create(&_OracleTest.TransactOpts, dbName, primaryKey)
}

// Get is a paid mutator transaction binding the contract method 0x7dc01df6.
//
// Solidity: function Get(string dbName, bytes recordID, string callBack) returns()
func (_OracleTest *OracleTestTransactor) Get(opts *bind.TransactOpts, dbName string, recordID []byte, callBack string) (*types.Transaction, error) {
	return _OracleTest.contract.Transact(opts, "Get", dbName, recordID, callBack)
}

// Get is a paid mutator transaction binding the contract method 0x7dc01df6.
//
// Solidity: function Get(string dbName, bytes recordID, string callBack) returns()
func (_OracleTest *OracleTestSession) Get(dbName string, recordID []byte, callBack string) (*types.Transaction, error) {
	return _OracleTest.Contract.Get(&_OracleTest.TransactOpts, dbName, recordID, callBack)
}

// Get is a paid mutator transaction binding the contract method 0x7dc01df6.
//
// Solidity: function Get(string dbName, bytes recordID, string callBack) returns()
func (_OracleTest *OracleTestTransactorSession) Get(dbName string, recordID []byte, callBack string) (*types.Transaction, error) {
	return _OracleTest.Contract.Get(&_OracleTest.TransactOpts, dbName, recordID, callBack)
}

// Index is a paid mutator transaction binding the contract method 0x7abc4b4c.
//
// Solidity: function Index(string dbName, string Key) returns()
func (_OracleTest *OracleTestTransactor) Index(opts *bind.TransactOpts, dbName string, Key string) (*types.Transaction, error) {
	return _OracleTest.contract.Transact(opts, "Index", dbName, Key)
}

// Index is a paid mutator transaction binding the contract method 0x7abc4b4c.
//
// Solidity: function Index(string dbName, string Key) returns()
func (_OracleTest *OracleTestSession) Index(dbName string, Key string) (*types.Transaction, error) {
	return _OracleTest.Contract.Index(&_OracleTest.TransactOpts, dbName, Key)
}

// Index is a paid mutator transaction binding the contract method 0x7abc4b4c.
//
// Solidity: function Index(string dbName, string Key) returns()
func (_OracleTest *OracleTestTransactorSession) Index(dbName string, Key string) (*types.Transaction, error) {
	return _OracleTest.Contract.Index(&_OracleTest.TransactOpts, dbName, Key)
}

// Put is a paid mutator transaction binding the contract method 0x36acab85.
//
// Solidity: function Put(string dbName, bytes data) returns()
func (_OracleTest *OracleTestTransactor) Put(opts *bind.TransactOpts, dbName string, data []byte) (*types.Transaction, error) {
	return _OracleTest.contract.Transact(opts, "Put", dbName, data)
}

// Put is a paid mutator transaction binding the contract method 0x36acab85.
//
// Solidity: function Put(string dbName, bytes data) returns()
func (_OracleTest *OracleTestSession) Put(dbName string, data []byte) (*types.Transaction, error) {
	return _OracleTest.Contract.Put(&_OracleTest.TransactOpts, dbName, data)
}

// Put is a paid mutator transaction binding the contract method 0x36acab85.
//
// Solidity: function Put(string dbName, bytes data) returns()
func (_OracleTest *OracleTestTransactorSession) Put(dbName string, data []byte) (*types.Transaction, error) {
	return _OracleTest.Contract.Put(&_OracleTest.TransactOpts, dbName, data)
}

// Search is a paid mutator transaction binding the contract method 0xffa6088a.
//
// Solidity: function Search(string dbName, (string,string,int64,uint64,bytes,bool,string,string) Val, string Method, string callBack) returns()
func (_OracleTest *OracleTestTransactor) Search(opts *bind.TransactOpts, dbName string, Val OracleInterfaceSearchController, Method string, callBack string) (*types.Transaction, error) {
	return _OracleTest.contract.Transact(opts, "Search", dbName, Val, Method, callBack)
}

// Search is a paid mutator transaction binding the contract method 0xffa6088a.
//
// Solidity: function Search(string dbName, (string,string,int64,uint64,bytes,bool,string,string) Val, string Method, string callBack) returns()
func (_OracleTest *OracleTestSession) Search(dbName string, Val OracleInterfaceSearchController, Method string, callBack string) (*types.Transaction, error) {
	return _OracleTest.Contract.Search(&_OracleTest.TransactOpts, dbName, Val, Method, callBack)
}

// Search is a paid mutator transaction binding the contract method 0xffa6088a.
//
// Solidity: function Search(string dbName, (string,string,int64,uint64,bytes,bool,string,string) Val, string Method, string callBack) returns()
func (_OracleTest *OracleTestTransactorSession) Search(dbName string, Val OracleInterfaceSearchController, Method string, callBack string) (*types.Transaction, error) {
	return _OracleTest.Contract.Search(&_OracleTest.TransactOpts, dbName, Val, Method, callBack)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address add) returns()
func (_OracleTest *OracleTestTransactor) SetOracle(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _OracleTest.contract.Transact(opts, "setOracle", add)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address add) returns()
func (_OracleTest *OracleTestSession) SetOracle(add common.Address) (*types.Transaction, error) {
	return _OracleTest.Contract.SetOracle(&_OracleTest.TransactOpts, add)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address add) returns()
func (_OracleTest *OracleTestTransactorSession) SetOracle(add common.Address) (*types.Transaction, error) {
	return _OracleTest.Contract.SetOracle(&_OracleTest.TransactOpts, add)
}

// OracleTestCatchDataIterator is returned from FilterCatchData and is used to iterate over the raw logs and unpacked data for CatchData events raised by the OracleTest contract.
type OracleTestCatchDataIterator struct {
	Event *OracleTestCatchData // Event containing the contract specifics and raw log

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
func (it *OracleTestCatchDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTestCatchData)
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
		it.Event = new(OracleTestCatchData)
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
func (it *OracleTestCatchDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTestCatchDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTestCatchData represents a CatchData event raised by the OracleTest contract.
type OracleTestCatchData struct {
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCatchData is a free log retrieval operation binding the contract event 0x4a2b9b80018c64b0e457c10afc077611e20cfdb20be120b8e33903d368cf9e92.
//
// Solidity: event CatchData(bytes data)
func (_OracleTest *OracleTestFilterer) FilterCatchData(opts *bind.FilterOpts) (*OracleTestCatchDataIterator, error) {

	logs, sub, err := _OracleTest.contract.FilterLogs(opts, "CatchData")
	if err != nil {
		return nil, err
	}
	return &OracleTestCatchDataIterator{contract: _OracleTest.contract, event: "CatchData", logs: logs, sub: sub}, nil
}

// WatchCatchData is a free log subscription operation binding the contract event 0x4a2b9b80018c64b0e457c10afc077611e20cfdb20be120b8e33903d368cf9e92.
//
// Solidity: event CatchData(bytes data)
func (_OracleTest *OracleTestFilterer) WatchCatchData(opts *bind.WatchOpts, sink chan<- *OracleTestCatchData) (event.Subscription, error) {

	logs, sub, err := _OracleTest.contract.WatchLogs(opts, "CatchData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTestCatchData)
				if err := _OracleTest.contract.UnpackLog(event, "CatchData", log); err != nil {
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

// ParseCatchData is a log parse operation binding the contract event 0x4a2b9b80018c64b0e457c10afc077611e20cfdb20be120b8e33903d368cf9e92.
//
// Solidity: event CatchData(bytes data)
func (_OracleTest *OracleTestFilterer) ParseCatchData(log types.Log) (*OracleTestCatchData, error) {
	event := new(OracleTestCatchData)
	if err := _OracleTest.contract.UnpackLog(event, "CatchData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
