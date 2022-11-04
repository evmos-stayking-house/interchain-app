// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

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

// AbisMetaData contains all meta data concerning the Abis contract.
var AbisMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"Unstake\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getAccruedValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyUnbondedToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unbonded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AbisABI is the input ABI used to generate the binding from.
// Deprecated: Use AbisMetaData.ABI instead.
var AbisABI = AbisMetaData.ABI

// Abis is an auto generated Go binding around an Ethereum contract.
type Abis struct {
	AbisCaller     // Read-only binding to the contract
	AbisTransactor // Write-only binding to the contract
	AbisFilterer   // Log filterer for contract events
}

// AbisCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbisSession struct {
	Contract     *Abis             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbisCallerSession struct {
	Contract *AbisCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbisTransactorSession struct {
	Contract     *AbisTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbisRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbisRaw struct {
	Contract *Abis // Generic contract binding to access the raw methods on
}

// AbisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbisCallerRaw struct {
	Contract *AbisCaller // Generic read-only contract binding to access the raw methods on
}

// AbisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbisTransactorRaw struct {
	Contract *AbisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbis creates a new instance of Abis, bound to a specific deployed contract.
func NewAbis(address common.Address, backend bind.ContractBackend) (*Abis, error) {
	contract, err := bindAbis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abis{AbisCaller: AbisCaller{contract: contract}, AbisTransactor: AbisTransactor{contract: contract}, AbisFilterer: AbisFilterer{contract: contract}}, nil
}

// NewAbisCaller creates a new read-only instance of Abis, bound to a specific deployed contract.
func NewAbisCaller(address common.Address, caller bind.ContractCaller) (*AbisCaller, error) {
	contract, err := bindAbis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbisCaller{contract: contract}, nil
}

// NewAbisTransactor creates a new write-only instance of Abis, bound to a specific deployed contract.
func NewAbisTransactor(address common.Address, transactor bind.ContractTransactor) (*AbisTransactor, error) {
	contract, err := bindAbis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbisTransactor{contract: contract}, nil
}

// NewAbisFilterer creates a new log filterer instance of Abis, bound to a specific deployed contract.
func NewAbisFilterer(address common.Address, filterer bind.ContractFilterer) (*AbisFilterer, error) {
	contract, err := bindAbis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbisFilterer{contract: contract}, nil
}

// bindAbis binds a generic wrapper to an already deployed contract.
func bindAbis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbisABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abis *AbisRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abis.Contract.AbisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abis *AbisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abis.Contract.AbisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abis *AbisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abis.Contract.AbisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abis *AbisCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abis *AbisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abis *AbisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abis.Contract.contract.Transact(opts, method, params...)
}

