package ink

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

var ErrContractReverted = errors.New("contract reverted: the specific error information is returned")

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
		util.LogWithPurple("[        contract ]", addres.Hex())
		util.LogWithPurple("[          origin ]", origin.ToHexString())
		util.LogWithPurple("[            args ]", "0x"+hex.EncodeToString(inputBt))
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
				err = errors.New("DryRun: Module Error: " + info.Name)
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
		util.LogWithYellow("[         RefTime ]", gas_limit.RefTime.Int64())
		util.LogWithYellow("[       ProofSize ]", gas_limit.ProofSize.Int64())
		util.LogWithYellow("[    DepositLimit ]", storage_deposit_limit.Int.String())
	}

	storageDepositLimit := big.NewInt(0)
	if storage_deposit_limit.Int != nil {
		storageDepositLimit = storage_deposit_limit.Int
	}

	runtimeCall := revive.MakeCallCall(
		addres,
		types.NewUCompact(amount.Int),
		gtypes.Weight{
			RefTime:   gas_limit.RefTime,
			ProofSize: gas_limit.ProofSize,
		},
		types.NewUCompact(storageDepositLimit),
		inputBt,
	)

	call, err := (runtimeCall).AsCall()
	if err != nil {
		return errors.New("(runtimeCall).AsCall() error: " + err.Error())
	}

	return client.SignAndSubmit(signer, call, true)
}

func CallOfTransaction(
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
		util.LogWithYellow("[        method ]", contractInput.Selector)
		util.LogWithYellow("[        origin ]", "0x"+hex.EncodeToString(signer.Public()))
		util.LogWithYellow("[          args ]", "0x"+hex.EncodeToString(inputBt))
		util.LogWithYellow("[       RefTime ]", gas_limit.RefTime.Int64())
		util.LogWithYellow("[     ProofSize ]", gas_limit.ProofSize.Int64())
		util.LogWithYellow("[  DepositLimit ]", storage_deposit_limit.Int.String())
		fmt.Println()
	}

	storageDepositLimit := big.NewInt(0)
	if storage_deposit_limit.Int != nil {
		storageDepositLimit = storage_deposit_limit.Int
	}

	runtimeCall := revive.MakeCallCall(
		addres,
		types.NewUCompact(amount.Int),
		gtypes.Weight{
			RefTime:   gas_limit.RefTime,
			ProofSize: gas_limit.ProofSize,
		},
		types.NewUCompact(storageDepositLimit),
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

func DefaultParamWithOrigin(origin types.AccountID) DryRunCallParams {
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

// Call param of Call
type DeployParams struct {
	Client *ChainClient
	Signer SignerType
	Code   util.InkCode
	Salt   util.Option[[32]byte]
}

func (c *ChainClient) UploadInkCode(code []byte, signer SignerType) (*types.H256, error) {
	resultWrap := util.Result[util.UploadResult, gtypes.DispatchError]{}
	origin := signer.AccountID()
	err := c.CallRuntimeApi(
		"ReviveApi",
		"upload_code",
		[]any{
			origin, code, util.NewNone[types.U128](),
		},
		&resultWrap,
	)
	if err != nil {
		return nil, errors.New("CallRuntimeApi: " + err.Error())
	}

	result, err := resultWrap.UnWrap()
	if err != nil {
		return nil, errors.New("UnWrap: " + err.Error())
	}

	runtimeCall := revive.MakeUploadCodeCall(code, types.NewUCompact(result.Deposit.Int))
	call, err := (runtimeCall).AsCall()
	if err != nil {
		return nil, errors.New("(runtimeCall).AsCall() error: " + err.Error())
	}

	err = c.SignAndSubmit(signer, call, true)
	if err != nil {
		return nil, errors.New("SignAndSubmit error: " + err.Error())
	}

	return &result.CodeHash, nil
}

func (c *ChainClient) DeployContract(code util.InkCode, signer SignerType, payAmount types.U128, args util.InkContractInput, salt util.Option[[32]byte]) (*types.H160, error) {
	resultWrap := util.ContractInitResult{}
	origin := signer.AccountID()

	argData, err := args.Encode()
	if err != nil {
		return nil, errors.New("args.Encode: " + err.Error())
	}

	err = c.CallRuntimeApi(
		"ReviveApi",
		"instantiate",
		[]any{
			origin,
			payAmount,
			util.NewNone[types.Weight](),
			util.NewNone[types.U128](),
			code,
			argData,
			salt,
		},
		&resultWrap,
	)
	if err != nil {
		return nil, errors.New("CallRuntimeApi: " + err.Error())
	}

	/// check runtime_api error
	if resultWrap.Result.IsErr {
		if resultWrap.Result.E.IsModule {
			merr := resultWrap.Result.E.AsModuleField0
			info, ierr := c.GetErrorInfo(merr.Index, merr.Error)
			if ierr == nil {
				err = errors.New("DryRun: Module Error: " + info.Name)
			} else {
				err = errors.New("DryRun: unknown Module Error")
			}
			return nil, err
		}
		bt, _ := json.Marshal(resultWrap.Result.E)
		return nil, errors.New(string(bt))
	}

	result, err := resultWrap.Result.UnWrap()
	if err != nil {
		return nil, errors.New("UnWrap: " + err.Error())
	}

	// 判断是否执行错误
	if result.Result.Flags == 1 {
		return nil, ErrContractReverted
	}

	// init salt
	gsalt := gtypes.OptionTByteArray321{
		IsNone: true,
	}
	if !salt.IsNone() {
		gsalt = gtypes.OptionTByteArray321{
			IsNone:       false,
			IsSome:       true,
			AsSomeField0: salt.V,
		}
	}

	if c.Debug {
		util.LogWithYellow("[ Deploy origin ]", origin.ToHexString())
		util.LogWithYellow("[         value ]", payAmount)
		util.LogWithYellow("[          args ]", "0x"+hex.EncodeToString(argData))
		util.LogWithYellow("[       RefTime ]", resultWrap.GasRequired.RefTime.Int64())
		util.LogWithYellow("[     ProofSize ]", resultWrap.GasRequired.ProofSize.Int64())
		util.LogWithYellow("[  DepositLimit ]", resultWrap.StorageDeposit.AsChargeField0.Int.String())
		fmt.Println("")
	}

	var runtimeCall gtypes.RuntimeCall
	if code.Upload != nil {
		runtimeCall = revive.MakeInstantiateWithCodeCall(
			types.NewUCompact(payAmount.Int),
			resultWrap.GasRequired,
			types.NewUCompact(resultWrap.StorageDeposit.AsChargeField0.Int),
			*code.Upload,
			argData,
			gsalt,
		)
	} else if code.Existing != nil {
		runtimeCall = revive.MakeInstantiateCall(
			types.NewUCompact(payAmount.Int),
			resultWrap.GasRequired,
			types.NewUCompact(resultWrap.StorageDeposit.AsChargeField0.Int),
			*code.Existing,
			argData,
			gsalt,
		)
	}
	call, err := (runtimeCall).AsCall()
	if err != nil {
		return nil, errors.New("(runtimeCall).AsCall() error: " + err.Error())
	}

	// submit call
	err = c.SignAndSubmit(signer, call, true)
	if err != nil {
		return nil, errors.New("SignAndSubmit error: " + err.Error())
	}

	return &result.Addr, nil
}
