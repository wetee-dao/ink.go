package tokens

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types1 "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for TotalIssuance
//
//	The total issuance of a token type.
func MakeTotalIssuanceStorageKey(currencyId0 types.CurrencyId) (types1.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(currencyId0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types1.CreateStorageKey(&types.Meta, "Tokens", "TotalIssuance", byteArgs...)
}

var TotalIssuanceResultDefaultBytes, _ = hex.DecodeString("00000000000000000000000000000000")

func GetTotalIssuance(state state.State, bhash types1.Hash, currencyId0 types.CurrencyId) (ret types1.U128, err error) {
	key, err := MakeTotalIssuanceStorageKey(currencyId0)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(TotalIssuanceResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetTotalIssuanceLatest(state state.State, currencyId0 types.CurrencyId) (ret types1.U128, err error) {
	key, err := MakeTotalIssuanceStorageKey(currencyId0)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(TotalIssuanceResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for Locks
//
//	Any liquidity locks of a token type under an account.
//	NOTE: Should only be accessed when setting, changing and freeing a lock.
func MakeLocksStorageKey(tupleOfByteArray32CurrencyId0 [32]byte, tupleOfByteArray32CurrencyId1 types.CurrencyId) (types1.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfByteArray32CurrencyId0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfByteArray32CurrencyId1)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types1.CreateStorageKey(&types.Meta, "Tokens", "Locks", byteArgs...)
}

var LocksResultDefaultBytes, _ = hex.DecodeString("00")

func GetLocks(state state.State, bhash types1.Hash, tupleOfByteArray32CurrencyId0 [32]byte, tupleOfByteArray32CurrencyId1 types.CurrencyId) (ret []types.BalanceLock1, err error) {
	key, err := MakeLocksStorageKey(tupleOfByteArray32CurrencyId0, tupleOfByteArray32CurrencyId1)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(LocksResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetLocksLatest(state state.State, tupleOfByteArray32CurrencyId0 [32]byte, tupleOfByteArray32CurrencyId1 types.CurrencyId) (ret []types.BalanceLock1, err error) {
	key, err := MakeLocksStorageKey(tupleOfByteArray32CurrencyId0, tupleOfByteArray32CurrencyId1)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(LocksResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for Accounts
//
//	The balance of a token type under an account.
//
//	NOTE: If the total is ever zero, decrease account ref account.
//
//	NOTE: This is only used in the case that this module is used to store
//	balances.
func MakeAccountsStorageKey(tupleOfByteArray32CurrencyId0 [32]byte, tupleOfByteArray32CurrencyId1 types.CurrencyId) (types1.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfByteArray32CurrencyId0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfByteArray32CurrencyId1)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types1.CreateStorageKey(&types.Meta, "Tokens", "Accounts", byteArgs...)
}

var AccountsResultDefaultBytes, _ = hex.DecodeString("000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

func GetAccounts(state state.State, bhash types1.Hash, tupleOfByteArray32CurrencyId0 [32]byte, tupleOfByteArray32CurrencyId1 types.CurrencyId) (ret types.AccountData1, err error) {
	key, err := MakeAccountsStorageKey(tupleOfByteArray32CurrencyId0, tupleOfByteArray32CurrencyId1)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(AccountsResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetAccountsLatest(state state.State, tupleOfByteArray32CurrencyId0 [32]byte, tupleOfByteArray32CurrencyId1 types.CurrencyId) (ret types.AccountData1, err error) {
	key, err := MakeAccountsStorageKey(tupleOfByteArray32CurrencyId0, tupleOfByteArray32CurrencyId1)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(AccountsResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for Reserves
//
//	Named reserves on some account balances.
func MakeReservesStorageKey(tupleOfByteArray32CurrencyId0 [32]byte, tupleOfByteArray32CurrencyId1 types.CurrencyId) (types1.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfByteArray32CurrencyId0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfByteArray32CurrencyId1)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types1.CreateStorageKey(&types.Meta, "Tokens", "Reserves", byteArgs...)
}

var ReservesResultDefaultBytes, _ = hex.DecodeString("00")

func GetReserves(state state.State, bhash types1.Hash, tupleOfByteArray32CurrencyId0 [32]byte, tupleOfByteArray32CurrencyId1 types.CurrencyId) (ret []types.ReserveDataReserveIdentifierByteArray8, err error) {
	key, err := MakeReservesStorageKey(tupleOfByteArray32CurrencyId0, tupleOfByteArray32CurrencyId1)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ReservesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetReservesLatest(state state.State, tupleOfByteArray32CurrencyId0 [32]byte, tupleOfByteArray32CurrencyId1 types.CurrencyId) (ret []types.ReserveDataReserveIdentifierByteArray8, err error) {
	key, err := MakeReservesStorageKey(tupleOfByteArray32CurrencyId0, tupleOfByteArray32CurrencyId1)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ReservesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
