// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stayking

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

// StaykingMetaData contains all meta data concerning the Stayking contract.
var StaykingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"distributed\",\"type\":\"uint256\"}],\"name\":\"Accrue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"equity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtInBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"AddPosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"AddVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"ChangeDelegator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"killer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"equity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtInBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"Kill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"}],\"name\":\"PositionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"equity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtInBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"RemovePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"Unstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDebtInBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"killFactorBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservedBps\",\"type\":\"uint256\"}],\"name\":\"UpdateConfigs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"UpdateVault\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uEVMOS_\",\"type\":\"address\"}],\"name\":\"__Stayking_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"}],\"name\":\"accrue\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"debtToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"extraDebtInBase\",\"type\":\"uint256\"}],\"name\":\"addDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"debtToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"extraEquity\",\"type\":\"uint256\"}],\"name\":\"addEquity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"debtToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"equity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtInBase\",\"type\":\"uint256\"}],\"name\":\"addPosition\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"amountToShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegator\",\"type\":\"address\"}],\"name\":\"changeDelegator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"debtToken\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"equityInBaseChanged\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"debtInBaseChanged\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"repaidDebt\",\"type\":\"uint256\"}],\"name\":\"changePosition\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"debtAmountOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"}],\"name\":\"getAccruedValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"debtToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"}],\"name\":\"isKillable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"healthy\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"debtToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"}],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"killFactorBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDebtInBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"positionIdOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtToken\",\"type\":\"address\"}],\"name\":\"positionInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"positionValueInBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtInBase\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"positionsLengthOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"debtToken\",\"type\":\"address\"}],\"name\":\"removePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"debtToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repaidDebt\",\"type\":\"uint256\"}],\"name\":\"repayDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"debtToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minRepaid\",\"type\":\"uint256\"}],\"name\":\"repayDebtInBase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservedBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservedPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"shareToAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenToVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalDebtOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uEVMOS\",\"outputs\":[{\"internalType\":\"contractIUnbondedEvmos\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDebtInBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_killFactorBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reservedBps\",\"type\":\"uint256\"}],\"name\":\"updateConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"updateVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedKiller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StaykingABI is the input ABI used to generate the binding from.
// Deprecated: Use StaykingMetaData.ABI instead.
var StaykingABI = StaykingMetaData.ABI

// Stayking is an auto generated Go binding around an Ethereum contract.
type Stayking struct {
	StaykingCaller     // Read-only binding to the contract
	StaykingTransactor // Write-only binding to the contract
	StaykingFilterer   // Log filterer for contract events
}

// StaykingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaykingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaykingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaykingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaykingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaykingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaykingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaykingSession struct {
	Contract     *Stayking         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StaykingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaykingCallerSession struct {
	Contract *StaykingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StaykingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaykingTransactorSession struct {
	Contract     *StaykingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StaykingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaykingRaw struct {
	Contract *Stayking // Generic contract binding to access the raw methods on
}

// StaykingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaykingCallerRaw struct {
	Contract *StaykingCaller // Generic read-only contract binding to access the raw methods on
}

// StaykingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaykingTransactorRaw struct {
	Contract *StaykingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStayking creates a new instance of Stayking, bound to a specific deployed contract.
func NewStayking(address common.Address, backend bind.ContractBackend) (*Stayking, error) {
	contract, err := bindStayking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stayking{StaykingCaller: StaykingCaller{contract: contract}, StaykingTransactor: StaykingTransactor{contract: contract}, StaykingFilterer: StaykingFilterer{contract: contract}}, nil
}

// NewStaykingCaller creates a new read-only instance of Stayking, bound to a specific deployed contract.
func NewStaykingCaller(address common.Address, caller bind.ContractCaller) (*StaykingCaller, error) {
	contract, err := bindStayking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaykingCaller{contract: contract}, nil
}

// NewStaykingTransactor creates a new write-only instance of Stayking, bound to a specific deployed contract.
func NewStaykingTransactor(address common.Address, transactor bind.ContractTransactor) (*StaykingTransactor, error) {
	contract, err := bindStayking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaykingTransactor{contract: contract}, nil
}

// NewStaykingFilterer creates a new log filterer instance of Stayking, bound to a specific deployed contract.
func NewStaykingFilterer(address common.Address, filterer bind.ContractFilterer) (*StaykingFilterer, error) {
	contract, err := bindStayking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaykingFilterer{contract: contract}, nil
}

// bindStayking binds a generic wrapper to an already deployed contract.
func bindStayking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StaykingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stayking *StaykingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stayking.Contract.StaykingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stayking *StaykingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stayking.Contract.StaykingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stayking *StaykingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stayking.Contract.StaykingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stayking *StaykingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stayking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stayking *StaykingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stayking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stayking *StaykingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stayking.Contract.contract.Transact(opts, method, params...)
}

