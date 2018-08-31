// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// MergedMinerValidatorABI is the input ABI used to generate the binding from.
const MergedMinerValidatorABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"submitBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastBlockSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SUBMISSIONREWARD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashedBlocks\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rtI\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setRTI\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumbers\",\"type\":\"uint256[]\"}],\"name\":\"bulkClaimReward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"goodNightSweetPrince\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BLOCKREWARD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"name\":\"number\",\"type\":\"uint256\"},{\"name\":\"coinbase\",\"type\":\"address\"},{\"name\":\"state\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_coinbase\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_submitter\",\"type\":\"address\"}],\"name\":\"BlockInformationSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_claimer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_blockNumbers\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"_totalReward\",\"type\":\"uint256\"}],\"name\":\"MergedMinedRewardClaimed\",\"type\":\"event\"}]"

// MergedMinerValidatorBin is the compiled bytecode used for deploying new contracts.
const MergedMinerValidatorBin = `608060405234801561001057600080fd5b50604051602080610fd2833981016040525161002a610122565b60018054600160a060020a031916600160a060020a0384161781556040805160608101825243815241602082015291908201524360028181556000918252600460209081526040928390208451815590840151600182018054600160a060020a031916600160a060020a039092169190911780825593850151949550859491939192909160a060020a60ff02191690740100000000000000000000000000000000000000009084908111156100db57fe5b0217905550506040805160008152905143925041917f607bba4a16235877e634164f23916760adfb2e9a596e6e1f691347e944a6bcc5919081900360200190a35050610140565b60408051606081018252600080825260208201819052909182015290565b610e838061014f6000396000f3006080604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166325ceb4b281146100c95780632722692c146100f25780633218ebcc1461011957806391323e501461012e5780639d76ea58146101bb578063be72ab59146101ec578063deb78ab614610201578063e2bfcb4214610222578063ed2d1d9e14610242578063f2477f7c14610257578063f25b3f991461026c578063f851a440146102bf578063ffa1ad74146102d4575b600080fd5b3480156100d557600080fd5b506100de6102e9565b604080519115158252519081900360200190f35b3480156100fe57600080fd5b50610107610614565b60408051918252519081900360200190f35b34801561012557600080fd5b5061010761061a565b34801561013a57600080fd5b50610146600435610626565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610180578181015183820152602001610168565b50505050905090810190601f1680156101ad5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101c757600080fd5b506101d06106c1565b60408051600160a060020a039092168252519081900360200190f35b3480156101f857600080fd5b506101d06106d0565b34801561020d57600080fd5b506100de600160a060020a03600435166106df565b34801561022e57600080fd5b506100de600480356024810191013561080b565b34801561024e57600080fd5b506100de610bab565b34801561026357600080fd5b50610107610c39565b34801561027857600080fd5b50610284600435610c45565b60408051848152600160a060020a03841660208201529081018260028111156102a957fe5b60ff168152602001935050505060405180910390f35b3480156102cb57600080fd5b506101d0610c84565b3480156102e057600080fd5b50610146610c93565b60006102f3610e39565b436000808281526004602052604090206001015474010000000000000000000000000000000000000000900460ff16600281111561032d57fe5b14610382576040805160e560020a62461bcd02815260206004820152601960248201527f626c6f636b207374617465206d75737420626520656d70747900000000000000604482015290519081900360640190fd5b60025443908111610403576040805160e560020a62461bcd02815260206004820152603860248201527f756e61626c6520746f207375626d697420696e666f726d6174696f6e20666f7260448201527f20616c7265616479207375626d697474656420626c6f636b0000000000000000606482015290519081900360840190fd5b6040805160608101825243808252416020808401918252600184860181815260028581556000958652600490935295909320845181559151928201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0390941693909317808455945193975087949193929174ff0000000000000000000000000000000000000000191690740100000000000000000000000000000000000000009084908111156104af57fe5b02179055505060408051338152905143925041917f607bba4a16235877e634164f23916760adfb2e9a596e6e1f691347e944a6bcc5919081900360200190a3600354604080517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526706f05b59d3b2000060248201529051600160a060020a03909216916340c10f19916044808201926020929091908290030181600087803b15801561056257600080fd5b505af1158015610576573d6000803e3d6000fd5b505050506040513d602081101561058c57600080fd5b5051151561060a576040805160e560020a62461bcd02815260206004820152602c60248201527f6661696c656420746f207472616e736665722072657761726420746f20626c6f60448201527f636b207375626d69747465720000000000000000000000000000000000000000606482015290519081900360840190fd5b6001935050505090565b60025481565b6706f05b59d3b2000081565b60056020908152600091825260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156106b95780601f1061068e576101008083540402835291602001916106b9565b820191906000526020600020905b81548152906001019060200180831161069c57829003601f168201915b505050505081565b600054600160a060020a031681565b600354600160a060020a031681565b600154600090600160a060020a0316331461076a576040805160e560020a62461bcd02815260206004820152602660248201527f6f6e6c7920616e2061646d696e2063616e20696e766f6b65207468697320667560448201527f6e6374696f6e0000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600054600160a060020a0316156107cb576040805160e560020a62461bcd02815260206004820152601d60248201527f746f6b656e2061646472657373206d757374206e6f7420626520736574000000604482015290519081900360640190fd5b5060008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1991821681179092556003805490911690911790556001919050565b600354604080517fe69575140000000000000000000000000000000000000000000000000000000081529051600092839283923092600160a060020a03169163e6957514916004808301926020929190829003018186803b15801561086f57600080fd5b505afa158015610883573d6000803e3d6000fd5b505050506040513d602081101561089957600080fd5b5051600160a060020a03161461091f576040805160e560020a62461bcd02815260206004820152603f60248201527f6d6572676564206d696e657220636f6e7472616374206f6e2072746320746f6b60448201527f656e206d7573742062652073657420746f207468697320636f6e747261637400606482015290519081900360840190fd5b601484111561099e576040805160e560020a62461bcd02815260206004820152602760248201527f63616e206f6e6c7920636c61696d20757020746f20323020726577617264732060448201527f6174206f6e636500000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b5060005b838110156109df576109d56109c88686848181106109bc57fe5b90506020020135610cca565b839063ffffffff610e2016565b91506001016109a2565b8484604051808383602002808284376040805191909301819003812088825292519295503394507fb1edf40adb1abba7ef019c51048fb2da24246512f4e7c1dd75061e727298c2fb9350829003602001919050a360008211610ab1576040805160e560020a62461bcd02815260206004820152602a60248201527f746f74616c20636f696e7320746f206d696e74206d757374206265206772656160448201527f746572207468616e203000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600354604080517f40c10f19000000000000000000000000000000000000000000000000000000008152336004820152602481018590529051600160a060020a03909216916340c10f19916044808201926020929091908290030181600087803b158015610b1e57600080fd5b505af1158015610b32573d6000803e3d6000fd5b505050506040513d6020811015610b4857600080fd5b50511515610ba0576040805160e560020a62461bcd02815260206004820152601560248201527f756e61626c6520746f206d696e7420746f6b656e730000000000000000000000604482015290519081900360640190fd5b506001949350505050565b600154600090600160a060020a03163314610c36576040805160e560020a62461bcd02815260206004820152602660248201527f6f6e6c7920616e2061646d696e2063616e20696e766f6b65207468697320667560448201527f6e6374696f6e0000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b33ff5b670429d069189e000081565b60046020526000908152604090208054600190910154600160a060020a0381169074010000000000000000000000000000000000000000900460ff1683565b600154600160a060020a031681565b60408051808201909152600a81527f70726f64756374696f6e00000000000000000000000000000000000000000000602082015281565b6000818152600460205260408120600101548290600160a060020a03163314610d3d576040805160e560020a62461bcd02815260206004820152601760248201527f73656e646572206d75737420626520636f696e62617365000000000000000000604482015290519081900360640190fd5b82600160008281526004602052604090206001015474010000000000000000000000000000000000000000900460ff166002811115610d7857fe5b14610dcd576040805160e560020a62461bcd02815260206004820152601d60248201527f626c6f636b207374617465206d757374206265207375626d6974746564000000604482015290519081900360640190fd5b5050506000908152600460205260409020600101805474ff0000000000000000000000000000000000000000191674020000000000000000000000000000000000000000179055670429d069189e000090565b600082820183811015610e3257600080fd5b9392505050565b604080516060810182526000808252602082018190529091820152905600a165627a7a723058203d06b189cf6c727a0a52eddf594a856c8a60057cf1c14aa44dad422ee6bba2770029`

