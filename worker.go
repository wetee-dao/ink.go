package client

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/wetee-dao/go-sdk/gen/weteeworker"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"

	gtypes "github.com/wetee-dao/go-sdk/gen/types"
)

// Worker
type Worker struct {
	Client *ChainClient
	Signer *signature.KeyringPair
}

// 集群注册
// Cluster register
func (w *Worker) ClusterRegister(name string, ip []uint8, port uint32, level uint8, untilFinalized bool) error {
	runtimeCall := weteeworker.MakeClusterRegisterCall(
		[]byte(name),
		[]gtypes.Ip{
			{
				Ipv4: gtypes.OptionTUint32{IsNone: false, IsSome: true, AsSomeField0: 100},
				Ipv6: gtypes.OptionTU128{IsNone: false, IsSome: true, AsSomeField0: types.NewU128(*big.NewInt(100))},
			},
		},
		port,
		level,
	)

	return w.Client.SignAndSubmit(w.Signer, runtimeCall, untilFinalized)
}

// 集群抵押
// Cluster mortgage
func (w *Worker) ClusterMortgage(id uint64, cpu uint32, mem uint32, disk uint32, deposit uint64, untilFinalized bool) error {
	d := big.NewInt(0)
	d.SetUint64(deposit)
	runtimeCall := weteeworker.MakeClusterMortgageCall(
		id,
		cpu,
		mem,
		disk,
		types.UCompact(*d),
	)

	return w.Client.SignAndSubmit(w.Signer, runtimeCall, untilFinalized)
}

func (w *Worker) ClusterWithdrawal(id gtypes.WorkId, val int64, untilFinalized bool) error {
	runtimeCall := weteeworker.MakeClusterWithdrawalCall(
		id,
		types.NewU128(*big.NewInt(val)),
	)

	return w.Client.SignAndSubmit(w.Signer, runtimeCall, untilFinalized)
}

func (w *Worker) ClusterUnmortgage(clusterID uint64, id uint64, untilFinalized bool) error {
	runtimeCall := weteeworker.MakeClusterUnmortgageCall(
		clusterID,
		id,
	)

	return w.Client.SignAndSubmit(w.Signer, runtimeCall, untilFinalized)
}

func (w *Worker) ClusterStop(clusterID uint64, untilFinalized bool) error {
	runtimeCall := weteeworker.MakeClusterStopCall(
		clusterID,
	)

	return w.Client.SignAndSubmit(w.Signer, runtimeCall, untilFinalized)
}

func (w *Worker) Getk8sClusterAccounts(publey []byte) (uint64, error) {
	if len(publey) != 32 {
		return 0, errors.New("publey length error")
	}

	var mss [32]byte
	copy(mss[:], publey)

	res, ok, err := weteeworker.GetK8sClusterAccountsLatest(w.Client.Api.RPC.State, mss)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errors.New("GetK8sClusterAccountsLatest => not start")
	}
	return res, nil
}

func (w *Worker) GetClusterContracts(clusterID uint64, at *types.Hash) ([]ContractStateWrap, error) {
	var pallet, method = "WeteeWorker", "ClusterContracts"
	set, err := w.Client.QueryDoubleMapAll(pallet, method, clusterID, at)
	if err != nil {
		return nil, err
	}

	var list []ContractStateWrap = make([]ContractStateWrap, 0, len(set))
	for _, elem := range set {
		for _, change := range elem.Changes {
			var cs gtypes.ClusterContractState
			// key := change.StorageKey
			// prefix, err := w.Client.GetDoubleMapPrefixKey(pallet, method, clusterID)
			// if err != nil {
			// 	fmt.Println(err)
			// 	continue
			// }

			// key = key[len(prefix):]
			// fmt.Println(key, len(key))

			// hashers, err := w.Client.GetHashers(pallet, method)
			// if err != nil {
			// 	return nil, err
			// }

			if err := codec.Decode(change.StorageData, &cs); err != nil {
				fmt.Println(err)
				continue
			}
			// head, _ := w.Client.Api.RPC.Chain.GetHeader(elem.Block)
			list = append(list, ContractStateWrap{
				BlockHash:     elem.Block.Hex(),
				ContractState: &cs,
			})
		}
	}

	fmt.Println(err)
	return list, nil
}

func (w *Worker) ClusterProofUpload(id uint64, proof []byte, untilFinalized bool) error {
	runtimeCall := weteeworker.MakeClusterProofUploadCall(
		id,
		proof,
	)
	return w.Client.SignAndSubmit(w.Signer, runtimeCall, untilFinalized)
}

func (w *Worker) WorkProofUpload(workId gtypes.WorkId, logHash []byte, crHash []byte, cr gtypes.Cr, pubkey []byte, untilFinalized bool) error {
	hasHash := false
	if len(logHash) > 0 && len(crHash) > 0 {
		hasHash = true
	}
	hasReport := false
	if len(pubkey) > 0 {
		hasReport = true
	}
	runtimeCall := weteeworker.MakeWorkProofUploadCall(
		workId,
		gtypes.OptionTProofOfWork{
			IsNone: !hasHash,
			IsSome: hasHash,
			AsSomeField0: gtypes.ProofOfWork{
				LogHash: logHash,
				CrHash:  crHash,
				Cr:      cr,
			},
		},
		gtypes.OptionTByteSlice{
			IsNone:       !hasReport,
			IsSome:       hasReport,
			AsSomeField0: pubkey,
		},
	)
	return w.Client.SignAndSubmit(w.Signer, runtimeCall, untilFinalized)
}

func (w *Worker) GetStage() (uint32, error) {
	return weteeworker.GetStageLatest(w.Client.Api.RPC.State)
}

func (w *Worker) GetWorkContract(workId gtypes.WorkId, id uint64) (*gtypes.ContractState, error) {
	res, ok, err := weteeworker.GetWorkContractStateLatest(w.Client.Api.RPC.State, workId, id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("GetAppIdAccountsLatest => not ok")
	}
	return &res, nil

}

type ContractStateWrap struct {
	BlockHash     string
	ContractState *gtypes.ClusterContractState
}
