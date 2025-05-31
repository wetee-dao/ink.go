package assets

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for Asset
//
//	Details of an asset.
func MakeAssetStorageKey(uint320 uint32) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Assets", "Asset", byteArgs...)
}
func GetAsset(state state.State, bhash types.Hash, uint320 uint32) (ret types1.AssetDetails, isSome bool, err error) {
	key, err := MakeAssetStorageKey(uint320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetAssetLatest(state state.State, uint320 uint32) (ret types1.AssetDetails, isSome bool, err error) {
	key, err := MakeAssetStorageKey(uint320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for Account
//
//	The holdings of a specific account for a specific asset.
func MakeAccountStorageKey(tupleOfUint32ByteArray320 uint32, tupleOfUint32ByteArray321 [32]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfUint32ByteArray320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfUint32ByteArray321)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Assets", "Account", byteArgs...)
}
func GetAccount(state state.State, bhash types.Hash, tupleOfUint32ByteArray320 uint32, tupleOfUint32ByteArray321 [32]byte) (ret types1.AssetAccount, isSome bool, err error) {
	key, err := MakeAccountStorageKey(tupleOfUint32ByteArray320, tupleOfUint32ByteArray321)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetAccountLatest(state state.State, tupleOfUint32ByteArray320 uint32, tupleOfUint32ByteArray321 [32]byte) (ret types1.AssetAccount, isSome bool, err error) {
	key, err := MakeAccountStorageKey(tupleOfUint32ByteArray320, tupleOfUint32ByteArray321)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for Approvals
//
//	Approved balance transfers. First balance is the amount approved for transfer. Second
//	is the amount of `T::Currency` reserved for storing this.
//	First key is the asset ID, second key is the owner and third key is the delegate.
func MakeApprovalsStorageKey(tuple1220 uint32, tuple1221 [32]byte, tuple1222 [32]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tuple1220)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tuple1221)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tuple1222)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Assets", "Approvals", byteArgs...)
}
func GetApprovals(state state.State, bhash types.Hash, tuple1220 uint32, tuple1221 [32]byte, tuple1222 [32]byte) (ret types1.Approval, isSome bool, err error) {
	key, err := MakeApprovalsStorageKey(tuple1220, tuple1221, tuple1222)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetApprovalsLatest(state state.State, tuple1220 uint32, tuple1221 [32]byte, tuple1222 [32]byte) (ret types1.Approval, isSome bool, err error) {
	key, err := MakeApprovalsStorageKey(tuple1220, tuple1221, tuple1222)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for Metadata
//
//	Metadata of an asset.
func MakeMetadataStorageKey(uint320 uint32) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Assets", "Metadata", byteArgs...)
}

var MetadataResultDefaultBytes, _ = hex.DecodeString("0000000000000000000000000000000000000000")

func GetMetadata(state state.State, bhash types.Hash, uint320 uint32) (ret types1.AssetMetadata, err error) {
	key, err := MakeMetadataStorageKey(uint320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(MetadataResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetMetadataLatest(state state.State, uint320 uint32) (ret types1.AssetMetadata, err error) {
	key, err := MakeMetadataStorageKey(uint320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(MetadataResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for NextAssetId id={{false [4]}}
//
//	The asset ID enforced for the next asset creation, if any present. Otherwise, this storage
//	item has no effect.
//
//	This can be useful for setting up constraints for IDs of the new assets. For example, by
//	providing an initial [`NextAssetId`] and using the [`crate::AutoIncAssetId`] callback, an
//	auto-increment model can be applied to all new asset IDs.
//
//	The initial next asset ID can be set using the [`GenesisConfig`] or the
//	[SetNextAssetId](`migration::next_asset_id::SetNextAssetId`) migration.
func MakeNextAssetIdStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Assets", "NextAssetId")
}
func GetNextAssetId(state state.State, bhash types.Hash) (ret uint32, isSome bool, err error) {
	key, err := MakeNextAssetIdStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetNextAssetIdLatest(state state.State) (ret uint32, isSome bool, err error) {
	key, err := MakeNextAssetIdStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