// GetAccruedValue is a free data retrieval call binding the contract method 0xe4175aee.
//
// Solidity: function getAccruedValue(uint256 amount) view returns(uint256)
func (_Abis *AbisCaller) GetAccruedValue(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "getAccruedValue", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccruedValue is a free data retrieval call binding the contract method 0xe4175aee.
//
// Solidity: function getAccruedValue(uint256 amount) view returns(uint256)
func (_Abis *AbisSession) GetAccruedValue(amount *big.Int) (*big.Int, error) {
	return _Abis.Contract.GetAccruedValue(&_Abis.CallOpts, amount)
}

// GetAccruedValue is a free data retrieval call binding the contract method 0xe4175aee.
//
// Solidity: function getAccruedValue(uint256 amount) view returns(uint256)
func (_Abis *AbisCallerSession) GetAccruedValue(amount *big.Int) (*big.Int, error) {
	return _Abis.Contract.GetAccruedValue(&_Abis.CallOpts, amount)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Abis *AbisCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Abis *AbisSession) Retrieve() (*big.Int, error) {
	return _Abis.Contract.Retrieve(&_Abis.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Abis *AbisCallerSession) Retrieve() (*big.Int, error) {
	return _Abis.Contract.Retrieve(&_Abis.CallOpts)
}

// Unbonded is a free data retrieval call binding the contract method 0xc144cd11.
//
// Solidity: function unbonded() view returns(uint256)
func (_Abis *AbisCaller) Unbonded(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abis.contract.Call(opts, &out, "unbonded")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Unbonded is a free data retrieval call binding the contract method 0xc144cd11.
//
// Solidity: function unbonded() view returns(uint256)
func (_Abis *AbisSession) Unbonded() (*big.Int, error) {
	return _Abis.Contract.Unbonded(&_Abis.CallOpts)
}

// Unbonded is a free data retrieval call binding the contract method 0xc144cd11.
//
// Solidity: function unbonded() view returns(uint256)
func (_Abis *AbisCallerSession) Unbonded() (*big.Int, error) {
	return _Abis.Contract.Unbonded(&_Abis.CallOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 amount) returns()
func (_Abis *AbisTransactor) Delegate(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Abis.contract.Transact(opts, "delegate", amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 amount) returns()
func (_Abis *AbisSession) Delegate(amount *big.Int) (*types.Transaction, error) {
	return _Abis.Contract.Delegate(&_Abis.TransactOpts, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 amount) returns()
func (_Abis *AbisTransactorSession) Delegate(amount *big.Int) (*types.Transaction, error) {
	return _Abis.Contract.Delegate(&_Abis.TransactOpts, amount)
}

// SupplyUnbondedToken is a paid mutator transaction binding the contract method 0xe158c1ac.
//
// Solidity: function supplyUnbondedToken() payable returns()
func (_Abis *AbisTransactor) SupplyUnbondedToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abis.contract.Transact(opts, "supplyUnbondedToken")
}

// SupplyUnbondedToken is a paid mutator transaction binding the contract method 0xe158c1ac.
//
// Solidity: function supplyUnbondedToken() payable returns()
func (_Abis *AbisSession) SupplyUnbondedToken() (*types.Transaction, error) {
	return _Abis.Contract.SupplyUnbondedToken(&_Abis.TransactOpts)
}

// SupplyUnbondedToken is a paid mutator transaction binding the contract method 0xe158c1ac.
//
// Solidity: function supplyUnbondedToken() payable returns()
func (_Abis *AbisTransactorSession) SupplyUnbondedToken() (*types.Transaction, error) {
	return _Abis.Contract.SupplyUnbondedToken(&_Abis.TransactOpts)
}

// Undelegate is a paid mutator transaction binding the contract method 0x6c68c0e1.
//
// Solidity: function undelegate(uint256 amount) returns()
func (_Abis *AbisTransactor) Undelegate(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Abis.contract.Transact(opts, "undelegate", amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x6c68c0e1.
//
// Solidity: function undelegate(uint256 amount) returns()
func (_Abis *AbisSession) Undelegate(amount *big.Int) (*types.Transaction, error) {
	return _Abis.Contract.Undelegate(&_Abis.TransactOpts, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x6c68c0e1.
//
// Solidity: function undelegate(uint256 amount) returns()
func (_Abis *AbisTransactorSession) Undelegate(amount *big.Int) (*types.Transaction, error) {
	return _Abis.Contract.Undelegate(&_Abis.TransactOpts, amount)
}

// AbisStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the Abis contract.
type AbisStakeIterator struct {
	Event *AbisStake // Event containing the contract specifics and raw log

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
func (it *AbisStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbisStake)
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
		it.Event = new(AbisStake)
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
func (it *AbisStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbisStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbisStake represents a Stake event raised by the Abis contract.
type AbisStake struct {
	Delegator common.Address
	User      common.Address
	Amount    *big.Int
	Share     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address indexed delegator, address indexed user, uint256 amount, uint256 share)
func (_Abis *AbisFilterer) FilterStake(opts *bind.FilterOpts, delegator []common.Address, user []common.Address) (*AbisStakeIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abis.contract.FilterLogs(opts, "Stake", delegatorRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AbisStakeIterator{contract: _Abis.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address indexed delegator, address indexed user, uint256 amount, uint256 share)
func (_Abis *AbisFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *AbisStake, delegator []common.Address, user []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abis.contract.WatchLogs(opts, "Stake", delegatorRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbisStake)
				if err := _Abis.contract.UnpackLog(event, "Stake", log); err != nil {
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
func (_Abis *AbisFilterer) ParseStake(log types.Log) (*AbisStake, error) {
	event := new(AbisStake)
	if err := _Abis.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbisUnstakeIterator is returned from FilterUnstake and is used to iterate over the raw logs and unpacked data for Unstake events raised by the Abis contract.
type AbisUnstakeIterator struct {
	Event *AbisUnstake // Event containing the contract specifics and raw log

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
func (it *AbisUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbisUnstake)
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
		it.Event = new(AbisUnstake)
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
func (it *AbisUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbisUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbisUnstake represents a Unstake event raised by the Abis contract.
type AbisUnstake struct {
	Delegator common.Address
	User      common.Address
	Amount    *big.Int
	Share     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstake is a free log retrieval operation binding the contract event 0x18edd09e80386cd99df397e2e0d87d2bb259423eae08645e776321a36fe680ef.
//
// Solidity: event Unstake(address indexed delegator, address indexed user, uint256 amount, uint256 share)
func (_Abis *AbisFilterer) FilterUnstake(opts *bind.FilterOpts, delegator []common.Address, user []common.Address) (*AbisUnstakeIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abis.contract.FilterLogs(opts, "Unstake", delegatorRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AbisUnstakeIterator{contract: _Abis.contract, event: "Unstake", logs: logs, sub: sub}, nil
}

// WatchUnstake is a free log subscription operation binding the contract event 0x18edd09e80386cd99df397e2e0d87d2bb259423eae08645e776321a36fe680ef.
//
// Solidity: event Unstake(address indexed delegator, address indexed user, uint256 amount, uint256 share)
func (_Abis *AbisFilterer) WatchUnstake(opts *bind.WatchOpts, sink chan<- *AbisUnstake, delegator []common.Address, user []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abis.contract.WatchLogs(opts, "Unstake", delegatorRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbisUnstake)
				if err := _Abis.contract.UnpackLog(event, "Unstake", log); err != nil {
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
func (_Abis *AbisFilterer) ParseUnstake(log types.Log) (*AbisUnstake, error) {
	event := new(AbisUnstake)
	if err := _Abis.contract.UnpackLog(event, "Unstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
