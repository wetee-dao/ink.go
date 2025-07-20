package util

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"

	gtypes "github.com/wetee-dao/ink.go/pallet/types"
)

type AccountId = types.AccountID

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

type ContractResult struct {
	GasConsumed    gtypes.Weight
	GasRequired    gtypes.Weight
	StorageDeposit gtypes.StorageDeposit
	Result         Result[gtypes.ExecReturnValue, gtypes.DispatchError]
}

type UploadResult struct {
	CodeHash types.H256
	Deposit  types.U128
}

type ContractInitResult struct {
	GasConsumed    gtypes.Weight
	GasRequired    gtypes.Weight
	StorageDeposit gtypes.StorageDeposit
	Result         Result[InitReturnValue, gtypes.DispatchError]
}

type InitReturnValue struct {
	Result gtypes.ExecReturnValue
	Addr   types.H160
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
