// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xo

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

// XoMetaData contains all meta data concerning the Xo contract.
var XoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"postId\",\"type\":\"string\"}],\"name\":\"NewGoodVibes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"myId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"targetId\",\"type\":\"string\"}],\"name\":\"NewMutualLike\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaidDM\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poster\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"postId\",\"type\":\"string\"}],\"name\":\"Post\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"postId\",\"type\":\"string\"}],\"name\":\"SBTUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"streak\",\"type\":\"uint256\"}],\"name\":\"SaveStreak\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"streak\",\"type\":\"uint256\"}],\"name\":\"Streak\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"swipedUserId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"cardId\",\"type\":\"string\"}],\"name\":\"Swiped\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"postId\",\"type\":\"string\"}],\"name\":\"goodVibes\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"swipedUserId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cardId\",\"type\":\"string\"}],\"name\":\"like\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"myId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetId\",\"type\":\"string\"}],\"name\":\"mutualLiked\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"paidDM\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"postId\",\"type\":\"string\"}],\"name\":\"post\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"userStreak\",\"type\":\"uint256\"}],\"name\":\"saveStreak\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenId\",\"type\":\"string\"}],\"name\":\"sbtUpdated\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"userStreak\",\"type\":\"uint256\"}],\"name\":\"streak\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// XoABI is the input ABI used to generate the binding from.
// Deprecated: Use XoMetaData.ABI instead.
var XoABI = XoMetaData.ABI

// Xo is an auto generated Go binding around an Ethereum contract.
type Xo struct {
	XoCaller     // Read-only binding to the contract
	XoTransactor // Write-only binding to the contract
	XoFilterer   // Log filterer for contract events
}

