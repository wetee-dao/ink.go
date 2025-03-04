package polkadotxcm

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for QueryCounter id={{false [12]}}
//
//	The latest available query index.
func MakeQueryCounterStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "QueryCounter")
}

var QueryCounterResultDefaultBytes, _ = hex.DecodeString("0000000000000000")

func GetQueryCounter(state state.State, bhash types.Hash) (ret uint64, err error) {
	key, err := MakeQueryCounterStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(QueryCounterResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetQueryCounterLatest(state state.State) (ret uint64, err error) {
	key, err := MakeQueryCounterStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(QueryCounterResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for Queries
//
//	The ongoing queries.
func MakeQueriesStorageKey(uint640 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "Queries", byteArgs...)
}
func GetQueries(state state.State, bhash types.Hash, uint640 uint64) (ret types1.QueryStatus, isSome bool, err error) {
	key, err := MakeQueriesStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetQueriesLatest(state state.State, uint640 uint64) (ret types1.QueryStatus, isSome bool, err error) {
	key, err := MakeQueriesStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for AssetTraps
//
//	The existing asset traps.
//
//	Key is the blake2 256 hash of (origin, versioned `Assets`) pair. Value is the number of
//	times this pair has been trapped (usually just 1 if it exists at all).
func MakeAssetTrapsStorageKey(byteArray320 [32]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteArray320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "AssetTraps", byteArgs...)
}

var AssetTrapsResultDefaultBytes, _ = hex.DecodeString("00000000")

func GetAssetTraps(state state.State, bhash types.Hash, byteArray320 [32]byte) (ret uint32, err error) {
	key, err := MakeAssetTrapsStorageKey(byteArray320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(AssetTrapsResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetAssetTrapsLatest(state state.State, byteArray320 [32]byte) (ret uint32, err error) {
	key, err := MakeAssetTrapsStorageKey(byteArray320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(AssetTrapsResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for SafeXcmVersion id={{false [4]}}
//
//	Default version to encode XCM when latest version of destination is unknown. If `None`,
//	then the destinations whose XCM version is unknown are considered unreachable.
func MakeSafeXcmVersionStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "SafeXcmVersion")
}
func GetSafeXcmVersion(state state.State, bhash types.Hash) (ret uint32, isSome bool, err error) {
	key, err := MakeSafeXcmVersionStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetSafeXcmVersionLatest(state state.State) (ret uint32, isSome bool, err error) {
	key, err := MakeSafeXcmVersionStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for SupportedVersion
//
//	The Latest versions that we know various locations support.
func MakeSupportedVersionStorageKey(tupleOfUint32VersionedLocation0 uint32, tupleOfUint32VersionedLocation1 types1.VersionedLocation) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfUint32VersionedLocation0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfUint32VersionedLocation1)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "SupportedVersion", byteArgs...)
}
func GetSupportedVersion(state state.State, bhash types.Hash, tupleOfUint32VersionedLocation0 uint32, tupleOfUint32VersionedLocation1 types1.VersionedLocation) (ret uint32, isSome bool, err error) {
	key, err := MakeSupportedVersionStorageKey(tupleOfUint32VersionedLocation0, tupleOfUint32VersionedLocation1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetSupportedVersionLatest(state state.State, tupleOfUint32VersionedLocation0 uint32, tupleOfUint32VersionedLocation1 types1.VersionedLocation) (ret uint32, isSome bool, err error) {
	key, err := MakeSupportedVersionStorageKey(tupleOfUint32VersionedLocation0, tupleOfUint32VersionedLocation1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for VersionNotifiers
//
//	All locations that we have requested version notifications from.
func MakeVersionNotifiersStorageKey(tupleOfUint32VersionedLocation0 uint32, tupleOfUint32VersionedLocation1 types1.VersionedLocation) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfUint32VersionedLocation0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfUint32VersionedLocation1)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "VersionNotifiers", byteArgs...)
}
func GetVersionNotifiers(state state.State, bhash types.Hash, tupleOfUint32VersionedLocation0 uint32, tupleOfUint32VersionedLocation1 types1.VersionedLocation) (ret uint64, isSome bool, err error) {
	key, err := MakeVersionNotifiersStorageKey(tupleOfUint32VersionedLocation0, tupleOfUint32VersionedLocation1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetVersionNotifiersLatest(state state.State, tupleOfUint32VersionedLocation0 uint32, tupleOfUint32VersionedLocation1 types1.VersionedLocation) (ret uint64, isSome bool, err error) {
	key, err := MakeVersionNotifiersStorageKey(tupleOfUint32VersionedLocation0, tupleOfUint32VersionedLocation1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for VersionNotifyTargets
//
//	The target locations that are subscribed to our version changes, as well as the most recent
//	of our versions we informed them of.
func MakeVersionNotifyTargetsStorageKey(tupleOfUint32VersionedLocation0 uint32, tupleOfUint32VersionedLocation1 types1.VersionedLocation) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfUint32VersionedLocation0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfUint32VersionedLocation1)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "VersionNotifyTargets", byteArgs...)
}
func GetVersionNotifyTargets(state state.State, bhash types.Hash, tupleOfUint32VersionedLocation0 uint32, tupleOfUint32VersionedLocation1 types1.VersionedLocation) (ret types1.Tuple416, isSome bool, err error) {
	key, err := MakeVersionNotifyTargetsStorageKey(tupleOfUint32VersionedLocation0, tupleOfUint32VersionedLocation1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetVersionNotifyTargetsLatest(state state.State, tupleOfUint32VersionedLocation0 uint32, tupleOfUint32VersionedLocation1 types1.VersionedLocation) (ret types1.Tuple416, isSome bool, err error) {
	key, err := MakeVersionNotifyTargetsStorageKey(tupleOfUint32VersionedLocation0, tupleOfUint32VersionedLocation1)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for VersionDiscoveryQueue id={{false [417]}}
//
//	Destinations whose latest XCM version we would like to know. Duplicates not allowed, and
//	the `u32` counter is the number of times that a send to the destination has been attempted,
//	which is used as a prioritization.
func MakeVersionDiscoveryQueueStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "VersionDiscoveryQueue")
}

var VersionDiscoveryQueueResultDefaultBytes, _ = hex.DecodeString("00")

func GetVersionDiscoveryQueue(state state.State, bhash types.Hash) (ret []types1.TupleOfVersionedLocationUint32, err error) {
	key, err := MakeVersionDiscoveryQueueStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(VersionDiscoveryQueueResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetVersionDiscoveryQueueLatest(state state.State) (ret []types1.TupleOfVersionedLocationUint32, err error) {
	key, err := MakeVersionDiscoveryQueueStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(VersionDiscoveryQueueResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for CurrentMigration id={{false [420]}}
//
//	The current migration's stage, if any.
func MakeCurrentMigrationStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "CurrentMigration")
}
func GetCurrentMigration(state state.State, bhash types.Hash) (ret types1.VersionMigrationStage, isSome bool, err error) {
	key, err := MakeCurrentMigrationStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetCurrentMigrationLatest(state state.State) (ret types1.VersionMigrationStage, isSome bool, err error) {
	key, err := MakeCurrentMigrationStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for RemoteLockedFungibles
//
//	Fungible assets which we know are locked on a remote chain.
func MakeRemoteLockedFungiblesStorageKey(tuple4210 uint32, tuple4211 [32]byte, tuple4212 types1.VersionedAssetId) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tuple4210)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tuple4211)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tuple4212)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "RemoteLockedFungibles", byteArgs...)
}
func GetRemoteLockedFungibles(state state.State, bhash types.Hash, tuple4210 uint32, tuple4211 [32]byte, tuple4212 types1.VersionedAssetId) (ret types1.RemoteLockedFungibleRecord, isSome bool, err error) {
	key, err := MakeRemoteLockedFungiblesStorageKey(tuple4210, tuple4211, tuple4212)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetRemoteLockedFungiblesLatest(state state.State, tuple4210 uint32, tuple4211 [32]byte, tuple4212 types1.VersionedAssetId) (ret types1.RemoteLockedFungibleRecord, isSome bool, err error) {
	key, err := MakeRemoteLockedFungiblesStorageKey(tuple4210, tuple4211, tuple4212)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for LockedFungibles
//
//	Fungible assets which we know are locked on this chain.
func MakeLockedFungiblesStorageKey(byteArray320 [32]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteArray320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "LockedFungibles", byteArgs...)
}
func GetLockedFungibles(state state.State, bhash types.Hash, byteArray320 [32]byte) (ret []types1.TupleOfU128VersionedLocation, isSome bool, err error) {
	key, err := MakeLockedFungiblesStorageKey(byteArray320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetLockedFungiblesLatest(state state.State, byteArray320 [32]byte) (ret []types1.TupleOfU128VersionedLocation, isSome bool, err error) {
	key, err := MakeLockedFungiblesStorageKey(byteArray320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for XcmExecutionSuspended id={{false [8]}}
//
//	Global suspension state of the XCM executor.
func MakeXcmExecutionSuspendedStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "XcmExecutionSuspended")
}

var XcmExecutionSuspendedResultDefaultBytes, _ = hex.DecodeString("00")

func GetXcmExecutionSuspended(state state.State, bhash types.Hash) (ret bool, err error) {
	key, err := MakeXcmExecutionSuspendedStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(XcmExecutionSuspendedResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetXcmExecutionSuspendedLatest(state state.State) (ret bool, err error) {
	key, err := MakeXcmExecutionSuspendedStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(XcmExecutionSuspendedResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for ShouldRecordXcm id={{false [8]}}
//
//	Whether or not incoming XCMs (both executed locally and received) should be recorded.
//	Only one XCM program will be recorded at a time.
//	This is meant to be used in runtime APIs, and it's advised it stays false
//	for all other use cases, so as to not degrade regular performance.
//
//	Only relevant if this pallet is being used as the [`xcm_executor::traits::RecordXcm`]
//	implementation in the XCM executor configuration.
func MakeShouldRecordXcmStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "ShouldRecordXcm")
}

var ShouldRecordXcmResultDefaultBytes, _ = hex.DecodeString("00")

func GetShouldRecordXcm(state state.State, bhash types.Hash) (ret bool, err error) {
	key, err := MakeShouldRecordXcmStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ShouldRecordXcmResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetShouldRecordXcmLatest(state state.State) (ret bool, err error) {
	key, err := MakeShouldRecordXcmStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ShouldRecordXcmResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for RecordedXcm id={{false [65]}}
//
//	If [`ShouldRecordXcm`] is set to true, then the last XCM program executed locally
//	will be stored here.
//	Runtime APIs can fetch the XCM that was executed by accessing this value.
//
//	Only relevant if this pallet is being used as the [`xcm_executor::traits::RecordXcm`]
//	implementation in the XCM executor configuration.
func MakeRecordedXcmStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "PolkadotXcm", "RecordedXcm")
}
func GetRecordedXcm(state state.State, bhash types.Hash) (ret []types1.Instruction, isSome bool, err error) {
	key, err := MakeRecordedXcmStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetRecordedXcmLatest(state state.State) (ret []types1.Instruction, isSome bool, err error) {
	key, err := MakeRecordedXcmStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
