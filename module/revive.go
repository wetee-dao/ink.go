package module

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"

	chain "github.com/wetee-dao/go-sdk"
	"github.com/wetee-dao/go-sdk/pallet/revive"
	gtypes "github.com/wetee-dao/go-sdk/pallet/types"
	"github.com/wetee-dao/go-sdk/util"
)

// Revive module
type Revive struct {
	Client *chain.ChainClient
}

// Get balance of h160
func (r *Revive) BalanceOfH160(address string) (types.U128, error) {
	balance := types.NewU128(*big.NewInt(0))
	bt, err := util.H160HexToBt(address)
	if err != nil {
		return types.U128{}, err
	}
	err = r.Client.CallRuntimeApi("ReviveApi", "balance", []any{bt}, &balance)
	if err != nil {
		return types.U128{}, err
	}

	return balance, nil
}

// Get block gas limit
func (r *Revive) InkBlockGasLimit(address [32]byte) error {
	balance := types.NewU128(*big.NewInt(0))
	err := r.Client.CallRuntimeApi("ReviveApi", "block_gas_limit", []any{}, &balance)
	return err
}

// Query ink contract data
func (r *Revive) QueryInk(
	origin types.AccountID,
	amount types.U128,
	gas_limit util.Option[types.Weight],
	storage_deposit_limit util.Option[types.U128],
	contract string,
	contractInput util.InkContractInput,
	returnValue any,
) error {
	result, err := r.TryCallInk(origin, amount, gas_limit, storage_deposit_limit, contract, contractInput)
	if err != nil {
		return err
	}

	// util.PrintJson(result)
	return scale.NewDecoder(bytes.NewReader(result.Data)).Decode(returnValue)
}

// try call ink contract
func (r *Revive) TryCallInk(
	origin types.AccountID,
	amount types.U128,
	gas_limit util.Option[types.Weight],
	storage_deposit_limit util.Option[types.U128],
	contract string,
	contractInput util.InkContractInput,
) (*gtypes.ExecReturnValue, error) {
	contractAddress, err := util.H160HexToBt(contract)
	if err != nil {
		return nil, errors.New("H160HexToBt: " + err.Error())
	}

	inputBt, err := contractInput.Encode()
	if err != nil {
		return nil, errors.New("contractInput.Encode: " + err.Error())
	}

	result := util.ContractResult{}
	err = r.Client.CallRuntimeApi(
		"ReviveApi",
		"call",
		[]any{
			origin,
			contractAddress,
			amount,
			gas_limit,
			storage_deposit_limit,
			inputBt,
		},
		&result,
	)

	util.LogWithRed("TryCall contract", contract)
	util.LogWithRed("TryCall   origin", origin.ToHexString())
	util.LogWithRed("TryCall     args", "0x"+hex.EncodeToString(inputBt))
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
			err = errors.New("TryCall: REVERT")
			fmt.Println(returnValue.Data)
			returnValue = nil
		}
	}

	return returnValue, err
}

// call blockchain submit transaction
func (r *Revive) CallInk(
	signer *chain.Signer,
	origin types.AccountID,
	amount types.U128,
	gas_limit types.Weight,
	storage_deposit_limit types.U128,
	contract string,
	contractInput util.InkContractInput,
) error {
	contractAddress, err := util.H160HexToBt(contract)
	if err != nil {
		return errors.New("H160HexToBt: " + err.Error())
	}

	inputBt, err := contractInput.Encode()
	if err != nil {
		return errors.New("contractInput.Encode: " + err.Error())
	}

	// r.Client.Api.
	call := revive.MakeCallCall(
		contractAddress,
		types.NewUCompact(amount.Int),
		gtypes.Weight{
			RefTime:   gas_limit.RefTime,
			ProofSize: gas_limit.ProofSize,
		},
		types.NewUCompact(storage_deposit_limit.Int),
		inputBt,
	)

	return r.Client.SignAndSubmit(signer, call, true)
}
