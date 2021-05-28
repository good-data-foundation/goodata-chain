// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interfaces

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

// GoodDataContractABI is the input ABI used to generate the binding from.
const GoodDataContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"uuid\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"DORegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"queryUUID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"doUUIDs\",\"type\":\"bytes\"}],\"name\":\"DOsOfQuery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"uuid\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"LogAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"uuid\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"MPCRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"queryUUID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"mpcUUIDs\",\"type\":\"bytes\"}],\"name\":\"MPCsOfQuery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"predictionUUID\",\"type\":\"bytes\"}],\"name\":\"PredictionFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"queryUUID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"predictionUUID\",\"type\":\"bytes\"}],\"name\":\"PredictionSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"uuid\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"QCRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"queryUUID\",\"type\":\"bytes\"}],\"name\":\"QueryFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"queryUUID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"doUUID\",\"type\":\"bytes\"}],\"name\":\"QueryFinishedInOneDO\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"queryUUID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"qcUUID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"dataSet\",\"type\":\"int64\"}],\"name\":\"QuerySubmitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"uuid\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"addLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"queryUUID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"doUUIDs\",\"type\":\"bytes\"}],\"name\":\"assignQueryToDOs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"queryUUID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"mpcUUIDs\",\"type\":\"bytes\"}],\"name\":\"assignQueryToMPCs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"predictionUUID\",\"type\":\"bytes\"}],\"name\":\"finishPrediction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"queryUUID\",\"type\":\"bytes\"}],\"name\":\"finishQuery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"queryUUID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"doUUID\",\"type\":\"bytes\"}],\"name\":\"finishQueryInOneDO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"uuid\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"registerDO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"uuid\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"registerMPC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"uuid\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"registerQC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"queryUUID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"predictionUUID\",\"type\":\"bytes\"}],\"name\":\"submitPrediction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"queryUUID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"qcUUID\",\"type\":\"bytes\"},{\"internalType\":\"int64\",\"name\":\"dataSet\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"submitQuery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GoodDataContract is an auto generated Go binding around an Ethereum contract.
type GoodDataContract struct {
	GoodDataContractCaller     // Read-only binding to the contract
	GoodDataContractTransactor // Write-only binding to the contract
	GoodDataContractFilterer   // Log filterer for contract events
}

// GoodDataContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type GoodDataContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoodDataContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GoodDataContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoodDataContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GoodDataContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoodDataContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GoodDataContractSession struct {
	Contract     *GoodDataContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GoodDataContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GoodDataContractCallerSession struct {
	Contract *GoodDataContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// GoodDataContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GoodDataContractTransactorSession struct {
	Contract     *GoodDataContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// GoodDataContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type GoodDataContractRaw struct {
	Contract *GoodDataContract // Generic contract binding to access the raw methods on
}

// GoodDataContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GoodDataContractCallerRaw struct {
	Contract *GoodDataContractCaller // Generic read-only contract binding to access the raw methods on
}

// GoodDataContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GoodDataContractTransactorRaw struct {
	Contract *GoodDataContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGoodDataContract creates a new instance of GoodDataContract, bound to a specific deployed contract.
func NewGoodDataContract(address common.Address, backend bind.ContractBackend) (*GoodDataContract, error) {
	contract, err := bindGoodDataContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GoodDataContract{GoodDataContractCaller: GoodDataContractCaller{contract: contract}, GoodDataContractTransactor: GoodDataContractTransactor{contract: contract}, GoodDataContractFilterer: GoodDataContractFilterer{contract: contract}}, nil
}

// NewGoodDataContractCaller creates a new read-only instance of GoodDataContract, bound to a specific deployed contract.
func NewGoodDataContractCaller(address common.Address, caller bind.ContractCaller) (*GoodDataContractCaller, error) {
	contract, err := bindGoodDataContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GoodDataContractCaller{contract: contract}, nil
}

// NewGoodDataContractTransactor creates a new write-only instance of GoodDataContract, bound to a specific deployed contract.
func NewGoodDataContractTransactor(address common.Address, transactor bind.ContractTransactor) (*GoodDataContractTransactor, error) {
	contract, err := bindGoodDataContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GoodDataContractTransactor{contract: contract}, nil
}

// NewGoodDataContractFilterer creates a new log filterer instance of GoodDataContract, bound to a specific deployed contract.
func NewGoodDataContractFilterer(address common.Address, filterer bind.ContractFilterer) (*GoodDataContractFilterer, error) {
	contract, err := bindGoodDataContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GoodDataContractFilterer{contract: contract}, nil
}

// bindGoodDataContract binds a generic wrapper to an already deployed contract.
func bindGoodDataContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GoodDataContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoodDataContract *GoodDataContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GoodDataContract.Contract.GoodDataContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoodDataContract *GoodDataContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoodDataContract.Contract.GoodDataContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoodDataContract *GoodDataContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoodDataContract.Contract.GoodDataContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoodDataContract *GoodDataContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GoodDataContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoodDataContract *GoodDataContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoodDataContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoodDataContract *GoodDataContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoodDataContract.Contract.contract.Transact(opts, method, params...)
}

// AddLog is a paid mutator transaction binding the contract method 0x403a10fe.
//
// Solidity: function addLog(bytes uuid, bytes payload) returns()
func (_GoodDataContract *GoodDataContractTransactor) AddLog(opts *bind.TransactOpts, uuid []byte, payload []byte) (*types.Transaction, error) {
	return _GoodDataContract.contract.Transact(opts, "addLog", uuid, payload)
}

// AddLog is a paid mutator transaction binding the contract method 0x403a10fe.
//
// Solidity: function addLog(bytes uuid, bytes payload) returns()
func (_GoodDataContract *GoodDataContractSession) AddLog(uuid []byte, payload []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.AddLog(&_GoodDataContract.TransactOpts, uuid, payload)
}

// AddLog is a paid mutator transaction binding the contract method 0x403a10fe.
//
// Solidity: function addLog(bytes uuid, bytes payload) returns()
func (_GoodDataContract *GoodDataContractTransactorSession) AddLog(uuid []byte, payload []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.AddLog(&_GoodDataContract.TransactOpts, uuid, payload)
}

