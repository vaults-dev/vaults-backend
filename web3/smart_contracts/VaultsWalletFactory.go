// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smart_contracts

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

// VaultsWalletFactoryMetaData contains all meta data concerning the VaultsWalletFactory contract.
var VaultsWalletFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"accountImplementation\",\"outputs\":[{\"internalType\":\"contractVaultsWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"contractVaultsWallet\",\"name\":\"ret\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VaultsWalletFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultsWalletFactoryMetaData.ABI instead.
var VaultsWalletFactoryABI = VaultsWalletFactoryMetaData.ABI

// VaultsWalletFactory is an auto generated Go binding around an Ethereum contract.
type VaultsWalletFactory struct {
	VaultsWalletFactoryCaller     // Read-only binding to the contract
	VaultsWalletFactoryTransactor // Write-only binding to the contract
	VaultsWalletFactoryFilterer   // Log filterer for contract events
}

// VaultsWalletFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultsWalletFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultsWalletFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultsWalletFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultsWalletFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultsWalletFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultsWalletFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultsWalletFactorySession struct {
	Contract     *VaultsWalletFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VaultsWalletFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultsWalletFactoryCallerSession struct {
	Contract *VaultsWalletFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// VaultsWalletFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultsWalletFactoryTransactorSession struct {
	Contract     *VaultsWalletFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// VaultsWalletFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultsWalletFactoryRaw struct {
	Contract *VaultsWalletFactory // Generic contract binding to access the raw methods on
}

// VaultsWalletFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultsWalletFactoryCallerRaw struct {
	Contract *VaultsWalletFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VaultsWalletFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultsWalletFactoryTransactorRaw struct {
	Contract *VaultsWalletFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultsWalletFactory creates a new instance of VaultsWalletFactory, bound to a specific deployed contract.
func NewVaultsWalletFactory(address common.Address, backend bind.ContractBackend) (*VaultsWalletFactory, error) {
	contract, err := bindVaultsWalletFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultsWalletFactory{VaultsWalletFactoryCaller: VaultsWalletFactoryCaller{contract: contract}, VaultsWalletFactoryTransactor: VaultsWalletFactoryTransactor{contract: contract}, VaultsWalletFactoryFilterer: VaultsWalletFactoryFilterer{contract: contract}}, nil
}

// NewVaultsWalletFactoryCaller creates a new read-only instance of VaultsWalletFactory, bound to a specific deployed contract.
func NewVaultsWalletFactoryCaller(address common.Address, caller bind.ContractCaller) (*VaultsWalletFactoryCaller, error) {
	contract, err := bindVaultsWalletFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultsWalletFactoryCaller{contract: contract}, nil
}

// NewVaultsWalletFactoryTransactor creates a new write-only instance of VaultsWalletFactory, bound to a specific deployed contract.
func NewVaultsWalletFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultsWalletFactoryTransactor, error) {
	contract, err := bindVaultsWalletFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultsWalletFactoryTransactor{contract: contract}, nil
}

// NewVaultsWalletFactoryFilterer creates a new log filterer instance of VaultsWalletFactory, bound to a specific deployed contract.
func NewVaultsWalletFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultsWalletFactoryFilterer, error) {
	contract, err := bindVaultsWalletFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultsWalletFactoryFilterer{contract: contract}, nil
}

// bindVaultsWalletFactory binds a generic wrapper to an already deployed contract.
func bindVaultsWalletFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultsWalletFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultsWalletFactory *VaultsWalletFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultsWalletFactory.Contract.VaultsWalletFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultsWalletFactory *VaultsWalletFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultsWalletFactory.Contract.VaultsWalletFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultsWalletFactory *VaultsWalletFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultsWalletFactory.Contract.VaultsWalletFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultsWalletFactory *VaultsWalletFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultsWalletFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultsWalletFactory *VaultsWalletFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultsWalletFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultsWalletFactory *VaultsWalletFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultsWalletFactory.Contract.contract.Transact(opts, method, params...)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_VaultsWalletFactory *VaultsWalletFactoryCaller) AccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultsWalletFactory.contract.Call(opts, &out, "accountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_VaultsWalletFactory *VaultsWalletFactorySession) AccountImplementation() (common.Address, error) {
	return _VaultsWalletFactory.Contract.AccountImplementation(&_VaultsWalletFactory.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_VaultsWalletFactory *VaultsWalletFactoryCallerSession) AccountImplementation() (common.Address, error) {
	return _VaultsWalletFactory.Contract.AccountImplementation(&_VaultsWalletFactory.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_VaultsWalletFactory *VaultsWalletFactoryCaller) GetAddress(opts *bind.CallOpts, owner common.Address, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VaultsWalletFactory.contract.Call(opts, &out, "getAddress", owner, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_VaultsWalletFactory *VaultsWalletFactorySession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _VaultsWalletFactory.Contract.GetAddress(&_VaultsWalletFactory.CallOpts, owner, salt)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_VaultsWalletFactory *VaultsWalletFactoryCallerSession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _VaultsWalletFactory.Contract.GetAddress(&_VaultsWalletFactory.CallOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_VaultsWalletFactory *VaultsWalletFactoryTransactor) CreateAccount(opts *bind.TransactOpts, owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _VaultsWalletFactory.contract.Transact(opts, "createAccount", owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_VaultsWalletFactory *VaultsWalletFactorySession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _VaultsWalletFactory.Contract.CreateAccount(&_VaultsWalletFactory.TransactOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_VaultsWalletFactory *VaultsWalletFactoryTransactorSession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _VaultsWalletFactory.Contract.CreateAccount(&_VaultsWalletFactory.TransactOpts, owner, salt)
}