// AmountToShare is a free data retrieval call binding the contract method 0x2d4047e5.
//
// Solidity: function amountToShare(uint256 amount) view returns(uint256)
func (_Stayking *StaykingCaller) AmountToShare(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "amountToShare", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountToShare is a free data retrieval call binding the contract method 0x2d4047e5.
//
// Solidity: function amountToShare(uint256 amount) view returns(uint256)
func (_Stayking *StaykingSession) AmountToShare(amount *big.Int) (*big.Int, error) {
	return _Stayking.Contract.AmountToShare(&_Stayking.CallOpts, amount)
}

// AmountToShare is a free data retrieval call binding the contract method 0x2d4047e5.
//
// Solidity: function amountToShare(uint256 amount) view returns(uint256)
func (_Stayking *StaykingCallerSession) AmountToShare(amount *big.Int) (*big.Int, error) {
	return _Stayking.Contract.AmountToShare(&_Stayking.CallOpts, amount)
}

// DebtAmountOf is a free data retrieval call binding the contract method 0x27dfaf92.
//
// Solidity: function debtAmountOf(address user, address vault) view returns(uint256)
func (_Stayking *StaykingCaller) DebtAmountOf(opts *bind.CallOpts, user common.Address, vault common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "debtAmountOf", user, vault)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtAmountOf is a free data retrieval call binding the contract method 0x27dfaf92.
//
// Solidity: function debtAmountOf(address user, address vault) view returns(uint256)
func (_Stayking *StaykingSession) DebtAmountOf(user common.Address, vault common.Address) (*big.Int, error) {
	return _Stayking.Contract.DebtAmountOf(&_Stayking.CallOpts, user, vault)
}

// DebtAmountOf is a free data retrieval call binding the contract method 0x27dfaf92.
//
// Solidity: function debtAmountOf(address user, address vault) view returns(uint256)
func (_Stayking *StaykingCallerSession) DebtAmountOf(user common.Address, vault common.Address) (*big.Int, error) {
	return _Stayking.Contract.DebtAmountOf(&_Stayking.CallOpts, user, vault)
}

// Delegator is a free data retrieval call binding the contract method 0xce9b7930.
//
// Solidity: function delegator() view returns(address)
func (_Stayking *StaykingCaller) Delegator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "delegator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegator is a free data retrieval call binding the contract method 0xce9b7930.
//
// Solidity: function delegator() view returns(address)
func (_Stayking *StaykingSession) Delegator() (common.Address, error) {
	return _Stayking.Contract.Delegator(&_Stayking.CallOpts)
}

// Delegator is a free data retrieval call binding the contract method 0xce9b7930.
//
// Solidity: function delegator() view returns(address)
func (_Stayking *StaykingCallerSession) Delegator() (common.Address, error) {
	return _Stayking.Contract.Delegator(&_Stayking.CallOpts)
}

// GetAccruedValue is a free data retrieval call binding the contract method 0xe4175aee.
//
// Solidity: function getAccruedValue(uint256 totalStaked) view returns(uint256)
func (_Stayking *StaykingCaller) GetAccruedValue(opts *bind.CallOpts, totalStaked *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "getAccruedValue", totalStaked)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccruedValue is a free data retrieval call binding the contract method 0xe4175aee.
//
// Solidity: function getAccruedValue(uint256 totalStaked) view returns(uint256)
func (_Stayking *StaykingSession) GetAccruedValue(totalStaked *big.Int) (*big.Int, error) {
	return _Stayking.Contract.GetAccruedValue(&_Stayking.CallOpts, totalStaked)
}

// GetAccruedValue is a free data retrieval call binding the contract method 0xe4175aee.
//
// Solidity: function getAccruedValue(uint256 totalStaked) view returns(uint256)
func (_Stayking *StaykingCallerSession) GetAccruedValue(totalStaked *big.Int) (*big.Int, error) {
	return _Stayking.Contract.GetAccruedValue(&_Stayking.CallOpts, totalStaked)
}

// IsKillable is a free data retrieval call binding the contract method 0x5ff02531.
//
// Solidity: function isKillable(address debtToken, uint256 positionId) view returns(bool healthy)
func (_Stayking *StaykingCaller) IsKillable(opts *bind.CallOpts, debtToken common.Address, positionId *big.Int) (bool, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "isKillable", debtToken, positionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKillable is a free data retrieval call binding the contract method 0x5ff02531.
//
// Solidity: function isKillable(address debtToken, uint256 positionId) view returns(bool healthy)
func (_Stayking *StaykingSession) IsKillable(debtToken common.Address, positionId *big.Int) (bool, error) {
	return _Stayking.Contract.IsKillable(&_Stayking.CallOpts, debtToken, positionId)
}

// IsKillable is a free data retrieval call binding the contract method 0x5ff02531.
//
// Solidity: function isKillable(address debtToken, uint256 positionId) view returns(bool healthy)
func (_Stayking *StaykingCallerSession) IsKillable(debtToken common.Address, positionId *big.Int) (bool, error) {
	return _Stayking.Contract.IsKillable(&_Stayking.CallOpts, debtToken, positionId)
}

// KillFactorBps is a free data retrieval call binding the contract method 0x28f55a6c.
//
// Solidity: function killFactorBps() view returns(uint256)
func (_Stayking *StaykingCaller) KillFactorBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "killFactorBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KillFactorBps is a free data retrieval call binding the contract method 0x28f55a6c.
//
// Solidity: function killFactorBps() view returns(uint256)
func (_Stayking *StaykingSession) KillFactorBps() (*big.Int, error) {
	return _Stayking.Contract.KillFactorBps(&_Stayking.CallOpts)
}

// KillFactorBps is a free data retrieval call binding the contract method 0x28f55a6c.
//
// Solidity: function killFactorBps() view returns(uint256)
func (_Stayking *StaykingCallerSession) KillFactorBps() (*big.Int, error) {
	return _Stayking.Contract.KillFactorBps(&_Stayking.CallOpts)
}

// MinDebtInBase is a free data retrieval call binding the contract method 0x178483a7.
//
// Solidity: function minDebtInBase() view returns(uint256)
func (_Stayking *StaykingCaller) MinDebtInBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "minDebtInBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDebtInBase is a free data retrieval call binding the contract method 0x178483a7.
//
// Solidity: function minDebtInBase() view returns(uint256)
func (_Stayking *StaykingSession) MinDebtInBase() (*big.Int, error) {
	return _Stayking.Contract.MinDebtInBase(&_Stayking.CallOpts)
}

// MinDebtInBase is a free data retrieval call binding the contract method 0x178483a7.
//
// Solidity: function minDebtInBase() view returns(uint256)
func (_Stayking *StaykingCallerSession) MinDebtInBase() (*big.Int, error) {
	return _Stayking.Contract.MinDebtInBase(&_Stayking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stayking *StaykingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stayking *StaykingSession) Owner() (common.Address, error) {
	return _Stayking.Contract.Owner(&_Stayking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stayking *StaykingCallerSession) Owner() (common.Address, error) {
	return _Stayking.Contract.Owner(&_Stayking.CallOpts)
}

// PositionIdOf is a free data retrieval call binding the contract method 0xbb94d3db.
//
// Solidity: function positionIdOf(address , address ) view returns(uint256)
func (_Stayking *StaykingCaller) PositionIdOf(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "positionIdOf", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PositionIdOf is a free data retrieval call binding the contract method 0xbb94d3db.
//
// Solidity: function positionIdOf(address , address ) view returns(uint256)
func (_Stayking *StaykingSession) PositionIdOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stayking.Contract.PositionIdOf(&_Stayking.CallOpts, arg0, arg1)
}

// PositionIdOf is a free data retrieval call binding the contract method 0xbb94d3db.
//
// Solidity: function positionIdOf(address , address ) view returns(uint256)
func (_Stayking *StaykingCallerSession) PositionIdOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stayking.Contract.PositionIdOf(&_Stayking.CallOpts, arg0, arg1)
}

// PositionInfo is a free data retrieval call binding the contract method 0x5066f3b5.
//
// Solidity: function positionInfo(address user, address debtToken) view returns(uint256 positionValueInBase, uint256 debtInBase)
func (_Stayking *StaykingCaller) PositionInfo(opts *bind.CallOpts, user common.Address, debtToken common.Address) (struct {
	PositionValueInBase *big.Int
	DebtInBase          *big.Int
}, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "positionInfo", user, debtToken)

	outstruct := new(struct {
		PositionValueInBase *big.Int
		DebtInBase          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PositionValueInBase = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DebtInBase = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PositionInfo is a free data retrieval call binding the contract method 0x5066f3b5.
//
// Solidity: function positionInfo(address user, address debtToken) view returns(uint256 positionValueInBase, uint256 debtInBase)
func (_Stayking *StaykingSession) PositionInfo(user common.Address, debtToken common.Address) (struct {
	PositionValueInBase *big.Int
	DebtInBase          *big.Int
}, error) {
	return _Stayking.Contract.PositionInfo(&_Stayking.CallOpts, user, debtToken)
}

// PositionInfo is a free data retrieval call binding the contract method 0x5066f3b5.
//
// Solidity: function positionInfo(address user, address debtToken) view returns(uint256 positionValueInBase, uint256 debtInBase)
func (_Stayking *StaykingCallerSession) PositionInfo(user common.Address, debtToken common.Address) (struct {
	PositionValueInBase *big.Int
	DebtInBase          *big.Int
}, error) {
	return _Stayking.Contract.PositionInfo(&_Stayking.CallOpts, user, debtToken)
}

// Positions is a free data retrieval call binding the contract method 0xc1be6677.
//
// Solidity: function positions(address , uint256 ) view returns(address user, uint256 share)
func (_Stayking *StaykingCaller) Positions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	User  common.Address
	Share *big.Int
}, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "positions", arg0, arg1)

	outstruct := new(struct {
		User  common.Address
		Share *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Share = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0xc1be6677.
//
// Solidity: function positions(address , uint256 ) view returns(address user, uint256 share)
func (_Stayking *StaykingSession) Positions(arg0 common.Address, arg1 *big.Int) (struct {
	User  common.Address
	Share *big.Int
}, error) {
	return _Stayking.Contract.Positions(&_Stayking.CallOpts, arg0, arg1)
}

// Positions is a free data retrieval call binding the contract method 0xc1be6677.
//
// Solidity: function positions(address , uint256 ) view returns(address user, uint256 share)
func (_Stayking *StaykingCallerSession) Positions(arg0 common.Address, arg1 *big.Int) (struct {
	User  common.Address
	Share *big.Int
}, error) {
	return _Stayking.Contract.Positions(&_Stayking.CallOpts, arg0, arg1)
}

// PositionsLengthOf is a free data retrieval call binding the contract method 0x51dabb75.
//
// Solidity: function positionsLengthOf(address ) view returns(uint256)
func (_Stayking *StaykingCaller) PositionsLengthOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "positionsLengthOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PositionsLengthOf is a free data retrieval call binding the contract method 0x51dabb75.
//
// Solidity: function positionsLengthOf(address ) view returns(uint256)
func (_Stayking *StaykingSession) PositionsLengthOf(arg0 common.Address) (*big.Int, error) {
	return _Stayking.Contract.PositionsLengthOf(&_Stayking.CallOpts, arg0)
}

// PositionsLengthOf is a free data retrieval call binding the contract method 0x51dabb75.
//
// Solidity: function positionsLengthOf(address ) view returns(uint256)
func (_Stayking *StaykingCallerSession) PositionsLengthOf(arg0 common.Address) (*big.Int, error) {
	return _Stayking.Contract.PositionsLengthOf(&_Stayking.CallOpts, arg0)
}

// ReservedBps is a free data retrieval call binding the contract method 0x52e83e52.
//
// Solidity: function reservedBps() view returns(uint256)
func (_Stayking *StaykingCaller) ReservedBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "reservedBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedBps is a free data retrieval call binding the contract method 0x52e83e52.
//
// Solidity: function reservedBps() view returns(uint256)
func (_Stayking *StaykingSession) ReservedBps() (*big.Int, error) {
	return _Stayking.Contract.ReservedBps(&_Stayking.CallOpts)
}

// ReservedBps is a free data retrieval call binding the contract method 0x52e83e52.
//
// Solidity: function reservedBps() view returns(uint256)
func (_Stayking *StaykingCallerSession) ReservedBps() (*big.Int, error) {
	return _Stayking.Contract.ReservedBps(&_Stayking.CallOpts)
}

// ReservedPool is a free data retrieval call binding the contract method 0xf0e63edf.
//
// Solidity: function reservedPool() view returns(uint256)
func (_Stayking *StaykingCaller) ReservedPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "reservedPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedPool is a free data retrieval call binding the contract method 0xf0e63edf.
//
// Solidity: function reservedPool() view returns(uint256)
func (_Stayking *StaykingSession) ReservedPool() (*big.Int, error) {
	return _Stayking.Contract.ReservedPool(&_Stayking.CallOpts)
}

// ReservedPool is a free data retrieval call binding the contract method 0xf0e63edf.
//
// Solidity: function reservedPool() view returns(uint256)
func (_Stayking *StaykingCallerSession) ReservedPool() (*big.Int, error) {
	return _Stayking.Contract.ReservedPool(&_Stayking.CallOpts)
}

// ShareToAmount is a free data retrieval call binding the contract method 0x2d6f8013.
//
// Solidity: function shareToAmount(uint256 share) view returns(uint256)
func (_Stayking *StaykingCaller) ShareToAmount(opts *bind.CallOpts, share *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "shareToAmount", share)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShareToAmount is a free data retrieval call binding the contract method 0x2d6f8013.
//
// Solidity: function shareToAmount(uint256 share) view returns(uint256)
func (_Stayking *StaykingSession) ShareToAmount(share *big.Int) (*big.Int, error) {
	return _Stayking.Contract.ShareToAmount(&_Stayking.CallOpts, share)
}

// ShareToAmount is a free data retrieval call binding the contract method 0x2d6f8013.
//
// Solidity: function shareToAmount(uint256 share) view returns(uint256)
func (_Stayking *StaykingCallerSession) ShareToAmount(share *big.Int) (*big.Int, error) {
	return _Stayking.Contract.ShareToAmount(&_Stayking.CallOpts, share)
}

// TokenToVault is a free data retrieval call binding the contract method 0x0c7e1725.
//
// Solidity: function tokenToVault(address ) view returns(address)
func (_Stayking *StaykingCaller) TokenToVault(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "tokenToVault", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenToVault is a free data retrieval call binding the contract method 0x0c7e1725.
//
// Solidity: function tokenToVault(address ) view returns(address)
func (_Stayking *StaykingSession) TokenToVault(arg0 common.Address) (common.Address, error) {
	return _Stayking.Contract.TokenToVault(&_Stayking.CallOpts, arg0)
}

// TokenToVault is a free data retrieval call binding the contract method 0x0c7e1725.
//
// Solidity: function tokenToVault(address ) view returns(address)
func (_Stayking *StaykingCallerSession) TokenToVault(arg0 common.Address) (common.Address, error) {
	return _Stayking.Contract.TokenToVault(&_Stayking.CallOpts, arg0)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_Stayking *StaykingCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_Stayking *StaykingSession) TotalAmount() (*big.Int, error) {
	return _Stayking.Contract.TotalAmount(&_Stayking.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_Stayking *StaykingCallerSession) TotalAmount() (*big.Int, error) {
	return _Stayking.Contract.TotalAmount(&_Stayking.CallOpts)
}

// TotalDebtOf is a free data retrieval call binding the contract method 0x9f2b2833.
//
// Solidity: function totalDebtOf(address ) view returns(uint256)
func (_Stayking *StaykingCaller) TotalDebtOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "totalDebtOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebtOf is a free data retrieval call binding the contract method 0x9f2b2833.
//
// Solidity: function totalDebtOf(address ) view returns(uint256)
func (_Stayking *StaykingSession) TotalDebtOf(arg0 common.Address) (*big.Int, error) {
	return _Stayking.Contract.TotalDebtOf(&_Stayking.CallOpts, arg0)
}

// TotalDebtOf is a free data retrieval call binding the contract method 0x9f2b2833.
//
// Solidity: function totalDebtOf(address ) view returns(uint256)
func (_Stayking *StaykingCallerSession) TotalDebtOf(arg0 common.Address) (*big.Int, error) {
	return _Stayking.Contract.TotalDebtOf(&_Stayking.CallOpts, arg0)
}

// TotalShare is a free data retrieval call binding the contract method 0x026c4207.
//
// Solidity: function totalShare() view returns(uint256)
func (_Stayking *StaykingCaller) TotalShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "totalShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShare is a free data retrieval call binding the contract method 0x026c4207.
//
// Solidity: function totalShare() view returns(uint256)
func (_Stayking *StaykingSession) TotalShare() (*big.Int, error) {
	return _Stayking.Contract.TotalShare(&_Stayking.CallOpts)
}

// TotalShare is a free data retrieval call binding the contract method 0x026c4207.
//
// Solidity: function totalShare() view returns(uint256)
func (_Stayking *StaykingCallerSession) TotalShare() (*big.Int, error) {
	return _Stayking.Contract.TotalShare(&_Stayking.CallOpts)
}

// UEVMOS is a free data retrieval call binding the contract method 0xafd06b34.
//
// Solidity: function uEVMOS() view returns(address)
func (_Stayking *StaykingCaller) UEVMOS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "uEVMOS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UEVMOS is a free data retrieval call binding the contract method 0xafd06b34.
//
// Solidity: function uEVMOS() view returns(address)
func (_Stayking *StaykingSession) UEVMOS() (common.Address, error) {
	return _Stayking.Contract.UEVMOS(&_Stayking.CallOpts)
}

// UEVMOS is a free data retrieval call binding the contract method 0xafd06b34.
//
// Solidity: function uEVMOS() view returns(address)
func (_Stayking *StaykingCallerSession) UEVMOS() (common.Address, error) {
	return _Stayking.Contract.UEVMOS(&_Stayking.CallOpts)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address)
func (_Stayking *StaykingCaller) Vaults(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "vaults", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address)
func (_Stayking *StaykingSession) Vaults(arg0 *big.Int) (common.Address, error) {
	return _Stayking.Contract.Vaults(&_Stayking.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address)
func (_Stayking *StaykingCallerSession) Vaults(arg0 *big.Int) (common.Address, error) {
	return _Stayking.Contract.Vaults(&_Stayking.CallOpts, arg0)
}

// WhitelistedKiller is a free data retrieval call binding the contract method 0xa3381d28.
//
// Solidity: function whitelistedKiller(address ) view returns(bool)
func (_Stayking *StaykingCaller) WhitelistedKiller(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Stayking.contract.Call(opts, &out, "whitelistedKiller", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedKiller is a free data retrieval call binding the contract method 0xa3381d28.
//
// Solidity: function whitelistedKiller(address ) view returns(bool)
func (_Stayking *StaykingSession) WhitelistedKiller(arg0 common.Address) (bool, error) {
	return _Stayking.Contract.WhitelistedKiller(&_Stayking.CallOpts, arg0)
}

// WhitelistedKiller is a free data retrieval call binding the contract method 0xa3381d28.
//
// Solidity: function whitelistedKiller(address ) view returns(bool)
func (_Stayking *StaykingCallerSession) WhitelistedKiller(arg0 common.Address) (bool, error) {
	return _Stayking.Contract.WhitelistedKiller(&_Stayking.CallOpts, arg0)
}

// StaykingInit is a paid mutator transaction binding the contract method 0xb2a2a96b.
//
// Solidity: function __Stayking_init(address delegator_, address uEVMOS_) returns()
func (_Stayking *StaykingTransactor) StaykingInit(opts *bind.TransactOpts, delegator_ common.Address, uEVMOS_ common.Address) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "__Stayking_init", delegator_, uEVMOS_)
}

// StaykingInit is a paid mutator transaction binding the contract method 0xb2a2a96b.
//
// Solidity: function __Stayking_init(address delegator_, address uEVMOS_) returns()
func (_Stayking *StaykingSession) StaykingInit(delegator_ common.Address, uEVMOS_ common.Address) (*types.Transaction, error) {
	return _Stayking.Contract.StaykingInit(&_Stayking.TransactOpts, delegator_, uEVMOS_)
}

// StaykingInit is a paid mutator transaction binding the contract method 0xb2a2a96b.
//
// Solidity: function __Stayking_init(address delegator_, address uEVMOS_) returns()
func (_Stayking *StaykingTransactorSession) StaykingInit(delegator_ common.Address, uEVMOS_ common.Address) (*types.Transaction, error) {
	return _Stayking.Contract.StaykingInit(&_Stayking.TransactOpts, delegator_, uEVMOS_)
}

// Accrue is a paid mutator transaction binding the contract method 0x744f4cf6.
//
// Solidity: function accrue(uint256 totalStaked) payable returns()
func (_Stayking *StaykingTransactor) Accrue(opts *bind.TransactOpts, totalStaked *big.Int) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "accrue", totalStaked)
}

// Accrue is a paid mutator transaction binding the contract method 0x744f4cf6.
//
// Solidity: function accrue(uint256 totalStaked) payable returns()
func (_Stayking *StaykingSession) Accrue(totalStaked *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.Accrue(&_Stayking.TransactOpts, totalStaked)
}

// Accrue is a paid mutator transaction binding the contract method 0x744f4cf6.
//
// Solidity: function accrue(uint256 totalStaked) payable returns()
func (_Stayking *StaykingTransactorSession) Accrue(totalStaked *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.Accrue(&_Stayking.TransactOpts, totalStaked)
}

// AddDebt is a paid mutator transaction binding the contract method 0x34c389ef.
//
// Solidity: function addDebt(address debtToken, uint256 extraDebtInBase) returns()
func (_Stayking *StaykingTransactor) AddDebt(opts *bind.TransactOpts, debtToken common.Address, extraDebtInBase *big.Int) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "addDebt", debtToken, extraDebtInBase)
}

// AddDebt is a paid mutator transaction binding the contract method 0x34c389ef.
//
// Solidity: function addDebt(address debtToken, uint256 extraDebtInBase) returns()
func (_Stayking *StaykingSession) AddDebt(debtToken common.Address, extraDebtInBase *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.AddDebt(&_Stayking.TransactOpts, debtToken, extraDebtInBase)
}

// AddDebt is a paid mutator transaction binding the contract method 0x34c389ef.
//
// Solidity: function addDebt(address debtToken, uint256 extraDebtInBase) returns()
func (_Stayking *StaykingTransactorSession) AddDebt(debtToken common.Address, extraDebtInBase *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.AddDebt(&_Stayking.TransactOpts, debtToken, extraDebtInBase)
}

// AddEquity is a paid mutator transaction binding the contract method 0xef6b2479.
//
// Solidity: function addEquity(address debtToken, uint256 extraEquity) payable returns()
func (_Stayking *StaykingTransactor) AddEquity(opts *bind.TransactOpts, debtToken common.Address, extraEquity *big.Int) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "addEquity", debtToken, extraEquity)
}

// AddEquity is a paid mutator transaction binding the contract method 0xef6b2479.
//
// Solidity: function addEquity(address debtToken, uint256 extraEquity) payable returns()
func (_Stayking *StaykingSession) AddEquity(debtToken common.Address, extraEquity *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.AddEquity(&_Stayking.TransactOpts, debtToken, extraEquity)
}

// AddEquity is a paid mutator transaction binding the contract method 0xef6b2479.
//
// Solidity: function addEquity(address debtToken, uint256 extraEquity) payable returns()
func (_Stayking *StaykingTransactorSession) AddEquity(debtToken common.Address, extraEquity *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.AddEquity(&_Stayking.TransactOpts, debtToken, extraEquity)
}

// AddPosition is a paid mutator transaction binding the contract method 0xcff07fd1.
//
// Solidity: function addPosition(address debtToken, uint256 equity, uint256 debtInBase) payable returns()
func (_Stayking *StaykingTransactor) AddPosition(opts *bind.TransactOpts, debtToken common.Address, equity *big.Int, debtInBase *big.Int) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "addPosition", debtToken, equity, debtInBase)
}

