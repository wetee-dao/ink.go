package client

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/wetee-dao/ink.go/pallet/revive"
	gtypes "github.com/wetee-dao/ink.go/pallet/types"
	"github.com/wetee-dao/ink.go/util"
)

// Revive module
type Ink interface {
	Client() *ChainClient
	ContractAddress() types.H160
}

func DryRun[T any](
	contractIns Ink,
	origin types.AccountID,
	amount types.U128,
	gas_limit util.Option[types.Weight],
	storage_deposit_limit util.Option[types.U128],
	contractInput util.InkContractInput,
) (*T, error) {
	inputBt, err := contractInput.Encode()
	if err != nil {
		return nil, errors.New("contractInput.Encode: " + err.Error())
	}

	client := contractIns.Client()
	addres := contractIns.ContractAddress()

	if client.Debug {
		util.LogWithPurple("[ TryCall contract ]", addres.Hex())
		util.LogWithPurple("[ TryCall   method ]", contractInput.Selector)
		util.LogWithPurple("[ TryCall   origin ]", origin.ToHexString())
		util.LogWithPurple("[ TryCall     args ]", "0x"+hex.EncodeToString(inputBt))
		util.LogWithPurple("[ TryCall gaslimit ]", gas_limit)
		util.LogWithPurple("[ TryCall  storage ]", storage_deposit_limit)
		fmt.Println("")
	}

	result := util.ContractResult{}
	err = client.CallRuntimeApi(
		"ReviveApi",
		"call",
		[]any{
			origin,
			addres,
			amount,
			gas_limit,
			storage_deposit_limit,
			inputBt,
		},
		&result,
	)
	if err != nil {
		return nil, errors.New("CallRuntimeApi: " + err.Error())
	}

	var returnValue *gtypes.ExecReturnValue
	if result.Result.IsErr {
		if result.Result.E.IsModule {
			merr := result.Result.E.AsModuleField0
			info, ierr := client.GetErrorInfo(merr.Index, merr.Error)
			if ierr == nil {
				err = errors.New("TryCall: Module Error" + info.Name)
			} else {
				err = errors.New("TryCall: unknown Module Error")
			}
			return nil, err
		}
		bt, _ := json.Marshal(result.Result.E)
		return nil, errors.New(string(bt))
	}

	returnValue = &result.Result.V
	data := new(T)
	scale.NewDecoder(bytes.NewReader(returnValue.Data)).Decode(data)

	// 判断是否执行错误
	if returnValue.Flags == 1 {
		return data, errors.New("contract reverted: the specific error information is in the second returned")
	}

	return data, nil
}

func Call(
	contractIns Ink,
	signer *Signer,
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
		util.LogWithYellow("[ Call   origin ]", "0x"+hex.EncodeToString(signer.PublicKey))
		util.LogWithYellow("[ Call     args ]", "0x"+hex.EncodeToString(inputBt))
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

	return client.SignAndSubmit(signer, call, false)
}

type DryRunCallParams struct {
	Origin              types.AccountID
	Amount              types.U128
	GasLimit            util.Option[types.Weight]
	StorageDepositLimit util.Option[types.U128]
}

type CallParams struct {
	Signer              *Signer
	Amount              types.U128
	GasLimit            types.Weight
	StorageDepositLimit types.U128
}
