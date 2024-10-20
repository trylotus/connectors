// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Adai

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ADAIABI is the input ABI used to generate the binding from.
const ADAIABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// ADAI is an auto generated Go binding around an Ethereum contract.
type ADAI struct {
	ADAICaller     // Read-only binding to the contract
	ADAITransactor // Write-only binding to the contract
	ADAIFilterer   // Log filterer for contract events
}

// ADAICaller is an auto generated read-only Go binding around an Ethereum contract.
type ADAICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADAITransactor is an auto generated write-only Go binding around an Ethereum contract.
type ADAITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADAIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ADAIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADAISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ADAISession struct {
	Contract     *ADAI             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ADAICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ADAICallerSession struct {
	Contract *ADAICaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ADAITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ADAITransactorSession struct {
	Contract     *ADAITransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ADAIRaw is an auto generated low-level Go binding around an Ethereum contract.
type ADAIRaw struct {
	Contract *ADAI // Generic contract binding to access the raw methods on
}

// ADAICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ADAICallerRaw struct {
	Contract *ADAICaller // Generic read-only contract binding to access the raw methods on
}

// ADAITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ADAITransactorRaw struct {
	Contract *ADAITransactor // Generic write-only contract binding to access the raw methods on
}

// NewADAI creates a new instance of ADAI, bound to a specific deployed contract.
func NewADAI(address common.Address, backend bind.ContractBackend) (*ADAI, error) {
	contract, err := bindADAI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ADAI{ADAICaller: ADAICaller{contract: contract}, ADAITransactor: ADAITransactor{contract: contract}, ADAIFilterer: ADAIFilterer{contract: contract}}, nil
}

// NewADAICaller creates a new read-only instance of ADAI, bound to a specific deployed contract.
func NewADAICaller(address common.Address, caller bind.ContractCaller) (*ADAICaller, error) {
	contract, err := bindADAI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ADAICaller{contract: contract}, nil
}

// NewADAITransactor creates a new write-only instance of ADAI, bound to a specific deployed contract.
func NewADAITransactor(address common.Address, transactor bind.ContractTransactor) (*ADAITransactor, error) {
	contract, err := bindADAI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ADAITransactor{contract: contract}, nil
}

// NewADAIFilterer creates a new log filterer instance of ADAI, bound to a specific deployed contract.
func NewADAIFilterer(address common.Address, filterer bind.ContractFilterer) (*ADAIFilterer, error) {
	contract, err := bindADAI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ADAIFilterer{contract: contract}, nil
}

// bindADAI binds a generic wrapper to an already deployed contract.
func bindADAI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ADAIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ADAI *ADAIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ADAI.Contract.ADAICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ADAI *ADAIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADAI.Contract.ADAITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ADAI *ADAIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ADAI.Contract.ADAITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ADAI *ADAICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ADAI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ADAI *ADAITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADAI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ADAI *ADAITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ADAI.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_ADAI *ADAITransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADAI.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_ADAI *ADAISession) Admin() (*types.Transaction, error) {
	return _ADAI.Contract.Admin(&_ADAI.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_ADAI *ADAITransactorSession) Admin() (*types.Transaction, error) {
	return _ADAI.Contract.Admin(&_ADAI.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_ADAI *ADAITransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADAI.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_ADAI *ADAISession) Implementation() (*types.Transaction, error) {
	return _ADAI.Contract.Implementation(&_ADAI.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_ADAI *ADAITransactorSession) Implementation() (*types.Transaction, error) {
	return _ADAI.Contract.Implementation(&_ADAI.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_ADAI *ADAITransactor) Initialize(opts *bind.TransactOpts, _logic common.Address, _data []byte) (*types.Transaction, error) {
	return _ADAI.contract.Transact(opts, "initialize", _logic, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_ADAI *ADAISession) Initialize(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _ADAI.Contract.Initialize(&_ADAI.TransactOpts, _logic, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_ADAI *ADAITransactorSession) Initialize(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _ADAI.Contract.Initialize(&_ADAI.TransactOpts, _logic, _data)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ADAI *ADAITransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ADAI.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ADAI *ADAISession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ADAI.Contract.UpgradeTo(&_ADAI.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ADAI *ADAITransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ADAI.Contract.UpgradeTo(&_ADAI.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ADAI *ADAITransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ADAI.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ADAI *ADAISession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ADAI.Contract.UpgradeToAndCall(&_ADAI.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ADAI *ADAITransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ADAI.Contract.UpgradeToAndCall(&_ADAI.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ADAI *ADAITransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ADAI.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ADAI *ADAISession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ADAI.Contract.Fallback(&_ADAI.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ADAI *ADAITransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ADAI.Contract.Fallback(&_ADAI.TransactOpts, calldata)
}

// ADAIUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ADAI contract.
type ADAIUpgradedIterator struct {
	Event *ADAIUpgraded // Event containing the contract specifics and raw log

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
func (it *ADAIUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ADAIUpgraded)
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
		it.Event = new(ADAIUpgraded)
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
func (it *ADAIUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ADAIUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ADAIUpgraded represents a Upgraded event raised by the ADAI contract.
type ADAIUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ADAI *ADAIFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ADAIUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ADAI.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ADAIUpgradedIterator{contract: _ADAI.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ADAI *ADAIFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ADAIUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ADAI.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ADAIUpgraded)
				if err := _ADAI.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ADAI *ADAIFilterer) ParseUpgraded(log types.Log) (*ADAIUpgraded, error) {
	event := new(ADAIUpgraded)
	if err := _ADAI.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
