package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"

	gtypes "github.com/wetee-dao/ink.go/pallet/types"
)

type AccountId = types.AccountID
type Bytes = []byte
type BlockNumber = uint32

// type AccountId [32]byte
func NewAccountID(pubkey []byte) types.AccountID {
	var bt [32]byte
	copy(bt[:], pubkey)

	return types.AccountID(bt)
}

// Ink Contract input
type InkContractInput struct {
	/// The selector for the smart contract execution.
	Selector string
	/// The arguments of the smart contract execution.
	Args []any
}

func (e *InkContractInput) Encode() ([]byte, error) {
	var buffer bytes.Buffer
	encoder := scale.NewEncoder(&buffer)
	selector := FuncToSelector(e.Selector)
	err := encoder.Encode(selector)
	if err != nil {
		return nil, err
	}

	for _, arg := range e.Args {
		err = encoder.Encode(arg)
		if err != nil {
			return nil, errors.New("Arg.Encode: " + err.Error())
		}
	}

	return buffer.Bytes(), nil
}

// ContractResult 与 pallet-revive ContractResult<R, Balance> 对齐（call 返回值）
// 字段顺序与 ContractInitResult 一致，使用 gtypes.Weight
type ContractResult struct {
	WeightConsumed    gtypes.Weight
	WeightRequired    gtypes.Weight
	StorageDeposit    StorageDeposit
	MaxStorageDeposit StorageDeposit
	GasConsumed       types.U128
	Result            Result[ExecReturnValue, gtypes.DispatchError]
}

type UploadResult struct {
	CodeHash types.H256
	Deposit  types.U128
}

// ContractInitResult 与 pallet-revive ReviveApi::instantiate 返回的 ContractResult<InstantiateReturnValue, Balance> 对齐
// 字段顺序: weight_consumed, weight_required, storage_deposit, max_storage_deposit, gas_consumed, result
// 使用 gtypes.Weight（与链元数据一致，可能为 Compact 编码）以正确对齐
type ContractInitResult struct {
	WeightConsumed    gtypes.Weight
	WeightRequired    gtypes.Weight
	StorageDeposit    StorageDeposit
	MaxStorageDeposit StorageDeposit
	GasConsumed       types.U128
	Result            Result[InitReturnValue, gtypes.DispatchError]
}

// InitReturnValue 与 pallet-revive InstantiateReturnValue 对齐（result: ExecReturnValue, addr: H160）
type InitReturnValue struct {
	Result    ExecReturnValue
	AccountID types.H160
}

type InkCode struct {
	Upload   *[]byte
	Existing *types.H256
}

func (ty InkCode) Encode(encoder scale.Encoder) (err error) {
	if ty.Upload != nil {
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.Upload)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.Existing != nil {
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.Existing)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("unrecognized enum")
}

func (ty *InkCode) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0:
		ty.Upload = new([]byte)
		err = decoder.Decode(ty.Upload)
		if err != nil {
			return err
		}
		return

	case 1:
		ty.Existing = new(types.H256)
		err = decoder.Decode(ty.Upload)
		if err != nil {
			return err
		}
		return
	default:
		return fmt.Errorf("unrecognized enum")
	}
}

type NullTuple struct{}

// Generated PalletRevivePrimitivesStorageDeposit with id=335
type StorageDeposit struct {
	IsRefund       bool
	AsRefundField0 types.U128
	IsCharge       bool
	AsChargeField0 types.U128
}

func (ty StorageDeposit) Encode(encoder scale.Encoder) (err error) {
	if ty.IsRefund {
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
		err = encoder.Encode(ty.AsRefundField0)
		if err != nil {
			return err
		}
		return nil
	}
	if ty.IsCharge {
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}
		err = encoder.Encode(ty.AsChargeField0)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("Unrecognized variant")
}
func (ty *StorageDeposit) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0:
		ty.IsRefund = true
		err = decoder.Decode(&ty.AsRefundField0)
		if err != nil {
			return err
		}
		return
	case 1:
		ty.IsCharge = true
		err = decoder.Decode(&ty.AsChargeField0)
		if err != nil {
			return err
		}
		return
	default:
		return fmt.Errorf("Unrecognized variant")
	}
}
func (ty *StorageDeposit) Variant() (uint8, error) {
	if ty.IsRefund {
		return 0, nil
	}
	if ty.IsCharge {
		return 1, nil
	}
	return 0, fmt.Errorf("No variant detected")
}
func (ty StorageDeposit) MarshalJSON() ([]byte, error) {
	if ty.IsRefund {
		m := map[string]interface{}{"StorageDeposit::Refund": ty.AsRefundField0}
		return json.Marshal(m)
	}
	if ty.IsCharge {
		m := map[string]interface{}{"StorageDeposit::Charge": ty.AsChargeField0}
		return json.Marshal(m)
	}
	return nil, fmt.Errorf("No variant detected")
}

type ExecReturnValue struct {
	// Field 0 with TypeId=334
	Flags uint32
	// Field 1 with TypeId=14
	Data []byte
}
