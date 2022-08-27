// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unbonded_evmos

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

// UnbondedEvmosMetaData contains all meta data concerning the UnbondedEvmos contract.
var UnbondedEvmosMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedIndex\",\"type\":\"uint256\"}],\"name\":\"Lock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Supply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returned\",\"type\":\"uint256\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unbondingInterval\",\"type\":\"uint256\"}],\"name\":\"UpdateConfigs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateMinterStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unbondingInterval_\",\"type\":\"uint256\"}],\"name\":\"__UnbondedEvmos_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUnlockable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unlockable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedId\",\"type\":\"uint256\"}],\"name\":\"isKillable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedId\",\"type\":\"uint256\"}],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUnbondedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastUnlocked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"locks\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"received\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locksLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtShare\",\"type\":\"uint256\"}],\"name\":\"mintLockedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyUnbondedToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unbondingInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minRepaid\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unbondingInterval\",\"type\":\"uint256\"}],\"name\":\"updateConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"updateMinterStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UnbondedEvmosABI is the input ABI used to generate the binding from.
// Deprecated: Use UnbondedEvmosMetaData.ABI instead.
var UnbondedEvmosABI = UnbondedEvmosMetaData.ABI

// UnbondedEvmos is an auto generated Go binding around an Ethereum contract.
type UnbondedEvmos struct {
	UnbondedEvmosCaller     // Read-only binding to the contract
	UnbondedEvmosTransactor // Write-only binding to the contract
	UnbondedEvmosFilterer   // Log filterer for contract events
}

// UnbondedEvmosCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnbondedEvmosCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnbondedEvmosTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnbondedEvmosTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnbondedEvmosFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnbondedEvmosFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnbondedEvmosSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnbondedEvmosSession struct {
	Contract     *UnbondedEvmos    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnbondedEvmosCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnbondedEvmosCallerSession struct {
	Contract *UnbondedEvmosCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UnbondedEvmosTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnbondedEvmosTransactorSession struct {
	Contract     *UnbondedEvmosTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UnbondedEvmosRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnbondedEvmosRaw struct {
	Contract *UnbondedEvmos // Generic contract binding to access the raw methods on
}

// UnbondedEvmosCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnbondedEvmosCallerRaw struct {
	Contract *UnbondedEvmosCaller // Generic read-only contract binding to access the raw methods on
}

// UnbondedEvmosTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnbondedEvmosTransactorRaw struct {
	Contract *UnbondedEvmosTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnbondedEvmos creates a new instance of UnbondedEvmos, bound to a specific deployed contract.
func NewUnbondedEvmos(address common.Address, backend bind.ContractBackend) (*UnbondedEvmos, error) {
	contract, err := bindUnbondedEvmos(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnbondedEvmos{UnbondedEvmosCaller: UnbondedEvmosCaller{contract: contract}, UnbondedEvmosTransactor: UnbondedEvmosTransactor{contract: contract}, UnbondedEvmosFilterer: UnbondedEvmosFilterer{contract: contract}}, nil
}

// NewUnbondedEvmosCaller creates a new read-only instance of UnbondedEvmos, bound to a specific deployed contract.
func NewUnbondedEvmosCaller(address common.Address, caller bind.ContractCaller) (*UnbondedEvmosCaller, error) {
	contract, err := bindUnbondedEvmos(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnbondedEvmosCaller{contract: contract}, nil
}

// NewUnbondedEvmosTransactor creates a new write-only instance of UnbondedEvmos, bound to a specific deployed contract.
func NewUnbondedEvmosTransactor(address common.Address, transactor bind.ContractTransactor) (*UnbondedEvmosTransactor, error) {
	contract, err := bindUnbondedEvmos(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnbondedEvmosTransactor{contract: contract}, nil
}

// NewUnbondedEvmosFilterer creates a new log filterer instance of UnbondedEvmos, bound to a specific deployed contract.
func NewUnbondedEvmosFilterer(address common.Address, filterer bind.ContractFilterer) (*UnbondedEvmosFilterer, error) {
	contract, err := bindUnbondedEvmos(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnbondedEvmosFilterer{contract: contract}, nil
}

// bindUnbondedEvmos binds a generic wrapper to an already deployed contract.
func bindUnbondedEvmos(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnbondedEvmosABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnbondedEvmos *UnbondedEvmosRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnbondedEvmos.Contract.UnbondedEvmosCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnbondedEvmos *UnbondedEvmosRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.UnbondedEvmosTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnbondedEvmos *UnbondedEvmosRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.UnbondedEvmosTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnbondedEvmos *UnbondedEvmosCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnbondedEvmos.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnbondedEvmos *UnbondedEvmosTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnbondedEvmos *UnbondedEvmosTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UnbondedEvmos.Contract.BalanceOf(&_UnbondedEvmos.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UnbondedEvmos.Contract.BalanceOf(&_UnbondedEvmos.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UnbondedEvmos *UnbondedEvmosCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UnbondedEvmos *UnbondedEvmosSession) Decimals() (uint8, error) {
	return _UnbondedEvmos.Contract.Decimals(&_UnbondedEvmos.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) Decimals() (uint8, error) {
	return _UnbondedEvmos.Contract.Decimals(&_UnbondedEvmos.CallOpts)
}

// GetUnlockable is a free data retrieval call binding the contract method 0xc8630178.
//
// Solidity: function getUnlockable(address account) view returns(uint256 unlockable, uint256 debt)
func (_UnbondedEvmos *UnbondedEvmosCaller) GetUnlockable(opts *bind.CallOpts, account common.Address) (struct {
	Unlockable *big.Int
	Debt       *big.Int
}, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "getUnlockable", account)

	outstruct := new(struct {
		Unlockable *big.Int
		Debt       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Unlockable = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Debt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUnlockable is a free data retrieval call binding the contract method 0xc8630178.
//
// Solidity: function getUnlockable(address account) view returns(uint256 unlockable, uint256 debt)
func (_UnbondedEvmos *UnbondedEvmosSession) GetUnlockable(account common.Address) (struct {
	Unlockable *big.Int
	Debt       *big.Int
}, error) {
	return _UnbondedEvmos.Contract.GetUnlockable(&_UnbondedEvmos.CallOpts, account)
}

// GetUnlockable is a free data retrieval call binding the contract method 0xc8630178.
//
// Solidity: function getUnlockable(address account) view returns(uint256 unlockable, uint256 debt)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) GetUnlockable(account common.Address) (struct {
	Unlockable *big.Int
	Debt       *big.Int
}, error) {
	return _UnbondedEvmos.Contract.GetUnlockable(&_UnbondedEvmos.CallOpts, account)
}

// IsKillable is a free data retrieval call binding the contract method 0x97e7381e.
//
// Solidity: function isKillable(uint256 lockedId) view returns(bool)
func (_UnbondedEvmos *UnbondedEvmosCaller) IsKillable(opts *bind.CallOpts, lockedId *big.Int) (bool, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "isKillable", lockedId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKillable is a free data retrieval call binding the contract method 0x97e7381e.
//
// Solidity: function isKillable(uint256 lockedId) view returns(bool)
func (_UnbondedEvmos *UnbondedEvmosSession) IsKillable(lockedId *big.Int) (bool, error) {
	return _UnbondedEvmos.Contract.IsKillable(&_UnbondedEvmos.CallOpts, lockedId)
}

// IsKillable is a free data retrieval call binding the contract method 0x97e7381e.
//
// Solidity: function isKillable(uint256 lockedId) view returns(bool)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) IsKillable(lockedId *big.Int) (bool, error) {
	return _UnbondedEvmos.Contract.IsKillable(&_UnbondedEvmos.CallOpts, lockedId)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_UnbondedEvmos *UnbondedEvmosCaller) IsMinter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "isMinter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_UnbondedEvmos *UnbondedEvmosSession) IsMinter(arg0 common.Address) (bool, error) {
	return _UnbondedEvmos.Contract.IsMinter(&_UnbondedEvmos.CallOpts, arg0)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) IsMinter(arg0 common.Address) (bool, error) {
	return _UnbondedEvmos.Contract.IsMinter(&_UnbondedEvmos.CallOpts, arg0)
}

// LastUnbondedAt is a free data retrieval call binding the contract method 0x9c6b2896.
//
// Solidity: function lastUnbondedAt() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCaller) LastUnbondedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "lastUnbondedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUnbondedAt is a free data retrieval call binding the contract method 0x9c6b2896.
//
// Solidity: function lastUnbondedAt() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosSession) LastUnbondedAt() (*big.Int, error) {
	return _UnbondedEvmos.Contract.LastUnbondedAt(&_UnbondedEvmos.CallOpts)
}

// LastUnbondedAt is a free data retrieval call binding the contract method 0x9c6b2896.
//
// Solidity: function lastUnbondedAt() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) LastUnbondedAt() (*big.Int, error) {
	return _UnbondedEvmos.Contract.LastUnbondedAt(&_UnbondedEvmos.CallOpts)
}

// LockedOf is a free data retrieval call binding the contract method 0xa5f1e282.
//
// Solidity: function lockedOf(address ) view returns(uint256 lastUnlocked)
func (_UnbondedEvmos *UnbondedEvmosCaller) LockedOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "lockedOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedOf is a free data retrieval call binding the contract method 0xa5f1e282.
//
// Solidity: function lockedOf(address ) view returns(uint256 lastUnlocked)
func (_UnbondedEvmos *UnbondedEvmosSession) LockedOf(arg0 common.Address) (*big.Int, error) {
	return _UnbondedEvmos.Contract.LockedOf(&_UnbondedEvmos.CallOpts, arg0)
}

// LockedOf is a free data retrieval call binding the contract method 0xa5f1e282.
//
// Solidity: function lockedOf(address ) view returns(uint256 lastUnlocked)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) LockedOf(arg0 common.Address) (*big.Int, error) {
	return _UnbondedEvmos.Contract.LockedOf(&_UnbondedEvmos.CallOpts, arg0)
}

// Locks is a free data retrieval call binding the contract method 0xf4dadc61.
//
// Solidity: function locks(uint256 ) view returns(bool received, address account, address vault, uint256 amount, uint256 debtShare, uint256 unlockedAt)
func (_UnbondedEvmos *UnbondedEvmosCaller) Locks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Received   bool
	Account    common.Address
	Vault      common.Address
	Amount     *big.Int
	DebtShare  *big.Int
	UnlockedAt *big.Int
}, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "locks", arg0)

	outstruct := new(struct {
		Received   bool
		Account    common.Address
		Vault      common.Address
		Amount     *big.Int
		DebtShare  *big.Int
		UnlockedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Received = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Account = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Vault = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DebtShare = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UnlockedAt = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Locks is a free data retrieval call binding the contract method 0xf4dadc61.
//
// Solidity: function locks(uint256 ) view returns(bool received, address account, address vault, uint256 amount, uint256 debtShare, uint256 unlockedAt)
func (_UnbondedEvmos *UnbondedEvmosSession) Locks(arg0 *big.Int) (struct {
	Received   bool
	Account    common.Address
	Vault      common.Address
	Amount     *big.Int
	DebtShare  *big.Int
	UnlockedAt *big.Int
}, error) {
	return _UnbondedEvmos.Contract.Locks(&_UnbondedEvmos.CallOpts, arg0)
}

// Locks is a free data retrieval call binding the contract method 0xf4dadc61.
//
// Solidity: function locks(uint256 ) view returns(bool received, address account, address vault, uint256 amount, uint256 debtShare, uint256 unlockedAt)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) Locks(arg0 *big.Int) (struct {
	Received   bool
	Account    common.Address
	Vault      common.Address
	Amount     *big.Int
	DebtShare  *big.Int
	UnlockedAt *big.Int
}, error) {
	return _UnbondedEvmos.Contract.Locks(&_UnbondedEvmos.CallOpts, arg0)
}

// LocksLength is a free data retrieval call binding the contract method 0x5bfc5a6e.
//
// Solidity: function locksLength() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCaller) LocksLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "locksLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocksLength is a free data retrieval call binding the contract method 0x5bfc5a6e.
//
// Solidity: function locksLength() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosSession) LocksLength() (*big.Int, error) {
	return _UnbondedEvmos.Contract.LocksLength(&_UnbondedEvmos.CallOpts)
}

// LocksLength is a free data retrieval call binding the contract method 0x5bfc5a6e.
//
// Solidity: function locksLength() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) LocksLength() (*big.Int, error) {
	return _UnbondedEvmos.Contract.LocksLength(&_UnbondedEvmos.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UnbondedEvmos *UnbondedEvmosCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UnbondedEvmos *UnbondedEvmosSession) Name() (string, error) {
	return _UnbondedEvmos.Contract.Name(&_UnbondedEvmos.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) Name() (string, error) {
	return _UnbondedEvmos.Contract.Name(&_UnbondedEvmos.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnbondedEvmos *UnbondedEvmosCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnbondedEvmos *UnbondedEvmosSession) Owner() (common.Address, error) {
	return _UnbondedEvmos.Contract.Owner(&_UnbondedEvmos.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) Owner() (common.Address, error) {
	return _UnbondedEvmos.Contract.Owner(&_UnbondedEvmos.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UnbondedEvmos *UnbondedEvmosCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UnbondedEvmos *UnbondedEvmosSession) Symbol() (string, error) {
	return _UnbondedEvmos.Contract.Symbol(&_UnbondedEvmos.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) Symbol() (string, error) {
	return _UnbondedEvmos.Contract.Symbol(&_UnbondedEvmos.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosSession) TotalAmount() (*big.Int, error) {
	return _UnbondedEvmos.Contract.TotalAmount(&_UnbondedEvmos.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) TotalAmount() (*big.Int, error) {
	return _UnbondedEvmos.Contract.TotalAmount(&_UnbondedEvmos.CallOpts)
}

// TotalShare is a free data retrieval call binding the contract method 0x026c4207.
//
// Solidity: function totalShare() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCaller) TotalShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "totalShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShare is a free data retrieval call binding the contract method 0x026c4207.
//
// Solidity: function totalShare() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosSession) TotalShare() (*big.Int, error) {
	return _UnbondedEvmos.Contract.TotalShare(&_UnbondedEvmos.CallOpts)
}

// TotalShare is a free data retrieval call binding the contract method 0x026c4207.
//
// Solidity: function totalShare() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) TotalShare() (*big.Int, error) {
	return _UnbondedEvmos.Contract.TotalShare(&_UnbondedEvmos.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosSession) TotalSupply() (*big.Int, error) {
	return _UnbondedEvmos.Contract.TotalSupply(&_UnbondedEvmos.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) TotalSupply() (*big.Int, error) {
	return _UnbondedEvmos.Contract.TotalSupply(&_UnbondedEvmos.CallOpts)
}

// UnbondingInterval is a free data retrieval call binding the contract method 0xa4086a8f.
//
// Solidity: function unbondingInterval() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCaller) UnbondingInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnbondedEvmos.contract.Call(opts, &out, "unbondingInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondingInterval is a free data retrieval call binding the contract method 0xa4086a8f.
//
// Solidity: function unbondingInterval() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosSession) UnbondingInterval() (*big.Int, error) {
	return _UnbondedEvmos.Contract.UnbondingInterval(&_UnbondedEvmos.CallOpts)
}

// UnbondingInterval is a free data retrieval call binding the contract method 0xa4086a8f.
//
// Solidity: function unbondingInterval() view returns(uint256)
func (_UnbondedEvmos *UnbondedEvmosCallerSession) UnbondingInterval() (*big.Int, error) {
	return _UnbondedEvmos.Contract.UnbondingInterval(&_UnbondedEvmos.CallOpts)
}

// UnbondedEvmosInit is a paid mutator transaction binding the contract method 0xcc9bc4af.
//
// Solidity: function __UnbondedEvmos_init(uint256 unbondingInterval_) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactor) UnbondedEvmosInit(opts *bind.TransactOpts, unbondingInterval_ *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.contract.Transact(opts, "__UnbondedEvmos_init", unbondingInterval_)
}

// UnbondedEvmosInit is a paid mutator transaction binding the contract method 0xcc9bc4af.
//
// Solidity: function __UnbondedEvmos_init(uint256 unbondingInterval_) returns()
func (_UnbondedEvmos *UnbondedEvmosSession) UnbondedEvmosInit(unbondingInterval_ *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.UnbondedEvmosInit(&_UnbondedEvmos.TransactOpts, unbondingInterval_)
}

// UnbondedEvmosInit is a paid mutator transaction binding the contract method 0xcc9bc4af.
//
// Solidity: function __UnbondedEvmos_init(uint256 unbondingInterval_) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactorSession) UnbondedEvmosInit(unbondingInterval_ *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.UnbondedEvmosInit(&_UnbondedEvmos.TransactOpts, unbondingInterval_)
}

// Kill is a paid mutator transaction binding the contract method 0xd29a0025.
//
// Solidity: function kill(uint256 lockedId) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactor) Kill(opts *bind.TransactOpts, lockedId *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.contract.Transact(opts, "kill", lockedId)
}

// Kill is a paid mutator transaction binding the contract method 0xd29a0025.
//
// Solidity: function kill(uint256 lockedId) returns()
func (_UnbondedEvmos *UnbondedEvmosSession) Kill(lockedId *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.Kill(&_UnbondedEvmos.TransactOpts, lockedId)
}

// Kill is a paid mutator transaction binding the contract method 0xd29a0025.
//
// Solidity: function kill(uint256 lockedId) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactorSession) Kill(lockedId *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.Kill(&_UnbondedEvmos.TransactOpts, lockedId)
}

// MintLockedToken is a paid mutator transaction binding the contract method 0x633235d0.
//
// Solidity: function mintLockedToken(address to, address vault, uint256 amount, uint256 debtShare) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactor) MintLockedToken(opts *bind.TransactOpts, to common.Address, vault common.Address, amount *big.Int, debtShare *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.contract.Transact(opts, "mintLockedToken", to, vault, amount, debtShare)
}

// MintLockedToken is a paid mutator transaction binding the contract method 0x633235d0.
//
// Solidity: function mintLockedToken(address to, address vault, uint256 amount, uint256 debtShare) returns()
func (_UnbondedEvmos *UnbondedEvmosSession) MintLockedToken(to common.Address, vault common.Address, amount *big.Int, debtShare *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.MintLockedToken(&_UnbondedEvmos.TransactOpts, to, vault, amount, debtShare)
}

// MintLockedToken is a paid mutator transaction binding the contract method 0x633235d0.
//
// Solidity: function mintLockedToken(address to, address vault, uint256 amount, uint256 debtShare) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactorSession) MintLockedToken(to common.Address, vault common.Address, amount *big.Int, debtShare *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.MintLockedToken(&_UnbondedEvmos.TransactOpts, to, vault, amount, debtShare)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnbondedEvmos *UnbondedEvmosTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnbondedEvmos.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnbondedEvmos *UnbondedEvmosSession) RenounceOwnership() (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.RenounceOwnership(&_UnbondedEvmos.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnbondedEvmos *UnbondedEvmosTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.RenounceOwnership(&_UnbondedEvmos.TransactOpts)
}

// SupplyUnbondedToken is a paid mutator transaction binding the contract method 0xe158c1ac.
//
// Solidity: function supplyUnbondedToken() payable returns()
func (_UnbondedEvmos *UnbondedEvmosTransactor) SupplyUnbondedToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnbondedEvmos.contract.Transact(opts, "supplyUnbondedToken")
}

