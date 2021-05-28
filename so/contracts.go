package so

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/so/interfaces"
)

func init() {
	vm.GDPrecompiledContracts[interfaces.GoodDataSrvAddr] = &GoodDataContract{}
}

//GetEventByName ...
func GetEventByName(abi abi.ABI, name string) *abi.Event {
	for _, e := range abi.Events {
		if e.Name == name {
			return &e
		}
	}
	return nil
}

// AddLog add log into chain.
func AddLog(evm *vm.EVM, contract *vm.Contract, ev *abi.Event, args ...interface{}) error {
	if len(ev.Inputs) != len(args) {
		return fmt.Errorf("Event %s expect %d arguments, but got %d", ev.Name, len(ev.Inputs), len(args))
	}
	tlog := types.Log{Address: contract.Address()}
	topics := []common.Hash{ev.ID()}

	nargs := []interface{}{}
	for j, a := range ev.Inputs {
		if a.Indexed {
			if data, err := abi.Arguments([]abi.Argument{a}).Pack(args[j]); err == nil {
				topics = append(topics, common.BytesToHash(data))
			} else {
				return fmt.Errorf("Event %s invalid log argument: %v %v %v", ev.Name, a.Type.Kind, a.Type.Type, a.Type.Size)
			}
		} else {
			nargs = append(nargs, args[j])
		}
	}
	if data, err := ev.Inputs.NonIndexed().PackValues(nargs); err == nil {
		tlog.Data = data
	} else {
		return fmt.Errorf("Event %s invalid log arguments", ev.Name)
	}
	tlog.Topics = topics
	tlog.BlockNumber = evm.BlockNumber.Uint64()

	evm.StateDB.AddLog(&tlog)
	return nil
}

// GoodDataContract implements vm.GDPrecompiledContract interface.
// This struct's method will be precompiled.
type GoodDataContract struct {
	evm      *vm.EVM
	contract *vm.Contract
}

//RequiredGas - testnet returns 0, no extra fees are needed
func (gd *GoodDataContract) RequiredGas(intput []byte) uint64 {
	return 0
}

