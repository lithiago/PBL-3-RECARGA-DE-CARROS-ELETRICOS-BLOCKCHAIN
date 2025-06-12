// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ReservationLedgerReservation is an auto generated low-level Go binding around an user-defined struct.
type ReservationLedgerReservation struct {
	ExternalId   [32]byte
	CarId        string
	PostoId      string
	Timestamp    *big.Int
	ChargeAmount *big.Int
	Status       uint8
}

// ReservationLedgerMetaData contains all meta data concerning the ReservationLedger contract.
var ReservationLedgerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"cancelReservation\",\"inputs\":[{\"name\":\"_externalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeRecharge\",\"inputs\":[{\"name\":\"_externalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getReservation\",\"inputs\":[{\"name\":\"_externalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structReservationLedger.Reservation\",\"components\":[{\"name\":\"externalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"carId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"postoId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chargeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumReservationLedger.Status\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"makeReservation\",\"inputs\":[{\"name\":\"_externalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_carId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_postoId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processPayment\",\"inputs\":[{\"name\":\"_externalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reservationHistory\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reservations\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"externalId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"carId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"postoId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chargeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumReservationLedger.Status\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"PaymentProcessed\",\"inputs\":[{\"name\":\"externalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RechargeCompleted\",\"inputs\":[{\"name\":\"externalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReservationCancelled\",\"inputs\":[{\"name\":\"externalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReservationMade\",\"inputs\":[{\"name\":\"externalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"carId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"postoId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]",
}

// ReservationLedgerABI is the input ABI used to generate the binding from.
// Deprecated: Use ReservationLedgerMetaData.ABI instead.
var ReservationLedgerABI = ReservationLedgerMetaData.ABI

// ReservationLedger is an auto generated Go binding around an Ethereum contract.
type ReservationLedger struct {
	ReservationLedgerCaller     // Read-only binding to the contract
	ReservationLedgerTransactor // Write-only binding to the contract
	ReservationLedgerFilterer   // Log filterer for contract events
}

// ReservationLedgerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReservationLedgerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReservationLedgerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReservationLedgerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReservationLedgerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReservationLedgerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReservationLedgerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReservationLedgerSession struct {
	Contract     *ReservationLedger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReservationLedgerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReservationLedgerCallerSession struct {
	Contract *ReservationLedgerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ReservationLedgerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReservationLedgerTransactorSession struct {
	Contract     *ReservationLedgerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ReservationLedgerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReservationLedgerRaw struct {
	Contract *ReservationLedger // Generic contract binding to access the raw methods on
}

// ReservationLedgerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReservationLedgerCallerRaw struct {
	Contract *ReservationLedgerCaller // Generic read-only contract binding to access the raw methods on
}

// ReservationLedgerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReservationLedgerTransactorRaw struct {
	Contract *ReservationLedgerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReservationLedger creates a new instance of ReservationLedger, bound to a specific deployed contract.
func NewReservationLedger(address common.Address, backend bind.ContractBackend) (*ReservationLedger, error) {
	contract, err := bindReservationLedger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReservationLedger{ReservationLedgerCaller: ReservationLedgerCaller{contract: contract}, ReservationLedgerTransactor: ReservationLedgerTransactor{contract: contract}, ReservationLedgerFilterer: ReservationLedgerFilterer{contract: contract}}, nil
}

// NewReservationLedgerCaller creates a new read-only instance of ReservationLedger, bound to a specific deployed contract.
func NewReservationLedgerCaller(address common.Address, caller bind.ContractCaller) (*ReservationLedgerCaller, error) {
	contract, err := bindReservationLedger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReservationLedgerCaller{contract: contract}, nil
}

// NewReservationLedgerTransactor creates a new write-only instance of ReservationLedger, bound to a specific deployed contract.
func NewReservationLedgerTransactor(address common.Address, transactor bind.ContractTransactor) (*ReservationLedgerTransactor, error) {
	contract, err := bindReservationLedger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReservationLedgerTransactor{contract: contract}, nil
}

// NewReservationLedgerFilterer creates a new log filterer instance of ReservationLedger, bound to a specific deployed contract.
func NewReservationLedgerFilterer(address common.Address, filterer bind.ContractFilterer) (*ReservationLedgerFilterer, error) {
	contract, err := bindReservationLedger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReservationLedgerFilterer{contract: contract}, nil
}

// bindReservationLedger binds a generic wrapper to an already deployed contract.
func bindReservationLedger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReservationLedgerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReservationLedger *ReservationLedgerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReservationLedger.Contract.ReservationLedgerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReservationLedger *ReservationLedgerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReservationLedger.Contract.ReservationLedgerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReservationLedger *ReservationLedgerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReservationLedger.Contract.ReservationLedgerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReservationLedger *ReservationLedgerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReservationLedger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReservationLedger *ReservationLedgerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReservationLedger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReservationLedger *ReservationLedgerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReservationLedger.Contract.contract.Transact(opts, method, params...)
}

