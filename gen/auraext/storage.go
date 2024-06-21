package auraext

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/gen/types"
)

// Make a storage key for Authorities id={{false [351]}}
//
//	Serves as cache for the authorities.
//
//	The authorities in AuRa are overwritten in `on_initialize` when we switch to a new session,
//	but we require the old authorities to verify the seal when validating a PoV. This will
//	always be updated to the latest AuRa authorities in `on_finalize`.
func MakeAuthoritiesStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "AuraExt", "Authorities")
}

var AuthoritiesResultDefaultBytes, _ = hex.DecodeString("00")

func GetAuthorities(state state.State, bhash types.Hash) (ret [][32]byte, err error) {
	key, err := MakeAuthoritiesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(AuthoritiesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetAuthoritiesLatest(state state.State) (ret [][32]byte, err error) {
	key, err := MakeAuthoritiesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(AuthoritiesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for SlotInfo id={{false [354]}}
//
//	Current slot paired with a number of authored blocks.
//
//	Updated on each block initialization.
func MakeSlotInfoStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "AuraExt", "SlotInfo")
}
func GetSlotInfo(state state.State, bhash types.Hash) (ret types1.TupleOfUint64Uint32, isSome bool, err error) {
	key, err := MakeSlotInfoStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetSlotInfoLatest(state state.State) (ret types1.TupleOfUint64Uint32, isSome bool, err error) {
	key, err := MakeSlotInfoStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