// AddPosition is a paid mutator transaction binding the contract method 0xcff07fd1.
//
// Solidity: function addPosition(address debtToken, uint256 equity, uint256 debtInBase) payable returns()
func (_Stayking *StaykingSession) AddPosition(debtToken common.Address, equity *big.Int, debtInBase *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.AddPosition(&_Stayking.TransactOpts, debtToken, equity, debtInBase)
}

// AddPosition is a paid mutator transaction binding the contract method 0xcff07fd1.
//
// Solidity: function addPosition(address debtToken, uint256 equity, uint256 debtInBase) payable returns()
func (_Stayking *StaykingTransactorSession) AddPosition(debtToken common.Address, equity *big.Int, debtInBase *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.AddPosition(&_Stayking.TransactOpts, debtToken, equity, debtInBase)
}

// ChangeDelegator is a paid mutator transaction binding the contract method 0xb1ae5ed3.
//
// Solidity: function changeDelegator(address _delegator) returns()
func (_Stayking *StaykingTransactor) ChangeDelegator(opts *bind.TransactOpts, _delegator common.Address) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "changeDelegator", _delegator)
}

// ChangeDelegator is a paid mutator transaction binding the contract method 0xb1ae5ed3.
//
// Solidity: function changeDelegator(address _delegator) returns()
func (_Stayking *StaykingSession) ChangeDelegator(_delegator common.Address) (*types.Transaction, error) {
	return _Stayking.Contract.ChangeDelegator(&_Stayking.TransactOpts, _delegator)
}