// SupplyUnbondedToken is a paid mutator transaction binding the contract method 0xe158c1ac.
//
// Solidity: function supplyUnbondedToken() payable returns()
func (_UnbondedEvmos *UnbondedEvmosSession) SupplyUnbondedToken() (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.SupplyUnbondedToken(&_UnbondedEvmos.TransactOpts)
}

// SupplyUnbondedToken is a paid mutator transaction binding the contract method 0xe158c1ac.
//
// Solidity: function supplyUnbondedToken() payable returns()
func (_UnbondedEvmos *UnbondedEvmosTransactorSession) SupplyUnbondedToken() (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.SupplyUnbondedToken(&_UnbondedEvmos.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UnbondedEvmos.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnbondedEvmos *UnbondedEvmosSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.TransferOwnership(&_UnbondedEvmos.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.TransferOwnership(&_UnbondedEvmos.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 minRepaid) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactor) Unlock(opts *bind.TransactOpts, minRepaid *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.contract.Transact(opts, "unlock", minRepaid)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 minRepaid) returns()
func (_UnbondedEvmos *UnbondedEvmosSession) Unlock(minRepaid *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.Unlock(&_UnbondedEvmos.TransactOpts, minRepaid)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 minRepaid) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactorSession) Unlock(minRepaid *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.Unlock(&_UnbondedEvmos.TransactOpts, minRepaid)
}

// UpdateConfigs is a paid mutator transaction binding the contract method 0x79564ace.
//
// Solidity: function updateConfigs(uint256 _unbondingInterval) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactor) UpdateConfigs(opts *bind.TransactOpts, _unbondingInterval *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.contract.Transact(opts, "updateConfigs", _unbondingInterval)
}

