package revive

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/ink.go/pallet/types"
)

// Make a storage key for PristineCode
//
//	A mapping from a contract's code hash to its code.
//	The code's size is bounded by [`crate::limits::BLOB_BYTES`] for PVM and
//	[`revm::primitives::eip170::MAX_CODE_SIZE`] for EVM bytecode.
func MakePristineCodeStorageKey(byteArray320 [32]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteArray320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Revive", "PristineCode", byteArgs...)
}
func GetPristineCode(state state.State, bhash types.Hash, byteArray320 [32]byte) (ret []byte, isSome bool, err error) {
	key, err := MakePristineCodeStorageKey(byteArray320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetPristineCodeLatest(state state.State, byteArray320 [32]byte) (ret []byte, isSome bool, err error) {
	key, err := MakePristineCodeStorageKey(byteArray320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for CodeInfoOf
//
//	A mapping from a contract's code hash to its code info.
func MakeCodeInfoOfStorageKey(byteArray320 [32]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteArray320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Revive", "CodeInfoOf", byteArgs...)
}
func GetCodeInfoOf(state state.State, bhash types.Hash, byteArray320 [32]byte) (ret types1.CodeInfo, isSome bool, err error) {
	key, err := MakeCodeInfoOfStorageKey(byteArray320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetCodeInfoOfLatest(state state.State, byteArray320 [32]byte) (ret types1.CodeInfo, isSome bool, err error) {
	key, err := MakeCodeInfoOfStorageKey(byteArray320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for AccountInfoOf
//
//	The data associated to a contract or externally owned account.
func MakeAccountInfoOfStorageKey(byteArray200 [20]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteArray200)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Revive", "AccountInfoOf", byteArgs...)
}
func GetAccountInfoOf(state state.State, bhash types.Hash, byteArray200 [20]byte) (ret types1.AccountInfo1, isSome bool, err error) {
	key, err := MakeAccountInfoOfStorageKey(byteArray200)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetAccountInfoOfLatest(state state.State, byteArray200 [20]byte) (ret types1.AccountInfo1, isSome bool, err error) {
	key, err := MakeAccountInfoOfStorageKey(byteArray200)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for ImmutableDataOf
//
//	The immutable data associated with a given account.
func MakeImmutableDataOfStorageKey(byteArray200 [20]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteArray200)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Revive", "ImmutableDataOf", byteArgs...)
}
func GetImmutableDataOf(state state.State, bhash types.Hash, byteArray200 [20]byte) (ret []byte, isSome bool, err error) {
	key, err := MakeImmutableDataOfStorageKey(byteArray200)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetImmutableDataOfLatest(state state.State, byteArray200 [20]byte) (ret []byte, isSome bool, err error) {
	key, err := MakeImmutableDataOfStorageKey(byteArray200)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for DeletionQueue
//
//	Evicted contracts that await child trie deletion.
//
//	Child trie deletion is a heavy operation depending on the amount of storage items
//	stored in said trie. Therefore this operation is performed lazily in `on_idle`.
func MakeDeletionQueueStorageKey(uint320 uint32) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Revive", "DeletionQueue", byteArgs...)
}
func GetDeletionQueue(state state.State, bhash types.Hash, uint320 uint32) (ret []byte, isSome bool, err error) {
	key, err := MakeDeletionQueueStorageKey(uint320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetDeletionQueueLatest(state state.State, uint320 uint32) (ret []byte, isSome bool, err error) {
	key, err := MakeDeletionQueueStorageKey(uint320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for DeletionQueueCounter id={{false [119]}}
//
//	A pair of monotonic counters used to track the latest contract marked for deletion
//	and the latest deleted contract in queue.
func MakeDeletionQueueCounterStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Revive", "DeletionQueueCounter")
}

var DeletionQueueCounterResultDefaultBytes, _ = hex.DecodeString("0000000000000000")

func GetDeletionQueueCounter(state state.State, bhash types.Hash) (ret types1.DeletionQueueManager, err error) {
	key, err := MakeDeletionQueueCounterStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(DeletionQueueCounterResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetDeletionQueueCounterLatest(state state.State) (ret types1.DeletionQueueManager, err error) {
	key, err := MakeDeletionQueueCounterStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(DeletionQueueCounterResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for OriginalAccount
//
//	Map a Ethereum address to its original `AccountId32`.
//
//	When deriving a `H160` from an `AccountId32` we use a hash function. In order to
//	reconstruct the original account we need to store the reverse mapping here.
//	Register your `AccountId32` using [`Pallet::map_account`] in order to
//	use it with this pallet.
func MakeOriginalAccountStorageKey(byteArray200 [20]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteArray200)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Revive", "OriginalAccount", byteArgs...)
}
func GetOriginalAccount(state state.State, bhash types.Hash, byteArray200 [20]byte) (ret [32]byte, isSome bool, err error) {
	key, err := MakeOriginalAccountStorageKey(byteArray200)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetOriginalAccountLatest(state state.State, byteArray200 [20]byte) (ret [32]byte, isSome bool, err error) {
	key, err := MakeOriginalAccountStorageKey(byteArray200)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for EthereumBlock id={{false [120]}}
//
//	The current Ethereum block that is stored in the `on_finalize` method.
//
//	# Note
//
//	This could be further optimized into the future to store only the minimum
//	information needed to reconstruct the Ethereum block at the RPC level.
//
//	Since the block is convenient to have around, and the extra details are capped
//	by a few hashes and the vector of transaction hashes, we store the block here.
func MakeEthereumBlockStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Revive", "EthereumBlock")
}

var EthereumBlockResultDefaultBytes, _ = hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

func GetEthereumBlock(state state.State, bhash types.Hash) (ret types1.Block, err error) {
	key, err := MakeEthereumBlockStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(EthereumBlockResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetEthereumBlockLatest(state state.State) (ret types1.Block, err error) {
	key, err := MakeEthereumBlockStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(EthereumBlockResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for BlockHash
//
//	Mapping for block number and hashes.
//
//	The maximum number of elements stored is capped by the block hash count `BLOCK_HASH_COUNT`.
func MakeBlockHashStorageKey(uint320 uint32) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Revive", "BlockHash", byteArgs...)
}

var BlockHashResultDefaultBytes, _ = hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")

func GetBlockHash(state state.State, bhash types.Hash, uint320 uint32) (ret [32]byte, err error) {
	key, err := MakeBlockHashStorageKey(uint320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(BlockHashResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetBlockHashLatest(state state.State, uint320 uint32) (ret [32]byte, err error) {
	key, err := MakeBlockHashStorageKey(uint320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(BlockHashResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for ReceiptInfoData id={{false [148]}}
//
//	The details needed to reconstruct the receipt info offchain.
//
//	This contains valuable information about the gas used by the transaction.
//
//	NOTE: The item is unbound and should therefore never be read on chain.
//	It could otherwise inflate the PoV size of a block.
func MakeReceiptInfoDataStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Revive", "ReceiptInfoData")
}

var ReceiptInfoDataResultDefaultBytes, _ = hex.DecodeString("00")

func GetReceiptInfoData(state state.State, bhash types.Hash) (ret []types1.ReceiptGasInfo, err error) {
	key, err := MakeReceiptInfoDataStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ReceiptInfoDataResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetReceiptInfoDataLatest(state state.State) (ret []types1.ReceiptGasInfo, err error) {
	key, err := MakeReceiptInfoDataStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ReceiptInfoDataResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for EthBlockBuilderIR id={{false [150]}}
//
//	Incremental ethereum block builder.
func MakeEthBlockBuilderIRStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Revive", "EthBlockBuilderIR")
}

var EthBlockBuilderIRResultDefaultBytes, _ = hex.DecodeString("0000000000000000000100000000000000000000000000000000010000000000000000743ba40b000000000000000000000000000000000000000000000000000000ffffffffffffffff0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

func GetEthBlockBuilderIR(state state.State, bhash types.Hash) (ret types1.EthereumBlockBuilderIR, err error) {
	key, err := MakeEthBlockBuilderIRStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(EthBlockBuilderIRResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetEthBlockBuilderIRLatest(state state.State) (ret types1.EthereumBlockBuilderIR, err error) {
	key, err := MakeEthBlockBuilderIRStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(EthBlockBuilderIRResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for EthBlockBuilderFirstValues id={{false [153]}}
//
//	The first transaction and receipt of the ethereum block.
//
//	These values are moved out of the `EthBlockBuilderIR` to avoid serializing and
//	deserializing them on every transaction. Instead, they are loaded when needed.
func MakeEthBlockBuilderFirstValuesStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Revive", "EthBlockBuilderFirstValues")
}

var EthBlockBuilderFirstValuesResultDefaultBytes, _ = hex.DecodeString("00")

func GetEthBlockBuilderFirstValues(state state.State, bhash types.Hash) (ret types1.OptionTTupleOfByteSliceByteSlice, err error) {
	key, err := MakeEthBlockBuilderFirstValuesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(EthBlockBuilderFirstValuesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetEthBlockBuilderFirstValuesLatest(state state.State) (ret types1.OptionTTupleOfByteSliceByteSlice, err error) {
	key, err := MakeEthBlockBuilderFirstValuesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(EthBlockBuilderFirstValuesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for DebugSettingsOf id={{false [154]}}
//
//	Debugging settings that can be configured when DebugEnabled config is true.
func MakeDebugSettingsOfStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Revive", "DebugSettingsOf")
}

var DebugSettingsOfResultDefaultBytes, _ = hex.DecodeString("00000000")

func GetDebugSettingsOf(state state.State, bhash types.Hash) (ret types1.DebugSettings, err error) {
	key, err := MakeDebugSettingsOfStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(DebugSettingsOfResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetDebugSettingsOfLatest(state state.State) (ret types1.DebugSettings, err error) {
	key, err := MakeDebugSettingsOfStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(DebugSettingsOfResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
