package weteebridge

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for NextId id={{false [6]}}
func MakeNextIdStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "WeTEEBridge", "NextId")
}

var NextIdResultDefaultBytes, _ = hex.DecodeString("00000000000000000000000000000000")

func GetNextId(state state.State, bhash types.Hash) (ret types.U128, err error) {
	key, err := MakeNextIdStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(NextIdResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetNextIdLatest(state state.State) (ret types.U128, err error) {
	key, err := MakeNextIdStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(NextIdResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for TEECalls
func MakeTEECallsStorageKey(u1280 types.U128) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(u1280)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEBridge", "TEECalls", byteArgs...)
}
func GetTEECalls(state state.State, bhash types.Hash, u1280 types.U128) (ret types1.TEECall, isSome bool, err error) {
	key, err := MakeTEECallsStorageKey(u1280)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetTEECallsLatest(state state.State, u1280 types.U128) (ret types1.TEECall, isSome bool, err error) {
	key, err := MakeTEECallsStorageKey(u1280)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