// UpdateConfigs is a paid mutator transaction binding the contract method 0x79564ace.
//
// Solidity: function updateConfigs(uint256 _unbondingInterval) returns()
func (_UnbondedEvmos *UnbondedEvmosSession) UpdateConfigs(_unbondingInterval *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.UpdateConfigs(&_UnbondedEvmos.TransactOpts, _unbondingInterval)
}

// UpdateConfigs is a paid mutator transaction binding the contract method 0x79564ace.
//
// Solidity: function updateConfigs(uint256 _unbondingInterval) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactorSession) UpdateConfigs(_unbondingInterval *big.Int) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.UpdateConfigs(&_UnbondedEvmos.TransactOpts, _unbondingInterval)
}

// UpdateMinterStatus is a paid mutator transaction binding the contract method 0xefc8f3c8.
//
// Solidity: function updateMinterStatus(address account, bool status) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactor) UpdateMinterStatus(opts *bind.TransactOpts, account common.Address, status bool) (*types.Transaction, error) {
	return _UnbondedEvmos.contract.Transact(opts, "updateMinterStatus", account, status)
}

// UpdateMinterStatus is a paid mutator transaction binding the contract method 0xefc8f3c8.
//
// Solidity: function updateMinterStatus(address account, bool status) returns()
func (_UnbondedEvmos *UnbondedEvmosSession) UpdateMinterStatus(account common.Address, status bool) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.UpdateMinterStatus(&_UnbondedEvmos.TransactOpts, account, status)
}