// AssignQueryToDOs is a paid mutator transaction binding the contract method 0xecd2cc22.
//
// Solidity: function assignQueryToDOs(bytes queryUUID, bytes doUUIDs) returns()
func (_GoodDataContract *GoodDataContractTransactor) AssignQueryToDOs(opts *bind.TransactOpts, queryUUID []byte, doUUIDs []byte) (*types.Transaction, error) {
	return _GoodDataContract.contract.Transact(opts, "assignQueryToDOs", queryUUID, doUUIDs)
}

// AssignQueryToDOs is a paid mutator transaction binding the contract method 0xecd2cc22.
//
// Solidity: function assignQueryToDOs(bytes queryUUID, bytes doUUIDs) returns()
func (_GoodDataContract *GoodDataContractSession) AssignQueryToDOs(queryUUID []byte, doUUIDs []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.AssignQueryToDOs(&_GoodDataContract.TransactOpts, queryUUID, doUUIDs)
}

// AssignQueryToDOs is a paid mutator transaction binding the contract method 0xecd2cc22.
//
// Solidity: function assignQueryToDOs(bytes queryUUID, bytes doUUIDs) returns()
func (_GoodDataContract *GoodDataContractTransactorSession) AssignQueryToDOs(queryUUID []byte, doUUIDs []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.AssignQueryToDOs(&_GoodDataContract.TransactOpts, queryUUID, doUUIDs)
}

// AssignQueryToMPCs is a paid mutator transaction binding the contract method 0xbc742822.
//
// Solidity: function assignQueryToMPCs(bytes queryUUID, bytes mpcUUIDs) returns()
func (_GoodDataContract *GoodDataContractTransactor) AssignQueryToMPCs(opts *bind.TransactOpts, queryUUID []byte, mpcUUIDs []byte) (*types.Transaction, error) {
	return _GoodDataContract.contract.Transact(opts, "assignQueryToMPCs", queryUUID, mpcUUIDs)
}

// AssignQueryToMPCs is a paid mutator transaction binding the contract method 0xbc742822.
//
// Solidity: function assignQueryToMPCs(bytes queryUUID, bytes mpcUUIDs) returns()
func (_GoodDataContract *GoodDataContractSession) AssignQueryToMPCs(queryUUID []byte, mpcUUIDs []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.AssignQueryToMPCs(&_GoodDataContract.TransactOpts, queryUUID, mpcUUIDs)
}

// AssignQueryToMPCs is a paid mutator transaction binding the contract method 0xbc742822.
//
// Solidity: function assignQueryToMPCs(bytes queryUUID, bytes mpcUUIDs) returns()
func (_GoodDataContract *GoodDataContractTransactorSession) AssignQueryToMPCs(queryUUID []byte, mpcUUIDs []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.AssignQueryToMPCs(&_GoodDataContract.TransactOpts, queryUUID, mpcUUIDs)
}

// FinishPrediction is a paid mutator transaction binding the contract method 0x5ff9ca6b.
//
// Solidity: function finishPrediction(bytes predictionUUID) returns()
func (_GoodDataContract *GoodDataContractTransactor) FinishPrediction(opts *bind.TransactOpts, predictionUUID []byte) (*types.Transaction, error) {
	return _GoodDataContract.contract.Transact(opts, "finishPrediction", predictionUUID)
}

// FinishPrediction is a paid mutator transaction binding the contract method 0x5ff9ca6b.
//
// Solidity: function finishPrediction(bytes predictionUUID) returns()
func (_GoodDataContract *GoodDataContractSession) FinishPrediction(predictionUUID []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.FinishPrediction(&_GoodDataContract.TransactOpts, predictionUUID)
}

// FinishPrediction is a paid mutator transaction binding the contract method 0x5ff9ca6b.
//
// Solidity: function finishPrediction(bytes predictionUUID) returns()
func (_GoodDataContract *GoodDataContractTransactorSession) FinishPrediction(predictionUUID []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.FinishPrediction(&_GoodDataContract.TransactOpts, predictionUUID)
}

// FinishQuery is a paid mutator transaction binding the contract method 0x2b30e748.
//
// Solidity: function finishQuery(bytes queryUUID) returns()
func (_GoodDataContract *GoodDataContractTransactor) FinishQuery(opts *bind.TransactOpts, queryUUID []byte) (*types.Transaction, error) {
	return _GoodDataContract.contract.Transact(opts, "finishQuery", queryUUID)
}

// FinishQuery is a paid mutator transaction binding the contract method 0x2b30e748.
//
// Solidity: function finishQuery(bytes queryUUID) returns()
func (_GoodDataContract *GoodDataContractSession) FinishQuery(queryUUID []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.FinishQuery(&_GoodDataContract.TransactOpts, queryUUID)
}

// FinishQuery is a paid mutator transaction binding the contract method 0x2b30e748.
//
// Solidity: function finishQuery(bytes queryUUID) returns()
func (_GoodDataContract *GoodDataContractTransactorSession) FinishQuery(queryUUID []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.FinishQuery(&_GoodDataContract.TransactOpts, queryUUID)
}

// FinishQueryInOneDO is a paid mutator transaction binding the contract method 0x98bccea0.
//
// Solidity: function finishQueryInOneDO(bytes queryUUID, bytes doUUID) returns()
func (_GoodDataContract *GoodDataContractTransactor) FinishQueryInOneDO(opts *bind.TransactOpts, queryUUID []byte, doUUID []byte) (*types.Transaction, error) {
	return _GoodDataContract.contract.Transact(opts, "finishQueryInOneDO", queryUUID, doUUID)
}

// FinishQueryInOneDO is a paid mutator transaction binding the contract method 0x98bccea0.
//
// Solidity: function finishQueryInOneDO(bytes queryUUID, bytes doUUID) returns()
func (_GoodDataContract *GoodDataContractSession) FinishQueryInOneDO(queryUUID []byte, doUUID []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.FinishQueryInOneDO(&_GoodDataContract.TransactOpts, queryUUID, doUUID)
}

// FinishQueryInOneDO is a paid mutator transaction binding the contract method 0x98bccea0.
//
// Solidity: function finishQueryInOneDO(bytes queryUUID, bytes doUUID) returns()
func (_GoodDataContract *GoodDataContractTransactorSession) FinishQueryInOneDO(queryUUID []byte, doUUID []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.FinishQueryInOneDO(&_GoodDataContract.TransactOpts, queryUUID, doUUID)
}

