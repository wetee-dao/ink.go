package client

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"

	"github.com/wetee-dao/go-sdk/pallet/revive"
	gtypes "github.com/wetee-dao/go-sdk/pallet/types"
	"github.com/wetee-dao/go-sdk/util"
)

// Revive module
type Revive struct {
	Client  *ChainClient
	Abi     *util.InkAbi
	Address types.H160
	Debug   bool
}

func NewRevive(client *ChainClient, address types.H160, abiRaw []byte) (*Revive, error) {
	abi, err := util.InitAbi(abiRaw)
	if err != nil {
		return nil, err
	}

	return &Revive{
		Client:  client,
		Abi:     abi,
		Address: address,
		Debug:   client.Debug,
	}, nil
}

// Query ink contract data
func (r *Revive) Query(
	origin types.AccountID,
	amount types.U128,
	gas_limit util.Option[types.Weight],
	storage_deposit_limit util.Option[types.U128],
	contractInput util.InkContractInput,
	returnValue any,
) error {
	result, err := r.DryRun(origin, amount, gas_limit, storage_deposit_limit, contractInput)
	if err != nil {
		return err
	}

	if result.Return.Data == nil {
		return errors.New("result data is nil")
	}

	return scale.NewDecoder(bytes.NewReader(result.Return.Data[1:])).Decode(returnValue)
}

// Query ink contract data
func QueryInk[T any](
	ABI *Revive,
	origin types.AccountID,
	amount types.U128,
	gas_limit util.Option[types.Weight],
	storage_deposit_limit util.Option[types.U128],
	contractInput util.InkContractInput,
) (*T, error) {
	returnValue := new(T)
	result, err := ABI.DryRun(origin, amount, gas_limit, storage_deposit_limit, contractInput)
	if err != nil {
		return nil, err
	}

	if result.Return.Data == nil {
		return nil, errors.New("result data is nil")
	}

	err = scale.NewDecoder(bytes.NewReader(result.Return.Data[1:])).Decode(returnValue)

	return returnValue, err
}

// try call ink contract
func (r *Revive) DryRun(
	origin types.AccountID,
	amount types.U128,
	gas_limit util.Option[types.Weight],
	storage_deposit_limit util.Option[types.U128],
	contractInput util.InkContractInput,
) (*util.DryRunResult, error) {
	selector := util.FuncToSelector(contractInput.Selector)
	f, err := r.GetArgsFromABI(selector)
	if err != nil {
		return nil, err
	}

	if len(f.Args) != len(contractInput.Args) {
		return nil, errors.New("Call args length not match, ABI expect: " + strconv.Itoa(len(f.Args)) + " actual: " + strconv.Itoa(len(contractInput.Args)))
	}

	inputBt, err := contractInput.Encode()
	if err != nil {
		return nil, errors.New("contractInput.Encode: " + err.Error())
	}

	if r.Debug {
		util.LogWithPurple("[ TryCall contract ]", r.Address.Hex())
		util.LogWithPurple("[ TryCall   method ]", contractInput.Selector)
		util.LogWithPurple("[ TryCall   origin ]", origin.ToHexString())
		util.LogWithPurple("[ TryCall     args ]", "0x"+hex.EncodeToString(inputBt))
		util.LogWithPurple("[ TryCall gaslimit ]", gas_limit)
		util.LogWithPurple("[ TryCall  storage ]", storage_deposit_limit)
		fmt.Println("")
	}

	result := util.ContractResult{}
	err = r.Client.CallRuntimeApi(
		"ReviveApi",
		"call",
		[]any{
			origin,
			r.Address,
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
	if result.Result.IsErr && result.Result.E.IsModule {
		merr := result.Result.E.AsModuleField0
		info, ierr := r.Client.GetErrorInfo(merr.Index, merr.Error)
		if ierr == nil {
			err = errors.New("TryCall: Module Error" + info.Name)
		} else {
			err = errors.New("TryCall: unknown Module Error")
		}
	} else {
		returnValue = &result.Result.V
		if returnValue.Flags == 1 {
			info := ""
			if len(returnValue.Data) > 2 && returnValue.Data[1] == 1 {
				info = r.GetErrorFromABI(returnValue.Data[2])
			}

			err = errors.New("TryCall: Contract Reverted" + info)
			returnValue = nil
		}
	}

	return &util.DryRunResult{
		GasConsumed:    result.GasConsumed,
		GasRequired:    result.GasRequired,
		StorageDeposit: result.StorageDeposit,
		Return:         returnValue,
	}, err
}

// Dryrun ink contract
func DryRunInk(
	ABI *Revive,
	origin types.AccountID,
	amount types.U128,
	gas_limit util.Option[types.Weight],
	storage_deposit_limit util.Option[types.U128],
	contractInput util.InkContractInput,
) (*util.DryRunResult, error) {
	return ABI.DryRun(origin, amount, gas_limit, storage_deposit_limit, contractInput)
}

// Call blockchain submit ink transaction
func CallInk(
	ABI *Revive,
	signer *Signer,
	amount types.U128,
	gas_limit types.Weight,
	storage_deposit_limit types.U128,
	contractInput util.InkContractInput,
) error {
	return ABI.Call(signer, amount, gas_limit, storage_deposit_limit, contractInput)
}

// call blockchain submit transaction
func (r *Revive) Call(
	signer *Signer,
	amount types.U128,
	gas_limit types.Weight,
	storage_deposit_limit types.U128,
	contractInput util.InkContractInput,
) error {
	_, err := r.DryRun(
		types.AccountID(signer.PublicKey),
		amount,
		util.NewSome(gas_limit),
		util.NewSome(storage_deposit_limit),
		contractInput,
	)

	if err != nil {
		return err
	}

	inputBt, err := contractInput.Encode()
	if err != nil {
		return errors.New("contractInput.Encode: " + err.Error())
	}

	if r.Debug {
		util.LogWithYellow("[ Call contract ]", r.Address.Hex())
		util.LogWithYellow("[ Call   method ]", contractInput.Selector)
		util.LogWithYellow("[ Call   origin ]", "0x"+hex.EncodeToString(signer.PublicKey))
		util.LogWithYellow("[ Call     args ]", "0x"+hex.EncodeToString(inputBt))
		fmt.Println("")
	}

	call := revive.MakeCallCall(
		r.Address,
		types.NewUCompact(amount.Int),
		gtypes.Weight{
			RefTime:   gas_limit.RefTime,
			ProofSize: gas_limit.ProofSize,
		},
		types.NewUCompact(storage_deposit_limit.Int),
		inputBt,
	)

	return r.Client.SignAndSubmit(signer, call, false)
}

// Get error info from ABI
func (r *Revive) GetErrorFromABI(index uint8) string {
	var errors = []util.SubVariant{}
	for _, t := range r.Abi.Types {
		if len(t.Type.Path) == 0 {
			continue
		}

		if t.Type.Path[len(t.Type.Path)-1] == "Error" && t.Type.Def.Variant != nil {
			errors = t.Type.Def.Variant.Variants
		}
	}

	if index >= uint8(len(errors)) {
		return " -> unknown"
	}

	return " -> " + errors[index].Name
}

// Get error info from ABI
func (r *Revive) GetArgsFromABI(selector [4]byte) (*util.Message, error) {
	for _, t := range r.Abi.Spec.Messages {
		if t.Selector == "0x"+hex.EncodeToString(selector[:]) {
			return &t, nil
		}
	}

	return nil, errors.New("Message 0x" + hex.EncodeToString(selector[:]) + " not found in ABI")
}
