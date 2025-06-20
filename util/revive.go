package util

import (
	"bytes"
	"errors"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"

	gtypes "github.com/wetee-dao/ink.go/pallet/types"
)

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

type DryRunResult struct {
	GasConsumed    gtypes.Weight
	GasRequired    gtypes.Weight
	StorageDeposit gtypes.StorageDeposit
	Return         *gtypes.ExecReturnValue
}

type NullTuple struct{}
