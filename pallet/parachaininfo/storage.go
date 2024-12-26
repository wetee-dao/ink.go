package parachaininfo

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for ParachainId id={{false [122]}}
func MakeParachainIdStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainInfo", "ParachainId")
}

var ParachainIdResultDefaultBytes, _ = hex.DecodeString("64000000")

func GetParachainId(state state.State, bhash types.Hash) (ret uint32, err error) {
	key, err := MakeParachainIdStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ParachainIdResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetParachainIdLatest(state state.State) (ret uint32, err error) {
	key, err := MakeParachainIdStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ParachainIdResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