// DeployMergedMinerValidator deploys a new Ethereum contract, binding an instance of MergedMinerValidator to it.
func DeployMergedMinerValidator(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address) (common.Address, *types.Transaction, *MergedMinerValidator, error) {
	parsed, err := abi.JSON(strings.NewReader(MergedMinerValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MergedMinerValidatorBin), backend, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MergedMinerValidator{MergedMinerValidatorCaller: MergedMinerValidatorCaller{contract: contract}, MergedMinerValidatorTransactor: MergedMinerValidatorTransactor{contract: contract}, MergedMinerValidatorFilterer: MergedMinerValidatorFilterer{contract: contract}}, nil
}

// MergedMinerValidator is an auto generated Go binding around an Ethereum contract.
type MergedMinerValidator struct {
	MergedMinerValidatorCaller     // Read-only binding to the contract
	MergedMinerValidatorTransactor // Write-only binding to the contract
	MergedMinerValidatorFilterer   // Log filterer for contract events
}

// MergedMinerValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MergedMinerValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MergedMinerValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MergedMinerValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MergedMinerValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MergedMinerValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MergedMinerValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MergedMinerValidatorSession struct {
	Contract     *MergedMinerValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MergedMinerValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MergedMinerValidatorCallerSession struct {
	Contract *MergedMinerValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// MergedMinerValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MergedMinerValidatorTransactorSession struct {
	Contract     *MergedMinerValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MergedMinerValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MergedMinerValidatorRaw struct {
	Contract *MergedMinerValidator // Generic contract binding to access the raw methods on
}

// MergedMinerValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MergedMinerValidatorCallerRaw struct {
	Contract *MergedMinerValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// MergedMinerValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MergedMinerValidatorTransactorRaw struct {
	Contract *MergedMinerValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMergedMinerValidator creates a new instance of MergedMinerValidator, bound to a specific deployed contract.
func NewMergedMinerValidator(address common.Address, backend bind.ContractBackend) (*MergedMinerValidator, error) {
	contract, err := bindMergedMinerValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MergedMinerValidator{MergedMinerValidatorCaller: MergedMinerValidatorCaller{contract: contract}, MergedMinerValidatorTransactor: MergedMinerValidatorTransactor{contract: contract}, MergedMinerValidatorFilterer: MergedMinerValidatorFilterer{contract: contract}}, nil
}

// NewMergedMinerValidatorCaller creates a new read-only instance of MergedMinerValidator, bound to a specific deployed contract.
func NewMergedMinerValidatorCaller(address common.Address, caller bind.ContractCaller) (*MergedMinerValidatorCaller, error) {
	contract, err := bindMergedMinerValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MergedMinerValidatorCaller{contract: contract}, nil
}

// NewMergedMinerValidatorTransactor creates a new write-only instance of MergedMinerValidator, bound to a specific deployed contract.
func NewMergedMinerValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*MergedMinerValidatorTransactor, error) {
	contract, err := bindMergedMinerValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MergedMinerValidatorTransactor{contract: contract}, nil
}

// NewMergedMinerValidatorFilterer creates a new log filterer instance of MergedMinerValidator, bound to a specific deployed contract.
func NewMergedMinerValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*MergedMinerValidatorFilterer, error) {
	contract, err := bindMergedMinerValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MergedMinerValidatorFilterer{contract: contract}, nil
}