// ChangeDelegator is a paid mutator transaction binding the contract method 0xb1ae5ed3.
//
// Solidity: function changeDelegator(address _delegator) returns()
func (_Stayking *StaykingTransactorSession) ChangeDelegator(_delegator common.Address) (*types.Transaction, error) {
	return _Stayking.Contract.ChangeDelegator(&_Stayking.TransactOpts, _delegator)
}

// ChangePosition is a paid mutator transaction binding the contract method 0xbed06b0b.
//
// Solidity: function changePosition(address debtToken, int256 equityInBaseChanged, int256 debtInBaseChanged, uint256 repaidDebt) payable returns()
func (_Stayking *StaykingTransactor) ChangePosition(opts *bind.TransactOpts, debtToken common.Address, equityInBaseChanged *big.Int, debtInBaseChanged *big.Int, repaidDebt *big.Int) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "changePosition", debtToken, equityInBaseChanged, debtInBaseChanged, repaidDebt)
}

// ChangePosition is a paid mutator transaction binding the contract method 0xbed06b0b.
//
// Solidity: function changePosition(address debtToken, int256 equityInBaseChanged, int256 debtInBaseChanged, uint256 repaidDebt) payable returns()
func (_Stayking *StaykingSession) ChangePosition(debtToken common.Address, equityInBaseChanged *big.Int, debtInBaseChanged *big.Int, repaidDebt *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.ChangePosition(&_Stayking.TransactOpts, debtToken, equityInBaseChanged, debtInBaseChanged, repaidDebt)
}

// ChangePosition is a paid mutator transaction binding the contract method 0xbed06b0b.
//
// Solidity: function changePosition(address debtToken, int256 equityInBaseChanged, int256 debtInBaseChanged, uint256 repaidDebt) payable returns()
func (_Stayking *StaykingTransactorSession) ChangePosition(debtToken common.Address, equityInBaseChanged *big.Int, debtInBaseChanged *big.Int, repaidDebt *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.ChangePosition(&_Stayking.TransactOpts, debtToken, equityInBaseChanged, debtInBaseChanged, repaidDebt)
}

// Kill is a paid mutator transaction binding the contract method 0x92707802.
//
// Solidity: function kill(address debtToken, uint256 positionId) returns()
func (_Stayking *StaykingTransactor) Kill(opts *bind.TransactOpts, debtToken common.Address, positionId *big.Int) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "kill", debtToken, positionId)
}

// Kill is a paid mutator transaction binding the contract method 0x92707802.
//
// Solidity: function kill(address debtToken, uint256 positionId) returns()
func (_Stayking *StaykingSession) Kill(debtToken common.Address, positionId *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.Kill(&_Stayking.TransactOpts, debtToken, positionId)
}

// Kill is a paid mutator transaction binding the contract method 0x92707802.
//
// Solidity: function kill(address debtToken, uint256 positionId) returns()
func (_Stayking *StaykingTransactorSession) Kill(debtToken common.Address, positionId *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.Kill(&_Stayking.TransactOpts, debtToken, positionId)
}

// RemovePosition is a paid mutator transaction binding the contract method 0xdf5ca1f8.
//
// Solidity: function removePosition(address debtToken) returns()
func (_Stayking *StaykingTransactor) RemovePosition(opts *bind.TransactOpts, debtToken common.Address) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "removePosition", debtToken)
}

// RemovePosition is a paid mutator transaction binding the contract method 0xdf5ca1f8.
//
// Solidity: function removePosition(address debtToken) returns()
func (_Stayking *StaykingSession) RemovePosition(debtToken common.Address) (*types.Transaction, error) {
	return _Stayking.Contract.RemovePosition(&_Stayking.TransactOpts, debtToken)
}

