package asset

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for ChainID id={{false [4]}}
func MakeChainIDStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Asset", "ChainID")
}

var ChainIDResultDefaultBytes, _ = hex.DecodeString("00000000")

func GetChainID(state state.State, bhash types.Hash) (ret uint32, err error) {
	key, err := MakeChainIDStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ChainIDResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetChainIDLatest(state state.State) (ret uint32, err error) {
	key, err := MakeChainIDStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ChainIDResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for AssetsInfo
func MakeAssetsInfoStorageKey(currencyId0 types1.CurrencyId) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(currencyId0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Asset", "AssetsInfo", byteArgs...)
}
func GetAssetsInfo(state state.State, bhash types.Hash, currencyId0 types1.CurrencyId) (ret types1.AssetInfo, isSome bool, err error) {
	key, err := MakeAssetsInfoStorageKey(currencyId0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetAssetsInfoLatest(state state.State, currencyId0 types1.CurrencyId) (ret types1.AssetInfo, isSome bool, err error) {
	key, err := MakeAssetsInfoStorageKey(currencyId0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for Symbols
func MakeSymbolsStorageKey(byteSlice0 []byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteSlice0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Asset", "Symbols", byteArgs...)
}
func GetSymbols(state state.State, bhash types.Hash, byteSlice0 []byte) (ret types1.CurrencyId, isSome bool, err error) {
	key, err := MakeSymbolsStorageKey(byteSlice0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetSymbolsLatest(state state.State, byteSlice0 []byte) (ret types1.CurrencyId, isSome bool, err error) {
	key, err := MakeSymbolsStorageKey(byteSlice0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for ParaMaps
func MakeParaMapsStorageKey(tupleOfUint32ByteSlice0 uint32, tupleOfUint32ByteSlice1 []byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfUint32ByteSlice0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfUint32ByteSlice1)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Asset", "ParaMaps", byteArgs...)
}
func GetParaMaps(state state.State, bhash types.Hash, tupleOfUint32ByteSlice0 uint32, tupleOfUint32ByteSlice1 []byte) (ret types1.CurrencyId, isSome bool, err error) {
	key, err := MakeParaMapsStorageKey(tupleOfUint32ByteSlice0, tupleOfUint32ByteSlice1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetParaMapsLatest(state state.State, tupleOfUint32ByteSlice0 uint32, tupleOfUint32ByteSlice1 []byte) (ret types1.CurrencyId, isSome bool, err error) {
	key, err := MakeParaMapsStorageKey(tupleOfUint32ByteSlice0, tupleOfUint32ByteSlice1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