//Run entry of GoodDataContract contract
func (gd *GoodDataContract) Run(evm *vm.EVM, contract *vm.Contract, input []byte) ([]byte, error) {
	if len(input) < 4 {
		return nil, nil
	}
	if evm.StateDB.GetBalance(contract.Address()).Cmp(big.NewInt(0)) == 0 {
		//Set 1 wei to avoid state not saving to chain db
		evm.StateDB.AddBalance(contract.Address(), big.NewInt(1))
	}
	gd.evm = evm
	gd.contract = contract

	method, err := interfaces.GdSrvAbi.MethodById(input[:4])
	if err == nil {
		params, err := method.Inputs.UnpackValues(input[4:])
		if err != nil {
			return nil, fmt.Errorf("method inputs unpackValues err: %v", err)
		}
		switch method.Name {
		// QSC
		case "registerQC":
			err := interfaces.ValidateParamsType(params, []interface{}{[]byte{}, []byte{}})
			if err != nil {
				return nil, fmt.Errorf("registerQC params err: %v", err)
			}
			err = gd.eventQCRegistered(params[0].([]byte), params[1].([]byte))
			if err != nil {
				return nil, err
			}
		case "assignQueryToMPCs":
			err := interfaces.ValidateParamsType(params, []interface{}{[]byte{}, []byte{}})
			if err != nil {
				return nil, fmt.Errorf("assignQueryToMPCs params err: %v", err)
			}
			err = gd.eventAssignQueryToMPCs(params[0].([]byte), params[1].([]byte))
			if err != nil {
				return nil, err
			}
		case "assignQueryToDOs":
			err := interfaces.ValidateParamsType(params, []interface{}{[]byte{}, []byte{}})
			if err != nil {
				return nil, fmt.Errorf("assignQueryToDOs params err: %v", err)
			}
			err = gd.eventAssignQueryToDOs(params[0].([]byte), params[1].([]byte))
			if err != nil {
				return nil, err
			}
		case "submitQuery":
			err := interfaces.ValidateParamsType(
				[]interface{}{params[0], params[1], params[2]},
				[]interface{}{[]byte{}, []byte{}, int64(0)})
			if err != nil {
				return nil, fmt.Errorf("submitQuery params err: %v", err)
			}
			err = gd.eventQuerySubmitted(params[0].([]byte), params[1].([]byte), params[2].(int64))
			if err != nil {
				return nil, err
			}
		case "finishQueryInOneDO":
			err := interfaces.ValidateParamsType(params, []interface{}{[]byte{}, []byte{}})
			if err != nil {
				return nil, fmt.Errorf("finishQueryInOneDO params err: %v", err)
			}
			err = gd.eventQueryFinishedInOneDO(params[0].([]byte), params[1].([]byte))
			if err != nil {
				return nil, err
			}
		case "finishQuery":
			err := interfaces.ValidateParamsType(params, []interface{}{[]byte{}})
			if err != nil {
				return nil, fmt.Errorf("finishQuery params err: %v", err)
			}
			err = gd.eventQueryFinished(params[0].([]byte))
			if err != nil {
				return nil, err
			}
		case "submitPrediction":
			err := interfaces.ValidateParamsType(
				[]interface{}{params[0], params[1]},
				[]interface{}{[]byte{}, []byte{}})
			if err != nil {
				return nil, fmt.Errorf("submitPrediction params err: %v", err)
			}
			err = gd.eventPredictionSubmitted(params[0].([]byte), params[1].([]byte))
			if err != nil {
				return nil, err
			}
		case "finishPrediction":
			err := interfaces.ValidateParamsType(
				[]interface{}{params[0]},
				[]interface{}{[]byte{}})
			if err != nil {
				return nil, fmt.Errorf("finishPrediction params err: %v", err)
			}
			err = gd.eventPredictionFinished(params[0].([]byte))
			if err != nil {
				return nil, err
			}
		// DSC
		case "registerDO":
			err := interfaces.ValidateParamsType(params, []interface{}{[]byte{}, []byte{}})
			if err != nil {
				return nil, fmt.Errorf("registerDO params err: %v", err)
			}
			err = gd.eventDORegistered(params[0].([]byte), params[1].([]byte))
			if err != nil {
				return nil, err
			}
		// MSC
		case "registerMPC":
			err := interfaces.ValidateParamsType(params, []interface{}{[]byte{}, []byte{}})
			if err != nil {
				return nil, fmt.Errorf("registerMSC params err: %v", err)
			}
			err = gd.eventMPCRegistered(params[0].([]byte), params[1].([]byte))
			if err != nil {
				return nil, err
			}
		// Other
		case "addLog":
			err := interfaces.ValidateParamsType(params, []interface{}{[]byte{}, []byte{}})
			if err != nil {
				return nil, fmt.Errorf("addLog params err: %v", err)
			}
			err = gd.eventLogAdded(params[0].([]byte), params[1].([]byte))
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}

func (gd *GoodDataContract) eventQCRegistered(uuid, publicKey []byte) error {
	return AddLog(gd.evm, gd.contract,
		GetEventByName(interfaces.GdSrvAbi, "QCRegistered"), uuid, publicKey)
}

func (gd *GoodDataContract) eventAssignQueryToMPCs(queryUUID, mpcUUIDs []byte) error {
	return AddLog(gd.evm, gd.contract,
		GetEventByName(interfaces.GdSrvAbi, "MPCsOfQuery"), queryUUID, mpcUUIDs)
}

func (gd *GoodDataContract) eventAssignQueryToDOs(queryUUID, doUUIDs []byte) error {
	return AddLog(gd.evm, gd.contract,
		GetEventByName(interfaces.GdSrvAbi, "DOsOfQuery"), queryUUID, doUUIDs)
}

func (gd *GoodDataContract) eventQuerySubmitted(queryUUID, qcUUID []byte, dataSet int64) error {
	return AddLog(gd.evm, gd.contract,
		GetEventByName(interfaces.GdSrvAbi, "QuerySubmitted"), queryUUID, qcUUID, dataSet)
}

func (gd *GoodDataContract) eventQueryFinishedInOneDO(queryUUID, doUUID []byte) error {
	return AddLog(gd.evm, gd.contract,
		GetEventByName(interfaces.GdSrvAbi, "QueryFinishedInOneDO"), queryUUID, doUUID)
}

func (gd *GoodDataContract) eventQueryFinished(queryUUID []byte) error {
	return AddLog(gd.evm, gd.contract,
		GetEventByName(interfaces.GdSrvAbi, "QueryFinished"), queryUUID)
}

func (gd *GoodDataContract) eventPredictionSubmitted(queryUUID, predictionUUID []byte) error {
	return AddLog(gd.evm, gd.contract,
		GetEventByName(interfaces.GdSrvAbi, "PredictionSubmitted"), queryUUID, predictionUUID)
}

func (gd *GoodDataContract) eventPredictionFinished(predictionUUID []byte) error {
	return AddLog(gd.evm, gd.contract,
		GetEventByName(interfaces.GdSrvAbi, "PredictionFinished"), predictionUUID)
}

func (gd *GoodDataContract) eventDORegistered(uuid, publicKey []byte) error {
	return AddLog(gd.evm, gd.contract,
		GetEventByName(interfaces.GdSrvAbi, "DORegistered"), uuid, publicKey)
}

func (gd *GoodDataContract) eventMPCRegistered(uuid, publicKey []byte) error {
	return AddLog(gd.evm, gd.contract,
		GetEventByName(interfaces.GdSrvAbi, "MPCRegistered"), uuid, publicKey)
}

func (gd *GoodDataContract) eventLogAdded(uuid, payload []byte) error {
	return AddLog(gd.evm, gd.contract,
		GetEventByName(interfaces.GdSrvAbi, "LogAdded"), uuid, payload)
}

//DataStore evm data store
type DataStore struct {
	Db   vm.StateDB
	Addr common.Address
}

var lengthSuffix = "length"
var indexSuffix = "index"

//Exists - check data has been saved or not
func (ds *DataStore) Exists(hash common.Hash) bool {
	lhash := common.Hash(sha256.Sum256(append(append([]byte(nil), hash[:]...), lengthSuffix...)))
	state := ds.Db.GetState(ds.Addr, lhash)

	return state != common.Hash{}
}

func (ds *DataStore) Write(hash common.Hash, data []byte) {
	lhash := common.Hash(sha256.Sum256(append(append([]byte(nil), hash[:]...), lengthSuffix...)))
	length := len(data)
	for offset := 0; offset < length; offset += 32 {
		ihash := common.Hash(sha256.Sum256(append(append([]byte(nil), hash.Bytes()...),
			[]byte(fmt.Sprintf("%d_%s", offset/32, indexSuffix))...)))
		end := offset + 32
		if end > length {
			end = length
		}
		ds.Db.SetState(ds.Addr, ihash, common.BytesToHash(data[offset:end]))
	}
	ds.Db.SetState(ds.Addr, lhash, common.BigToHash(big.NewInt(int64(length))))
}

func (ds *DataStore) Read(hash common.Hash) []byte {
	lhash := common.Hash(sha256.Sum256(append(append([]byte(nil), hash[:]...), lengthSuffix...)))
	length := int(ds.Db.GetState(ds.Addr, lhash).Big().Int64())
	data := make([]byte, length)
	for offset := 0; offset < length; offset += 32 {
		ihash := common.Hash(sha256.Sum256(append(append([]byte(nil), hash.Bytes()...),
			[]byte(fmt.Sprintf("%d_%s", offset/32, indexSuffix))...)))
		val := ds.Db.GetState(ds.Addr, ihash)
		end := offset + 32
		if end > length {
			end = length
			copy(data[offset:end], val.Bytes()[32-end%32:])
		} else {
			copy(data[offset:end], val.Bytes())
		}
	}
	return data
}