// bindMergedMinerValidator binds a generic wrapper to an already deployed contract.
func bindMergedMinerValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MergedMinerValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MergedMinerValidator *MergedMinerValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MergedMinerValidator.Contract.MergedMinerValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MergedMinerValidator *MergedMinerValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MergedMinerValidator.Contract.MergedMinerValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MergedMinerValidator *MergedMinerValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MergedMinerValidator.Contract.MergedMinerValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MergedMinerValidator *MergedMinerValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MergedMinerValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MergedMinerValidator *MergedMinerValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MergedMinerValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MergedMinerValidator *MergedMinerValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MergedMinerValidator.Contract.contract.Transact(opts, method, params...)
}

// BLOCKREWARD is a free data retrieval call binding the contract method 0xf2477f7c.
//
// Solidity: function BLOCKREWARD() constant returns(uint256)
func (_MergedMinerValidator *MergedMinerValidatorCaller) BLOCKREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MergedMinerValidator.contract.Call(opts, out, "BLOCKREWARD")
	return *ret0, err
}

// BLOCKREWARD is a free data retrieval call binding the contract method 0xf2477f7c.
//
// Solidity: function BLOCKREWARD() constant returns(uint256)
func (_MergedMinerValidator *MergedMinerValidatorSession) BLOCKREWARD() (*big.Int, error) {
	return _MergedMinerValidator.Contract.BLOCKREWARD(&_MergedMinerValidator.CallOpts)
}

