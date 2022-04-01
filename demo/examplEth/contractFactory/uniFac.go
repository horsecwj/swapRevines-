// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniFac

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
)

// UniFacMetaData contains all meta data concerning the UniFac contract.
var UniFacMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UniFacABI is the input ABI used to generate the binding from.
// Deprecated: Use UniFacMetaData.ABI instead.
var UniFacABI = UniFacMetaData.ABI

// UniFac is an auto generated Go binding around an Ethereum contract.
type UniFac struct {
	UniFacCaller     // Read-only binding to the contract
	UniFacTransactor // Write-only binding to the contract
	UniFacFilterer   // Log filterer for contract events
}

// UniFacCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniFacCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniFacTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniFacTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniFacFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniFacFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniFacSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniFacSession struct {
	Contract     *UniFac           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniFacCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniFacCallerSession struct {
	Contract *UniFacCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UniFacTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniFacTransactorSession struct {
	Contract     *UniFacTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniFacRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniFacRaw struct {
	Contract *UniFac // Generic contract binding to access the raw methods on
}

// UniFacCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniFacCallerRaw struct {
	Contract *UniFacCaller // Generic read-only contract binding to access the raw methods on
}

// UniFacTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniFacTransactorRaw struct {
	Contract *UniFacTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniFac creates a new instance of UniFac, bound to a specific deployed contract.
func NewUniFac(address common.Address, backend bind.ContractBackend) (*UniFac, error) {
	contract, err := bindUniFac(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniFac{UniFacCaller: UniFacCaller{contract: contract}, UniFacTransactor: UniFacTransactor{contract: contract}, UniFacFilterer: UniFacFilterer{contract: contract}}, nil
}

// NewUniFacCaller creates a new read-only instance of UniFac, bound to a specific deployed contract.
func NewUniFacCaller(address common.Address, caller bind.ContractCaller) (*UniFacCaller, error) {
	contract, err := bindUniFac(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniFacCaller{contract: contract}, nil
}

// NewUniFacTransactor creates a new write-only instance of UniFac, bound to a specific deployed contract.
func NewUniFacTransactor(address common.Address, transactor bind.ContractTransactor) (*UniFacTransactor, error) {
	contract, err := bindUniFac(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniFacTransactor{contract: contract}, nil
}

// NewUniFacFilterer creates a new log filterer instance of UniFac, bound to a specific deployed contract.
func NewUniFacFilterer(address common.Address, filterer bind.ContractFilterer) (*UniFacFilterer, error) {
	contract, err := bindUniFac(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniFacFilterer{contract: contract}, nil
}

// bindUniFac binds a generic wrapper to an already deployed contract.
func bindUniFac(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniFacABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniFac *UniFacRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniFac.Contract.UniFacCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniFac *UniFacRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniFac.Contract.UniFacTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniFac *UniFacRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniFac.Contract.UniFacTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniFac *UniFacCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniFac.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniFac *UniFacTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniFac.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniFac *UniFacTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniFac.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_UniFac *UniFacCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UniFac.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_UniFac *UniFacSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _UniFac.Contract.AllPairs(&_UniFac.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_UniFac *UniFacCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _UniFac.Contract.AllPairs(&_UniFac.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_UniFac *UniFacCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniFac.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_UniFac *UniFacSession) AllPairsLength() (*big.Int, error) {
	return _UniFac.Contract.AllPairsLength(&_UniFac.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_UniFac *UniFacCallerSession) AllPairsLength() (*big.Int, error) {
	return _UniFac.Contract.AllPairsLength(&_UniFac.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_UniFac *UniFacCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniFac.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_UniFac *UniFacSession) FeeTo() (common.Address, error) {
	return _UniFac.Contract.FeeTo(&_UniFac.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_UniFac *UniFacCallerSession) FeeTo() (common.Address, error) {
	return _UniFac.Contract.FeeTo(&_UniFac.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_UniFac *UniFacCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniFac.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_UniFac *UniFacSession) FeeToSetter() (common.Address, error) {
	return _UniFac.Contract.FeeToSetter(&_UniFac.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_UniFac *UniFacCallerSession) FeeToSetter() (common.Address, error) {
	return _UniFac.Contract.FeeToSetter(&_UniFac.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_UniFac *UniFacCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _UniFac.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_UniFac *UniFacSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _UniFac.Contract.GetPair(&_UniFac.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_UniFac *UniFacCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _UniFac.Contract.GetPair(&_UniFac.CallOpts, arg0, arg1)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_UniFac *UniFacTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _UniFac.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_UniFac *UniFacSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _UniFac.Contract.CreatePair(&_UniFac.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_UniFac *UniFacTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _UniFac.Contract.CreatePair(&_UniFac.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_UniFac *UniFacTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _UniFac.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_UniFac *UniFacSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _UniFac.Contract.SetFeeTo(&_UniFac.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_UniFac *UniFacTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _UniFac.Contract.SetFeeTo(&_UniFac.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_UniFac *UniFacTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _UniFac.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_UniFac *UniFacSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _UniFac.Contract.SetFeeToSetter(&_UniFac.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_UniFac *UniFacTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _UniFac.Contract.SetFeeToSetter(&_UniFac.TransactOpts, _feeToSetter)
}

// UniFacPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the UniFac contract.
type UniFacPairCreatedIterator struct {
	Event *UniFacPairCreated // Event containing the contract specifics and raw log

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
func (it *UniFacPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniFacPairCreated)
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
		it.Event = new(UniFacPairCreated)
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
func (it *UniFacPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniFacPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniFacPairCreated represents a PairCreated event raised by the UniFac contract.
type UniFacPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_UniFac *UniFacFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*UniFacPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _UniFac.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &UniFacPairCreatedIterator{contract: _UniFac.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_UniFac *UniFacFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *UniFacPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _UniFac.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniFacPairCreated)
				if err := _UniFac.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_UniFac *UniFacFilterer) ParsePairCreated(log types.Log) (*UniFacPairCreated, error) {
	event := new(UniFacPairCreated)
	if err := _UniFac.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
