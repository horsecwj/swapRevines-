// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package socket

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

// SocketMetaData contains all meta data concerning the Socket contract.
var SocketMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
}

// SocketABI is the input ABI used to generate the binding from.
// Deprecated: Use SocketMetaData.ABI instead.
var SocketABI = SocketMetaData.ABI

// Socket is an auto generated Go binding around an Ethereum contract.
type Socket struct {
	SocketCaller     // Read-only binding to the contract
	SocketTransactor // Write-only binding to the contract
	SocketFilterer   // Log filterer for contract events
}

// SocketCaller is an auto generated read-only Go binding around an Ethereum contract.
type SocketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SocketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SocketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SocketSession struct {
	Contract     *Socket           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SocketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SocketCallerSession struct {
	Contract *SocketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SocketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SocketTransactorSession struct {
	Contract     *SocketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SocketRaw is an auto generated low-level Go binding around an Ethereum contract.
type SocketRaw struct {
	Contract *Socket // Generic contract binding to access the raw methods on
}

// SocketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SocketCallerRaw struct {
	Contract *SocketCaller // Generic read-only contract binding to access the raw methods on
}

// SocketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SocketTransactorRaw struct {
	Contract *SocketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSocket creates a new instance of Socket, bound to a specific deployed contract.
func NewSocket(address common.Address, backend bind.ContractBackend) (*Socket, error) {
	contract, err := bindSocket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Socket{SocketCaller: SocketCaller{contract: contract}, SocketTransactor: SocketTransactor{contract: contract}, SocketFilterer: SocketFilterer{contract: contract}}, nil
}

// NewSocketCaller creates a new read-only instance of Socket, bound to a specific deployed contract.
func NewSocketCaller(address common.Address, caller bind.ContractCaller) (*SocketCaller, error) {
	contract, err := bindSocket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SocketCaller{contract: contract}, nil
}

// NewSocketTransactor creates a new write-only instance of Socket, bound to a specific deployed contract.
func NewSocketTransactor(address common.Address, transactor bind.ContractTransactor) (*SocketTransactor, error) {
	contract, err := bindSocket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SocketTransactor{contract: contract}, nil
}

// NewSocketFilterer creates a new log filterer instance of Socket, bound to a specific deployed contract.
func NewSocketFilterer(address common.Address, filterer bind.ContractFilterer) (*SocketFilterer, error) {
	contract, err := bindSocket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SocketFilterer{contract: contract}, nil
}

// bindSocket binds a generic wrapper to an already deployed contract.
func bindSocket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SocketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Socket *SocketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Socket.Contract.SocketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Socket *SocketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Socket.Contract.SocketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Socket *SocketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Socket.Contract.SocketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Socket *SocketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Socket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Socket *SocketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Socket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Socket *SocketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Socket.Contract.contract.Transact(opts, method, params...)
}

// SocketApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Socket contract.
type SocketApprovalIterator struct {
	Event *SocketApproval // Event containing the contract specifics and raw log

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
func (it *SocketApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocketApproval)
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
		it.Event = new(SocketApproval)
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
func (it *SocketApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocketApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocketApproval represents a Approval event raised by the Socket contract.
type SocketApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Socket *SocketFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*SocketApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Socket.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SocketApprovalIterator{contract: _Socket.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Socket *SocketFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SocketApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Socket.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocketApproval)
				if err := _Socket.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Socket *SocketFilterer) ParseApproval(log types.Log) (*SocketApproval, error) {
	event := new(SocketApproval)
	if err := _Socket.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocketTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Socket contract.
type SocketTransferIterator struct {
	Event *SocketTransfer // Event containing the contract specifics and raw log

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
func (it *SocketTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocketTransfer)
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
		it.Event = new(SocketTransfer)
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
func (it *SocketTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocketTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocketTransfer represents a Transfer event raised by the Socket contract.
type SocketTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Socket *SocketFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SocketTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Socket.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SocketTransferIterator{contract: _Socket.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Socket *SocketFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SocketTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Socket.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocketTransfer)
				if err := _Socket.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Socket *SocketFilterer) ParseTransfer(log types.Log) (*SocketTransfer, error) {
	event := new(SocketTransfer)
	if err := _Socket.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
