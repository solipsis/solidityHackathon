// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SpawnABI is the input ABI used to generate the binding from.
const SpawnABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"createIssue\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressLUT\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"issues\",\"outputs\":[{\"name\":\"threshold\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"payable\":true,\"type\":\"fallback\"}]"

// SpawnBin is the compiled bytecode used for deploying new contracts.
const SpawnBin = `0x6060604052341561000c57fe5b5b6106238061001c6000396000f3006060604052361561005f5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638615d8558114610068578063949d225d146100dc578063aca49018146100fe578063e9a5551c1461012d575b6100665b5b565b005b341561007057fe5b6100c0600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284375094965050933593506101d892505050565b60408051600160a060020a039092168252519081900360200190f35b34156100e457fe5b6100ec610294565b60408051918252519081900360200190f35b341561010657fe5b6100c060043561029b565b60408051600160a060020a039092168252519081900360200190f35b341561013557fe5b610149600160a060020a03600435166102cd565b604080518381526020810182815283546002600019610100600184161502019091160492820183905290916060830190849080156101c85780601f1061019d576101008083540402835291602001916101c8565b820191906000526020600020905b8154815290600101906020018083116101ab57829003601f168201915b5050935050505060405180910390f35b60006000826101e56102e4565b90815260405190819003602001906000f08015156101ff57fe5b600160a060020a038116600090815260208181526040909120865192935061022f926001909101918701906102f4565b50600160a060020a03811660009081526020819052604090208390556001805480820161025c8382610373565b916000526020600020900160005b8154600160a060020a038086166101009390930a92830292021916179055509050805b5092915050565b6001545b90565b60018054829081106102a957fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b600060208190529081526040902080549060010182565b604051610239806103bf83390190565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061033557805160ff1916838001178555610362565b82800160010185558215610362579182015b82811115610362578251825591602001919060010190610347565b5b5061036f92915061039d565b5090565b8154818355818115116103975760008381526020902061039791810190830161039d565b5b505050565b61029891905b8082111561036f57600081556001016103a3565b5090565b9056006060604052341561000c57fe5b60405160208061023983398101604052515b600180546000839055600160a060020a03191633600160a060020a03161760a060020a60ff02191690555b505b6101df8061005a6000396000f3006060604052361561005f5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631922ff3981146100b157806342cde4e8146100c75780638da5cb5b146100e9578063f3a504f214610115575b6100af5b600054600160a060020a03301631106100ac576001805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790555b5b565b005b6100af600160a060020a0360043516610139565b005b34156100cf57fe5b6100d761017d565b60408051918252519081900360200190f35b34156100f157fe5b6100f9610183565b60408051600160a060020a039092168252519081900360200190f35b341561011d57fe5b610125610192565b604080519115158252519081900360200190f35b60015433600160a060020a0390811691161461017957604051600160a060020a0380831691309091163180156108fc02916000818181858888f150505050505b5b50565b60005481565b600154600160a060020a031681565b60015474010000000000000000000000000000000000000000900460ff16815600a165627a7a723058208f363b32654b5396ae379deeb6e2f3240a1fa6bc7cf1819f61dc5cc1865021ec0029a165627a7a723058203fd8c695e7d214830479c2df082c5d9ec48afa64086f8c06fc1dc44048dc6d3c0029`