// BLOCKREWARD is a free data retrieval call binding the contract method 0xf2477f7c.
//
// Solidity: function BLOCKREWARD() constant returns(uint256)
func (_MergedMinerValidator *MergedMinerValidatorCallerSession) BLOCKREWARD() (*big.Int, error) {
	return _MergedMinerValidator.Contract.BLOCKREWARD(&_MergedMinerValidator.CallOpts)
}

// SUBMISSIONREWARD is a free data retrieval call binding the contract method 0x3218ebcc.
//
// Solidity: function SUBMISSIONREWARD() constant returns(uint256)
func (_MergedMinerValidator *MergedMinerValidatorCaller) SUBMISSIONREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MergedMinerValidator.contract.Call(opts, out, "SUBMISSIONREWARD")
	return *ret0, err
}

// SUBMISSIONREWARD is a free data retrieval call binding the contract method 0x3218ebcc.
//
// Solidity: function SUBMISSIONREWARD() constant returns(uint256)
func (_MergedMinerValidator *MergedMinerValidatorSession) SUBMISSIONREWARD() (*big.Int, error) {
	return _MergedMinerValidator.Contract.SUBMISSIONREWARD(&_MergedMinerValidator.CallOpts)
}

// SUBMISSIONREWARD is a free data retrieval call binding the contract method 0x3218ebcc.
//
// Solidity: function SUBMISSIONREWARD() constant returns(uint256)
func (_MergedMinerValidator *MergedMinerValidatorCallerSession) SUBMISSIONREWARD() (*big.Int, error) {
	return _MergedMinerValidator.Contract.SUBMISSIONREWARD(&_MergedMinerValidator.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_MergedMinerValidator *MergedMinerValidatorCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MergedMinerValidator.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_MergedMinerValidator *MergedMinerValidatorSession) VERSION() (string, error) {
	return _MergedMinerValidator.Contract.VERSION(&_MergedMinerValidator.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_MergedMinerValidator *MergedMinerValidatorCallerSession) VERSION() (string, error) {
	return _MergedMinerValidator.Contract.VERSION(&_MergedMinerValidator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_MergedMinerValidator *MergedMinerValidatorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MergedMinerValidator.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_MergedMinerValidator *MergedMinerValidatorSession) Admin() (common.Address, error) {
	return _MergedMinerValidator.Contract.Admin(&_MergedMinerValidator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_MergedMinerValidator *MergedMinerValidatorCallerSession) Admin() (common.Address, error) {
	return _MergedMinerValidator.Contract.Admin(&_MergedMinerValidator.CallOpts)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks( uint256) constant returns(number uint256, coinbase address, state uint8)
func (_MergedMinerValidator *MergedMinerValidatorCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Number   *big.Int
	Coinbase common.Address
	State    uint8
}, error) {
	ret := new(struct {
		Number   *big.Int
		Coinbase common.Address
		State    uint8
	})
	out := ret
	err := _MergedMinerValidator.contract.Call(opts, out, "blocks", arg0)
	return *ret, err
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks( uint256) constant returns(number uint256, coinbase address, state uint8)
func (_MergedMinerValidator *MergedMinerValidatorSession) Blocks(arg0 *big.Int) (struct {
	Number   *big.Int
	Coinbase common.Address
	State    uint8
}, error) {
	return _MergedMinerValidator.Contract.Blocks(&_MergedMinerValidator.CallOpts, arg0)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks( uint256) constant returns(number uint256, coinbase address, state uint8)
func (_MergedMinerValidator *MergedMinerValidatorCallerSession) Blocks(arg0 *big.Int) (struct {
	Number   *big.Int
	Coinbase common.Address
	State    uint8
}, error) {
	return _MergedMinerValidator.Contract.Blocks(&_MergedMinerValidator.CallOpts, arg0)
}

// HashedBlocks is a free data retrieval call binding the contract method 0x91323e50.
//
// Solidity: function hashedBlocks( uint256) constant returns(bytes)
func (_MergedMinerValidator *MergedMinerValidatorCaller) HashedBlocks(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _MergedMinerValidator.contract.Call(opts, out, "hashedBlocks", arg0)
	return *ret0, err
}

// HashedBlocks is a free data retrieval call binding the contract method 0x91323e50.
//
// Solidity: function hashedBlocks( uint256) constant returns(bytes)
func (_MergedMinerValidator *MergedMinerValidatorSession) HashedBlocks(arg0 *big.Int) ([]byte, error) {
	return _MergedMinerValidator.Contract.HashedBlocks(&_MergedMinerValidator.CallOpts, arg0)
}

// HashedBlocks is a free data retrieval call binding the contract method 0x91323e50.
//
// Solidity: function hashedBlocks( uint256) constant returns(bytes)
func (_MergedMinerValidator *MergedMinerValidatorCallerSession) HashedBlocks(arg0 *big.Int) ([]byte, error) {
	return _MergedMinerValidator.Contract.HashedBlocks(&_MergedMinerValidator.CallOpts, arg0)
}

// LastBlockSet is a free data retrieval call binding the contract method 0x2722692c.
//
// Solidity: function lastBlockSet() constant returns(uint256)
func (_MergedMinerValidator *MergedMinerValidatorCaller) LastBlockSet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MergedMinerValidator.contract.Call(opts, out, "lastBlockSet")
	return *ret0, err
}

// LastBlockSet is a free data retrieval call binding the contract method 0x2722692c.
//
// Solidity: function lastBlockSet() constant returns(uint256)
func (_MergedMinerValidator *MergedMinerValidatorSession) LastBlockSet() (*big.Int, error) {
	return _MergedMinerValidator.Contract.LastBlockSet(&_MergedMinerValidator.CallOpts)
}

// LastBlockSet is a free data retrieval call binding the contract method 0x2722692c.
//
// Solidity: function lastBlockSet() constant returns(uint256)
func (_MergedMinerValidator *MergedMinerValidatorCallerSession) LastBlockSet() (*big.Int, error) {
	return _MergedMinerValidator.Contract.LastBlockSet(&_MergedMinerValidator.CallOpts)
}

// RtI is a free data retrieval call binding the contract method 0xbe72ab59.
//
// Solidity: function rtI() constant returns(address)
func (_MergedMinerValidator *MergedMinerValidatorCaller) RtI(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MergedMinerValidator.contract.Call(opts, out, "rtI")
	return *ret0, err
}

// RtI is a free data retrieval call binding the contract method 0xbe72ab59.
//
// Solidity: function rtI() constant returns(address)
func (_MergedMinerValidator *MergedMinerValidatorSession) RtI() (common.Address, error) {
	return _MergedMinerValidator.Contract.RtI(&_MergedMinerValidator.CallOpts)
}

// RtI is a free data retrieval call binding the contract method 0xbe72ab59.
//
// Solidity: function rtI() constant returns(address)
func (_MergedMinerValidator *MergedMinerValidatorCallerSession) RtI() (common.Address, error) {
	return _MergedMinerValidator.Contract.RtI(&_MergedMinerValidator.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_MergedMinerValidator *MergedMinerValidatorCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MergedMinerValidator.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_MergedMinerValidator *MergedMinerValidatorSession) TokenAddress() (common.Address, error) {
	return _MergedMinerValidator.Contract.TokenAddress(&_MergedMinerValidator.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_MergedMinerValidator *MergedMinerValidatorCallerSession) TokenAddress() (common.Address, error) {
	return _MergedMinerValidator.Contract.TokenAddress(&_MergedMinerValidator.CallOpts)
}

// BulkClaimReward is a paid mutator transaction binding the contract method 0xe2bfcb42.
//
// Solidity: function bulkClaimReward(_blockNumbers uint256[]) returns(bool)
func (_MergedMinerValidator *MergedMinerValidatorTransactor) BulkClaimReward(opts *bind.TransactOpts, _blockNumbers []*big.Int) (*types.Transaction, error) {
	return _MergedMinerValidator.contract.Transact(opts, "bulkClaimReward", _blockNumbers)
}

// BulkClaimReward is a paid mutator transaction binding the contract method 0xe2bfcb42.
//
// Solidity: function bulkClaimReward(_blockNumbers uint256[]) returns(bool)
func (_MergedMinerValidator *MergedMinerValidatorSession) BulkClaimReward(_blockNumbers []*big.Int) (*types.Transaction, error) {
	return _MergedMinerValidator.Contract.BulkClaimReward(&_MergedMinerValidator.TransactOpts, _blockNumbers)
}

// BulkClaimReward is a paid mutator transaction binding the contract method 0xe2bfcb42.
//
// Solidity: function bulkClaimReward(_blockNumbers uint256[]) returns(bool)
func (_MergedMinerValidator *MergedMinerValidatorTransactorSession) BulkClaimReward(_blockNumbers []*big.Int) (*types.Transaction, error) {
	return _MergedMinerValidator.Contract.BulkClaimReward(&_MergedMinerValidator.TransactOpts, _blockNumbers)
}

// GoodNightSweetPrince is a paid mutator transaction binding the contract method 0xed2d1d9e.
//
// Solidity: function goodNightSweetPrince() returns(bool)
func (_MergedMinerValidator *MergedMinerValidatorTransactor) GoodNightSweetPrince(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MergedMinerValidator.contract.Transact(opts, "goodNightSweetPrince")
}

// GoodNightSweetPrince is a paid mutator transaction binding the contract method 0xed2d1d9e.
//
// Solidity: function goodNightSweetPrince() returns(bool)
func (_MergedMinerValidator *MergedMinerValidatorSession) GoodNightSweetPrince() (*types.Transaction, error) {
	return _MergedMinerValidator.Contract.GoodNightSweetPrince(&_MergedMinerValidator.TransactOpts)
}

// GoodNightSweetPrince is a paid mutator transaction binding the contract method 0xed2d1d9e.
//
// Solidity: function goodNightSweetPrince() returns(bool)
func (_MergedMinerValidator *MergedMinerValidatorTransactorSession) GoodNightSweetPrince() (*types.Transaction, error) {
	return _MergedMinerValidator.Contract.GoodNightSweetPrince(&_MergedMinerValidator.TransactOpts)
}

// SetRTI is a paid mutator transaction binding the contract method 0xdeb78ab6.
//
// Solidity: function setRTI(_tokenAddress address) returns(bool)
func (_MergedMinerValidator *MergedMinerValidatorTransactor) SetRTI(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _MergedMinerValidator.contract.Transact(opts, "setRTI", _tokenAddress)
}

// SetRTI is a paid mutator transaction binding the contract method 0xdeb78ab6.
//
// Solidity: function setRTI(_tokenAddress address) returns(bool)
func (_MergedMinerValidator *MergedMinerValidatorSession) SetRTI(_tokenAddress common.Address) (*types.Transaction, error) {
	return _MergedMinerValidator.Contract.SetRTI(&_MergedMinerValidator.TransactOpts, _tokenAddress)
}

// SetRTI is a paid mutator transaction binding the contract method 0xdeb78ab6.
//
// Solidity: function setRTI(_tokenAddress address) returns(bool)
func (_MergedMinerValidator *MergedMinerValidatorTransactorSession) SetRTI(_tokenAddress common.Address) (*types.Transaction, error) {
	return _MergedMinerValidator.Contract.SetRTI(&_MergedMinerValidator.TransactOpts, _tokenAddress)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0x25ceb4b2.
//
// Solidity: function submitBlock() returns(bool)
func (_MergedMinerValidator *MergedMinerValidatorTransactor) SubmitBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MergedMinerValidator.contract.Transact(opts, "submitBlock")
}

// SubmitBlock is a paid mutator transaction binding the contract method 0x25ceb4b2.
//
// Solidity: function submitBlock() returns(bool)
func (_MergedMinerValidator *MergedMinerValidatorSession) SubmitBlock() (*types.Transaction, error) {
	return _MergedMinerValidator.Contract.SubmitBlock(&_MergedMinerValidator.TransactOpts)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0x25ceb4b2.
//
// Solidity: function submitBlock() returns(bool)
func (_MergedMinerValidator *MergedMinerValidatorTransactorSession) SubmitBlock() (*types.Transaction, error) {
	return _MergedMinerValidator.Contract.SubmitBlock(&_MergedMinerValidator.TransactOpts)
}

// MergedMinerValidatorBlockInformationSubmittedIterator is returned from FilterBlockInformationSubmitted and is used to iterate over the raw logs and unpacked data for BlockInformationSubmitted events raised by the MergedMinerValidator contract.
type MergedMinerValidatorBlockInformationSubmittedIterator struct {
	Event *MergedMinerValidatorBlockInformationSubmitted // Event containing the contract specifics and raw log

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
func (it *MergedMinerValidatorBlockInformationSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MergedMinerValidatorBlockInformationSubmitted)
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
		it.Event = new(MergedMinerValidatorBlockInformationSubmitted)
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
func (it *MergedMinerValidatorBlockInformationSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MergedMinerValidatorBlockInformationSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MergedMinerValidatorBlockInformationSubmitted represents a BlockInformationSubmitted event raised by the MergedMinerValidator contract.
type MergedMinerValidatorBlockInformationSubmitted struct {
	Coinbase    common.Address
	BlockNumber *big.Int
	Submitter   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockInformationSubmitted is a free log retrieval operation binding the contract event 0x607bba4a16235877e634164f23916760adfb2e9a596e6e1f691347e944a6bcc5.
//
// Solidity: e BlockInformationSubmitted(_coinbase indexed address, _blockNumber indexed uint256, _submitter address)
func (_MergedMinerValidator *MergedMinerValidatorFilterer) FilterBlockInformationSubmitted(opts *bind.FilterOpts, _coinbase []common.Address, _blockNumber []*big.Int) (*MergedMinerValidatorBlockInformationSubmittedIterator, error) {

	var _coinbaseRule []interface{}
	for _, _coinbaseItem := range _coinbase {
		_coinbaseRule = append(_coinbaseRule, _coinbaseItem)
	}
	var _blockNumberRule []interface{}
	for _, _blockNumberItem := range _blockNumber {
		_blockNumberRule = append(_blockNumberRule, _blockNumberItem)
	}

	logs, sub, err := _MergedMinerValidator.contract.FilterLogs(opts, "BlockInformationSubmitted", _coinbaseRule, _blockNumberRule)
	if err != nil {
		return nil, err
	}
	return &MergedMinerValidatorBlockInformationSubmittedIterator{contract: _MergedMinerValidator.contract, event: "BlockInformationSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockInformationSubmitted is a free log subscription operation binding the contract event 0x607bba4a16235877e634164f23916760adfb2e9a596e6e1f691347e944a6bcc5.
//
// Solidity: e BlockInformationSubmitted(_coinbase indexed address, _blockNumber indexed uint256, _submitter address)
func (_MergedMinerValidator *MergedMinerValidatorFilterer) WatchBlockInformationSubmitted(opts *bind.WatchOpts, sink chan<- *MergedMinerValidatorBlockInformationSubmitted, _coinbase []common.Address, _blockNumber []*big.Int) (event.Subscription, error) {

	var _coinbaseRule []interface{}
	for _, _coinbaseItem := range _coinbase {
		_coinbaseRule = append(_coinbaseRule, _coinbaseItem)
	}
	var _blockNumberRule []interface{}
	for _, _blockNumberItem := range _blockNumber {
		_blockNumberRule = append(_blockNumberRule, _blockNumberItem)
	}

	logs, sub, err := _MergedMinerValidator.contract.WatchLogs(opts, "BlockInformationSubmitted", _coinbaseRule, _blockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MergedMinerValidatorBlockInformationSubmitted)
				if err := _MergedMinerValidator.contract.UnpackLog(event, "BlockInformationSubmitted", log); err != nil {
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

// MergedMinerValidatorMergedMinedRewardClaimedIterator is returned from FilterMergedMinedRewardClaimed and is used to iterate over the raw logs and unpacked data for MergedMinedRewardClaimed events raised by the MergedMinerValidator contract.
type MergedMinerValidatorMergedMinedRewardClaimedIterator struct {
	Event *MergedMinerValidatorMergedMinedRewardClaimed // Event containing the contract specifics and raw log

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
func (it *MergedMinerValidatorMergedMinedRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MergedMinerValidatorMergedMinedRewardClaimed)
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
		it.Event = new(MergedMinerValidatorMergedMinedRewardClaimed)
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
func (it *MergedMinerValidatorMergedMinedRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MergedMinerValidatorMergedMinedRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MergedMinerValidatorMergedMinedRewardClaimed represents a MergedMinedRewardClaimed event raised by the MergedMinerValidator contract.
type MergedMinerValidatorMergedMinedRewardClaimed struct {
	Claimer      common.Address
	BlockNumbers []*big.Int
	TotalReward  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMergedMinedRewardClaimed is a free log retrieval operation binding the contract event 0xb1edf40adb1abba7ef019c51048fb2da24246512f4e7c1dd75061e727298c2fb.
//
// Solidity: e MergedMinedRewardClaimed(_claimer indexed address, _blockNumbers indexed uint256[], _totalReward uint256)
func (_MergedMinerValidator *MergedMinerValidatorFilterer) FilterMergedMinedRewardClaimed(opts *bind.FilterOpts, _claimer []common.Address, _blockNumbers [][]*big.Int) (*MergedMinerValidatorMergedMinedRewardClaimedIterator, error) {

	var _claimerRule []interface{}
	for _, _claimerItem := range _claimer {
		_claimerRule = append(_claimerRule, _claimerItem)
	}
	var _blockNumbersRule []interface{}
	for _, _blockNumbersItem := range _blockNumbers {
		_blockNumbersRule = append(_blockNumbersRule, _blockNumbersItem)
	}

	logs, sub, err := _MergedMinerValidator.contract.FilterLogs(opts, "MergedMinedRewardClaimed", _claimerRule, _blockNumbersRule)
	if err != nil {
		return nil, err
	}
	return &MergedMinerValidatorMergedMinedRewardClaimedIterator{contract: _MergedMinerValidator.contract, event: "MergedMinedRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchMergedMinedRewardClaimed is a free log subscription operation binding the contract event 0xb1edf40adb1abba7ef019c51048fb2da24246512f4e7c1dd75061e727298c2fb.
//
// Solidity: e MergedMinedRewardClaimed(_claimer indexed address, _blockNumbers indexed uint256[], _totalReward uint256)
func (_MergedMinerValidator *MergedMinerValidatorFilterer) WatchMergedMinedRewardClaimed(opts *bind.WatchOpts, sink chan<- *MergedMinerValidatorMergedMinedRewardClaimed, _claimer []common.Address, _blockNumbers [][]*big.Int) (event.Subscription, error) {

	var _claimerRule []interface{}
	for _, _claimerItem := range _claimer {
		_claimerRule = append(_claimerRule, _claimerItem)
	}
	var _blockNumbersRule []interface{}
	for _, _blockNumbersItem := range _blockNumbers {
		_blockNumbersRule = append(_blockNumbersRule, _blockNumbersItem)
	}

	logs, sub, err := _MergedMinerValidator.contract.WatchLogs(opts, "MergedMinedRewardClaimed", _claimerRule, _blockNumbersRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MergedMinerValidatorMergedMinedRewardClaimed)
				if err := _MergedMinerValidator.contract.UnpackLog(event, "MergedMinedRewardClaimed", log); err != nil {
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
