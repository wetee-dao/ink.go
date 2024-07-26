package weteeworker

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for K8sClusterAccounts
//
//	用户对应集群的信息
//	user's K8sCluster information
func MakeK8sClusterAccountsStorageKey(byteArray320 [32]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteArray320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "K8sClusterAccounts", byteArgs...)
}
func GetK8sClusterAccounts(state state.State, bhash types.Hash, byteArray320 [32]byte) (ret uint64, isSome bool, err error) {
	key, err := MakeK8sClusterAccountsStorageKey(byteArray320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetK8sClusterAccountsLatest(state state.State, byteArray320 [32]byte) (ret uint64, isSome bool, err error) {
	key, err := MakeK8sClusterAccountsStorageKey(byteArray320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for NextClusterId id={{false [12]}}
//
//	The id of the next cluster to be created.
//	获取下一个集群id
func MakeNextClusterIdStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "NextClusterId")
}

var NextClusterIdResultDefaultBytes, _ = hex.DecodeString("0100000000000000")

func GetNextClusterId(state state.State, bhash types.Hash) (ret uint64, err error) {
	key, err := MakeNextClusterIdStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(NextClusterIdResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetNextClusterIdLatest(state state.State) (ret uint64, err error) {
	key, err := MakeNextClusterIdStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(NextClusterIdResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for CodeMrenclave id={{false [201]}}
//
//	代码版本
func MakeCodeMrenclaveStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "CodeMrenclave")
}

var CodeMrenclaveResultDefaultBytes, _ = hex.DecodeString("00")

func GetCodeMrenclave(state state.State, bhash types.Hash) (ret []byte, err error) {
	key, err := MakeCodeMrenclaveStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CodeMrenclaveResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetCodeMrenclaveLatest(state state.State) (ret []byte, err error) {
	key, err := MakeCodeMrenclaveStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CodeMrenclaveResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for CodeMrsigner id={{false [201]}}
//
//	代码打包签名人
func MakeCodeMrsignerStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "CodeMrsigner")
}

var CodeMrsignerResultDefaultBytes, _ = hex.DecodeString("00")

func GetCodeMrsigner(state state.State, bhash types.Hash) (ret []byte, err error) {
	key, err := MakeCodeMrsignerStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CodeMrsignerResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetCodeMrsignerLatest(state state.State) (ret []byte, err error) {
	key, err := MakeCodeMrsignerStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CodeMrsignerResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for K8sClusters
//
//	集群信息
func MakeK8sClustersStorageKey(uint640 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "K8sClusters", byteArgs...)
}
func GetK8sClusters(state state.State, bhash types.Hash, uint640 uint64) (ret types1.K8sCluster, isSome bool, err error) {
	key, err := MakeK8sClustersStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetK8sClustersLatest(state state.State, uint640 uint64) (ret types1.K8sCluster, isSome bool, err error) {
	key, err := MakeK8sClustersStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for ProofOfClusters
//
//	集群工作量证明
//	K8sCluster proof of work
func MakeProofOfClustersStorageKey(uint640 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "ProofOfClusters", byteArgs...)
}
func GetProofOfClusters(state state.State, bhash types.Hash, uint640 uint64) (ret []byte, isSome bool, err error) {
	key, err := MakeProofOfClustersStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetProofOfClustersLatest(state state.State, uint640 uint64) (ret []byte, isSome bool, err error) {
	key, err := MakeProofOfClustersStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for Crs
//
//	计算资源 抵押/使用
//	computing resource
func MakeCrsStorageKey(uint640 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "Crs", byteArgs...)
}
func GetCrs(state state.State, bhash types.Hash, uint640 uint64) (ret types1.TupleOfComCrComCr, isSome bool, err error) {
	key, err := MakeCrsStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetCrsLatest(state state.State, uint640 uint64) (ret types1.TupleOfComCrComCr, isSome bool, err error) {
	key, err := MakeCrsStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for Scores
//
//	节点(评级,评分)
//	computing resource
func MakeScoresStorageKey(uint640 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "Scores", byteArgs...)
}
func GetScores(state state.State, bhash types.Hash, uint640 uint64) (ret types1.TupleOfByteByte, isSome bool, err error) {
	key, err := MakeScoresStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetScoresLatest(state state.State, uint640 uint64) (ret types1.TupleOfByteByte, isSome bool, err error) {
	key, err := MakeScoresStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for DepositPrices
//
//	抵押价格
//	deposit of computing resource
func MakeDepositPricesStorageKey(byte0 byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byte0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "DepositPrices", byteArgs...)
}
func GetDepositPrices(state state.State, bhash types.Hash, byte0 byte) (ret types1.DepositPrice, isSome bool, err error) {
	key, err := MakeDepositPricesStorageKey(byte0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetDepositPricesLatest(state state.State, byte0 byte) (ret types1.DepositPrice, isSome bool, err error) {
	key, err := MakeDepositPricesStorageKey(byte0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for Deposits
//
//	抵押信息
//	deposit of computing resource
func MakeDepositsStorageKey(tupleOfUint64Uint640 uint64, tupleOfUint64Uint641 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfUint64Uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfUint64Uint641)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "Deposits", byteArgs...)
}
func GetDeposits(state state.State, bhash types.Hash, tupleOfUint64Uint640 uint64, tupleOfUint64Uint641 uint64) (ret types1.Deposit, isSome bool, err error) {
	key, err := MakeDepositsStorageKey(tupleOfUint64Uint640, tupleOfUint64Uint641)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetDepositsLatest(state state.State, tupleOfUint64Uint640 uint64, tupleOfUint64Uint641 uint64) (ret types1.Deposit, isSome bool, err error) {
	key, err := MakeDepositsStorageKey(tupleOfUint64Uint640, tupleOfUint64Uint641)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for ClusterContracts
//
//	集群包含的智能合同
//	smart contract
func MakeClusterContractsStorageKey(tupleOfUint64WorkId0 uint64, tupleOfUint64WorkId1 types1.WorkId) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfUint64WorkId0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfUint64WorkId1)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "ClusterContracts", byteArgs...)
}
func GetClusterContracts(state state.State, bhash types.Hash, tupleOfUint64WorkId0 uint64, tupleOfUint64WorkId1 types1.WorkId) (ret types1.ClusterContractState, isSome bool, err error) {
	key, err := MakeClusterContractsStorageKey(tupleOfUint64WorkId0, tupleOfUint64WorkId1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetClusterContractsLatest(state state.State, tupleOfUint64WorkId0 uint64, tupleOfUint64WorkId1 types1.WorkId) (ret types1.ClusterContractState, isSome bool, err error) {
	key, err := MakeClusterContractsStorageKey(tupleOfUint64WorkId0, tupleOfUint64WorkId1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for WorkContracts
//
//	程序使用的智能合同 （节点id，解锁)
//	smart contract
func MakeWorkContractsStorageKey(workId0 types1.WorkId) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(workId0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "WorkContracts", byteArgs...)
}
func GetWorkContracts(state state.State, bhash types.Hash, workId0 types1.WorkId) (ret uint64, isSome bool, err error) {
	key, err := MakeWorkContractsStorageKey(workId0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetWorkContractsLatest(state state.State, workId0 types1.WorkId) (ret uint64, isSome bool, err error) {
	key, err := MakeWorkContractsStorageKey(workId0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for WorkContractState
//
//	程序使用的智能合同日志 （节点id，日志）
//	smart contract log
func MakeWorkContractStateStorageKey(tupleOfWorkIdUint640 types1.WorkId, tupleOfWorkIdUint641 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfWorkIdUint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfWorkIdUint641)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "WorkContractState", byteArgs...)
}
func GetWorkContractState(state state.State, bhash types.Hash, tupleOfWorkIdUint640 types1.WorkId, tupleOfWorkIdUint641 uint64) (ret types1.ContractState, isSome bool, err error) {
	key, err := MakeWorkContractStateStorageKey(tupleOfWorkIdUint640, tupleOfWorkIdUint641)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetWorkContractStateLatest(state state.State, tupleOfWorkIdUint640 types1.WorkId, tupleOfWorkIdUint641 uint64) (ret types1.ContractState, isSome bool, err error) {
	key, err := MakeWorkContractStateStorageKey(tupleOfWorkIdUint640, tupleOfWorkIdUint641)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for Stage id={{false [4]}}
//
//	Work 结算周期
//	Work settle period
func MakeStageStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "Stage")
}

var StageResultDefaultBytes, _ = hex.DecodeString("58020000")

func GetStage(state state.State, bhash types.Hash) (ret uint32, err error) {
	key, err := MakeStageStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(StageResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetStageLatest(state state.State) (ret uint32, err error) {
	key, err := MakeStageStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(StageResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for ProofsOfWork
//
//	工作任务工作量证明
//	proof of work of task
func MakeProofsOfWorkStorageKey(tupleOfWorkIdUint640 types1.WorkId, tupleOfWorkIdUint641 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfWorkIdUint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfWorkIdUint641)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "ProofsOfWork", byteArgs...)
}
func GetProofsOfWork(state state.State, bhash types.Hash, tupleOfWorkIdUint640 types1.WorkId, tupleOfWorkIdUint641 uint64) (ret types1.ProofOfWork, isSome bool, err error) {
	key, err := MakeProofsOfWorkStorageKey(tupleOfWorkIdUint640, tupleOfWorkIdUint641)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetProofsOfWorkLatest(state state.State, tupleOfWorkIdUint640 types1.WorkId, tupleOfWorkIdUint641 uint64) (ret types1.ProofOfWork, isSome bool, err error) {
	key, err := MakeProofsOfWorkStorageKey(tupleOfWorkIdUint640, tupleOfWorkIdUint641)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for ReportOfWork
//
//	work report
func MakeReportOfWorkStorageKey(workId0 types1.WorkId) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(workId0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "ReportOfWork", byteArgs...)
}
func GetReportOfWork(state state.State, bhash types.Hash, workId0 types1.WorkId) (ret []byte, isSome bool, err error) {
	key, err := MakeReportOfWorkStorageKey(workId0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetReportOfWorkLatest(state state.State, workId0 types1.WorkId) (ret []byte, isSome bool, err error) {
	key, err := MakeReportOfWorkStorageKey(workId0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for ReportOfWorkTime
//
//	work report
func MakeReportOfWorkTimeStorageKey(workId0 types1.WorkId) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(workId0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "ReportOfWorkTime", byteArgs...)
}
func GetReportOfWorkTime(state state.State, bhash types.Hash, workId0 types1.WorkId) (ret uint64, isSome bool, err error) {
	key, err := MakeReportOfWorkTimeStorageKey(workId0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetReportOfWorkTimeLatest(state state.State, workId0 types1.WorkId) (ret uint64, isSome bool, err error) {
	key, err := MakeReportOfWorkTimeStorageKey(workId0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for Reports
//
//	投诉信息
//	reports of work / cluster
func MakeReportsStorageKey(tupleOfUint64WorkId0 uint64, tupleOfUint64WorkId1 types1.WorkId) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfUint64WorkId0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfUint64WorkId1)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEWorker", "Reports", byteArgs...)
}
func GetReports(state state.State, bhash types.Hash, tupleOfUint64WorkId0 uint64, tupleOfUint64WorkId1 types1.WorkId) (ret []byte, isSome bool, err error) {
	key, err := MakeReportsStorageKey(tupleOfUint64WorkId0, tupleOfUint64WorkId1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetReportsLatest(state state.State, tupleOfUint64WorkId0 uint64, tupleOfUint64WorkId1 types1.WorkId) (ret []byte, isSome bool, err error) {
	key, err := MakeReportsStorageKey(tupleOfUint64WorkId0, tupleOfUint64WorkId1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