// RegisterDO is a paid mutator transaction binding the contract method 0xa15b8bf6.
//
// Solidity: function registerDO(bytes uuid, bytes publicKey) returns()
func (_GoodDataContract *GoodDataContractTransactor) RegisterDO(opts *bind.TransactOpts, uuid []byte, publicKey []byte) (*types.Transaction, error) {
	return _GoodDataContract.contract.Transact(opts, "registerDO", uuid, publicKey)
}

// RegisterDO is a paid mutator transaction binding the contract method 0xa15b8bf6.
//
// Solidity: function registerDO(bytes uuid, bytes publicKey) returns()
func (_GoodDataContract *GoodDataContractSession) RegisterDO(uuid []byte, publicKey []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.RegisterDO(&_GoodDataContract.TransactOpts, uuid, publicKey)
}

// RegisterDO is a paid mutator transaction binding the contract method 0xa15b8bf6.
//
// Solidity: function registerDO(bytes uuid, bytes publicKey) returns()
func (_GoodDataContract *GoodDataContractTransactorSession) RegisterDO(uuid []byte, publicKey []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.RegisterDO(&_GoodDataContract.TransactOpts, uuid, publicKey)
}

// RegisterMPC is a paid mutator transaction binding the contract method 0x064a7d3c.
//
// Solidity: function registerMPC(bytes uuid, bytes publicKey) returns()
func (_GoodDataContract *GoodDataContractTransactor) RegisterMPC(opts *bind.TransactOpts, uuid []byte, publicKey []byte) (*types.Transaction, error) {
	return _GoodDataContract.contract.Transact(opts, "registerMPC", uuid, publicKey)
}

// RegisterMPC is a paid mutator transaction binding the contract method 0x064a7d3c.
//
// Solidity: function registerMPC(bytes uuid, bytes publicKey) returns()
func (_GoodDataContract *GoodDataContractSession) RegisterMPC(uuid []byte, publicKey []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.RegisterMPC(&_GoodDataContract.TransactOpts, uuid, publicKey)
}

// RegisterMPC is a paid mutator transaction binding the contract method 0x064a7d3c.
//
// Solidity: function registerMPC(bytes uuid, bytes publicKey) returns()
func (_GoodDataContract *GoodDataContractTransactorSession) RegisterMPC(uuid []byte, publicKey []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.RegisterMPC(&_GoodDataContract.TransactOpts, uuid, publicKey)
}

// RegisterQC is a paid mutator transaction binding the contract method 0x0ab0421e.
//
// Solidity: function registerQC(bytes uuid, bytes publicKey) returns()
func (_GoodDataContract *GoodDataContractTransactor) RegisterQC(opts *bind.TransactOpts, uuid []byte, publicKey []byte) (*types.Transaction, error) {
	return _GoodDataContract.contract.Transact(opts, "registerQC", uuid, publicKey)
}

// RegisterQC is a paid mutator transaction binding the contract method 0x0ab0421e.
//
// Solidity: function registerQC(bytes uuid, bytes publicKey) returns()
func (_GoodDataContract *GoodDataContractSession) RegisterQC(uuid []byte, publicKey []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.RegisterQC(&_GoodDataContract.TransactOpts, uuid, publicKey)
}

// RegisterQC is a paid mutator transaction binding the contract method 0x0ab0421e.
//
// Solidity: function registerQC(bytes uuid, bytes publicKey) returns()
func (_GoodDataContract *GoodDataContractTransactorSession) RegisterQC(uuid []byte, publicKey []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.RegisterQC(&_GoodDataContract.TransactOpts, uuid, publicKey)
}

// SubmitPrediction is a paid mutator transaction binding the contract method 0x98e1d488.
//
// Solidity: function submitPrediction(bytes queryUUID, bytes predictionUUID) returns()
func (_GoodDataContract *GoodDataContractTransactor) SubmitPrediction(opts *bind.TransactOpts, queryUUID []byte, predictionUUID []byte) (*types.Transaction, error) {
	return _GoodDataContract.contract.Transact(opts, "submitPrediction", queryUUID, predictionUUID)
}

// SubmitPrediction is a paid mutator transaction binding the contract method 0x98e1d488.
//
// Solidity: function submitPrediction(bytes queryUUID, bytes predictionUUID) returns()
func (_GoodDataContract *GoodDataContractSession) SubmitPrediction(queryUUID []byte, predictionUUID []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.SubmitPrediction(&_GoodDataContract.TransactOpts, queryUUID, predictionUUID)
}

// SubmitPrediction is a paid mutator transaction binding the contract method 0x98e1d488.
//
// Solidity: function submitPrediction(bytes queryUUID, bytes predictionUUID) returns()
func (_GoodDataContract *GoodDataContractTransactorSession) SubmitPrediction(queryUUID []byte, predictionUUID []byte) (*types.Transaction, error) {
	return _GoodDataContract.Contract.SubmitPrediction(&_GoodDataContract.TransactOpts, queryUUID, predictionUUID)
}

// SubmitQuery is a paid mutator transaction binding the contract method 0x16aa61e2.
//
// Solidity: function submitQuery(bytes queryUUID, bytes qcUUID, int64 dataSet, uint256 fees) returns()
func (_GoodDataContract *GoodDataContractTransactor) SubmitQuery(opts *bind.TransactOpts, queryUUID []byte, qcUUID []byte, dataSet int64, fees *big.Int) (*types.Transaction, error) {
	return _GoodDataContract.contract.Transact(opts, "submitQuery", queryUUID, qcUUID, dataSet, fees)
}

// SubmitQuery is a paid mutator transaction binding the contract method 0x16aa61e2.
//
// Solidity: function submitQuery(bytes queryUUID, bytes qcUUID, int64 dataSet, uint256 fees) returns()
func (_GoodDataContract *GoodDataContractSession) SubmitQuery(queryUUID []byte, qcUUID []byte, dataSet int64, fees *big.Int) (*types.Transaction, error) {
	return _GoodDataContract.Contract.SubmitQuery(&_GoodDataContract.TransactOpts, queryUUID, qcUUID, dataSet, fees)
}