// DeploySpawn deploys a new Ethereum contract, binding an instance of Spawn to it.
func DeploySpawn(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Spawn, error) {
	parsed, err := abi.JSON(strings.NewReader(SpawnABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SpawnBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Spawn{SpawnCaller: SpawnCaller{contract: contract}, SpawnTransactor: SpawnTransactor{contract: contract}}, nil
}

// Spawn is an auto generated Go binding around an Ethereum contract.
type Spawn struct {
	SpawnCaller     // Read-only binding to the contract
	SpawnTransactor // Write-only binding to the contract
}

// SpawnCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpawnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpawnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpawnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpawnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpawnSession struct {
	Contract     *Spawn            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpawnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpawnCallerSession struct {
	Contract *SpawnCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SpawnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpawnTransactorSession struct {
	Contract     *SpawnTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpawnRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpawnRaw struct {
	Contract *Spawn // Generic contract binding to access the raw methods on
}

// SpawnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpawnCallerRaw struct {
	Contract *SpawnCaller // Generic read-only contract binding to access the raw methods on
}

// SpawnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpawnTransactorRaw struct {
	Contract *SpawnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpawn creates a new instance of Spawn, bound to a specific deployed contract.
func NewSpawn(address common.Address, backend bind.ContractBackend) (*Spawn, error) {
	contract, err := bindSpawn(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Spawn{SpawnCaller: SpawnCaller{contract: contract}, SpawnTransactor: SpawnTransactor{contract: contract}}, nil
}

// NewSpawnCaller creates a new read-only instance of Spawn, bound to a specific deployed contract.
func NewSpawnCaller(address common.Address, caller bind.ContractCaller) (*SpawnCaller, error) {
	contract, err := bindSpawn(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SpawnCaller{contract: contract}, nil
}

// NewSpawnTransactor creates a new write-only instance of Spawn, bound to a specific deployed contract.
func NewSpawnTransactor(address common.Address, transactor bind.ContractTransactor) (*SpawnTransactor, error) {
	contract, err := bindSpawn(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SpawnTransactor{contract: contract}, nil
}

// bindSpawn binds a generic wrapper to an already deployed contract.
func bindSpawn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SpawnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spawn *SpawnRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Spawn.Contract.SpawnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spawn *SpawnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spawn.Contract.SpawnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spawn *SpawnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spawn.Contract.SpawnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spawn *SpawnCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Spawn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spawn *SpawnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spawn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spawn *SpawnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spawn.Contract.contract.Transact(opts, method, params...)
}

// AddressLUT is a free data retrieval call binding the contract method 0xaca49018.
//
// Solidity: function addressLUT( uint256) constant returns(address)
func (_Spawn *SpawnCaller) AddressLUT(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Spawn.contract.Call(opts, out, "addressLUT", arg0)
	return *ret0, err
}

// AddressLUT is a free data retrieval call binding the contract method 0xaca49018.
//
// Solidity: function addressLUT( uint256) constant returns(address)
func (_Spawn *SpawnSession) AddressLUT(arg0 *big.Int) (common.Address, error) {
	return _Spawn.Contract.AddressLUT(&_Spawn.CallOpts, arg0)
}

// AddressLUT is a free data retrieval call binding the contract method 0xaca49018.
//
// Solidity: function addressLUT( uint256) constant returns(address)
func (_Spawn *SpawnCallerSession) AddressLUT(arg0 *big.Int) (common.Address, error) {
	return _Spawn.Contract.AddressLUT(&_Spawn.CallOpts, arg0)
}

// Issues is a free data retrieval call binding the contract method 0xe9a5551c.
//
// Solidity: function issues( address) constant returns(threshold uint256, id string)
func (_Spawn *SpawnCaller) Issues(opts *bind.CallOpts, arg0 common.Address) (struct {
	Threshold *big.Int
	Id        string
}, error) {
	ret := new(struct {
		Threshold *big.Int
		Id        string
	})
	out := ret
	err := _Spawn.contract.Call(opts, out, "issues", arg0)
	return *ret, err
}

// Issues is a free data retrieval call binding the contract method 0xe9a5551c.
//
// Solidity: function issues( address) constant returns(threshold uint256, id string)
func (_Spawn *SpawnSession) Issues(arg0 common.Address) (struct {
	Threshold *big.Int
	Id        string
}, error) {
	return _Spawn.Contract.Issues(&_Spawn.CallOpts, arg0)
}

// Issues is a free data retrieval call binding the contract method 0xe9a5551c.
//
// Solidity: function issues( address) constant returns(threshold uint256, id string)
func (_Spawn *SpawnCallerSession) Issues(arg0 common.Address) (struct {
	Threshold *big.Int
	Id        string
}, error) {
	return _Spawn.Contract.Issues(&_Spawn.CallOpts, arg0)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() constant returns(uint256)
func (_Spawn *SpawnCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Spawn.contract.Call(opts, out, "size")
	return *ret0, err
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() constant returns(uint256)
func (_Spawn *SpawnSession) Size() (*big.Int, error) {
	return _Spawn.Contract.Size(&_Spawn.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() constant returns(uint256)
func (_Spawn *SpawnCallerSession) Size() (*big.Int, error) {
	return _Spawn.Contract.Size(&_Spawn.CallOpts)
}

// CreateIssue is a paid mutator transaction binding the contract method 0x8615d855.
//
// Solidity: function createIssue(id string, threshold uint256) returns(address)
func (_Spawn *SpawnTransactor) CreateIssue(opts *bind.TransactOpts, id string, threshold *big.Int) (*types.Transaction, error) {
	return _Spawn.contract.Transact(opts, "createIssue", id, threshold)
}

// CreateIssue is a paid mutator transaction binding the contract method 0x8615d855.
//
// Solidity: function createIssue(id string, threshold uint256) returns(address)
func (_Spawn *SpawnSession) CreateIssue(id string, threshold *big.Int) (*types.Transaction, error) {
	return _Spawn.Contract.CreateIssue(&_Spawn.TransactOpts, id, threshold)
}

// CreateIssue is a paid mutator transaction binding the contract method 0x8615d855.
//
// Solidity: function createIssue(id string, threshold uint256) returns(address)
func (_Spawn *SpawnTransactorSession) CreateIssue(id string, threshold *big.Int) (*types.Transaction, error) {
	return _Spawn.Contract.CreateIssue(&_Spawn.TransactOpts, id, threshold)
}

// TransferABI is the input ABI used to generate the binding from.
const TransferABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"payOut\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"funded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"}]"

// TransferBin is the compiled bytecode used for deploying new contracts.
const TransferBin = `0x6060604052341561000c57fe5b60405160208061023983398101604052515b600180546000839055600160a060020a03191633600160a060020a03161760a060020a60ff02191690555b505b6101df8061005a6000396000f3006060604052361561005f5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631922ff3981146100b157806342cde4e8146100c75780638da5cb5b146100e9578063f3a504f214610115575b6100af5b600054600160a060020a03301631106100ac576001805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790555b5b565b005b6100af600160a060020a0360043516610139565b005b34156100cf57fe5b6100d761017d565b60408051918252519081900360200190f35b34156100f157fe5b6100f9610183565b60408051600160a060020a039092168252519081900360200190f35b341561011d57fe5b610125610192565b604080519115158252519081900360200190f35b60015433600160a060020a0390811691161461017957604051600160a060020a0380831691309091163180156108fc02916000818181858888f150505050505b5b50565b60005481565b600154600160a060020a031681565b60015474010000000000000000000000000000000000000000900460ff16815600a165627a7a723058208f363b32654b5396ae379deeb6e2f3240a1fa6bc7cf1819f61dc5cc1865021ec0029`

// DeployTransfer deploys a new Ethereum contract, binding an instance of Transfer to it.
func DeployTransfer(auth *bind.TransactOpts, backend bind.ContractBackend, _threshold *big.Int) (common.Address, *types.Transaction, *Transfer, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TransferBin), backend, _threshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Transfer{TransferCaller: TransferCaller{contract: contract}, TransferTransactor: TransferTransactor{contract: contract}}, nil
}

// Transfer is an auto generated Go binding around an Ethereum contract.
type Transfer struct {
	TransferCaller     // Read-only binding to the contract
	TransferTransactor // Write-only binding to the contract
}

// TransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferSession struct {
	Contract     *Transfer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferCallerSession struct {
	Contract *TransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferTransactorSession struct {
	Contract     *TransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferRaw struct {
	Contract *Transfer // Generic contract binding to access the raw methods on
}

// TransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferCallerRaw struct {
	Contract *TransferCaller // Generic read-only contract binding to access the raw methods on
}

// TransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferTransactorRaw struct {
	Contract *TransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransfer creates a new instance of Transfer, bound to a specific deployed contract.
func NewTransfer(address common.Address, backend bind.ContractBackend) (*Transfer, error) {
	contract, err := bindTransfer(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transfer{TransferCaller: TransferCaller{contract: contract}, TransferTransactor: TransferTransactor{contract: contract}}, nil
}

// NewTransferCaller creates a new read-only instance of Transfer, bound to a specific deployed contract.
func NewTransferCaller(address common.Address, caller bind.ContractCaller) (*TransferCaller, error) {
	contract, err := bindTransfer(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TransferCaller{contract: contract}, nil
}

// NewTransferTransactor creates a new write-only instance of Transfer, bound to a specific deployed contract.
func NewTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferTransactor, error) {
	contract, err := bindTransfer(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TransferTransactor{contract: contract}, nil
}

// bindTransfer binds a generic wrapper to an already deployed contract.
func bindTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.TransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transact(opts, method, params...)
}

// Funded is a free data retrieval call binding the contract method 0xf3a504f2.
//
// Solidity: function funded() constant returns(bool)
func (_Transfer *TransferCaller) Funded(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Transfer.contract.Call(opts, out, "funded")
	return *ret0, err
}

// Funded is a free data retrieval call binding the contract method 0xf3a504f2.
//
// Solidity: function funded() constant returns(bool)
func (_Transfer *TransferSession) Funded() (bool, error) {
	return _Transfer.Contract.Funded(&_Transfer.CallOpts)
}

// Funded is a free data retrieval call binding the contract method 0xf3a504f2.
//
// Solidity: function funded() constant returns(bool)
func (_Transfer *TransferCallerSession) Funded() (bool, error) {
	return _Transfer.Contract.Funded(&_Transfer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Transfer *TransferCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Transfer.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Transfer *TransferSession) Owner() (common.Address, error) {
	return _Transfer.Contract.Owner(&_Transfer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Transfer *TransferCallerSession) Owner() (common.Address, error) {
	return _Transfer.Contract.Owner(&_Transfer.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() constant returns(uint256)
func (_Transfer *TransferCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Transfer.contract.Call(opts, out, "threshold")
	return *ret0, err
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() constant returns(uint256)
func (_Transfer *TransferSession) Threshold() (*big.Int, error) {
	return _Transfer.Contract.Threshold(&_Transfer.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() constant returns(uint256)
func (_Transfer *TransferCallerSession) Threshold() (*big.Int, error) {
	return _Transfer.Contract.Threshold(&_Transfer.CallOpts)
}

// PayOut is a paid mutator transaction binding the contract method 0x1922ff39.
//
// Solidity: function payOut(addr address) returns()
func (_Transfer *TransferTransactor) PayOut(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Transfer.contract.Transact(opts, "payOut", addr)
}

// PayOut is a paid mutator transaction binding the contract method 0x1922ff39.
//
// Solidity: function payOut(addr address) returns()
func (_Transfer *TransferSession) PayOut(addr common.Address) (*types.Transaction, error) {
	return _Transfer.Contract.PayOut(&_Transfer.TransactOpts, addr)
}

// PayOut is a paid mutator transaction binding the contract method 0x1922ff39.
//
// Solidity: function payOut(addr address) returns()
func (_Transfer *TransferTransactorSession) PayOut(addr common.Address) (*types.Transaction, error) {
	return _Transfer.Contract.PayOut(&_Transfer.TransactOpts, addr)
}
