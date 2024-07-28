package weteebridge

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for NextId id={{false [12]}}
func MakeNextIdStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "WeTEEBridge", "NextId")
}

var NextIdResultDefaultBytes, _ = hex.DecodeString("0000000000000000")

func GetNextId(state state.State, bhash types.Hash) (ret uint64, err error) {
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
func GetNextIdLatest(state state.State) (ret uint64, err error) {
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
