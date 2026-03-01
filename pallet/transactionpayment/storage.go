package transactionpayment

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/ink.go/pallet/types"
)

// Make a storage key for NextFeeMultiplier id={{false [101]}}
func MakeNextFeeMultiplierStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "TransactionPayment", "NextFeeMultiplier")
}

var NextFeeMultiplierResultDefaultBytes, _ = hex.DecodeString("000064a7b3b6e00d0000000000000000")

func GetNextFeeMultiplier(state state.State, bhash types.Hash) (ret types.U128, err error) {
	key, err := MakeNextFeeMultiplierStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(NextFeeMultiplierResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetNextFeeMultiplierLatest(state state.State) (ret types.U128, err error) {
	key, err := MakeNextFeeMultiplierStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(NextFeeMultiplierResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for StorageVersion id={{false [102]}}
func MakeStorageVersionStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "TransactionPayment", "StorageVersion")
}

var StorageVersionResultDefaultBytes, _ = hex.DecodeString("00")

func GetStorageVersion(state state.State, bhash types.Hash) (ret types1.Releases, err error) {
	key, err := MakeStorageVersionStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(StorageVersionResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetStorageVersionLatest(state state.State) (ret types1.Releases, err error) {
	key, err := MakeStorageVersionStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(StorageVersionResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for TxPaymentCredit id={{false [103]}}
//
//	The `OnChargeTransaction` stores the withdrawn tx fee here.
//
//	Use `withdraw_txfee` and `remaining_txfee` to access from outside the crate.
func MakeTxPaymentCreditStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "TransactionPayment", "TxPaymentCredit")
}
func GetTxPaymentCredit(state state.State, bhash types.Hash) (ret types.U128, isSome bool, err error) {
	key, err := MakeTxPaymentCreditStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetTxPaymentCreditLatest(state state.State) (ret types.U128, isSome bool, err error) {
	key, err := MakeTxPaymentCreditStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