// GetReservation is a free data retrieval call binding the contract method 0xec68555f.
//
// Solidity: function getReservation(bytes32 _externalId) view returns((bytes32,string,string,uint256,uint256,uint8))
func (_ReservationLedger *ReservationLedgerCaller) GetReservation(opts *bind.CallOpts, _externalId [32]byte) (ReservationLedgerReservation, error) {
	var out []interface{}
	err := _ReservationLedger.contract.Call(opts, &out, "getReservation", _externalId)

	if err != nil {
		return *new(ReservationLedgerReservation), err
	}

	out0 := *abi.ConvertType(out[0], new(ReservationLedgerReservation)).(*ReservationLedgerReservation)

	return out0, err

}

// GetReservation is a free data retrieval call binding the contract method 0xec68555f.
//
// Solidity: function getReservation(bytes32 _externalId) view returns((bytes32,string,string,uint256,uint256,uint8))
func (_ReservationLedger *ReservationLedgerSession) GetReservation(_externalId [32]byte) (ReservationLedgerReservation, error) {
	return _ReservationLedger.Contract.GetReservation(&_ReservationLedger.CallOpts, _externalId)
}

// GetReservation is a free data retrieval call binding the contract method 0xec68555f.
//
// Solidity: function getReservation(bytes32 _externalId) view returns((bytes32,string,string,uint256,uint256,uint8))
func (_ReservationLedger *ReservationLedgerCallerSession) GetReservation(_externalId [32]byte) (ReservationLedgerReservation, error) {
	return _ReservationLedger.Contract.GetReservation(&_ReservationLedger.CallOpts, _externalId)
}