// RemovePosition is a paid mutator transaction binding the contract method 0xdf5ca1f8.
//
// Solidity: function removePosition(address debtToken) returns()
func (_Stayking *StaykingTransactorSession) RemovePosition(debtToken common.Address) (*types.Transaction, error) {
	return _Stayking.Contract.RemovePosition(&_Stayking.TransactOpts, debtToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stayking *StaykingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stayking *StaykingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stayking.Contract.RenounceOwnership(&_Stayking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stayking *StaykingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stayking.Contract.RenounceOwnership(&_Stayking.TransactOpts)
}

// RepayDebt is a paid mutator transaction binding the contract method 0x79e20bdc.
//
// Solidity: function repayDebt(address debtToken, uint256 repaidDebt) returns()
func (_Stayking *StaykingTransactor) RepayDebt(opts *bind.TransactOpts, debtToken common.Address, repaidDebt *big.Int) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "repayDebt", debtToken, repaidDebt)
}

// RepayDebt is a paid mutator transaction binding the contract method 0x79e20bdc.
//
// Solidity: function repayDebt(address debtToken, uint256 repaidDebt) returns()
func (_Stayking *StaykingSession) RepayDebt(debtToken common.Address, repaidDebt *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.RepayDebt(&_Stayking.TransactOpts, debtToken, repaidDebt)
}

// RepayDebt is a paid mutator transaction binding the contract method 0x79e20bdc.
//
// Solidity: function repayDebt(address debtToken, uint256 repaidDebt) returns()
func (_Stayking *StaykingTransactorSession) RepayDebt(debtToken common.Address, repaidDebt *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.RepayDebt(&_Stayking.TransactOpts, debtToken, repaidDebt)
}

// RepayDebtInBase is a paid mutator transaction binding the contract method 0xc6918a61.
//
// Solidity: function repayDebtInBase(address debtToken, uint256 minRepaid) payable returns()
func (_Stayking *StaykingTransactor) RepayDebtInBase(opts *bind.TransactOpts, debtToken common.Address, minRepaid *big.Int) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "repayDebtInBase", debtToken, minRepaid)
}

// RepayDebtInBase is a paid mutator transaction binding the contract method 0xc6918a61.
//
// Solidity: function repayDebtInBase(address debtToken, uint256 minRepaid) payable returns()
func (_Stayking *StaykingSession) RepayDebtInBase(debtToken common.Address, minRepaid *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.RepayDebtInBase(&_Stayking.TransactOpts, debtToken, minRepaid)
}

// RepayDebtInBase is a paid mutator transaction binding the contract method 0xc6918a61.
//
// Solidity: function repayDebtInBase(address debtToken, uint256 minRepaid) payable returns()
func (_Stayking *StaykingTransactorSession) RepayDebtInBase(debtToken common.Address, minRepaid *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.RepayDebtInBase(&_Stayking.TransactOpts, debtToken, minRepaid)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stayking *StaykingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stayking *StaykingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stayking.Contract.TransferOwnership(&_Stayking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stayking *StaykingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stayking.Contract.TransferOwnership(&_Stayking.TransactOpts, newOwner)
}

// UpdateConfigs is a paid mutator transaction binding the contract method 0x4f9e3805.
//
// Solidity: function updateConfigs(uint256 _minDebtInBase, uint256 _killFactorBps, uint256 _reservedBps) returns()
func (_Stayking *StaykingTransactor) UpdateConfigs(opts *bind.TransactOpts, _minDebtInBase *big.Int, _killFactorBps *big.Int, _reservedBps *big.Int) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "updateConfigs", _minDebtInBase, _killFactorBps, _reservedBps)
}

// UpdateConfigs is a paid mutator transaction binding the contract method 0x4f9e3805.
//
// Solidity: function updateConfigs(uint256 _minDebtInBase, uint256 _killFactorBps, uint256 _reservedBps) returns()
func (_Stayking *StaykingSession) UpdateConfigs(_minDebtInBase *big.Int, _killFactorBps *big.Int, _reservedBps *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.UpdateConfigs(&_Stayking.TransactOpts, _minDebtInBase, _killFactorBps, _reservedBps)
}

// UpdateConfigs is a paid mutator transaction binding the contract method 0x4f9e3805.
//
// Solidity: function updateConfigs(uint256 _minDebtInBase, uint256 _killFactorBps, uint256 _reservedBps) returns()
func (_Stayking *StaykingTransactorSession) UpdateConfigs(_minDebtInBase *big.Int, _killFactorBps *big.Int, _reservedBps *big.Int) (*types.Transaction, error) {
	return _Stayking.Contract.UpdateConfigs(&_Stayking.TransactOpts, _minDebtInBase, _killFactorBps, _reservedBps)
}

// UpdateVault is a paid mutator transaction binding the contract method 0x1b451d28.
//
// Solidity: function updateVault(address token, address vault) returns()
func (_Stayking *StaykingTransactor) UpdateVault(opts *bind.TransactOpts, token common.Address, vault common.Address) (*types.Transaction, error) {
	return _Stayking.contract.Transact(opts, "updateVault", token, vault)
}

// UpdateVault is a paid mutator transaction binding the contract method 0x1b451d28.
//
// Solidity: function updateVault(address token, address vault) returns()
func (_Stayking *StaykingSession) UpdateVault(token common.Address, vault common.Address) (*types.Transaction, error) {
	return _Stayking.Contract.UpdateVault(&_Stayking.TransactOpts, token, vault)
}

