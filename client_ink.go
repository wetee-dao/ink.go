package client

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/wetee-dao/ink.go/pallet/revive"
	gtypes "github.com/wetee-dao/ink.go/pallet/types"
	"github.com/wetee-dao/ink.go/util"
)

var ErrContractReverted = errors.New("contract reverted: the specific error information is in the second returned")

// Revive module
type Ink interface {
	Client() *ChainClient
	ContractAddress() types.H160
}

// Dry run contract
func DryRunInk[T any](
	contractIns Ink,
	origin types.AccountID,
	amount types.U128,
	gas_limit util.Option[types.Weight],
	storage_deposit_limit util.Option[types.U128],
	contractInput util.InkContractInput,
) (*T, *DryRunReturnGas, error) {
	inputBt, err := contractInput.Encode()
	if err != nil {
		return nil, nil, errors.New("contractInput.Encode: " + err.Error())
	}

	client := contractIns.Client()
	addres := contractIns.ContractAddress()

	if client.Debug {
		util.LogWithPurple("[ DryRun contract ]", addres.Hex())
		util.LogWithPurple("[ DryRun   method ]", contractInput.Selector)
		util.LogWithPurple("[ DryRun   origin ]", origin.ToHexString())
		util.LogWithPurple("[ DryRun     args ]", "0x"+hex.EncodeToString(inputBt))
		util.LogWithPurple("[ DryRun gaslimit ]", gas_limit)
		util.LogWithPurple("[ DryRun  storage ]", storage_deposit_limit)
		fmt.Println("")
	}

	result := util.ContractResult{}
	err = client.CallRuntimeApi(
		"ReviveApi",
		"call",
		[]any{
			origin, addres, amount, gas_limit, storage_deposit_limit, inputBt,
		},
		&result,
	)
	if err != nil {
		return nil, nil, errors.New("CallRuntimeApi: " + err.Error())
	}

	var returnValue *gtypes.ExecReturnValue
	if result.Result.IsErr {
		if result.Result.E.IsModule {
			merr := result.Result.E.AsModuleField0
			info, ierr := client.GetErrorInfo(merr.Index, merr.Error)
			if ierr == nil {
				err = errors.New("DryRun: Module Error" + info.Name)
			} else {
				err = errors.New("DryRun: unknown Module Error")
			}
			return nil, nil, err
		}
		bt, _ := json.Marshal(result.Result.E)
		return nil, nil, errors.New(string(bt))
	}

	// 获取返回值
	returnValue = &result.Result.V
	data := new(T)
	err = scale.NewDecoder(bytes.NewReader(returnValue.Data[1:])).Decode(data)
	if err != nil {
		return nil, nil, errors.New("DryRun scale.NewDecoder.Decode: " + err.Error())
	}

	// 判断是否执行错误
	if returnValue.Flags == 1 {
		return data, nil, ErrContractReverted
	}

	var storageDeposit types.U128
	if result.StorageDeposit.IsCharge {
		storageDeposit = result.StorageDeposit.AsChargeField0
	}

	return data, &DryRunReturnGas{
		GasConsumed: types.Weight{
			RefTime:   result.GasConsumed.RefTime,
			ProofSize: result.GasConsumed.ProofSize,
		},
		GasRequired: types.Weight{
			RefTime:   result.GasRequired.RefTime,
			ProofSize: result.GasRequired.ProofSize,
		},
		StorageDeposit: storageDeposit,
	}, nil
}

// Call contract use substrate api
func CallInk(
	contractIns Ink,
	signer SignerType,
	amount types.U128,
	gas_limit types.Weight,
	storage_deposit_limit types.U128,
	contractInput util.InkContractInput,
) error {
	inputBt, err := contractInput.Encode()
	if err != nil {
		return errors.New("contractInput.Encode: " + err.Error())
	}

	client := contractIns.Client()
	addres := contractIns.ContractAddress()

	if client.Debug {
		util.LogWithYellow("[ Call contract ]", addres.Hex())
		util.LogWithYellow("[ Call   method ]", contractInput.Selector)
		util.LogWithYellow("[ Call   origin ]", "0x"+hex.EncodeToString(signer.Public()))
		util.LogWithYellow("[ Call     args ]", "0x"+hex.EncodeToString(inputBt))
		util.LogWithYellow("[       RefTime ]", gas_limit.RefTime.Int64())
		util.LogWithYellow("[     ProofSize ]", gas_limit.ProofSize.Int64())
		util.LogWithYellow("[  DepositLimit ]", storage_deposit_limit.Int.String())
		fmt.Println("")
	}

	runtimeCall := revive.MakeCallCall(
		addres,
		types.NewUCompact(amount.Int),
		gtypes.Weight{
			RefTime:   gas_limit.RefTime,
			ProofSize: gas_limit.ProofSize,
		},
		types.NewUCompact(storage_deposit_limit.Int),
		inputBt,
	)

	call, err := (runtimeCall).AsCall()
	if err != nil {
		return errors.New("(runtimeCall).AsCall() error: " + err.Error())
	}

	return client.SignAndSubmit(signer, call, true)
}

func TxCall(
	contractIns Ink,
	signer SignerType,
	amount types.U128,
	gas_limit types.Weight,
	storage_deposit_limit types.U128,
	contractInput util.InkContractInput,
) (*types.Call, error) {
	inputBt, err := contractInput.Encode()
	if err != nil {
		return nil, errors.New("contractInput.Encode: " + err.Error())
	}

	client := contractIns.Client()
	addres := contractIns.ContractAddress()

	if client.Debug {
		util.LogWithYellow("[ Call contract ]", addres.Hex())
		util.LogWithYellow("[ Call   method ]", contractInput.Selector)
		util.LogWithYellow("[ Call   origin ]", "0x"+hex.EncodeToString(signer.Public()))
		util.LogWithYellow("[ Call     args ]", "0x"+hex.EncodeToString(inputBt))
		util.LogWithYellow("[       RefTime ]", gas_limit.RefTime.Int64())
		util.LogWithYellow("[     ProofSize ]", gas_limit.ProofSize.Int64())
		util.LogWithYellow("[  DepositLimit ]", storage_deposit_limit.Int.String())
		fmt.Println("")
	}

	runtimeCall := revive.MakeCallCall(
		addres,
		types.NewUCompact(amount.Int),
		gtypes.Weight{
			RefTime:   gas_limit.RefTime,
			ProofSize: gas_limit.ProofSize,
		},
		types.NewUCompact(storage_deposit_limit.Int),
		inputBt,
	)

	call, err := (runtimeCall).AsCall()
	return &call, err
}

// DryRun param of DryRun
type DryRunCallParams struct {
	Origin              types.AccountID
	PayAmount           types.U128
	GasLimit            util.Option[types.Weight]
	StorageDepositLimit util.Option[types.U128]
}

func DefaultParamWithOragin(origin types.AccountID) DryRunCallParams {
	var defaultParam = DryRunCallParams{
		PayAmount:           types.NewU128(*big.NewInt(0)),
		GasLimit:            util.NewNone[types.Weight](),
		StorageDepositLimit: util.NewNone[types.U128](),
	}
	defaultParam.Origin = origin
	return defaultParam
}

// DryRun return gas consumed
type DryRunReturnGas struct {
	GasConsumed    types.Weight
	GasRequired    types.Weight
	StorageDeposit types.U128
}

// Call param of Call
type CallParams struct {
	Signer    SignerType
	PayAmount types.U128
}