// ReservationHistory is a free data retrieval call binding the contract method 0xc47b88ab.
//
// Solidity: function reservationHistory(uint256 ) view returns(bytes32)
func (_ReservationLedger *ReservationLedgerCaller) ReservationHistory(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ReservationLedger.contract.Call(opts, &out, "reservationHistory", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReservationHistory is a free data retrieval call binding the contract method 0xc47b88ab.
//
// Solidity: function reservationHistory(uint256 ) view returns(bytes32)
func (_ReservationLedger *ReservationLedgerSession) ReservationHistory(arg0 *big.Int) ([32]byte, error) {
	return _ReservationLedger.Contract.ReservationHistory(&_ReservationLedger.CallOpts, arg0)
}

// ReservationHistory is a free data retrieval call binding the contract method 0xc47b88ab.
//
// Solidity: function reservationHistory(uint256 ) view returns(bytes32)
func (_ReservationLedger *ReservationLedgerCallerSession) ReservationHistory(arg0 *big.Int) ([32]byte, error) {
	return _ReservationLedger.Contract.ReservationHistory(&_ReservationLedger.CallOpts, arg0)
}

// Reservations is a free data retrieval call binding the contract method 0xa1b416f9.
//
// Solidity: function reservations(bytes32 ) view returns(bytes32 externalId, string carId, string postoId, uint256 timestamp, uint256 chargeAmount, uint8 status)
func (_ReservationLedger *ReservationLedgerCaller) Reservations(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ExternalId   [32]byte
	CarId        string
	PostoId      string
	Timestamp    *big.Int
	ChargeAmount *big.Int
	Status       uint8
}, error) {
	var out []interface{}
	err := _ReservationLedger.contract.Call(opts, &out, "reservations", arg0)

	outstruct := new(struct {
		ExternalId   [32]byte
		CarId        string
		PostoId      string
		Timestamp    *big.Int
		ChargeAmount *big.Int
		Status       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ExternalId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.CarId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.PostoId = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ChargeAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// Reservations is a free data retrieval call binding the contract method 0xa1b416f9.
//
// Solidity: function reservations(bytes32 ) view returns(bytes32 externalId, string carId, string postoId, uint256 timestamp, uint256 chargeAmount, uint8 status)
func (_ReservationLedger *ReservationLedgerSession) Reservations(arg0 [32]byte) (struct {
	ExternalId   [32]byte
	CarId        string
	PostoId      string
	Timestamp    *big.Int
	ChargeAmount *big.Int
	Status       uint8
}, error) {
	return _ReservationLedger.Contract.Reservations(&_ReservationLedger.CallOpts, arg0)
}

// Reservations is a free data retrieval call binding the contract method 0xa1b416f9.
//
// Solidity: function reservations(bytes32 ) view returns(bytes32 externalId, string carId, string postoId, uint256 timestamp, uint256 chargeAmount, uint8 status)
func (_ReservationLedger *ReservationLedgerCallerSession) Reservations(arg0 [32]byte) (struct {
	ExternalId   [32]byte
	CarId        string
	PostoId      string
	Timestamp    *big.Int
	ChargeAmount *big.Int
	Status       uint8
}, error) {
	return _ReservationLedger.Contract.Reservations(&_ReservationLedger.CallOpts, arg0)
}

// CancelReservation is a paid mutator transaction binding the contract method 0xa2d6ca20.
//
// Solidity: function cancelReservation(bytes32 _externalId) returns()
func (_ReservationLedger *ReservationLedgerTransactor) CancelReservation(opts *bind.TransactOpts, _externalId [32]byte) (*types.Transaction, error) {
	return _ReservationLedger.contract.Transact(opts, "cancelReservation", _externalId)
}

// CancelReservation is a paid mutator transaction binding the contract method 0xa2d6ca20.
//
// Solidity: function cancelReservation(bytes32 _externalId) returns()
func (_ReservationLedger *ReservationLedgerSession) CancelReservation(_externalId [32]byte) (*types.Transaction, error) {
	return _ReservationLedger.Contract.CancelReservation(&_ReservationLedger.TransactOpts, _externalId)
}

// CancelReservation is a paid mutator transaction binding the contract method 0xa2d6ca20.
//
// Solidity: function cancelReservation(bytes32 _externalId) returns()
func (_ReservationLedger *ReservationLedgerTransactorSession) CancelReservation(_externalId [32]byte) (*types.Transaction, error) {
	return _ReservationLedger.Contract.CancelReservation(&_ReservationLedger.TransactOpts, _externalId)
}

// CompleteRecharge is a paid mutator transaction binding the contract method 0x68204f29.
//
// Solidity: function completeRecharge(bytes32 _externalId) returns()
func (_ReservationLedger *ReservationLedgerTransactor) CompleteRecharge(opts *bind.TransactOpts, _externalId [32]byte) (*types.Transaction, error) {
	return _ReservationLedger.contract.Transact(opts, "completeRecharge", _externalId)
}

// CompleteRecharge is a paid mutator transaction binding the contract method 0x68204f29.
//
// Solidity: function completeRecharge(bytes32 _externalId) returns()
func (_ReservationLedger *ReservationLedgerSession) CompleteRecharge(_externalId [32]byte) (*types.Transaction, error) {
	return _ReservationLedger.Contract.CompleteRecharge(&_ReservationLedger.TransactOpts, _externalId)
}

// CompleteRecharge is a paid mutator transaction binding the contract method 0x68204f29.
//
// Solidity: function completeRecharge(bytes32 _externalId) returns()
func (_ReservationLedger *ReservationLedgerTransactorSession) CompleteRecharge(_externalId [32]byte) (*types.Transaction, error) {
	return _ReservationLedger.Contract.CompleteRecharge(&_ReservationLedger.TransactOpts, _externalId)
}

// MakeReservation is a paid mutator transaction binding the contract method 0x1b898639.
//
// Solidity: function makeReservation(bytes32 _externalId, string _carId, string _postoId) returns()
func (_ReservationLedger *ReservationLedgerTransactor) MakeReservation(opts *bind.TransactOpts, _externalId [32]byte, _carId string, _postoId string) (*types.Transaction, error) {
	return _ReservationLedger.contract.Transact(opts, "makeReservation", _externalId, _carId, _postoId)
}

// MakeReservation is a paid mutator transaction binding the contract method 0x1b898639.
//
// Solidity: function makeReservation(bytes32 _externalId, string _carId, string _postoId) returns()
func (_ReservationLedger *ReservationLedgerSession) MakeReservation(_externalId [32]byte, _carId string, _postoId string) (*types.Transaction, error) {
	return _ReservationLedger.Contract.MakeReservation(&_ReservationLedger.TransactOpts, _externalId, _carId, _postoId)
}

// MakeReservation is a paid mutator transaction binding the contract method 0x1b898639.
//
// Solidity: function makeReservation(bytes32 _externalId, string _carId, string _postoId) returns()
func (_ReservationLedger *ReservationLedgerTransactorSession) MakeReservation(_externalId [32]byte, _carId string, _postoId string) (*types.Transaction, error) {
	return _ReservationLedger.Contract.MakeReservation(&_ReservationLedger.TransactOpts, _externalId, _carId, _postoId)
}

// ProcessPayment is a paid mutator transaction binding the contract method 0x571376de.
//
// Solidity: function processPayment(bytes32 _externalId, uint256 _amount) returns()
func (_ReservationLedger *ReservationLedgerTransactor) ProcessPayment(opts *bind.TransactOpts, _externalId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _ReservationLedger.contract.Transact(opts, "processPayment", _externalId, _amount)
}

// ProcessPayment is a paid mutator transaction binding the contract method 0x571376de.
//
// Solidity: function processPayment(bytes32 _externalId, uint256 _amount) returns()
func (_ReservationLedger *ReservationLedgerSession) ProcessPayment(_externalId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _ReservationLedger.Contract.ProcessPayment(&_ReservationLedger.TransactOpts, _externalId, _amount)
}

// ProcessPayment is a paid mutator transaction binding the contract method 0x571376de.
//
// Solidity: function processPayment(bytes32 _externalId, uint256 _amount) returns()
func (_ReservationLedger *ReservationLedgerTransactorSession) ProcessPayment(_externalId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _ReservationLedger.Contract.ProcessPayment(&_ReservationLedger.TransactOpts, _externalId, _amount)
}

// ReservationLedgerPaymentProcessedIterator is returned from FilterPaymentProcessed and is used to iterate over the raw logs and unpacked data for PaymentProcessed events raised by the ReservationLedger contract.
type ReservationLedgerPaymentProcessedIterator struct {
	Event *ReservationLedgerPaymentProcessed // Event containing the contract specifics and raw log

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
func (it *ReservationLedgerPaymentProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationLedgerPaymentProcessed)
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
		it.Event = new(ReservationLedgerPaymentProcessed)
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
func (it *ReservationLedgerPaymentProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationLedgerPaymentProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationLedgerPaymentProcessed represents a PaymentProcessed event raised by the ReservationLedger contract.
type ReservationLedgerPaymentProcessed struct {
	ExternalId [32]byte
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPaymentProcessed is a free log retrieval operation binding the contract event 0xd6950d49b769cdf8a136d8b2dd19b86a4582d0aee36c6348c8d14b8623925a7f.
//
// Solidity: event PaymentProcessed(bytes32 indexed externalId, uint256 amount)
func (_ReservationLedger *ReservationLedgerFilterer) FilterPaymentProcessed(opts *bind.FilterOpts, externalId [][32]byte) (*ReservationLedgerPaymentProcessedIterator, error) {

	var externalIdRule []interface{}
	for _, externalIdItem := range externalId {
		externalIdRule = append(externalIdRule, externalIdItem)
	}

	logs, sub, err := _ReservationLedger.contract.FilterLogs(opts, "PaymentProcessed", externalIdRule)
	if err != nil {
		return nil, err
	}
	return &ReservationLedgerPaymentProcessedIterator{contract: _ReservationLedger.contract, event: "PaymentProcessed", logs: logs, sub: sub}, nil
}

// WatchPaymentProcessed is a free log subscription operation binding the contract event 0xd6950d49b769cdf8a136d8b2dd19b86a4582d0aee36c6348c8d14b8623925a7f.
//
// Solidity: event PaymentProcessed(bytes32 indexed externalId, uint256 amount)
func (_ReservationLedger *ReservationLedgerFilterer) WatchPaymentProcessed(opts *bind.WatchOpts, sink chan<- *ReservationLedgerPaymentProcessed, externalId [][32]byte) (event.Subscription, error) {

	var externalIdRule []interface{}
	for _, externalIdItem := range externalId {
		externalIdRule = append(externalIdRule, externalIdItem)
	}

	logs, sub, err := _ReservationLedger.contract.WatchLogs(opts, "PaymentProcessed", externalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationLedgerPaymentProcessed)
				if err := _ReservationLedger.contract.UnpackLog(event, "PaymentProcessed", log); err != nil {
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

// ParsePaymentProcessed is a log parse operation binding the contract event 0xd6950d49b769cdf8a136d8b2dd19b86a4582d0aee36c6348c8d14b8623925a7f.
//
// Solidity: event PaymentProcessed(bytes32 indexed externalId, uint256 amount)
func (_ReservationLedger *ReservationLedgerFilterer) ParsePaymentProcessed(log types.Log) (*ReservationLedgerPaymentProcessed, error) {
	event := new(ReservationLedgerPaymentProcessed)
	if err := _ReservationLedger.contract.UnpackLog(event, "PaymentProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservationLedgerRechargeCompletedIterator is returned from FilterRechargeCompleted and is used to iterate over the raw logs and unpacked data for RechargeCompleted events raised by the ReservationLedger contract.
type ReservationLedgerRechargeCompletedIterator struct {
	Event *ReservationLedgerRechargeCompleted // Event containing the contract specifics and raw log

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
func (it *ReservationLedgerRechargeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationLedgerRechargeCompleted)
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
		it.Event = new(ReservationLedgerRechargeCompleted)
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
func (it *ReservationLedgerRechargeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationLedgerRechargeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationLedgerRechargeCompleted represents a RechargeCompleted event raised by the ReservationLedger contract.
type ReservationLedgerRechargeCompleted struct {
	ExternalId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRechargeCompleted is a free log retrieval operation binding the contract event 0x14e5ff07af041b5d4b7327ec36caf215e96b0929ba2477e89a24f255a845bf03.
//
// Solidity: event RechargeCompleted(bytes32 indexed externalId)
func (_ReservationLedger *ReservationLedgerFilterer) FilterRechargeCompleted(opts *bind.FilterOpts, externalId [][32]byte) (*ReservationLedgerRechargeCompletedIterator, error) {

	var externalIdRule []interface{}
	for _, externalIdItem := range externalId {
		externalIdRule = append(externalIdRule, externalIdItem)
	}

	logs, sub, err := _ReservationLedger.contract.FilterLogs(opts, "RechargeCompleted", externalIdRule)
	if err != nil {
		return nil, err
	}
	return &ReservationLedgerRechargeCompletedIterator{contract: _ReservationLedger.contract, event: "RechargeCompleted", logs: logs, sub: sub}, nil
}

// WatchRechargeCompleted is a free log subscription operation binding the contract event 0x14e5ff07af041b5d4b7327ec36caf215e96b0929ba2477e89a24f255a845bf03.
//
// Solidity: event RechargeCompleted(bytes32 indexed externalId)
func (_ReservationLedger *ReservationLedgerFilterer) WatchRechargeCompleted(opts *bind.WatchOpts, sink chan<- *ReservationLedgerRechargeCompleted, externalId [][32]byte) (event.Subscription, error) {

	var externalIdRule []interface{}
	for _, externalIdItem := range externalId {
		externalIdRule = append(externalIdRule, externalIdItem)
	}

	logs, sub, err := _ReservationLedger.contract.WatchLogs(opts, "RechargeCompleted", externalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationLedgerRechargeCompleted)
				if err := _ReservationLedger.contract.UnpackLog(event, "RechargeCompleted", log); err != nil {
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

// ParseRechargeCompleted is a log parse operation binding the contract event 0x14e5ff07af041b5d4b7327ec36caf215e96b0929ba2477e89a24f255a845bf03.
//
// Solidity: event RechargeCompleted(bytes32 indexed externalId)
func (_ReservationLedger *ReservationLedgerFilterer) ParseRechargeCompleted(log types.Log) (*ReservationLedgerRechargeCompleted, error) {
	event := new(ReservationLedgerRechargeCompleted)
	if err := _ReservationLedger.contract.UnpackLog(event, "RechargeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservationLedgerReservationCancelledIterator is returned from FilterReservationCancelled and is used to iterate over the raw logs and unpacked data for ReservationCancelled events raised by the ReservationLedger contract.
type ReservationLedgerReservationCancelledIterator struct {
	Event *ReservationLedgerReservationCancelled // Event containing the contract specifics and raw log

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
func (it *ReservationLedgerReservationCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationLedgerReservationCancelled)
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
		it.Event = new(ReservationLedgerReservationCancelled)
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
func (it *ReservationLedgerReservationCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationLedgerReservationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationLedgerReservationCancelled represents a ReservationCancelled event raised by the ReservationLedger contract.
type ReservationLedgerReservationCancelled struct {
	ExternalId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReservationCancelled is a free log retrieval operation binding the contract event 0x62aedca87cdbb250ff7a1ee919a4fc3cb2d7d0aad7bfdf913b027a2ba34e17a5.
//
// Solidity: event ReservationCancelled(bytes32 indexed externalId)
func (_ReservationLedger *ReservationLedgerFilterer) FilterReservationCancelled(opts *bind.FilterOpts, externalId [][32]byte) (*ReservationLedgerReservationCancelledIterator, error) {

	var externalIdRule []interface{}
	for _, externalIdItem := range externalId {
		externalIdRule = append(externalIdRule, externalIdItem)
	}

	logs, sub, err := _ReservationLedger.contract.FilterLogs(opts, "ReservationCancelled", externalIdRule)
	if err != nil {
		return nil, err
	}
	return &ReservationLedgerReservationCancelledIterator{contract: _ReservationLedger.contract, event: "ReservationCancelled", logs: logs, sub: sub}, nil
}

// WatchReservationCancelled is a free log subscription operation binding the contract event 0x62aedca87cdbb250ff7a1ee919a4fc3cb2d7d0aad7bfdf913b027a2ba34e17a5.
//
// Solidity: event ReservationCancelled(bytes32 indexed externalId)
func (_ReservationLedger *ReservationLedgerFilterer) WatchReservationCancelled(opts *bind.WatchOpts, sink chan<- *ReservationLedgerReservationCancelled, externalId [][32]byte) (event.Subscription, error) {

	var externalIdRule []interface{}
	for _, externalIdItem := range externalId {
		externalIdRule = append(externalIdRule, externalIdItem)
	}

	logs, sub, err := _ReservationLedger.contract.WatchLogs(opts, "ReservationCancelled", externalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationLedgerReservationCancelled)
				if err := _ReservationLedger.contract.UnpackLog(event, "ReservationCancelled", log); err != nil {
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

// ParseReservationCancelled is a log parse operation binding the contract event 0x62aedca87cdbb250ff7a1ee919a4fc3cb2d7d0aad7bfdf913b027a2ba34e17a5.
//
// Solidity: event ReservationCancelled(bytes32 indexed externalId)
func (_ReservationLedger *ReservationLedgerFilterer) ParseReservationCancelled(log types.Log) (*ReservationLedgerReservationCancelled, error) {
	event := new(ReservationLedgerReservationCancelled)
	if err := _ReservationLedger.contract.UnpackLog(event, "ReservationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservationLedgerReservationMadeIterator is returned from FilterReservationMade and is used to iterate over the raw logs and unpacked data for ReservationMade events raised by the ReservationLedger contract.
type ReservationLedgerReservationMadeIterator struct {
	Event *ReservationLedgerReservationMade // Event containing the contract specifics and raw log

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
func (it *ReservationLedgerReservationMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationLedgerReservationMade)
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
		it.Event = new(ReservationLedgerReservationMade)
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
func (it *ReservationLedgerReservationMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationLedgerReservationMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationLedgerReservationMade represents a ReservationMade event raised by the ReservationLedger contract.
type ReservationLedgerReservationMade struct {
	ExternalId [32]byte
	CarId      string
	PostoId    string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReservationMade is a free log retrieval operation binding the contract event 0x9c0cdc6273c92721e246d12f2be87890adfc5359298ffb6d42a33787dc6adde6.
//
// Solidity: event ReservationMade(bytes32 indexed externalId, string carId, string postoId)
func (_ReservationLedger *ReservationLedgerFilterer) FilterReservationMade(opts *bind.FilterOpts, externalId [][32]byte) (*ReservationLedgerReservationMadeIterator, error) {

	var externalIdRule []interface{}
	for _, externalIdItem := range externalId {
		externalIdRule = append(externalIdRule, externalIdItem)
	}

	logs, sub, err := _ReservationLedger.contract.FilterLogs(opts, "ReservationMade", externalIdRule)
	if err != nil {
		return nil, err
	}
	return &ReservationLedgerReservationMadeIterator{contract: _ReservationLedger.contract, event: "ReservationMade", logs: logs, sub: sub}, nil
}

// WatchReservationMade is a free log subscription operation binding the contract event 0x9c0cdc6273c92721e246d12f2be87890adfc5359298ffb6d42a33787dc6adde6.
//
// Solidity: event ReservationMade(bytes32 indexed externalId, string carId, string postoId)
func (_ReservationLedger *ReservationLedgerFilterer) WatchReservationMade(opts *bind.WatchOpts, sink chan<- *ReservationLedgerReservationMade, externalId [][32]byte) (event.Subscription, error) {

	var externalIdRule []interface{}
	for _, externalIdItem := range externalId {
		externalIdRule = append(externalIdRule, externalIdItem)
	}

	logs, sub, err := _ReservationLedger.contract.WatchLogs(opts, "ReservationMade", externalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationLedgerReservationMade)
				if err := _ReservationLedger.contract.UnpackLog(event, "ReservationMade", log); err != nil {
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

// ParseReservationMade is a log parse operation binding the contract event 0x9c0cdc6273c92721e246d12f2be87890adfc5359298ffb6d42a33787dc6adde6.
//
// Solidity: event ReservationMade(bytes32 indexed externalId, string carId, string postoId)
func (_ReservationLedger *ReservationLedgerFilterer) ParseReservationMade(log types.Log) (*ReservationLedgerReservationMade, error) {
	event := new(ReservationLedgerReservationMade)
	if err := _ReservationLedger.contract.UnpackLog(event, "ReservationMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