// SubmitQuery is a paid mutator transaction binding the contract method 0x16aa61e2.
//
// Solidity: function submitQuery(bytes queryUUID, bytes qcUUID, int64 dataSet, uint256 fees) returns()
func (_GoodDataContract *GoodDataContractTransactorSession) SubmitQuery(queryUUID []byte, qcUUID []byte, dataSet int64, fees *big.Int) (*types.Transaction, error) {
	return _GoodDataContract.Contract.SubmitQuery(&_GoodDataContract.TransactOpts, queryUUID, qcUUID, dataSet, fees)
}

// GoodDataContractDORegisteredIterator is returned from FilterDORegistered and is used to iterate over the raw logs and unpacked data for DORegistered events raised by the GoodDataContract contract.
type GoodDataContractDORegisteredIterator struct {
	Event *GoodDataContractDORegistered // Event containing the contract specifics and raw log

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
func (it *GoodDataContractDORegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodDataContractDORegistered)
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
		it.Event = new(GoodDataContractDORegistered)
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
func (it *GoodDataContractDORegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodDataContractDORegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodDataContractDORegistered represents a DORegistered event raised by the GoodDataContract contract.
type GoodDataContractDORegistered struct {
	Uuid      []byte
	PublicKey []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDORegistered is a free log retrieval operation binding the contract event 0x6d7f211c63019c81b6d1310ff9eb2ca9593e64cb007933fc42f72528e0928838.
//
// Solidity: event DORegistered(bytes uuid, bytes publicKey)
func (_GoodDataContract *GoodDataContractFilterer) FilterDORegistered(opts *bind.FilterOpts) (*GoodDataContractDORegisteredIterator, error) {

	logs, sub, err := _GoodDataContract.contract.FilterLogs(opts, "DORegistered")
	if err != nil {
		return nil, err
	}
	return &GoodDataContractDORegisteredIterator{contract: _GoodDataContract.contract, event: "DORegistered", logs: logs, sub: sub}, nil
}

// WatchDORegistered is a free log subscription operation binding the contract event 0x6d7f211c63019c81b6d1310ff9eb2ca9593e64cb007933fc42f72528e0928838.
//
// Solidity: event DORegistered(bytes uuid, bytes publicKey)
func (_GoodDataContract *GoodDataContractFilterer) WatchDORegistered(opts *bind.WatchOpts, sink chan<- *GoodDataContractDORegistered) (event.Subscription, error) {

	logs, sub, err := _GoodDataContract.contract.WatchLogs(opts, "DORegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodDataContractDORegistered)
				if err := _GoodDataContract.contract.UnpackLog(event, "DORegistered", log); err != nil {
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

// ParseDORegistered is a log parse operation binding the contract event 0x6d7f211c63019c81b6d1310ff9eb2ca9593e64cb007933fc42f72528e0928838.
//
// Solidity: event DORegistered(bytes uuid, bytes publicKey)
func (_GoodDataContract *GoodDataContractFilterer) ParseDORegistered(log types.Log) (*GoodDataContractDORegistered, error) {
	event := new(GoodDataContractDORegistered)
	if err := _GoodDataContract.contract.UnpackLog(event, "DORegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoodDataContractDOsOfQueryIterator is returned from FilterDOsOfQuery and is used to iterate over the raw logs and unpacked data for DOsOfQuery events raised by the GoodDataContract contract.
type GoodDataContractDOsOfQueryIterator struct {
	Event *GoodDataContractDOsOfQuery // Event containing the contract specifics and raw log

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
func (it *GoodDataContractDOsOfQueryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodDataContractDOsOfQuery)
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
		it.Event = new(GoodDataContractDOsOfQuery)
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
func (it *GoodDataContractDOsOfQueryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodDataContractDOsOfQueryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodDataContractDOsOfQuery represents a DOsOfQuery event raised by the GoodDataContract contract.
type GoodDataContractDOsOfQuery struct {
	QueryUUID []byte
	DoUUIDs   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDOsOfQuery is a free log retrieval operation binding the contract event 0x5afa6c090ab20346a4ace7c247aa417fdafbdfd14417ade41975f8f59af0f26f.
//
// Solidity: event DOsOfQuery(bytes queryUUID, bytes doUUIDs)
func (_GoodDataContract *GoodDataContractFilterer) FilterDOsOfQuery(opts *bind.FilterOpts) (*GoodDataContractDOsOfQueryIterator, error) {

	logs, sub, err := _GoodDataContract.contract.FilterLogs(opts, "DOsOfQuery")
	if err != nil {
		return nil, err
	}
	return &GoodDataContractDOsOfQueryIterator{contract: _GoodDataContract.contract, event: "DOsOfQuery", logs: logs, sub: sub}, nil
}

// WatchDOsOfQuery is a free log subscription operation binding the contract event 0x5afa6c090ab20346a4ace7c247aa417fdafbdfd14417ade41975f8f59af0f26f.
//
// Solidity: event DOsOfQuery(bytes queryUUID, bytes doUUIDs)
func (_GoodDataContract *GoodDataContractFilterer) WatchDOsOfQuery(opts *bind.WatchOpts, sink chan<- *GoodDataContractDOsOfQuery) (event.Subscription, error) {

	logs, sub, err := _GoodDataContract.contract.WatchLogs(opts, "DOsOfQuery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodDataContractDOsOfQuery)
				if err := _GoodDataContract.contract.UnpackLog(event, "DOsOfQuery", log); err != nil {
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

// ParseDOsOfQuery is a log parse operation binding the contract event 0x5afa6c090ab20346a4ace7c247aa417fdafbdfd14417ade41975f8f59af0f26f.
//
// Solidity: event DOsOfQuery(bytes queryUUID, bytes doUUIDs)
func (_GoodDataContract *GoodDataContractFilterer) ParseDOsOfQuery(log types.Log) (*GoodDataContractDOsOfQuery, error) {
	event := new(GoodDataContractDOsOfQuery)
	if err := _GoodDataContract.contract.UnpackLog(event, "DOsOfQuery", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoodDataContractLogAddedIterator is returned from FilterLogAdded and is used to iterate over the raw logs and unpacked data for LogAdded events raised by the GoodDataContract contract.
type GoodDataContractLogAddedIterator struct {
	Event *GoodDataContractLogAdded // Event containing the contract specifics and raw log

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
func (it *GoodDataContractLogAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodDataContractLogAdded)
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
		it.Event = new(GoodDataContractLogAdded)
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
func (it *GoodDataContractLogAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodDataContractLogAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodDataContractLogAdded represents a LogAdded event raised by the GoodDataContract contract.
type GoodDataContractLogAdded struct {
	Uuid    []byte
	Payload []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogAdded is a free log retrieval operation binding the contract event 0x2a105b504873e6cae2ea9e9eba00d2f1ded4557ca784f61d4729458967139860.
//
// Solidity: event LogAdded(bytes uuid, bytes payload)
func (_GoodDataContract *GoodDataContractFilterer) FilterLogAdded(opts *bind.FilterOpts) (*GoodDataContractLogAddedIterator, error) {

	logs, sub, err := _GoodDataContract.contract.FilterLogs(opts, "LogAdded")
	if err != nil {
		return nil, err
	}
	return &GoodDataContractLogAddedIterator{contract: _GoodDataContract.contract, event: "LogAdded", logs: logs, sub: sub}, nil
}

// WatchLogAdded is a free log subscription operation binding the contract event 0x2a105b504873e6cae2ea9e9eba00d2f1ded4557ca784f61d4729458967139860.
//
// Solidity: event LogAdded(bytes uuid, bytes payload)
func (_GoodDataContract *GoodDataContractFilterer) WatchLogAdded(opts *bind.WatchOpts, sink chan<- *GoodDataContractLogAdded) (event.Subscription, error) {

	logs, sub, err := _GoodDataContract.contract.WatchLogs(opts, "LogAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodDataContractLogAdded)
				if err := _GoodDataContract.contract.UnpackLog(event, "LogAdded", log); err != nil {
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

// ParseLogAdded is a log parse operation binding the contract event 0x2a105b504873e6cae2ea9e9eba00d2f1ded4557ca784f61d4729458967139860.
//
// Solidity: event LogAdded(bytes uuid, bytes payload)
func (_GoodDataContract *GoodDataContractFilterer) ParseLogAdded(log types.Log) (*GoodDataContractLogAdded, error) {
	event := new(GoodDataContractLogAdded)
	if err := _GoodDataContract.contract.UnpackLog(event, "LogAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoodDataContractMPCRegisteredIterator is returned from FilterMPCRegistered and is used to iterate over the raw logs and unpacked data for MPCRegistered events raised by the GoodDataContract contract.
type GoodDataContractMPCRegisteredIterator struct {
	Event *GoodDataContractMPCRegistered // Event containing the contract specifics and raw log

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
func (it *GoodDataContractMPCRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodDataContractMPCRegistered)
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
		it.Event = new(GoodDataContractMPCRegistered)
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
func (it *GoodDataContractMPCRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodDataContractMPCRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodDataContractMPCRegistered represents a MPCRegistered event raised by the GoodDataContract contract.
type GoodDataContractMPCRegistered struct {
	Uuid      []byte
	PublicKey []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMPCRegistered is a free log retrieval operation binding the contract event 0x7e95d85d8e6ec6cb0797372a3549bd1c3bd2de9e51f59e53499aca8877700534.
//
// Solidity: event MPCRegistered(bytes uuid, bytes publicKey)
func (_GoodDataContract *GoodDataContractFilterer) FilterMPCRegistered(opts *bind.FilterOpts) (*GoodDataContractMPCRegisteredIterator, error) {

	logs, sub, err := _GoodDataContract.contract.FilterLogs(opts, "MPCRegistered")
	if err != nil {
		return nil, err
	}
	return &GoodDataContractMPCRegisteredIterator{contract: _GoodDataContract.contract, event: "MPCRegistered", logs: logs, sub: sub}, nil
}

// WatchMPCRegistered is a free log subscription operation binding the contract event 0x7e95d85d8e6ec6cb0797372a3549bd1c3bd2de9e51f59e53499aca8877700534.
//
// Solidity: event MPCRegistered(bytes uuid, bytes publicKey)
func (_GoodDataContract *GoodDataContractFilterer) WatchMPCRegistered(opts *bind.WatchOpts, sink chan<- *GoodDataContractMPCRegistered) (event.Subscription, error) {

	logs, sub, err := _GoodDataContract.contract.WatchLogs(opts, "MPCRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodDataContractMPCRegistered)
				if err := _GoodDataContract.contract.UnpackLog(event, "MPCRegistered", log); err != nil {
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

// ParseMPCRegistered is a log parse operation binding the contract event 0x7e95d85d8e6ec6cb0797372a3549bd1c3bd2de9e51f59e53499aca8877700534.
//
// Solidity: event MPCRegistered(bytes uuid, bytes publicKey)
func (_GoodDataContract *GoodDataContractFilterer) ParseMPCRegistered(log types.Log) (*GoodDataContractMPCRegistered, error) {
	event := new(GoodDataContractMPCRegistered)
	if err := _GoodDataContract.contract.UnpackLog(event, "MPCRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoodDataContractMPCsOfQueryIterator is returned from FilterMPCsOfQuery and is used to iterate over the raw logs and unpacked data for MPCsOfQuery events raised by the GoodDataContract contract.
type GoodDataContractMPCsOfQueryIterator struct {
	Event *GoodDataContractMPCsOfQuery // Event containing the contract specifics and raw log

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
func (it *GoodDataContractMPCsOfQueryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodDataContractMPCsOfQuery)
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
		it.Event = new(GoodDataContractMPCsOfQuery)
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
func (it *GoodDataContractMPCsOfQueryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodDataContractMPCsOfQueryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodDataContractMPCsOfQuery represents a MPCsOfQuery event raised by the GoodDataContract contract.
type GoodDataContractMPCsOfQuery struct {
	QueryUUID []byte
	MpcUUIDs  []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMPCsOfQuery is a free log retrieval operation binding the contract event 0x5e85fb60a99bcdfd3a06c06e11221730f946b1cb94d93ff6a9b0792617a849b8.
//
// Solidity: event MPCsOfQuery(bytes queryUUID, bytes mpcUUIDs)
func (_GoodDataContract *GoodDataContractFilterer) FilterMPCsOfQuery(opts *bind.FilterOpts) (*GoodDataContractMPCsOfQueryIterator, error) {

	logs, sub, err := _GoodDataContract.contract.FilterLogs(opts, "MPCsOfQuery")
	if err != nil {
		return nil, err
	}
	return &GoodDataContractMPCsOfQueryIterator{contract: _GoodDataContract.contract, event: "MPCsOfQuery", logs: logs, sub: sub}, nil
}

// WatchMPCsOfQuery is a free log subscription operation binding the contract event 0x5e85fb60a99bcdfd3a06c06e11221730f946b1cb94d93ff6a9b0792617a849b8.
//
// Solidity: event MPCsOfQuery(bytes queryUUID, bytes mpcUUIDs)
func (_GoodDataContract *GoodDataContractFilterer) WatchMPCsOfQuery(opts *bind.WatchOpts, sink chan<- *GoodDataContractMPCsOfQuery) (event.Subscription, error) {

	logs, sub, err := _GoodDataContract.contract.WatchLogs(opts, "MPCsOfQuery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodDataContractMPCsOfQuery)
				if err := _GoodDataContract.contract.UnpackLog(event, "MPCsOfQuery", log); err != nil {
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

// ParseMPCsOfQuery is a log parse operation binding the contract event 0x5e85fb60a99bcdfd3a06c06e11221730f946b1cb94d93ff6a9b0792617a849b8.
//
// Solidity: event MPCsOfQuery(bytes queryUUID, bytes mpcUUIDs)
func (_GoodDataContract *GoodDataContractFilterer) ParseMPCsOfQuery(log types.Log) (*GoodDataContractMPCsOfQuery, error) {
	event := new(GoodDataContractMPCsOfQuery)
	if err := _GoodDataContract.contract.UnpackLog(event, "MPCsOfQuery", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoodDataContractPredictionFinishedIterator is returned from FilterPredictionFinished and is used to iterate over the raw logs and unpacked data for PredictionFinished events raised by the GoodDataContract contract.
type GoodDataContractPredictionFinishedIterator struct {
	Event *GoodDataContractPredictionFinished // Event containing the contract specifics and raw log

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
func (it *GoodDataContractPredictionFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodDataContractPredictionFinished)
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
		it.Event = new(GoodDataContractPredictionFinished)
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
func (it *GoodDataContractPredictionFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodDataContractPredictionFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodDataContractPredictionFinished represents a PredictionFinished event raised by the GoodDataContract contract.
type GoodDataContractPredictionFinished struct {
	PredictionUUID []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPredictionFinished is a free log retrieval operation binding the contract event 0xae1e63742daf7163b4edff7ae9bbed38a21d474b7335f4dfb53440ddfde151a3.
//
// Solidity: event PredictionFinished(bytes predictionUUID)
func (_GoodDataContract *GoodDataContractFilterer) FilterPredictionFinished(opts *bind.FilterOpts) (*GoodDataContractPredictionFinishedIterator, error) {

	logs, sub, err := _GoodDataContract.contract.FilterLogs(opts, "PredictionFinished")
	if err != nil {
		return nil, err
	}
	return &GoodDataContractPredictionFinishedIterator{contract: _GoodDataContract.contract, event: "PredictionFinished", logs: logs, sub: sub}, nil
}

// WatchPredictionFinished is a free log subscription operation binding the contract event 0xae1e63742daf7163b4edff7ae9bbed38a21d474b7335f4dfb53440ddfde151a3.
//
// Solidity: event PredictionFinished(bytes predictionUUID)
func (_GoodDataContract *GoodDataContractFilterer) WatchPredictionFinished(opts *bind.WatchOpts, sink chan<- *GoodDataContractPredictionFinished) (event.Subscription, error) {

	logs, sub, err := _GoodDataContract.contract.WatchLogs(opts, "PredictionFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodDataContractPredictionFinished)
				if err := _GoodDataContract.contract.UnpackLog(event, "PredictionFinished", log); err != nil {
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

// ParsePredictionFinished is a log parse operation binding the contract event 0xae1e63742daf7163b4edff7ae9bbed38a21d474b7335f4dfb53440ddfde151a3.
//
// Solidity: event PredictionFinished(bytes predictionUUID)
func (_GoodDataContract *GoodDataContractFilterer) ParsePredictionFinished(log types.Log) (*GoodDataContractPredictionFinished, error) {
	event := new(GoodDataContractPredictionFinished)
	if err := _GoodDataContract.contract.UnpackLog(event, "PredictionFinished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoodDataContractPredictionSubmittedIterator is returned from FilterPredictionSubmitted and is used to iterate over the raw logs and unpacked data for PredictionSubmitted events raised by the GoodDataContract contract.
type GoodDataContractPredictionSubmittedIterator struct {
	Event *GoodDataContractPredictionSubmitted // Event containing the contract specifics and raw log

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
func (it *GoodDataContractPredictionSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodDataContractPredictionSubmitted)
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
		it.Event = new(GoodDataContractPredictionSubmitted)
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
func (it *GoodDataContractPredictionSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodDataContractPredictionSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodDataContractPredictionSubmitted represents a PredictionSubmitted event raised by the GoodDataContract contract.
type GoodDataContractPredictionSubmitted struct {
	QueryUUID      []byte
	PredictionUUID []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPredictionSubmitted is a free log retrieval operation binding the contract event 0xae4dd2ed761989d525646c3d29fae9a4e1acccce89882640dd73896590fa4d50.
//
// Solidity: event PredictionSubmitted(bytes queryUUID, bytes predictionUUID)
func (_GoodDataContract *GoodDataContractFilterer) FilterPredictionSubmitted(opts *bind.FilterOpts) (*GoodDataContractPredictionSubmittedIterator, error) {

	logs, sub, err := _GoodDataContract.contract.FilterLogs(opts, "PredictionSubmitted")
	if err != nil {
		return nil, err
	}
	return &GoodDataContractPredictionSubmittedIterator{contract: _GoodDataContract.contract, event: "PredictionSubmitted", logs: logs, sub: sub}, nil
}

// WatchPredictionSubmitted is a free log subscription operation binding the contract event 0xae4dd2ed761989d525646c3d29fae9a4e1acccce89882640dd73896590fa4d50.
//
// Solidity: event PredictionSubmitted(bytes queryUUID, bytes predictionUUID)
func (_GoodDataContract *GoodDataContractFilterer) WatchPredictionSubmitted(opts *bind.WatchOpts, sink chan<- *GoodDataContractPredictionSubmitted) (event.Subscription, error) {

	logs, sub, err := _GoodDataContract.contract.WatchLogs(opts, "PredictionSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodDataContractPredictionSubmitted)
				if err := _GoodDataContract.contract.UnpackLog(event, "PredictionSubmitted", log); err != nil {
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

// ParsePredictionSubmitted is a log parse operation binding the contract event 0xae4dd2ed761989d525646c3d29fae9a4e1acccce89882640dd73896590fa4d50.
//
// Solidity: event PredictionSubmitted(bytes queryUUID, bytes predictionUUID)
func (_GoodDataContract *GoodDataContractFilterer) ParsePredictionSubmitted(log types.Log) (*GoodDataContractPredictionSubmitted, error) {
	event := new(GoodDataContractPredictionSubmitted)
	if err := _GoodDataContract.contract.UnpackLog(event, "PredictionSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoodDataContractQCRegisteredIterator is returned from FilterQCRegistered and is used to iterate over the raw logs and unpacked data for QCRegistered events raised by the GoodDataContract contract.
type GoodDataContractQCRegisteredIterator struct {
	Event *GoodDataContractQCRegistered // Event containing the contract specifics and raw log

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
func (it *GoodDataContractQCRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodDataContractQCRegistered)
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
		it.Event = new(GoodDataContractQCRegistered)
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
func (it *GoodDataContractQCRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodDataContractQCRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodDataContractQCRegistered represents a QCRegistered event raised by the GoodDataContract contract.
type GoodDataContractQCRegistered struct {
	Uuid      []byte
	PublicKey []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQCRegistered is a free log retrieval operation binding the contract event 0x7867aecee1277642c50d987030f5c42aa8d302249a4a2076c060abb514052610.
//
// Solidity: event QCRegistered(bytes uuid, bytes publicKey)
func (_GoodDataContract *GoodDataContractFilterer) FilterQCRegistered(opts *bind.FilterOpts) (*GoodDataContractQCRegisteredIterator, error) {

	logs, sub, err := _GoodDataContract.contract.FilterLogs(opts, "QCRegistered")
	if err != nil {
		return nil, err
	}
	return &GoodDataContractQCRegisteredIterator{contract: _GoodDataContract.contract, event: "QCRegistered", logs: logs, sub: sub}, nil
}

// WatchQCRegistered is a free log subscription operation binding the contract event 0x7867aecee1277642c50d987030f5c42aa8d302249a4a2076c060abb514052610.
//
// Solidity: event QCRegistered(bytes uuid, bytes publicKey)
func (_GoodDataContract *GoodDataContractFilterer) WatchQCRegistered(opts *bind.WatchOpts, sink chan<- *GoodDataContractQCRegistered) (event.Subscription, error) {

	logs, sub, err := _GoodDataContract.contract.WatchLogs(opts, "QCRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodDataContractQCRegistered)
				if err := _GoodDataContract.contract.UnpackLog(event, "QCRegistered", log); err != nil {
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

// ParseQCRegistered is a log parse operation binding the contract event 0x7867aecee1277642c50d987030f5c42aa8d302249a4a2076c060abb514052610.
//
// Solidity: event QCRegistered(bytes uuid, bytes publicKey)
func (_GoodDataContract *GoodDataContractFilterer) ParseQCRegistered(log types.Log) (*GoodDataContractQCRegistered, error) {
	event := new(GoodDataContractQCRegistered)
	if err := _GoodDataContract.contract.UnpackLog(event, "QCRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoodDataContractQueryFinishedIterator is returned from FilterQueryFinished and is used to iterate over the raw logs and unpacked data for QueryFinished events raised by the GoodDataContract contract.
type GoodDataContractQueryFinishedIterator struct {
	Event *GoodDataContractQueryFinished // Event containing the contract specifics and raw log

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
func (it *GoodDataContractQueryFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodDataContractQueryFinished)
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
		it.Event = new(GoodDataContractQueryFinished)
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
func (it *GoodDataContractQueryFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodDataContractQueryFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodDataContractQueryFinished represents a QueryFinished event raised by the GoodDataContract contract.
type GoodDataContractQueryFinished struct {
	QueryUUID []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQueryFinished is a free log retrieval operation binding the contract event 0xec239bd768b84bf03a45fd513ce7770b395f182525a856476e534cfbc45c030a.
//
// Solidity: event QueryFinished(bytes queryUUID)
func (_GoodDataContract *GoodDataContractFilterer) FilterQueryFinished(opts *bind.FilterOpts) (*GoodDataContractQueryFinishedIterator, error) {

	logs, sub, err := _GoodDataContract.contract.FilterLogs(opts, "QueryFinished")
	if err != nil {
		return nil, err
	}
	return &GoodDataContractQueryFinishedIterator{contract: _GoodDataContract.contract, event: "QueryFinished", logs: logs, sub: sub}, nil
}

// WatchQueryFinished is a free log subscription operation binding the contract event 0xec239bd768b84bf03a45fd513ce7770b395f182525a856476e534cfbc45c030a.
//
// Solidity: event QueryFinished(bytes queryUUID)
func (_GoodDataContract *GoodDataContractFilterer) WatchQueryFinished(opts *bind.WatchOpts, sink chan<- *GoodDataContractQueryFinished) (event.Subscription, error) {

	logs, sub, err := _GoodDataContract.contract.WatchLogs(opts, "QueryFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodDataContractQueryFinished)
				if err := _GoodDataContract.contract.UnpackLog(event, "QueryFinished", log); err != nil {
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

// ParseQueryFinished is a log parse operation binding the contract event 0xec239bd768b84bf03a45fd513ce7770b395f182525a856476e534cfbc45c030a.
//
// Solidity: event QueryFinished(bytes queryUUID)
func (_GoodDataContract *GoodDataContractFilterer) ParseQueryFinished(log types.Log) (*GoodDataContractQueryFinished, error) {
	event := new(GoodDataContractQueryFinished)
	if err := _GoodDataContract.contract.UnpackLog(event, "QueryFinished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoodDataContractQueryFinishedInOneDOIterator is returned from FilterQueryFinishedInOneDO and is used to iterate over the raw logs and unpacked data for QueryFinishedInOneDO events raised by the GoodDataContract contract.
type GoodDataContractQueryFinishedInOneDOIterator struct {
	Event *GoodDataContractQueryFinishedInOneDO // Event containing the contract specifics and raw log

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
func (it *GoodDataContractQueryFinishedInOneDOIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodDataContractQueryFinishedInOneDO)
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
		it.Event = new(GoodDataContractQueryFinishedInOneDO)
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
func (it *GoodDataContractQueryFinishedInOneDOIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodDataContractQueryFinishedInOneDOIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodDataContractQueryFinishedInOneDO represents a QueryFinishedInOneDO event raised by the GoodDataContract contract.
type GoodDataContractQueryFinishedInOneDO struct {
	QueryUUID []byte
	DoUUID    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQueryFinishedInOneDO is a free log retrieval operation binding the contract event 0xeec0bcfa77c271a48dafc4f835c9b0743ba44cdce9256e288dab09dcb98995db.
//
// Solidity: event QueryFinishedInOneDO(bytes queryUUID, bytes doUUID)
func (_GoodDataContract *GoodDataContractFilterer) FilterQueryFinishedInOneDO(opts *bind.FilterOpts) (*GoodDataContractQueryFinishedInOneDOIterator, error) {

	logs, sub, err := _GoodDataContract.contract.FilterLogs(opts, "QueryFinishedInOneDO")
	if err != nil {
		return nil, err
	}
	return &GoodDataContractQueryFinishedInOneDOIterator{contract: _GoodDataContract.contract, event: "QueryFinishedInOneDO", logs: logs, sub: sub}, nil
}

// WatchQueryFinishedInOneDO is a free log subscription operation binding the contract event 0xeec0bcfa77c271a48dafc4f835c9b0743ba44cdce9256e288dab09dcb98995db.
//
// Solidity: event QueryFinishedInOneDO(bytes queryUUID, bytes doUUID)
func (_GoodDataContract *GoodDataContractFilterer) WatchQueryFinishedInOneDO(opts *bind.WatchOpts, sink chan<- *GoodDataContractQueryFinishedInOneDO) (event.Subscription, error) {

	logs, sub, err := _GoodDataContract.contract.WatchLogs(opts, "QueryFinishedInOneDO")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodDataContractQueryFinishedInOneDO)
				if err := _GoodDataContract.contract.UnpackLog(event, "QueryFinishedInOneDO", log); err != nil {
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

// ParseQueryFinishedInOneDO is a log parse operation binding the contract event 0xeec0bcfa77c271a48dafc4f835c9b0743ba44cdce9256e288dab09dcb98995db.
//
// Solidity: event QueryFinishedInOneDO(bytes queryUUID, bytes doUUID)
func (_GoodDataContract *GoodDataContractFilterer) ParseQueryFinishedInOneDO(log types.Log) (*GoodDataContractQueryFinishedInOneDO, error) {
	event := new(GoodDataContractQueryFinishedInOneDO)
	if err := _GoodDataContract.contract.UnpackLog(event, "QueryFinishedInOneDO", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GoodDataContractQuerySubmittedIterator is returned from FilterQuerySubmitted and is used to iterate over the raw logs and unpacked data for QuerySubmitted events raised by the GoodDataContract contract.
type GoodDataContractQuerySubmittedIterator struct {
	Event *GoodDataContractQuerySubmitted // Event containing the contract specifics and raw log

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
func (it *GoodDataContractQuerySubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoodDataContractQuerySubmitted)
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
		it.Event = new(GoodDataContractQuerySubmitted)
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
func (it *GoodDataContractQuerySubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoodDataContractQuerySubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoodDataContractQuerySubmitted represents a QuerySubmitted event raised by the GoodDataContract contract.
type GoodDataContractQuerySubmitted struct {
	QueryUUID []byte
	QcUUID    []byte
	DataSet   int64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQuerySubmitted is a free log retrieval operation binding the contract event 0xe371a7b4ee8c7a088f0d1101193a4afb52e38d58fc487b77bc78ca7ddff4d6b3.
//
// Solidity: event QuerySubmitted(bytes queryUUID, bytes qcUUID, int64 dataSet)
func (_GoodDataContract *GoodDataContractFilterer) FilterQuerySubmitted(opts *bind.FilterOpts) (*GoodDataContractQuerySubmittedIterator, error) {

	logs, sub, err := _GoodDataContract.contract.FilterLogs(opts, "QuerySubmitted")
	if err != nil {
		return nil, err
	}
	return &GoodDataContractQuerySubmittedIterator{contract: _GoodDataContract.contract, event: "QuerySubmitted", logs: logs, sub: sub}, nil
}

// WatchQuerySubmitted is a free log subscription operation binding the contract event 0xe371a7b4ee8c7a088f0d1101193a4afb52e38d58fc487b77bc78ca7ddff4d6b3.
//
// Solidity: event QuerySubmitted(bytes queryUUID, bytes qcUUID, int64 dataSet)
func (_GoodDataContract *GoodDataContractFilterer) WatchQuerySubmitted(opts *bind.WatchOpts, sink chan<- *GoodDataContractQuerySubmitted) (event.Subscription, error) {

	logs, sub, err := _GoodDataContract.contract.WatchLogs(opts, "QuerySubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoodDataContractQuerySubmitted)
				if err := _GoodDataContract.contract.UnpackLog(event, "QuerySubmitted", log); err != nil {
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

// ParseQuerySubmitted is a log parse operation binding the contract event 0xe371a7b4ee8c7a088f0d1101193a4afb52e38d58fc487b77bc78ca7ddff4d6b3.
//
// Solidity: event QuerySubmitted(bytes queryUUID, bytes qcUUID, int64 dataSet)
func (_GoodDataContract *GoodDataContractFilterer) ParseQuerySubmitted(log types.Log) (*GoodDataContractQuerySubmitted, error) {
	event := new(GoodDataContractQuerySubmitted)
	if err := _GoodDataContract.contract.UnpackLog(event, "QuerySubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}