// XoCaller is an auto generated read-only Go binding around an Ethereum contract.
type XoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XoSession struct {
	Contract     *Xo               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XoCallerSession struct {
	Contract *XoCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// XoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XoTransactorSession struct {
	Contract     *XoTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XoRaw is an auto generated low-level Go binding around an Ethereum contract.
type XoRaw struct {
	Contract *Xo // Generic contract binding to access the raw methods on
}

// XoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XoCallerRaw struct {
	Contract *XoCaller // Generic read-only contract binding to access the raw methods on
}

// XoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XoTransactorRaw struct {
	Contract *XoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXo creates a new instance of Xo, bound to a specific deployed contract.
func NewXo(address common.Address, backend bind.ContractBackend) (*Xo, error) {
	contract, err := bindXo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Xo{XoCaller: XoCaller{contract: contract}, XoTransactor: XoTransactor{contract: contract}, XoFilterer: XoFilterer{contract: contract}}, nil
}

// NewXoCaller creates a new read-only instance of Xo, bound to a specific deployed contract.
func NewXoCaller(address common.Address, caller bind.ContractCaller) (*XoCaller, error) {
	contract, err := bindXo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XoCaller{contract: contract}, nil
}

// NewXoTransactor creates a new write-only instance of Xo, bound to a specific deployed contract.
func NewXoTransactor(address common.Address, transactor bind.ContractTransactor) (*XoTransactor, error) {
	contract, err := bindXo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XoTransactor{contract: contract}, nil
}

// NewXoFilterer creates a new log filterer instance of Xo, bound to a specific deployed contract.
func NewXoFilterer(address common.Address, filterer bind.ContractFilterer) (*XoFilterer, error) {
	contract, err := bindXo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XoFilterer{contract: contract}, nil
}

// bindXo binds a generic wrapper to an already deployed contract.
func bindXo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := XoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Xo *XoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Xo.Contract.XoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Xo *XoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xo.Contract.XoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Xo *XoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Xo.Contract.XoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Xo *XoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Xo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Xo *XoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Xo *XoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Xo.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Xo *XoCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Xo.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Xo *XoSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Xo.Contract.DEFAULTADMINROLE(&_Xo.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Xo *XoCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Xo.Contract.DEFAULTADMINROLE(&_Xo.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Xo *XoCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Xo.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Xo *XoSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Xo.Contract.GetRoleAdmin(&_Xo.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Xo *XoCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Xo.Contract.GetRoleAdmin(&_Xo.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Xo *XoCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Xo.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Xo *XoSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Xo.Contract.HasRole(&_Xo.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Xo *XoCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Xo.Contract.HasRole(&_Xo.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Xo *XoCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Xo.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Xo *XoSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Xo.Contract.SupportsInterface(&_Xo.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Xo *XoCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Xo.Contract.SupportsInterface(&_Xo.CallOpts, interfaceId)
}

// GoodVibes is a paid mutator transaction binding the contract method 0x85330ce4.
//
// Solidity: function goodVibes(string postId) payable returns()
func (_Xo *XoTransactor) GoodVibes(opts *bind.TransactOpts, postId string) (*types.Transaction, error) {
	return _Xo.contract.Transact(opts, "goodVibes", postId)
}

// GoodVibes is a paid mutator transaction binding the contract method 0x85330ce4.
//
// Solidity: function goodVibes(string postId) payable returns()
func (_Xo *XoSession) GoodVibes(postId string) (*types.Transaction, error) {
	return _Xo.Contract.GoodVibes(&_Xo.TransactOpts, postId)
}

// GoodVibes is a paid mutator transaction binding the contract method 0x85330ce4.
//
// Solidity: function goodVibes(string postId) payable returns()
func (_Xo *XoTransactorSession) GoodVibes(postId string) (*types.Transaction, error) {
	return _Xo.Contract.GoodVibes(&_Xo.TransactOpts, postId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Xo *XoTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Xo.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Xo *XoSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Xo.Contract.GrantRole(&_Xo.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Xo *XoTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Xo.Contract.GrantRole(&_Xo.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_Xo *XoTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Xo.contract.Transact(opts, "initialize", _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_Xo *XoSession) Initialize(_admin common.Address) (*types.Transaction, error) {
	return _Xo.Contract.Initialize(&_Xo.TransactOpts, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_Xo *XoTransactorSession) Initialize(_admin common.Address) (*types.Transaction, error) {
	return _Xo.Contract.Initialize(&_Xo.TransactOpts, _admin)
}

// Like is a paid mutator transaction binding the contract method 0x55678a32.
//
// Solidity: function like(string swipedUserId, string cardId) payable returns()
func (_Xo *XoTransactor) Like(opts *bind.TransactOpts, swipedUserId string, cardId string) (*types.Transaction, error) {
	return _Xo.contract.Transact(opts, "like", swipedUserId, cardId)
}

// Like is a paid mutator transaction binding the contract method 0x55678a32.
//
// Solidity: function like(string swipedUserId, string cardId) payable returns()
func (_Xo *XoSession) Like(swipedUserId string, cardId string) (*types.Transaction, error) {
	return _Xo.Contract.Like(&_Xo.TransactOpts, swipedUserId, cardId)
}

// Like is a paid mutator transaction binding the contract method 0x55678a32.
//
// Solidity: function like(string swipedUserId, string cardId) payable returns()
func (_Xo *XoTransactorSession) Like(swipedUserId string, cardId string) (*types.Transaction, error) {
	return _Xo.Contract.Like(&_Xo.TransactOpts, swipedUserId, cardId)
}

// MutualLiked is a paid mutator transaction binding the contract method 0x0f8314d1.
//
// Solidity: function mutualLiked(string myId, string targetId) payable returns()
func (_Xo *XoTransactor) MutualLiked(opts *bind.TransactOpts, myId string, targetId string) (*types.Transaction, error) {
	return _Xo.contract.Transact(opts, "mutualLiked", myId, targetId)
}

// MutualLiked is a paid mutator transaction binding the contract method 0x0f8314d1.
//
// Solidity: function mutualLiked(string myId, string targetId) payable returns()
func (_Xo *XoSession) MutualLiked(myId string, targetId string) (*types.Transaction, error) {
	return _Xo.Contract.MutualLiked(&_Xo.TransactOpts, myId, targetId)
}

// MutualLiked is a paid mutator transaction binding the contract method 0x0f8314d1.
//
// Solidity: function mutualLiked(string myId, string targetId) payable returns()
func (_Xo *XoTransactorSession) MutualLiked(myId string, targetId string) (*types.Transaction, error) {
	return _Xo.Contract.MutualLiked(&_Xo.TransactOpts, myId, targetId)
}

// PaidDM is a paid mutator transaction binding the contract method 0x1cec6e46.
//
// Solidity: function paidDM(string userId, string denom, uint256 amount) payable returns()
func (_Xo *XoTransactor) PaidDM(opts *bind.TransactOpts, userId string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _Xo.contract.Transact(opts, "paidDM", userId, denom, amount)
}

// PaidDM is a paid mutator transaction binding the contract method 0x1cec6e46.
//
// Solidity: function paidDM(string userId, string denom, uint256 amount) payable returns()
func (_Xo *XoSession) PaidDM(userId string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _Xo.Contract.PaidDM(&_Xo.TransactOpts, userId, denom, amount)
}

// PaidDM is a paid mutator transaction binding the contract method 0x1cec6e46.
//
// Solidity: function paidDM(string userId, string denom, uint256 amount) payable returns()
func (_Xo *XoTransactorSession) PaidDM(userId string, denom string, amount *big.Int) (*types.Transaction, error) {
	return _Xo.Contract.PaidDM(&_Xo.TransactOpts, userId, denom, amount)
}

// Post is a paid mutator transaction binding the contract method 0x8ee93cf3.
//
// Solidity: function post(string postId) payable returns()
func (_Xo *XoTransactor) Post(opts *bind.TransactOpts, postId string) (*types.Transaction, error) {
	return _Xo.contract.Transact(opts, "post", postId)
}

// Post is a paid mutator transaction binding the contract method 0x8ee93cf3.
//
// Solidity: function post(string postId) payable returns()
func (_Xo *XoSession) Post(postId string) (*types.Transaction, error) {
	return _Xo.Contract.Post(&_Xo.TransactOpts, postId)
}

// Post is a paid mutator transaction binding the contract method 0x8ee93cf3.
//
// Solidity: function post(string postId) payable returns()
func (_Xo *XoTransactorSession) Post(postId string) (*types.Transaction, error) {
	return _Xo.Contract.Post(&_Xo.TransactOpts, postId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Xo *XoTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Xo.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Xo *XoSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Xo.Contract.RenounceRole(&_Xo.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Xo *XoTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Xo.Contract.RenounceRole(&_Xo.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Xo *XoTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Xo.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Xo *XoSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Xo.Contract.RevokeRole(&_Xo.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Xo *XoTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Xo.Contract.RevokeRole(&_Xo.TransactOpts, role, account)
}

// SaveStreak is a paid mutator transaction binding the contract method 0x63636565.
//
// Solidity: function saveStreak(string userId, uint256 userStreak) payable returns()
func (_Xo *XoTransactor) SaveStreak(opts *bind.TransactOpts, userId string, userStreak *big.Int) (*types.Transaction, error) {
	return _Xo.contract.Transact(opts, "saveStreak", userId, userStreak)
}

// SaveStreak is a paid mutator transaction binding the contract method 0x63636565.
//
// Solidity: function saveStreak(string userId, uint256 userStreak) payable returns()
func (_Xo *XoSession) SaveStreak(userId string, userStreak *big.Int) (*types.Transaction, error) {
	return _Xo.Contract.SaveStreak(&_Xo.TransactOpts, userId, userStreak)
}

// SaveStreak is a paid mutator transaction binding the contract method 0x63636565.
//
// Solidity: function saveStreak(string userId, uint256 userStreak) payable returns()
func (_Xo *XoTransactorSession) SaveStreak(userId string, userStreak *big.Int) (*types.Transaction, error) {
	return _Xo.Contract.SaveStreak(&_Xo.TransactOpts, userId, userStreak)
}

// SbtUpdated is a paid mutator transaction binding the contract method 0x78fea6ac.
//
// Solidity: function sbtUpdated(string tokenId) payable returns()
func (_Xo *XoTransactor) SbtUpdated(opts *bind.TransactOpts, tokenId string) (*types.Transaction, error) {
	return _Xo.contract.Transact(opts, "sbtUpdated", tokenId)
}

// SbtUpdated is a paid mutator transaction binding the contract method 0x78fea6ac.
//
// Solidity: function sbtUpdated(string tokenId) payable returns()
func (_Xo *XoSession) SbtUpdated(tokenId string) (*types.Transaction, error) {
	return _Xo.Contract.SbtUpdated(&_Xo.TransactOpts, tokenId)
}

// SbtUpdated is a paid mutator transaction binding the contract method 0x78fea6ac.
//
// Solidity: function sbtUpdated(string tokenId) payable returns()
func (_Xo *XoTransactorSession) SbtUpdated(tokenId string) (*types.Transaction, error) {
	return _Xo.Contract.SbtUpdated(&_Xo.TransactOpts, tokenId)
}

// Streak is a paid mutator transaction binding the contract method 0x7d2ceabb.
//
// Solidity: function streak(string userId, uint256 userStreak) payable returns()
func (_Xo *XoTransactor) Streak(opts *bind.TransactOpts, userId string, userStreak *big.Int) (*types.Transaction, error) {
	return _Xo.contract.Transact(opts, "streak", userId, userStreak)
}

// Streak is a paid mutator transaction binding the contract method 0x7d2ceabb.
//
// Solidity: function streak(string userId, uint256 userStreak) payable returns()
func (_Xo *XoSession) Streak(userId string, userStreak *big.Int) (*types.Transaction, error) {
	return _Xo.Contract.Streak(&_Xo.TransactOpts, userId, userStreak)
}

// Streak is a paid mutator transaction binding the contract method 0x7d2ceabb.
//
// Solidity: function streak(string userId, uint256 userStreak) payable returns()
func (_Xo *XoTransactorSession) Streak(userId string, userStreak *big.Int) (*types.Transaction, error) {
	return _Xo.Contract.Streak(&_Xo.TransactOpts, userId, userStreak)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Xo *XoTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Xo.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Xo *XoSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Xo.Contract.Fallback(&_Xo.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Xo *XoTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Xo.Contract.Fallback(&_Xo.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Xo *XoTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xo.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Xo *XoSession) Receive() (*types.Transaction, error) {
	return _Xo.Contract.Receive(&_Xo.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Xo *XoTransactorSession) Receive() (*types.Transaction, error) {
	return _Xo.Contract.Receive(&_Xo.TransactOpts)
}

// XoInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Xo contract.
type XoInitializedIterator struct {
	Event *XoInitialized // Event containing the contract specifics and raw log

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
func (it *XoInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XoInitialized)
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
		it.Event = new(XoInitialized)
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
func (it *XoInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XoInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XoInitialized represents a Initialized event raised by the Xo contract.
type XoInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Xo *XoFilterer) FilterInitialized(opts *bind.FilterOpts) (*XoInitializedIterator, error) {

	logs, sub, err := _Xo.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &XoInitializedIterator{contract: _Xo.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Xo *XoFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *XoInitialized) (event.Subscription, error) {

	logs, sub, err := _Xo.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XoInitialized)
				if err := _Xo.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Xo *XoFilterer) ParseInitialized(log types.Log) (*XoInitialized, error) {
	event := new(XoInitialized)
	if err := _Xo.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XoNewGoodVibesIterator is returned from FilterNewGoodVibes and is used to iterate over the raw logs and unpacked data for NewGoodVibes events raised by the Xo contract.
type XoNewGoodVibesIterator struct {
	Event *XoNewGoodVibes // Event containing the contract specifics and raw log

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
func (it *XoNewGoodVibesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XoNewGoodVibes)
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
		it.Event = new(XoNewGoodVibes)
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
func (it *XoNewGoodVibesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XoNewGoodVibesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XoNewGoodVibes represents a NewGoodVibes event raised by the Xo contract.
type XoNewGoodVibes struct {
	PostId common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewGoodVibes is a free log retrieval operation binding the contract event 0xecd901994167328a5bb90355db2dffc91ae5188c0b89bbe190f837829ee9cb36.
//
// Solidity: event NewGoodVibes(string indexed postId)
func (_Xo *XoFilterer) FilterNewGoodVibes(opts *bind.FilterOpts, postId []string) (*XoNewGoodVibesIterator, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}

	logs, sub, err := _Xo.contract.FilterLogs(opts, "NewGoodVibes", postIdRule)
	if err != nil {
		return nil, err
	}
	return &XoNewGoodVibesIterator{contract: _Xo.contract, event: "NewGoodVibes", logs: logs, sub: sub}, nil
}

// WatchNewGoodVibes is a free log subscription operation binding the contract event 0xecd901994167328a5bb90355db2dffc91ae5188c0b89bbe190f837829ee9cb36.
//
// Solidity: event NewGoodVibes(string indexed postId)
func (_Xo *XoFilterer) WatchNewGoodVibes(opts *bind.WatchOpts, sink chan<- *XoNewGoodVibes, postId []string) (event.Subscription, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}

	logs, sub, err := _Xo.contract.WatchLogs(opts, "NewGoodVibes", postIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XoNewGoodVibes)
				if err := _Xo.contract.UnpackLog(event, "NewGoodVibes", log); err != nil {
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

// ParseNewGoodVibes is a log parse operation binding the contract event 0xecd901994167328a5bb90355db2dffc91ae5188c0b89bbe190f837829ee9cb36.
//
// Solidity: event NewGoodVibes(string indexed postId)
func (_Xo *XoFilterer) ParseNewGoodVibes(log types.Log) (*XoNewGoodVibes, error) {
	event := new(XoNewGoodVibes)
	if err := _Xo.contract.UnpackLog(event, "NewGoodVibes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XoNewMutualLikeIterator is returned from FilterNewMutualLike and is used to iterate over the raw logs and unpacked data for NewMutualLike events raised by the Xo contract.
type XoNewMutualLikeIterator struct {
	Event *XoNewMutualLike // Event containing the contract specifics and raw log

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
func (it *XoNewMutualLikeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XoNewMutualLike)
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
		it.Event = new(XoNewMutualLike)
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
func (it *XoNewMutualLikeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XoNewMutualLikeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XoNewMutualLike represents a NewMutualLike event raised by the Xo contract.
type XoNewMutualLike struct {
	MyId     common.Hash
	TargetId common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewMutualLike is a free log retrieval operation binding the contract event 0xdd4035769ac78355e41b12417a6907043084bf825f28985a3103288dd11b7f6c.
//
// Solidity: event NewMutualLike(string indexed myId, string indexed targetId)
func (_Xo *XoFilterer) FilterNewMutualLike(opts *bind.FilterOpts, myId []string, targetId []string) (*XoNewMutualLikeIterator, error) {

	var myIdRule []interface{}
	for _, myIdItem := range myId {
		myIdRule = append(myIdRule, myIdItem)
	}
	var targetIdRule []interface{}
	for _, targetIdItem := range targetId {
		targetIdRule = append(targetIdRule, targetIdItem)
	}

	logs, sub, err := _Xo.contract.FilterLogs(opts, "NewMutualLike", myIdRule, targetIdRule)
	if err != nil {
		return nil, err
	}
	return &XoNewMutualLikeIterator{contract: _Xo.contract, event: "NewMutualLike", logs: logs, sub: sub}, nil
}

// WatchNewMutualLike is a free log subscription operation binding the contract event 0xdd4035769ac78355e41b12417a6907043084bf825f28985a3103288dd11b7f6c.
//
// Solidity: event NewMutualLike(string indexed myId, string indexed targetId)
func (_Xo *XoFilterer) WatchNewMutualLike(opts *bind.WatchOpts, sink chan<- *XoNewMutualLike, myId []string, targetId []string) (event.Subscription, error) {

	var myIdRule []interface{}
	for _, myIdItem := range myId {
		myIdRule = append(myIdRule, myIdItem)
	}
	var targetIdRule []interface{}
	for _, targetIdItem := range targetId {
		targetIdRule = append(targetIdRule, targetIdItem)
	}

	logs, sub, err := _Xo.contract.WatchLogs(opts, "NewMutualLike", myIdRule, targetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XoNewMutualLike)
				if err := _Xo.contract.UnpackLog(event, "NewMutualLike", log); err != nil {
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

// ParseNewMutualLike is a log parse operation binding the contract event 0xdd4035769ac78355e41b12417a6907043084bf825f28985a3103288dd11b7f6c.
//
// Solidity: event NewMutualLike(string indexed myId, string indexed targetId)
func (_Xo *XoFilterer) ParseNewMutualLike(log types.Log) (*XoNewMutualLike, error) {
	event := new(XoNewMutualLike)
	if err := _Xo.contract.UnpackLog(event, "NewMutualLike", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XoPaidDMIterator is returned from FilterPaidDM and is used to iterate over the raw logs and unpacked data for PaidDM events raised by the Xo contract.
type XoPaidDMIterator struct {
	Event *XoPaidDM // Event containing the contract specifics and raw log

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
func (it *XoPaidDMIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XoPaidDM)
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
		it.Event = new(XoPaidDM)
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
func (it *XoPaidDMIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XoPaidDMIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XoPaidDM represents a PaidDM event raised by the Xo contract.
type XoPaidDM struct {
	UserId common.Hash
	Denom  common.Hash
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaidDM is a free log retrieval operation binding the contract event 0xeb457ead8a262717b1446ab539d804752c92e57463c5e345aea33f10d01cc4dc.
//
// Solidity: event PaidDM(string indexed userId, string indexed denom, uint256 amount)
func (_Xo *XoFilterer) FilterPaidDM(opts *bind.FilterOpts, userId []string, denom []string) (*XoPaidDMIterator, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}

	logs, sub, err := _Xo.contract.FilterLogs(opts, "PaidDM", userIdRule, denomRule)
	if err != nil {
		return nil, err
	}
	return &XoPaidDMIterator{contract: _Xo.contract, event: "PaidDM", logs: logs, sub: sub}, nil
}

// WatchPaidDM is a free log subscription operation binding the contract event 0xeb457ead8a262717b1446ab539d804752c92e57463c5e345aea33f10d01cc4dc.
//
// Solidity: event PaidDM(string indexed userId, string indexed denom, uint256 amount)
func (_Xo *XoFilterer) WatchPaidDM(opts *bind.WatchOpts, sink chan<- *XoPaidDM, userId []string, denom []string) (event.Subscription, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}

	logs, sub, err := _Xo.contract.WatchLogs(opts, "PaidDM", userIdRule, denomRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XoPaidDM)
				if err := _Xo.contract.UnpackLog(event, "PaidDM", log); err != nil {
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

// ParsePaidDM is a log parse operation binding the contract event 0xeb457ead8a262717b1446ab539d804752c92e57463c5e345aea33f10d01cc4dc.
//
// Solidity: event PaidDM(string indexed userId, string indexed denom, uint256 amount)
func (_Xo *XoFilterer) ParsePaidDM(log types.Log) (*XoPaidDM, error) {
	event := new(XoPaidDM)
	if err := _Xo.contract.UnpackLog(event, "PaidDM", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XoPostIterator is returned from FilterPost and is used to iterate over the raw logs and unpacked data for Post events raised by the Xo contract.
type XoPostIterator struct {
	Event *XoPost // Event containing the contract specifics and raw log

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
func (it *XoPostIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XoPost)
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
		it.Event = new(XoPost)
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
func (it *XoPostIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XoPostIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XoPost represents a Post event raised by the Xo contract.
type XoPost struct {
	Poster common.Address
	PostId common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPost is a free log retrieval operation binding the contract event 0x0e0e619c1bc12ccc26618d44ee0d9df23aa32f8a90381ab4d4ce768b70eb7b08.
//
// Solidity: event Post(address indexed poster, string indexed postId)
func (_Xo *XoFilterer) FilterPost(opts *bind.FilterOpts, poster []common.Address, postId []string) (*XoPostIterator, error) {

	var posterRule []interface{}
	for _, posterItem := range poster {
		posterRule = append(posterRule, posterItem)
	}
	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}

	logs, sub, err := _Xo.contract.FilterLogs(opts, "Post", posterRule, postIdRule)
	if err != nil {
		return nil, err
	}
	return &XoPostIterator{contract: _Xo.contract, event: "Post", logs: logs, sub: sub}, nil
}

// WatchPost is a free log subscription operation binding the contract event 0x0e0e619c1bc12ccc26618d44ee0d9df23aa32f8a90381ab4d4ce768b70eb7b08.
//
// Solidity: event Post(address indexed poster, string indexed postId)
func (_Xo *XoFilterer) WatchPost(opts *bind.WatchOpts, sink chan<- *XoPost, poster []common.Address, postId []string) (event.Subscription, error) {

	var posterRule []interface{}
	for _, posterItem := range poster {
		posterRule = append(posterRule, posterItem)
	}
	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}

	logs, sub, err := _Xo.contract.WatchLogs(opts, "Post", posterRule, postIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XoPost)
				if err := _Xo.contract.UnpackLog(event, "Post", log); err != nil {
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

// ParsePost is a log parse operation binding the contract event 0x0e0e619c1bc12ccc26618d44ee0d9df23aa32f8a90381ab4d4ce768b70eb7b08.
//
// Solidity: event Post(address indexed poster, string indexed postId)
func (_Xo *XoFilterer) ParsePost(log types.Log) (*XoPost, error) {
	event := new(XoPost)
	if err := _Xo.contract.UnpackLog(event, "Post", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XoRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Xo contract.
type XoRoleAdminChangedIterator struct {
	Event *XoRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *XoRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XoRoleAdminChanged)
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
		it.Event = new(XoRoleAdminChanged)
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
func (it *XoRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XoRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XoRoleAdminChanged represents a RoleAdminChanged event raised by the Xo contract.
type XoRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Xo *XoFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*XoRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Xo.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &XoRoleAdminChangedIterator{contract: _Xo.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Xo *XoFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *XoRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Xo.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XoRoleAdminChanged)
				if err := _Xo.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Xo *XoFilterer) ParseRoleAdminChanged(log types.Log) (*XoRoleAdminChanged, error) {
	event := new(XoRoleAdminChanged)
	if err := _Xo.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XoRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Xo contract.
type XoRoleGrantedIterator struct {
	Event *XoRoleGranted // Event containing the contract specifics and raw log

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
func (it *XoRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XoRoleGranted)
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
		it.Event = new(XoRoleGranted)
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
func (it *XoRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XoRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XoRoleGranted represents a RoleGranted event raised by the Xo contract.
type XoRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Xo *XoFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*XoRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Xo.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &XoRoleGrantedIterator{contract: _Xo.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Xo *XoFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *XoRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Xo.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XoRoleGranted)
				if err := _Xo.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Xo *XoFilterer) ParseRoleGranted(log types.Log) (*XoRoleGranted, error) {
	event := new(XoRoleGranted)
	if err := _Xo.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XoRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Xo contract.
type XoRoleRevokedIterator struct {
	Event *XoRoleRevoked // Event containing the contract specifics and raw log

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
func (it *XoRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XoRoleRevoked)
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
		it.Event = new(XoRoleRevoked)
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
func (it *XoRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XoRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XoRoleRevoked represents a RoleRevoked event raised by the Xo contract.
type XoRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Xo *XoFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*XoRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Xo.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &XoRoleRevokedIterator{contract: _Xo.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Xo *XoFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *XoRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Xo.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XoRoleRevoked)
				if err := _Xo.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Xo *XoFilterer) ParseRoleRevoked(log types.Log) (*XoRoleRevoked, error) {
	event := new(XoRoleRevoked)
	if err := _Xo.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XoSBTUpdatedIterator is returned from FilterSBTUpdated and is used to iterate over the raw logs and unpacked data for SBTUpdated events raised by the Xo contract.
type XoSBTUpdatedIterator struct {
	Event *XoSBTUpdated // Event containing the contract specifics and raw log

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
func (it *XoSBTUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XoSBTUpdated)
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
		it.Event = new(XoSBTUpdated)
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
func (it *XoSBTUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XoSBTUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XoSBTUpdated represents a SBTUpdated event raised by the Xo contract.
type XoSBTUpdated struct {
	PostId common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSBTUpdated is a free log retrieval operation binding the contract event 0xef2edb72de0ed8efc22951e3da0cd5fa5de0f86a8b73a79474528888f24a08c0.
//
// Solidity: event SBTUpdated(string indexed postId)
func (_Xo *XoFilterer) FilterSBTUpdated(opts *bind.FilterOpts, postId []string) (*XoSBTUpdatedIterator, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}

	logs, sub, err := _Xo.contract.FilterLogs(opts, "SBTUpdated", postIdRule)
	if err != nil {
		return nil, err
	}
	return &XoSBTUpdatedIterator{contract: _Xo.contract, event: "SBTUpdated", logs: logs, sub: sub}, nil
}

// WatchSBTUpdated is a free log subscription operation binding the contract event 0xef2edb72de0ed8efc22951e3da0cd5fa5de0f86a8b73a79474528888f24a08c0.
//
// Solidity: event SBTUpdated(string indexed postId)
func (_Xo *XoFilterer) WatchSBTUpdated(opts *bind.WatchOpts, sink chan<- *XoSBTUpdated, postId []string) (event.Subscription, error) {

	var postIdRule []interface{}
	for _, postIdItem := range postId {
		postIdRule = append(postIdRule, postIdItem)
	}

	logs, sub, err := _Xo.contract.WatchLogs(opts, "SBTUpdated", postIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XoSBTUpdated)
				if err := _Xo.contract.UnpackLog(event, "SBTUpdated", log); err != nil {
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

// ParseSBTUpdated is a log parse operation binding the contract event 0xef2edb72de0ed8efc22951e3da0cd5fa5de0f86a8b73a79474528888f24a08c0.
//
// Solidity: event SBTUpdated(string indexed postId)
func (_Xo *XoFilterer) ParseSBTUpdated(log types.Log) (*XoSBTUpdated, error) {
	event := new(XoSBTUpdated)
	if err := _Xo.contract.UnpackLog(event, "SBTUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XoSaveStreakIterator is returned from FilterSaveStreak and is used to iterate over the raw logs and unpacked data for SaveStreak events raised by the Xo contract.
type XoSaveStreakIterator struct {
	Event *XoSaveStreak // Event containing the contract specifics and raw log

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
func (it *XoSaveStreakIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XoSaveStreak)
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
		it.Event = new(XoSaveStreak)
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
func (it *XoSaveStreakIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XoSaveStreakIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XoSaveStreak represents a SaveStreak event raised by the Xo contract.
type XoSaveStreak struct {
	UserId common.Hash
	Streak *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSaveStreak is a free log retrieval operation binding the contract event 0x1d2df96e42c4c3f93d3981d88ce0d421908431e13bebb6ae932eaca845b39b4b.
//
// Solidity: event SaveStreak(string indexed userId, uint256 streak)
func (_Xo *XoFilterer) FilterSaveStreak(opts *bind.FilterOpts, userId []string) (*XoSaveStreakIterator, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _Xo.contract.FilterLogs(opts, "SaveStreak", userIdRule)
	if err != nil {
		return nil, err
	}
	return &XoSaveStreakIterator{contract: _Xo.contract, event: "SaveStreak", logs: logs, sub: sub}, nil
}

// WatchSaveStreak is a free log subscription operation binding the contract event 0x1d2df96e42c4c3f93d3981d88ce0d421908431e13bebb6ae932eaca845b39b4b.
//
// Solidity: event SaveStreak(string indexed userId, uint256 streak)
func (_Xo *XoFilterer) WatchSaveStreak(opts *bind.WatchOpts, sink chan<- *XoSaveStreak, userId []string) (event.Subscription, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _Xo.contract.WatchLogs(opts, "SaveStreak", userIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XoSaveStreak)
				if err := _Xo.contract.UnpackLog(event, "SaveStreak", log); err != nil {
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

// ParseSaveStreak is a log parse operation binding the contract event 0x1d2df96e42c4c3f93d3981d88ce0d421908431e13bebb6ae932eaca845b39b4b.
//
// Solidity: event SaveStreak(string indexed userId, uint256 streak)
func (_Xo *XoFilterer) ParseSaveStreak(log types.Log) (*XoSaveStreak, error) {
	event := new(XoSaveStreak)
	if err := _Xo.contract.UnpackLog(event, "SaveStreak", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XoStreakIterator is returned from FilterStreak and is used to iterate over the raw logs and unpacked data for Streak events raised by the Xo contract.
type XoStreakIterator struct {
	Event *XoStreak // Event containing the contract specifics and raw log

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
func (it *XoStreakIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XoStreak)
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
		it.Event = new(XoStreak)
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
func (it *XoStreakIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XoStreakIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XoStreak represents a Streak event raised by the Xo contract.
type XoStreak struct {
	UserId common.Hash
	Streak *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStreak is a free log retrieval operation binding the contract event 0x1efb0ae80ec0009194da5969b16d7d40e84d6665e6472ee6e883bfb15b9f53cb.
//
// Solidity: event Streak(string indexed userId, uint256 streak)
func (_Xo *XoFilterer) FilterStreak(opts *bind.FilterOpts, userId []string) (*XoStreakIterator, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _Xo.contract.FilterLogs(opts, "Streak", userIdRule)
	if err != nil {
		return nil, err
	}
	return &XoStreakIterator{contract: _Xo.contract, event: "Streak", logs: logs, sub: sub}, nil
}

// WatchStreak is a free log subscription operation binding the contract event 0x1efb0ae80ec0009194da5969b16d7d40e84d6665e6472ee6e883bfb15b9f53cb.
//
// Solidity: event Streak(string indexed userId, uint256 streak)
func (_Xo *XoFilterer) WatchStreak(opts *bind.WatchOpts, sink chan<- *XoStreak, userId []string) (event.Subscription, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _Xo.contract.WatchLogs(opts, "Streak", userIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XoStreak)
				if err := _Xo.contract.UnpackLog(event, "Streak", log); err != nil {
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

// ParseStreak is a log parse operation binding the contract event 0x1efb0ae80ec0009194da5969b16d7d40e84d6665e6472ee6e883bfb15b9f53cb.
//
// Solidity: event Streak(string indexed userId, uint256 streak)
func (_Xo *XoFilterer) ParseStreak(log types.Log) (*XoStreak, error) {
	event := new(XoStreak)
	if err := _Xo.contract.UnpackLog(event, "Streak", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XoSwipedIterator is returned from FilterSwiped and is used to iterate over the raw logs and unpacked data for Swiped events raised by the Xo contract.
type XoSwipedIterator struct {
	Event *XoSwiped // Event containing the contract specifics and raw log

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
func (it *XoSwipedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XoSwiped)
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
		it.Event = new(XoSwiped)
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
func (it *XoSwipedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XoSwipedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XoSwiped represents a Swiped event raised by the Xo contract.
type XoSwiped struct {
	Sender       common.Address
	SwipedUserId common.Hash
	CardId       common.Hash
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwiped is a free log retrieval operation binding the contract event 0x16d325539f487ad66cd94afa4858461e8f59e6b050c12a291efebd92563aa746.
//
// Solidity: event Swiped(address indexed sender, string indexed swipedUserId, string indexed cardId)
func (_Xo *XoFilterer) FilterSwiped(opts *bind.FilterOpts, sender []common.Address, swipedUserId []string, cardId []string) (*XoSwipedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var swipedUserIdRule []interface{}
	for _, swipedUserIdItem := range swipedUserId {
		swipedUserIdRule = append(swipedUserIdRule, swipedUserIdItem)
	}
	var cardIdRule []interface{}
	for _, cardIdItem := range cardId {
		cardIdRule = append(cardIdRule, cardIdItem)
	}

	logs, sub, err := _Xo.contract.FilterLogs(opts, "Swiped", senderRule, swipedUserIdRule, cardIdRule)
	if err != nil {
		return nil, err
	}
	return &XoSwipedIterator{contract: _Xo.contract, event: "Swiped", logs: logs, sub: sub}, nil
}

// WatchSwiped is a free log subscription operation binding the contract event 0x16d325539f487ad66cd94afa4858461e8f59e6b050c12a291efebd92563aa746.
//
// Solidity: event Swiped(address indexed sender, string indexed swipedUserId, string indexed cardId)
func (_Xo *XoFilterer) WatchSwiped(opts *bind.WatchOpts, sink chan<- *XoSwiped, sender []common.Address, swipedUserId []string, cardId []string) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var swipedUserIdRule []interface{}
	for _, swipedUserIdItem := range swipedUserId {
		swipedUserIdRule = append(swipedUserIdRule, swipedUserIdItem)
	}
	var cardIdRule []interface{}
	for _, cardIdItem := range cardId {
		cardIdRule = append(cardIdRule, cardIdItem)
	}

	logs, sub, err := _Xo.contract.WatchLogs(opts, "Swiped", senderRule, swipedUserIdRule, cardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XoSwiped)
				if err := _Xo.contract.UnpackLog(event, "Swiped", log); err != nil {
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

// ParseSwiped is a log parse operation binding the contract event 0x16d325539f487ad66cd94afa4858461e8f59e6b050c12a291efebd92563aa746.
//
// Solidity: event Swiped(address indexed sender, string indexed swipedUserId, string indexed cardId)
func (_Xo *XoFilterer) ParseSwiped(log types.Log) (*XoSwiped, error) {
	event := new(XoSwiped)
	if err := _Xo.contract.UnpackLog(event, "Swiped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