// UpdateVault is a paid mutator transaction binding the contract method 0x1b451d28.
//
// Solidity: function updateVault(address token, address vault) returns()
func (_Stayking *StaykingTransactorSession) UpdateVault(token common.Address, vault common.Address) (*types.Transaction, error) {
	return _Stayking.Contract.UpdateVault(&_Stayking.TransactOpts, token, vault)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Stayking *StaykingTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Stayking.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Stayking *StaykingSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Stayking.Contract.Fallback(&_Stayking.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Stayking *StaykingTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Stayking.Contract.Fallback(&_Stayking.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Stayking *StaykingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stayking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Stayking *StaykingSession) Receive() (*types.Transaction, error) {
	return _Stayking.Contract.Receive(&_Stayking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Stayking *StaykingTransactorSession) Receive() (*types.Transaction, error) {
	return _Stayking.Contract.Receive(&_Stayking.TransactOpts)
}

// StaykingAccrueIterator is returned from FilterAccrue and is used to iterate over the raw logs and unpacked data for Accrue events raised by the Stayking contract.
type StaykingAccrueIterator struct {
	Event *StaykingAccrue // Event containing the contract specifics and raw log

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
func (it *StaykingAccrueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaykingAccrue)
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
		it.Event = new(StaykingAccrue)
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
func (it *StaykingAccrueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaykingAccrueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaykingAccrue represents a Accrue event raised by the Stayking contract.
type StaykingAccrue struct {
	Delegator   common.Address
	TotalStaked *big.Int
	Distributed *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAccrue is a free log retrieval operation binding the contract event 0x68f604dad2502091469814ccdd58acac4592d7766a922eee86e6ab9965cfe949.
//
// Solidity: event Accrue(address indexed delegator, uint256 totalStaked, uint256 distributed)
func (_Stayking *StaykingFilterer) FilterAccrue(opts *bind.FilterOpts, delegator []common.Address) (*StaykingAccrueIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Stayking.contract.FilterLogs(opts, "Accrue", delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StaykingAccrueIterator{contract: _Stayking.contract, event: "Accrue", logs: logs, sub: sub}, nil
}

// WatchAccrue is a free log subscription operation binding the contract event 0x68f604dad2502091469814ccdd58acac4592d7766a922eee86e6ab9965cfe949.
//
// Solidity: event Accrue(address indexed delegator, uint256 totalStaked, uint256 distributed)
func (_Stayking *StaykingFilterer) WatchAccrue(opts *bind.WatchOpts, sink chan<- *StaykingAccrue, delegator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Stayking.contract.WatchLogs(opts, "Accrue", delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaykingAccrue)
				if err := _Stayking.contract.UnpackLog(event, "Accrue", log); err != nil {
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

// ParseAccrue is a log parse operation binding the contract event 0x68f604dad2502091469814ccdd58acac4592d7766a922eee86e6ab9965cfe949.
//
// Solidity: event Accrue(address indexed delegator, uint256 totalStaked, uint256 distributed)
func (_Stayking *StaykingFilterer) ParseAccrue(log types.Log) (*StaykingAccrue, error) {
	event := new(StaykingAccrue)
	if err := _Stayking.contract.UnpackLog(event, "Accrue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaykingAddPositionIterator is returned from FilterAddPosition and is used to iterate over the raw logs and unpacked data for AddPosition events raised by the Stayking contract.
type StaykingAddPositionIterator struct {
	Event *StaykingAddPosition // Event containing the contract specifics and raw log

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
func (it *StaykingAddPositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaykingAddPosition)
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
		it.Event = new(StaykingAddPosition)
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
func (it *StaykingAddPositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaykingAddPositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaykingAddPosition represents a AddPosition event raised by the Stayking contract.
type StaykingAddPosition struct {
	User       common.Address
	Vault      common.Address
	Equity     *big.Int
	DebtInBase *big.Int
	Debt       *big.Int
	Share      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddPosition is a free log retrieval operation binding the contract event 0x92b0fd5c6dbe6ec5b43a8afae4c587c1db47afcca43a21b7800bd848fcf8e496.
//
// Solidity: event AddPosition(address indexed user, address indexed vault, uint256 equity, uint256 debtInBase, uint256 debt, uint256 share)
func (_Stayking *StaykingFilterer) FilterAddPosition(opts *bind.FilterOpts, user []common.Address, vault []common.Address) (*StaykingAddPositionIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Stayking.contract.FilterLogs(opts, "AddPosition", userRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return &StaykingAddPositionIterator{contract: _Stayking.contract, event: "AddPosition", logs: logs, sub: sub}, nil
}

// WatchAddPosition is a free log subscription operation binding the contract event 0x92b0fd5c6dbe6ec5b43a8afae4c587c1db47afcca43a21b7800bd848fcf8e496.
//
// Solidity: event AddPosition(address indexed user, address indexed vault, uint256 equity, uint256 debtInBase, uint256 debt, uint256 share)
func (_Stayking *StaykingFilterer) WatchAddPosition(opts *bind.WatchOpts, sink chan<- *StaykingAddPosition, user []common.Address, vault []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Stayking.contract.WatchLogs(opts, "AddPosition", userRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaykingAddPosition)
				if err := _Stayking.contract.UnpackLog(event, "AddPosition", log); err != nil {
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

// ParseAddPosition is a log parse operation binding the contract event 0x92b0fd5c6dbe6ec5b43a8afae4c587c1db47afcca43a21b7800bd848fcf8e496.
//
// Solidity: event AddPosition(address indexed user, address indexed vault, uint256 equity, uint256 debtInBase, uint256 debt, uint256 share)
func (_Stayking *StaykingFilterer) ParseAddPosition(log types.Log) (*StaykingAddPosition, error) {
	event := new(StaykingAddPosition)
	if err := _Stayking.contract.UnpackLog(event, "AddPosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaykingAddVaultIterator is returned from FilterAddVault and is used to iterate over the raw logs and unpacked data for AddVault events raised by the Stayking contract.
type StaykingAddVaultIterator struct {
	Event *StaykingAddVault // Event containing the contract specifics and raw log

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
func (it *StaykingAddVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaykingAddVault)
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
		it.Event = new(StaykingAddVault)
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
func (it *StaykingAddVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaykingAddVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaykingAddVault represents a AddVault event raised by the Stayking contract.
type StaykingAddVault struct {
	Token common.Address
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddVault is a free log retrieval operation binding the contract event 0x7f03ac7449144076e2d9d5a237c3ff0d0108e0a2ebdd5213c4b0de6f4d4f6953.
//
// Solidity: event AddVault(address token, address vault)
func (_Stayking *StaykingFilterer) FilterAddVault(opts *bind.FilterOpts) (*StaykingAddVaultIterator, error) {

	logs, sub, err := _Stayking.contract.FilterLogs(opts, "AddVault")
	if err != nil {
		return nil, err
	}
	return &StaykingAddVaultIterator{contract: _Stayking.contract, event: "AddVault", logs: logs, sub: sub}, nil
}

// WatchAddVault is a free log subscription operation binding the contract event 0x7f03ac7449144076e2d9d5a237c3ff0d0108e0a2ebdd5213c4b0de6f4d4f6953.
//
// Solidity: event AddVault(address token, address vault)
func (_Stayking *StaykingFilterer) WatchAddVault(opts *bind.WatchOpts, sink chan<- *StaykingAddVault) (event.Subscription, error) {

	logs, sub, err := _Stayking.contract.WatchLogs(opts, "AddVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaykingAddVault)
				if err := _Stayking.contract.UnpackLog(event, "AddVault", log); err != nil {
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

// ParseAddVault is a log parse operation binding the contract event 0x7f03ac7449144076e2d9d5a237c3ff0d0108e0a2ebdd5213c4b0de6f4d4f6953.
//
// Solidity: event AddVault(address token, address vault)
func (_Stayking *StaykingFilterer) ParseAddVault(log types.Log) (*StaykingAddVault, error) {
	event := new(StaykingAddVault)
	if err := _Stayking.contract.UnpackLog(event, "AddVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaykingChangeDelegatorIterator is returned from FilterChangeDelegator and is used to iterate over the raw logs and unpacked data for ChangeDelegator events raised by the Stayking contract.
type StaykingChangeDelegatorIterator struct {
	Event *StaykingChangeDelegator // Event containing the contract specifics and raw log

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
func (it *StaykingChangeDelegatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaykingChangeDelegator)
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
		it.Event = new(StaykingChangeDelegator)
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
func (it *StaykingChangeDelegatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaykingChangeDelegatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaykingChangeDelegator represents a ChangeDelegator event raised by the Stayking contract.
type StaykingChangeDelegator struct {
	Delegator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangeDelegator is a free log retrieval operation binding the contract event 0xc46722a14b09e7b7e59991ec89fb4bb88ffed57f1b2187a65e959f5c0053716d.
//
// Solidity: event ChangeDelegator(address delegator)
func (_Stayking *StaykingFilterer) FilterChangeDelegator(opts *bind.FilterOpts) (*StaykingChangeDelegatorIterator, error) {

	logs, sub, err := _Stayking.contract.FilterLogs(opts, "ChangeDelegator")
	if err != nil {
		return nil, err
	}
	return &StaykingChangeDelegatorIterator{contract: _Stayking.contract, event: "ChangeDelegator", logs: logs, sub: sub}, nil
}

// WatchChangeDelegator is a free log subscription operation binding the contract event 0xc46722a14b09e7b7e59991ec89fb4bb88ffed57f1b2187a65e959f5c0053716d.
//
// Solidity: event ChangeDelegator(address delegator)
func (_Stayking *StaykingFilterer) WatchChangeDelegator(opts *bind.WatchOpts, sink chan<- *StaykingChangeDelegator) (event.Subscription, error) {

	logs, sub, err := _Stayking.contract.WatchLogs(opts, "ChangeDelegator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaykingChangeDelegator)
				if err := _Stayking.contract.UnpackLog(event, "ChangeDelegator", log); err != nil {
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

// ParseChangeDelegator is a log parse operation binding the contract event 0xc46722a14b09e7b7e59991ec89fb4bb88ffed57f1b2187a65e959f5c0053716d.
//
// Solidity: event ChangeDelegator(address delegator)
func (_Stayking *StaykingFilterer) ParseChangeDelegator(log types.Log) (*StaykingChangeDelegator, error) {
	event := new(StaykingChangeDelegator)
	if err := _Stayking.contract.UnpackLog(event, "ChangeDelegator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaykingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Stayking contract.
type StaykingInitializedIterator struct {
	Event *StaykingInitialized // Event containing the contract specifics and raw log

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
func (it *StaykingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaykingInitialized)
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
		it.Event = new(StaykingInitialized)
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
func (it *StaykingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaykingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaykingInitialized represents a Initialized event raised by the Stayking contract.
type StaykingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Stayking *StaykingFilterer) FilterInitialized(opts *bind.FilterOpts) (*StaykingInitializedIterator, error) {

	logs, sub, err := _Stayking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StaykingInitializedIterator{contract: _Stayking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Stayking *StaykingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StaykingInitialized) (event.Subscription, error) {

	logs, sub, err := _Stayking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaykingInitialized)
				if err := _Stayking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Stayking *StaykingFilterer) ParseInitialized(log types.Log) (*StaykingInitialized, error) {
	event := new(StaykingInitialized)
	if err := _Stayking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaykingKillIterator is returned from FilterKill and is used to iterate over the raw logs and unpacked data for Kill events raised by the Stayking contract.
type StaykingKillIterator struct {
	Event *StaykingKill // Event containing the contract specifics and raw log

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
func (it *StaykingKillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaykingKill)
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
		it.Event = new(StaykingKill)
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
func (it *StaykingKillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaykingKillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaykingKill represents a Kill event raised by the Stayking contract.
type StaykingKill struct {
	Killer     common.Address
	User       common.Address
	Vault      common.Address
	Equity     *big.Int
	DebtInBase *big.Int
	Debt       *big.Int
	Share      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterKill is a free log retrieval operation binding the contract event 0x62b067e31c5a2d2734bc19924c46c930e1b7063b5b473238f2f8bb7ab88a15b6.
//
// Solidity: event Kill(address indexed killer, address indexed user, address vault, uint256 equity, uint256 debtInBase, uint256 debt, uint256 share)
func (_Stayking *StaykingFilterer) FilterKill(opts *bind.FilterOpts, killer []common.Address, user []common.Address) (*StaykingKillIterator, error) {

	var killerRule []interface{}
	for _, killerItem := range killer {
		killerRule = append(killerRule, killerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stayking.contract.FilterLogs(opts, "Kill", killerRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StaykingKillIterator{contract: _Stayking.contract, event: "Kill", logs: logs, sub: sub}, nil
}

// WatchKill is a free log subscription operation binding the contract event 0x62b067e31c5a2d2734bc19924c46c930e1b7063b5b473238f2f8bb7ab88a15b6.
//
// Solidity: event Kill(address indexed killer, address indexed user, address vault, uint256 equity, uint256 debtInBase, uint256 debt, uint256 share)
func (_Stayking *StaykingFilterer) WatchKill(opts *bind.WatchOpts, sink chan<- *StaykingKill, killer []common.Address, user []common.Address) (event.Subscription, error) {

	var killerRule []interface{}
	for _, killerItem := range killer {
		killerRule = append(killerRule, killerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stayking.contract.WatchLogs(opts, "Kill", killerRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaykingKill)
				if err := _Stayking.contract.UnpackLog(event, "Kill", log); err != nil {
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

// ParseKill is a log parse operation binding the contract event 0x62b067e31c5a2d2734bc19924c46c930e1b7063b5b473238f2f8bb7ab88a15b6.
//
// Solidity: event Kill(address indexed killer, address indexed user, address vault, uint256 equity, uint256 debtInBase, uint256 debt, uint256 share)
func (_Stayking *StaykingFilterer) ParseKill(log types.Log) (*StaykingKill, error) {
	event := new(StaykingKill)
	if err := _Stayking.contract.UnpackLog(event, "Kill", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaykingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stayking contract.
type StaykingOwnershipTransferredIterator struct {
	Event *StaykingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StaykingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaykingOwnershipTransferred)
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
		it.Event = new(StaykingOwnershipTransferred)
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
func (it *StaykingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaykingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaykingOwnershipTransferred represents a OwnershipTransferred event raised by the Stayking contract.
type StaykingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stayking *StaykingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StaykingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stayking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StaykingOwnershipTransferredIterator{contract: _Stayking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stayking *StaykingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StaykingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stayking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaykingOwnershipTransferred)
				if err := _Stayking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Stayking *StaykingFilterer) ParseOwnershipTransferred(log types.Log) (*StaykingOwnershipTransferred, error) {
	event := new(StaykingOwnershipTransferred)
	if err := _Stayking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaykingPositionChangedIterator is returned from FilterPositionChanged and is used to iterate over the raw logs and unpacked data for PositionChanged events raised by the Stayking contract.
type StaykingPositionChangedIterator struct {
	Event *StaykingPositionChanged // Event containing the contract specifics and raw log

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
func (it *StaykingPositionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaykingPositionChanged)
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
		it.Event = new(StaykingPositionChanged)
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
func (it *StaykingPositionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaykingPositionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaykingPositionChanged represents a PositionChanged event raised by the Stayking contract.
type StaykingPositionChanged struct {
	User   common.Address
	Vault  common.Address
	Amount *big.Int
	Share  *big.Int
	Debt   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPositionChanged is a free log retrieval operation binding the contract event 0xf35a14688e4794e8022cd9e5c7ec6475a6095ef189dea9647fad0de771465b8a.
//
// Solidity: event PositionChanged(address indexed user, address indexed vault, uint256 amount, uint256 share, uint256 debt)
func (_Stayking *StaykingFilterer) FilterPositionChanged(opts *bind.FilterOpts, user []common.Address, vault []common.Address) (*StaykingPositionChangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Stayking.contract.FilterLogs(opts, "PositionChanged", userRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return &StaykingPositionChangedIterator{contract: _Stayking.contract, event: "PositionChanged", logs: logs, sub: sub}, nil
}

// WatchPositionChanged is a free log subscription operation binding the contract event 0xf35a14688e4794e8022cd9e5c7ec6475a6095ef189dea9647fad0de771465b8a.
//
// Solidity: event PositionChanged(address indexed user, address indexed vault, uint256 amount, uint256 share, uint256 debt)
func (_Stayking *StaykingFilterer) WatchPositionChanged(opts *bind.WatchOpts, sink chan<- *StaykingPositionChanged, user []common.Address, vault []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Stayking.contract.WatchLogs(opts, "PositionChanged", userRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaykingPositionChanged)
				if err := _Stayking.contract.UnpackLog(event, "PositionChanged", log); err != nil {
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

// ParsePositionChanged is a log parse operation binding the contract event 0xf35a14688e4794e8022cd9e5c7ec6475a6095ef189dea9647fad0de771465b8a.
//
// Solidity: event PositionChanged(address indexed user, address indexed vault, uint256 amount, uint256 share, uint256 debt)
func (_Stayking *StaykingFilterer) ParsePositionChanged(log types.Log) (*StaykingPositionChanged, error) {
	event := new(StaykingPositionChanged)
	if err := _Stayking.contract.UnpackLog(event, "PositionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaykingRemovePositionIterator is returned from FilterRemovePosition and is used to iterate over the raw logs and unpacked data for RemovePosition events raised by the Stayking contract.
type StaykingRemovePositionIterator struct {
	Event *StaykingRemovePosition // Event containing the contract specifics and raw log

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
func (it *StaykingRemovePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaykingRemovePosition)
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
		it.Event = new(StaykingRemovePosition)
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
func (it *StaykingRemovePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaykingRemovePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaykingRemovePosition represents a RemovePosition event raised by the Stayking contract.
type StaykingRemovePosition struct {
	User       common.Address
	Vault      common.Address
	Equity     *big.Int
	DebtInBase *big.Int
	Debt       *big.Int
	Share      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemovePosition is a free log retrieval operation binding the contract event 0x3cb91b4d61d4fefa96da17651a09b54f7b6751cb68d11134cf0efe0a49657968.
//
// Solidity: event RemovePosition(address indexed user, address indexed vault, uint256 equity, uint256 debtInBase, uint256 debt, uint256 share)
func (_Stayking *StaykingFilterer) FilterRemovePosition(opts *bind.FilterOpts, user []common.Address, vault []common.Address) (*StaykingRemovePositionIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Stayking.contract.FilterLogs(opts, "RemovePosition", userRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return &StaykingRemovePositionIterator{contract: _Stayking.contract, event: "RemovePosition", logs: logs, sub: sub}, nil
}

// WatchRemovePosition is a free log subscription operation binding the contract event 0x3cb91b4d61d4fefa96da17651a09b54f7b6751cb68d11134cf0efe0a49657968.
//
// Solidity: event RemovePosition(address indexed user, address indexed vault, uint256 equity, uint256 debtInBase, uint256 debt, uint256 share)
func (_Stayking *StaykingFilterer) WatchRemovePosition(opts *bind.WatchOpts, sink chan<- *StaykingRemovePosition, user []common.Address, vault []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Stayking.contract.WatchLogs(opts, "RemovePosition", userRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaykingRemovePosition)
				if err := _Stayking.contract.UnpackLog(event, "RemovePosition", log); err != nil {
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

// ParseRemovePosition is a log parse operation binding the contract event 0x3cb91b4d61d4fefa96da17651a09b54f7b6751cb68d11134cf0efe0a49657968.
//
// Solidity: event RemovePosition(address indexed user, address indexed vault, uint256 equity, uint256 debtInBase, uint256 debt, uint256 share)
func (_Stayking *StaykingFilterer) ParseRemovePosition(log types.Log) (*StaykingRemovePosition, error) {
	event := new(StaykingRemovePosition)
	if err := _Stayking.contract.UnpackLog(event, "RemovePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaykingStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the Stayking contract.
type StaykingStakeIterator struct {
	Event *StaykingStake // Event containing the contract specifics and raw log

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
func (it *StaykingStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaykingStake)
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
		it.Event = new(StaykingStake)
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
func (it *StaykingStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaykingStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaykingStake represents a Stake event raised by the Stayking contract.
type StaykingStake struct {
	Delegator common.Address
	User      common.Address
	Amount    *big.Int
	Share     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address indexed delegator, address indexed user, uint256 amount, uint256 share)
func (_Stayking *StaykingFilterer) FilterStake(opts *bind.FilterOpts, delegator []common.Address, user []common.Address) (*StaykingStakeIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stayking.contract.FilterLogs(opts, "Stake", delegatorRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StaykingStakeIterator{contract: _Stayking.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address indexed delegator, address indexed user, uint256 amount, uint256 share)
func (_Stayking *StaykingFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *StaykingStake, delegator []common.Address, user []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stayking.contract.WatchLogs(opts, "Stake", delegatorRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaykingStake)
				if err := _Stayking.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address indexed delegator, address indexed user, uint256 amount, uint256 share)
func (_Stayking *StaykingFilterer) ParseStake(log types.Log) (*StaykingStake, error) {
	event := new(StaykingStake)
	if err := _Stayking.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaykingUnstakeIterator is returned from FilterUnstake and is used to iterate over the raw logs and unpacked data for Unstake events raised by the Stayking contract.
type StaykingUnstakeIterator struct {
	Event *StaykingUnstake // Event containing the contract specifics and raw log

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
func (it *StaykingUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaykingUnstake)
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
		it.Event = new(StaykingUnstake)
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
func (it *StaykingUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaykingUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaykingUnstake represents a Unstake event raised by the Stayking contract.
type StaykingUnstake struct {
	Delegator common.Address
	User      common.Address
	Amount    *big.Int
	Share     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstake is a free log retrieval operation binding the contract event 0x18edd09e80386cd99df397e2e0d87d2bb259423eae08645e776321a36fe680ef.
//
// Solidity: event Unstake(address indexed delegator, address indexed user, uint256 amount, uint256 share)
func (_Stayking *StaykingFilterer) FilterUnstake(opts *bind.FilterOpts, delegator []common.Address, user []common.Address) (*StaykingUnstakeIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stayking.contract.FilterLogs(opts, "Unstake", delegatorRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StaykingUnstakeIterator{contract: _Stayking.contract, event: "Unstake", logs: logs, sub: sub}, nil
}

// WatchUnstake is a free log subscription operation binding the contract event 0x18edd09e80386cd99df397e2e0d87d2bb259423eae08645e776321a36fe680ef.
//
// Solidity: event Unstake(address indexed delegator, address indexed user, uint256 amount, uint256 share)
func (_Stayking *StaykingFilterer) WatchUnstake(opts *bind.WatchOpts, sink chan<- *StaykingUnstake, delegator []common.Address, user []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stayking.contract.WatchLogs(opts, "Unstake", delegatorRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaykingUnstake)
				if err := _Stayking.contract.UnpackLog(event, "Unstake", log); err != nil {
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

// ParseUnstake is a log parse operation binding the contract event 0x18edd09e80386cd99df397e2e0d87d2bb259423eae08645e776321a36fe680ef.
//
// Solidity: event Unstake(address indexed delegator, address indexed user, uint256 amount, uint256 share)
func (_Stayking *StaykingFilterer) ParseUnstake(log types.Log) (*StaykingUnstake, error) {
	event := new(StaykingUnstake)
	if err := _Stayking.contract.UnpackLog(event, "Unstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaykingUpdateConfigsIterator is returned from FilterUpdateConfigs and is used to iterate over the raw logs and unpacked data for UpdateConfigs events raised by the Stayking contract.
type StaykingUpdateConfigsIterator struct {
	Event *StaykingUpdateConfigs // Event containing the contract specifics and raw log

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
func (it *StaykingUpdateConfigsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaykingUpdateConfigs)
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
		it.Event = new(StaykingUpdateConfigs)
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
func (it *StaykingUpdateConfigsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaykingUpdateConfigsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaykingUpdateConfigs represents a UpdateConfigs event raised by the Stayking contract.
type StaykingUpdateConfigs struct {
	MinDebtInBase *big.Int
	KillFactorBps *big.Int
	ReservedBps   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateConfigs is a free log retrieval operation binding the contract event 0xe032360dc1790b6d0ffbd119645e134124ea15d2ab7d012f101fe1df3016a52c.
//
// Solidity: event UpdateConfigs(uint256 minDebtInBase, uint256 killFactorBps, uint256 reservedBps)
func (_Stayking *StaykingFilterer) FilterUpdateConfigs(opts *bind.FilterOpts) (*StaykingUpdateConfigsIterator, error) {

	logs, sub, err := _Stayking.contract.FilterLogs(opts, "UpdateConfigs")
	if err != nil {
		return nil, err
	}
	return &StaykingUpdateConfigsIterator{contract: _Stayking.contract, event: "UpdateConfigs", logs: logs, sub: sub}, nil
}

// WatchUpdateConfigs is a free log subscription operation binding the contract event 0xe032360dc1790b6d0ffbd119645e134124ea15d2ab7d012f101fe1df3016a52c.
//
// Solidity: event UpdateConfigs(uint256 minDebtInBase, uint256 killFactorBps, uint256 reservedBps)
func (_Stayking *StaykingFilterer) WatchUpdateConfigs(opts *bind.WatchOpts, sink chan<- *StaykingUpdateConfigs) (event.Subscription, error) {

	logs, sub, err := _Stayking.contract.WatchLogs(opts, "UpdateConfigs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaykingUpdateConfigs)
				if err := _Stayking.contract.UnpackLog(event, "UpdateConfigs", log); err != nil {
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

// ParseUpdateConfigs is a log parse operation binding the contract event 0xe032360dc1790b6d0ffbd119645e134124ea15d2ab7d012f101fe1df3016a52c.
//
// Solidity: event UpdateConfigs(uint256 minDebtInBase, uint256 killFactorBps, uint256 reservedBps)
func (_Stayking *StaykingFilterer) ParseUpdateConfigs(log types.Log) (*StaykingUpdateConfigs, error) {
	event := new(StaykingUpdateConfigs)
	if err := _Stayking.contract.UnpackLog(event, "UpdateConfigs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaykingUpdateVaultIterator is returned from FilterUpdateVault and is used to iterate over the raw logs and unpacked data for UpdateVault events raised by the Stayking contract.
type StaykingUpdateVaultIterator struct {
	Event *StaykingUpdateVault // Event containing the contract specifics and raw log

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
func (it *StaykingUpdateVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaykingUpdateVault)
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
		it.Event = new(StaykingUpdateVault)
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
func (it *StaykingUpdateVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaykingUpdateVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaykingUpdateVault represents a UpdateVault event raised by the Stayking contract.
type StaykingUpdateVault struct {
	Token common.Address
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateVault is a free log retrieval operation binding the contract event 0x4cc233b502e1e28f371b28114d4a8be87a93f330c18e3c339079fc498320fb8d.
//
// Solidity: event UpdateVault(address token, address vault)
func (_Stayking *StaykingFilterer) FilterUpdateVault(opts *bind.FilterOpts) (*StaykingUpdateVaultIterator, error) {

	logs, sub, err := _Stayking.contract.FilterLogs(opts, "UpdateVault")
	if err != nil {
		return nil, err
	}
	return &StaykingUpdateVaultIterator{contract: _Stayking.contract, event: "UpdateVault", logs: logs, sub: sub}, nil
}

// WatchUpdateVault is a free log subscription operation binding the contract event 0x4cc233b502e1e28f371b28114d4a8be87a93f330c18e3c339079fc498320fb8d.
//
// Solidity: event UpdateVault(address token, address vault)
func (_Stayking *StaykingFilterer) WatchUpdateVault(opts *bind.WatchOpts, sink chan<- *StaykingUpdateVault) (event.Subscription, error) {

	logs, sub, err := _Stayking.contract.WatchLogs(opts, "UpdateVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaykingUpdateVault)
				if err := _Stayking.contract.UnpackLog(event, "UpdateVault", log); err != nil {
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

// ParseUpdateVault is a log parse operation binding the contract event 0x4cc233b502e1e28f371b28114d4a8be87a93f330c18e3c339079fc498320fb8d.
//
// Solidity: event UpdateVault(address token, address vault)
func (_Stayking *StaykingFilterer) ParseUpdateVault(log types.Log) (*StaykingUpdateVault, error) {
	event := new(StaykingUpdateVault)
	if err := _Stayking.contract.UnpackLog(event, "UpdateVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