// UpdateMinterStatus is a paid mutator transaction binding the contract method 0xefc8f3c8.
//
// Solidity: function updateMinterStatus(address account, bool status) returns()
func (_UnbondedEvmos *UnbondedEvmosTransactorSession) UpdateMinterStatus(account common.Address, status bool) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.UpdateMinterStatus(&_UnbondedEvmos.TransactOpts, account, status)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UnbondedEvmos *UnbondedEvmosTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _UnbondedEvmos.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UnbondedEvmos *UnbondedEvmosSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.Fallback(&_UnbondedEvmos.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UnbondedEvmos *UnbondedEvmosTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.Fallback(&_UnbondedEvmos.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UnbondedEvmos *UnbondedEvmosTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnbondedEvmos.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UnbondedEvmos *UnbondedEvmosSession) Receive() (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.Receive(&_UnbondedEvmos.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UnbondedEvmos *UnbondedEvmosTransactorSession) Receive() (*types.Transaction, error) {
	return _UnbondedEvmos.Contract.Receive(&_UnbondedEvmos.TransactOpts)
}

// UnbondedEvmosInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UnbondedEvmos contract.
type UnbondedEvmosInitializedIterator struct {
	Event *UnbondedEvmosInitialized // Event containing the contract specifics and raw log

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
func (it *UnbondedEvmosInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnbondedEvmosInitialized)
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
		it.Event = new(UnbondedEvmosInitialized)
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
func (it *UnbondedEvmosInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnbondedEvmosInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnbondedEvmosInitialized represents a Initialized event raised by the UnbondedEvmos contract.
type UnbondedEvmosInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UnbondedEvmos *UnbondedEvmosFilterer) FilterInitialized(opts *bind.FilterOpts) (*UnbondedEvmosInitializedIterator, error) {

	logs, sub, err := _UnbondedEvmos.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UnbondedEvmosInitializedIterator{contract: _UnbondedEvmos.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UnbondedEvmos *UnbondedEvmosFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UnbondedEvmosInitialized) (event.Subscription, error) {

	logs, sub, err := _UnbondedEvmos.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnbondedEvmosInitialized)
				if err := _UnbondedEvmos.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UnbondedEvmos *UnbondedEvmosFilterer) ParseInitialized(log types.Log) (*UnbondedEvmosInitialized, error) {
	event := new(UnbondedEvmosInitialized)
	if err := _UnbondedEvmos.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnbondedEvmosLockIterator is returned from FilterLock and is used to iterate over the raw logs and unpacked data for Lock events raised by the UnbondedEvmos contract.
type UnbondedEvmosLockIterator struct {
	Event *UnbondedEvmosLock // Event containing the contract specifics and raw log

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
func (it *UnbondedEvmosLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnbondedEvmosLock)
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
		it.Event = new(UnbondedEvmosLock)
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
func (it *UnbondedEvmosLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnbondedEvmosLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnbondedEvmosLock represents a Lock event raised by the UnbondedEvmos contract.
type UnbondedEvmosLock struct {
	Account     common.Address
	Vault       common.Address
	LockedIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLock is a free log retrieval operation binding the contract event 0xec36c0364d931187a76cf66d7eee08fad0ec2e8b7458a8d8b26b36769d4d13f3.
//
// Solidity: event Lock(address account, address vault, uint256 lockedIndex)
func (_UnbondedEvmos *UnbondedEvmosFilterer) FilterLock(opts *bind.FilterOpts) (*UnbondedEvmosLockIterator, error) {

	logs, sub, err := _UnbondedEvmos.contract.FilterLogs(opts, "Lock")
	if err != nil {
		return nil, err
	}
	return &UnbondedEvmosLockIterator{contract: _UnbondedEvmos.contract, event: "Lock", logs: logs, sub: sub}, nil
}

// WatchLock is a free log subscription operation binding the contract event 0xec36c0364d931187a76cf66d7eee08fad0ec2e8b7458a8d8b26b36769d4d13f3.
//
// Solidity: event Lock(address account, address vault, uint256 lockedIndex)
func (_UnbondedEvmos *UnbondedEvmosFilterer) WatchLock(opts *bind.WatchOpts, sink chan<- *UnbondedEvmosLock) (event.Subscription, error) {

	logs, sub, err := _UnbondedEvmos.contract.WatchLogs(opts, "Lock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnbondedEvmosLock)
				if err := _UnbondedEvmos.contract.UnpackLog(event, "Lock", log); err != nil {
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

// ParseLock is a log parse operation binding the contract event 0xec36c0364d931187a76cf66d7eee08fad0ec2e8b7458a8d8b26b36769d4d13f3.
//
// Solidity: event Lock(address account, address vault, uint256 lockedIndex)
func (_UnbondedEvmos *UnbondedEvmosFilterer) ParseLock(log types.Log) (*UnbondedEvmosLock, error) {
	event := new(UnbondedEvmosLock)
	if err := _UnbondedEvmos.contract.UnpackLog(event, "Lock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnbondedEvmosOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UnbondedEvmos contract.
type UnbondedEvmosOwnershipTransferredIterator struct {
	Event *UnbondedEvmosOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UnbondedEvmosOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnbondedEvmosOwnershipTransferred)
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
		it.Event = new(UnbondedEvmosOwnershipTransferred)
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
func (it *UnbondedEvmosOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnbondedEvmosOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnbondedEvmosOwnershipTransferred represents a OwnershipTransferred event raised by the UnbondedEvmos contract.
type UnbondedEvmosOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnbondedEvmos *UnbondedEvmosFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UnbondedEvmosOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UnbondedEvmos.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UnbondedEvmosOwnershipTransferredIterator{contract: _UnbondedEvmos.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnbondedEvmos *UnbondedEvmosFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UnbondedEvmosOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UnbondedEvmos.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnbondedEvmosOwnershipTransferred)
				if err := _UnbondedEvmos.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnbondedEvmos *UnbondedEvmosFilterer) ParseOwnershipTransferred(log types.Log) (*UnbondedEvmosOwnershipTransferred, error) {
	event := new(UnbondedEvmosOwnershipTransferred)
	if err := _UnbondedEvmos.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnbondedEvmosSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the UnbondedEvmos contract.
type UnbondedEvmosSupplyIterator struct {
	Event *UnbondedEvmosSupply // Event containing the contract specifics and raw log

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
func (it *UnbondedEvmosSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnbondedEvmosSupply)
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
		it.Event = new(UnbondedEvmosSupply)
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
func (it *UnbondedEvmosSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnbondedEvmosSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnbondedEvmosSupply represents a Supply event raised by the UnbondedEvmos contract.
type UnbondedEvmosSupply struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x1862f918d5600ec0980589e8cc806b3c79b1e762fcbf44cc2947ba12499207eb.
//
// Solidity: event Supply(uint256 amount)
func (_UnbondedEvmos *UnbondedEvmosFilterer) FilterSupply(opts *bind.FilterOpts) (*UnbondedEvmosSupplyIterator, error) {

	logs, sub, err := _UnbondedEvmos.contract.FilterLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return &UnbondedEvmosSupplyIterator{contract: _UnbondedEvmos.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x1862f918d5600ec0980589e8cc806b3c79b1e762fcbf44cc2947ba12499207eb.
//
// Solidity: event Supply(uint256 amount)
func (_UnbondedEvmos *UnbondedEvmosFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *UnbondedEvmosSupply) (event.Subscription, error) {

	logs, sub, err := _UnbondedEvmos.contract.WatchLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnbondedEvmosSupply)
				if err := _UnbondedEvmos.contract.UnpackLog(event, "Supply", log); err != nil {
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

// ParseSupply is a log parse operation binding the contract event 0x1862f918d5600ec0980589e8cc806b3c79b1e762fcbf44cc2947ba12499207eb.
//
// Solidity: event Supply(uint256 amount)
func (_UnbondedEvmos *UnbondedEvmosFilterer) ParseSupply(log types.Log) (*UnbondedEvmosSupply, error) {
	event := new(UnbondedEvmosSupply)
	if err := _UnbondedEvmos.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnbondedEvmosUnlockIterator is returned from FilterUnlock and is used to iterate over the raw logs and unpacked data for Unlock events raised by the UnbondedEvmos contract.
type UnbondedEvmosUnlockIterator struct {
	Event *UnbondedEvmosUnlock // Event containing the contract specifics and raw log

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
func (it *UnbondedEvmosUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnbondedEvmosUnlock)
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
		it.Event = new(UnbondedEvmosUnlock)
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
func (it *UnbondedEvmosUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnbondedEvmosUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnbondedEvmosUnlock represents a Unlock event raised by the UnbondedEvmos contract.
type UnbondedEvmosUnlock struct {
	Account  common.Address
	Amount   *big.Int
	Returned *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnlock is a free log retrieval operation binding the contract event 0xf7870c5b224cbc19873599e46ccfc7103934650509b1af0c3ce90138377c2004.
//
// Solidity: event Unlock(address account, uint256 amount, uint256 returned)
func (_UnbondedEvmos *UnbondedEvmosFilterer) FilterUnlock(opts *bind.FilterOpts) (*UnbondedEvmosUnlockIterator, error) {

	logs, sub, err := _UnbondedEvmos.contract.FilterLogs(opts, "Unlock")
	if err != nil {
		return nil, err
	}
	return &UnbondedEvmosUnlockIterator{contract: _UnbondedEvmos.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0xf7870c5b224cbc19873599e46ccfc7103934650509b1af0c3ce90138377c2004.
//
// Solidity: event Unlock(address account, uint256 amount, uint256 returned)
func (_UnbondedEvmos *UnbondedEvmosFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *UnbondedEvmosUnlock) (event.Subscription, error) {

	logs, sub, err := _UnbondedEvmos.contract.WatchLogs(opts, "Unlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnbondedEvmosUnlock)
				if err := _UnbondedEvmos.contract.UnpackLog(event, "Unlock", log); err != nil {
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

// ParseUnlock is a log parse operation binding the contract event 0xf7870c5b224cbc19873599e46ccfc7103934650509b1af0c3ce90138377c2004.
//
// Solidity: event Unlock(address account, uint256 amount, uint256 returned)
func (_UnbondedEvmos *UnbondedEvmosFilterer) ParseUnlock(log types.Log) (*UnbondedEvmosUnlock, error) {
	event := new(UnbondedEvmosUnlock)
	if err := _UnbondedEvmos.contract.UnpackLog(event, "Unlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnbondedEvmosUpdateConfigsIterator is returned from FilterUpdateConfigs and is used to iterate over the raw logs and unpacked data for UpdateConfigs events raised by the UnbondedEvmos contract.
type UnbondedEvmosUpdateConfigsIterator struct {
	Event *UnbondedEvmosUpdateConfigs // Event containing the contract specifics and raw log

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
func (it *UnbondedEvmosUpdateConfigsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnbondedEvmosUpdateConfigs)
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
		it.Event = new(UnbondedEvmosUpdateConfigs)
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
func (it *UnbondedEvmosUpdateConfigsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnbondedEvmosUpdateConfigsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnbondedEvmosUpdateConfigs represents a UpdateConfigs event raised by the UnbondedEvmos contract.
type UnbondedEvmosUpdateConfigs struct {
	UnbondingInterval *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateConfigs is a free log retrieval operation binding the contract event 0x78a3671679b68721aaad9eb74535be0be119bd34c0efa671eb6ab3210d1fe257.
//
// Solidity: event UpdateConfigs(uint256 unbondingInterval)
func (_UnbondedEvmos *UnbondedEvmosFilterer) FilterUpdateConfigs(opts *bind.FilterOpts) (*UnbondedEvmosUpdateConfigsIterator, error) {

	logs, sub, err := _UnbondedEvmos.contract.FilterLogs(opts, "UpdateConfigs")
	if err != nil {
		return nil, err
	}
	return &UnbondedEvmosUpdateConfigsIterator{contract: _UnbondedEvmos.contract, event: "UpdateConfigs", logs: logs, sub: sub}, nil
}

// WatchUpdateConfigs is a free log subscription operation binding the contract event 0x78a3671679b68721aaad9eb74535be0be119bd34c0efa671eb6ab3210d1fe257.
//
// Solidity: event UpdateConfigs(uint256 unbondingInterval)
func (_UnbondedEvmos *UnbondedEvmosFilterer) WatchUpdateConfigs(opts *bind.WatchOpts, sink chan<- *UnbondedEvmosUpdateConfigs) (event.Subscription, error) {

	logs, sub, err := _UnbondedEvmos.contract.WatchLogs(opts, "UpdateConfigs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnbondedEvmosUpdateConfigs)
				if err := _UnbondedEvmos.contract.UnpackLog(event, "UpdateConfigs", log); err != nil {
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

// ParseUpdateConfigs is a log parse operation binding the contract event 0x78a3671679b68721aaad9eb74535be0be119bd34c0efa671eb6ab3210d1fe257.
//
// Solidity: event UpdateConfigs(uint256 unbondingInterval)
func (_UnbondedEvmos *UnbondedEvmosFilterer) ParseUpdateConfigs(log types.Log) (*UnbondedEvmosUpdateConfigs, error) {
	event := new(UnbondedEvmosUpdateConfigs)
	if err := _UnbondedEvmos.contract.UnpackLog(event, "UpdateConfigs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnbondedEvmosUpdateMinterStatusIterator is returned from FilterUpdateMinterStatus and is used to iterate over the raw logs and unpacked data for UpdateMinterStatus events raised by the UnbondedEvmos contract.
type UnbondedEvmosUpdateMinterStatusIterator struct {
	Event *UnbondedEvmosUpdateMinterStatus // Event containing the contract specifics and raw log

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
func (it *UnbondedEvmosUpdateMinterStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnbondedEvmosUpdateMinterStatus)
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
		it.Event = new(UnbondedEvmosUpdateMinterStatus)
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
func (it *UnbondedEvmosUpdateMinterStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnbondedEvmosUpdateMinterStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnbondedEvmosUpdateMinterStatus represents a UpdateMinterStatus event raised by the UnbondedEvmos contract.
type UnbondedEvmosUpdateMinterStatus struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinterStatus is a free log retrieval operation binding the contract event 0xf6c10b2804d0174249f23bd4c462fb3e65ea16bf2b6896effe8a09f8f6a8e3b0.
//
// Solidity: event UpdateMinterStatus(address account, bool status)
func (_UnbondedEvmos *UnbondedEvmosFilterer) FilterUpdateMinterStatus(opts *bind.FilterOpts) (*UnbondedEvmosUpdateMinterStatusIterator, error) {

	logs, sub, err := _UnbondedEvmos.contract.FilterLogs(opts, "UpdateMinterStatus")
	if err != nil {
		return nil, err
	}
	return &UnbondedEvmosUpdateMinterStatusIterator{contract: _UnbondedEvmos.contract, event: "UpdateMinterStatus", logs: logs, sub: sub}, nil
}

// WatchUpdateMinterStatus is a free log subscription operation binding the contract event 0xf6c10b2804d0174249f23bd4c462fb3e65ea16bf2b6896effe8a09f8f6a8e3b0.
//
// Solidity: event UpdateMinterStatus(address account, bool status)
func (_UnbondedEvmos *UnbondedEvmosFilterer) WatchUpdateMinterStatus(opts *bind.WatchOpts, sink chan<- *UnbondedEvmosUpdateMinterStatus) (event.Subscription, error) {

	logs, sub, err := _UnbondedEvmos.contract.WatchLogs(opts, "UpdateMinterStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnbondedEvmosUpdateMinterStatus)
				if err := _UnbondedEvmos.contract.UnpackLog(event, "UpdateMinterStatus", log); err != nil {
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

// ParseUpdateMinterStatus is a log parse operation binding the contract event 0xf6c10b2804d0174249f23bd4c462fb3e65ea16bf2b6896effe8a09f8f6a8e3b0.
//
// Solidity: event UpdateMinterStatus(address account, bool status)
func (_UnbondedEvmos *UnbondedEvmosFilterer) ParseUpdateMinterStatus(log types.Log) (*UnbondedEvmosUpdateMinterStatus, error) {
	event := new(UnbondedEvmosUpdateMinterStatus)
	if err := _UnbondedEvmos.contract.UnpackLog(event, "UpdateMinterStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnbondedEvmosWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the UnbondedEvmos contract.
type UnbondedEvmosWithdrawIterator struct {
	Event *UnbondedEvmosWithdraw // Event containing the contract specifics and raw log

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
func (it *UnbondedEvmosWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnbondedEvmosWithdraw)
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
		it.Event = new(UnbondedEvmosWithdraw)
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
func (it *UnbondedEvmosWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnbondedEvmosWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnbondedEvmosWithdraw represents a Withdraw event raised by the UnbondedEvmos contract.
type UnbondedEvmosWithdraw struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address account, uint256 amount)
func (_UnbondedEvmos *UnbondedEvmosFilterer) FilterWithdraw(opts *bind.FilterOpts) (*UnbondedEvmosWithdrawIterator, error) {

	logs, sub, err := _UnbondedEvmos.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &UnbondedEvmosWithdrawIterator{contract: _UnbondedEvmos.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address account, uint256 amount)
func (_UnbondedEvmos *UnbondedEvmosFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *UnbondedEvmosWithdraw) (event.Subscription, error) {

	logs, sub, err := _UnbondedEvmos.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnbondedEvmosWithdraw)
				if err := _UnbondedEvmos.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address account, uint256 amount)
func (_UnbondedEvmos *UnbondedEvmosFilterer) ParseWithdraw(log types.Log) (*UnbondedEvmosWithdraw, error) {
	event := new(UnbondedEvmosWithdraw)
	if err := _UnbondedEvmos.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
