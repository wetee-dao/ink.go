package unknowntokens

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types1 "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for ConcreteFungibleBalances
//
//	Concrete fungible balances under a given location and a concrete
//	fungible id.
//
//	double_map: who, asset_id => u128
func MakeConcreteFungibleBalancesStorageKey(tupleOfLocationLocation0 types.Location, tupleOfLocationLocation1 types.Location) (types1.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfLocationLocation0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfLocationLocation1)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types1.CreateStorageKey(&types.Meta, "UnknownTokens", "ConcreteFungibleBalances", byteArgs...)
}

var ConcreteFungibleBalancesResultDefaultBytes, _ = hex.DecodeString("00000000000000000000000000000000")

func GetConcreteFungibleBalances(state state.State, bhash types1.Hash, tupleOfLocationLocation0 types.Location, tupleOfLocationLocation1 types.Location) (ret types1.U128, err error) {
	key, err := MakeConcreteFungibleBalancesStorageKey(tupleOfLocationLocation0, tupleOfLocationLocation1)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ConcreteFungibleBalancesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetConcreteFungibleBalancesLatest(state state.State, tupleOfLocationLocation0 types.Location, tupleOfLocationLocation1 types.Location) (ret types1.U128, err error) {
	key, err := MakeConcreteFungibleBalancesStorageKey(tupleOfLocationLocation0, tupleOfLocationLocation1)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ConcreteFungibleBalancesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for AbstractFungibleBalances
//
//	Abstract fungible balances under a given location and a abstract
//	fungible id.
//
//	double_map: who, asset_id => u128
func MakeAbstractFungibleBalancesStorageKey(tupleOfLocationByteSlice0 types.Location, tupleOfLocationByteSlice1 []byte) (types1.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfLocationByteSlice0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfLocationByteSlice1)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types1.CreateStorageKey(&types.Meta, "UnknownTokens", "AbstractFungibleBalances", byteArgs...)
}

var AbstractFungibleBalancesResultDefaultBytes, _ = hex.DecodeString("00000000000000000000000000000000")

func GetAbstractFungibleBalances(state state.State, bhash types1.Hash, tupleOfLocationByteSlice0 types.Location, tupleOfLocationByteSlice1 []byte) (ret types1.U128, err error) {
	key, err := MakeAbstractFungibleBalancesStorageKey(tupleOfLocationByteSlice0, tupleOfLocationByteSlice1)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(AbstractFungibleBalancesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetAbstractFungibleBalancesLatest(state state.State, tupleOfLocationByteSlice0 types.Location, tupleOfLocationByteSlice1 []byte) (ret types1.U128, err error) {
	key, err := MakeAbstractFungibleBalancesStorageKey(tupleOfLocationByteSlice0, tupleOfLocationByteSlice1)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(AbstractFungibleBalancesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
