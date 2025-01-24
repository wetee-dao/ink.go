package session

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for Validators id={{false [41]}}
//
//	The current set of validators.
func MakeValidatorsStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Session", "Validators")
}

var ValidatorsResultDefaultBytes, _ = hex.DecodeString("00")

func GetValidators(state state.State, bhash types.Hash) (ret [][32]byte, err error) {
	key, err := MakeValidatorsStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ValidatorsResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetValidatorsLatest(state state.State) (ret [][32]byte, err error) {
	key, err := MakeValidatorsStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ValidatorsResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for CurrentIndex id={{false [4]}}
//
//	Current index of the session.
func MakeCurrentIndexStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Session", "CurrentIndex")
}

var CurrentIndexResultDefaultBytes, _ = hex.DecodeString("00000000")

func GetCurrentIndex(state state.State, bhash types.Hash) (ret uint32, err error) {
	key, err := MakeCurrentIndexStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CurrentIndexResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetCurrentIndexLatest(state state.State) (ret uint32, err error) {
	key, err := MakeCurrentIndexStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CurrentIndexResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for QueuedChanged id={{false [8]}}
//
//	True if the underlying economic identities or weighting behind the validators
//	has changed in the queued validator set.
func MakeQueuedChangedStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Session", "QueuedChanged")
}

var QueuedChangedResultDefaultBytes, _ = hex.DecodeString("00")

func GetQueuedChanged(state state.State, bhash types.Hash) (ret bool, err error) {
	key, err := MakeQueuedChangedStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(QueuedChangedResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetQueuedChangedLatest(state state.State) (ret bool, err error) {
	key, err := MakeQueuedChangedStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(QueuedChangedResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for QueuedKeys id={{false [382]}}
//
//	The queued keys for the next session. When the next session begins, these keys
//	will be used to determine the validator's session keys.
func MakeQueuedKeysStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Session", "QueuedKeys")
}

var QueuedKeysResultDefaultBytes, _ = hex.DecodeString("00")

func GetQueuedKeys(state state.State, bhash types.Hash) (ret []types1.TupleOfByteArray32ByteArray32, err error) {
	key, err := MakeQueuedKeysStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(QueuedKeysResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetQueuedKeysLatest(state state.State) (ret []types1.TupleOfByteArray32ByteArray32, err error) {
	key, err := MakeQueuedKeysStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(QueuedKeysResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for DisabledValidators id={{false [384]}}
//
//	Indices of disabled validators.
//
//	The vec is always kept sorted so that we can find whether a given validator is
//	disabled using binary search. It gets cleared when `on_session_ending` returns
//	a new set of identities.
func MakeDisabledValidatorsStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Session", "DisabledValidators")
}

var DisabledValidatorsResultDefaultBytes, _ = hex.DecodeString("00")

func GetDisabledValidators(state state.State, bhash types.Hash) (ret []uint32, err error) {
	key, err := MakeDisabledValidatorsStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(DisabledValidatorsResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetDisabledValidatorsLatest(state state.State) (ret []uint32, err error) {
	key, err := MakeDisabledValidatorsStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(DisabledValidatorsResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for NextKeys
//
//	The next session keys for a validator.
func MakeNextKeysStorageKey(byteArray320 [32]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteArray320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Session", "NextKeys", byteArgs...)
}
func GetNextKeys(state state.State, bhash types.Hash, byteArray320 [32]byte) (ret [32]byte, isSome bool, err error) {
	key, err := MakeNextKeysStorageKey(byteArray320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetNextKeysLatest(state state.State, byteArray320 [32]byte) (ret [32]byte, isSome bool, err error) {
	key, err := MakeNextKeysStorageKey(byteArray320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for KeyOwner
//
//	The owner of a key. The key is the `KeyTypeId` + the encoded key.
func MakeKeyOwnerStorageKey(tupleOfByteArray4ByteSlice0 [4]byte, tupleOfByteArray4ByteSlice1 []byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfByteArray4ByteSlice0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfByteArray4ByteSlice1)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Session", "KeyOwner", byteArgs...)
}
func GetKeyOwner(state state.State, bhash types.Hash, tupleOfByteArray4ByteSlice0 [4]byte, tupleOfByteArray4ByteSlice1 []byte) (ret [32]byte, isSome bool, err error) {
	key, err := MakeKeyOwnerStorageKey(tupleOfByteArray4ByteSlice0, tupleOfByteArray4ByteSlice1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetKeyOwnerLatest(state state.State, tupleOfByteArray4ByteSlice0 [4]byte, tupleOfByteArray4ByteSlice1 []byte) (ret [32]byte, isSome bool, err error) {
	key, err := MakeKeyOwnerStorageKey(tupleOfByteArray4ByteSlice0, tupleOfByteArray4ByteSlice1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
