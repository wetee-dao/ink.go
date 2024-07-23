package contracts

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/gen/types"
)

// Make a storage key for PristineCode
//
//	A mapping from a contract's code hash to its code.
func MakePristineCodeStorageKey(byteArray320 [32]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteArray320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Contracts", "PristineCode", byteArgs...)
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
	return types.CreateStorageKey(&types1.Meta, "Contracts", "CodeInfoOf", byteArgs...)
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

// Make a storage key for Nonce id={{false [12]}}
//
//	This is a **monotonic** counter incremented on contract instantiation.
//
//	This is used in order to generate unique trie ids for contracts.
//	The trie id of a new contract is calculated from hash(account_id, nonce).
//	The nonce is required because otherwise the following sequence would lead to
//	a possible collision of storage:
//
//	1. Create a new contract.
//	2. Terminate the contract.
//	3. Immediately recreate the contract with the same account_id.
//
//	This is bad because the contents of a trie are deleted lazily and there might be
//	storage of the old instantiation still in it when the new contract is created. Please
//	note that we can't replace the counter by the block number because the sequence above
//	can happen in the same block. We also can't keep the account counter in memory only
//	because storage is the only way to communicate across different extrinsics in the
//	same block.
//
//	# Note
//
//	Do not use it to determine the number of contracts. It won't be decremented if
//	a contract is destroyed.
func MakeNonceStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Contracts", "Nonce")
}

var NonceResultDefaultBytes, _ = hex.DecodeString("0000000000000000")

func GetNonce(state state.State, bhash types.Hash) (ret uint64, err error) {
	key, err := MakeNonceStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(NonceResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetNonceLatest(state state.State) (ret uint64, err error) {
	key, err := MakeNonceStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(NonceResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for ContractInfoOf
//
//	The code associated with a given account.
//
//	TWOX-NOTE: SAFE since `AccountId` is a secure hash.
func MakeContractInfoOfStorageKey(byteArray320 [32]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteArray320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Contracts", "ContractInfoOf", byteArgs...)
}
func GetContractInfoOf(state state.State, bhash types.Hash, byteArray320 [32]byte) (ret types1.ContractInfo, isSome bool, err error) {
	key, err := MakeContractInfoOfStorageKey(byteArray320)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetContractInfoOfLatest(state state.State, byteArray320 [32]byte) (ret types1.ContractInfo, isSome bool, err error) {
	key, err := MakeContractInfoOfStorageKey(byteArray320)
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
	return types.CreateStorageKey(&types1.Meta, "Contracts", "DeletionQueue", byteArgs...)
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

// Make a storage key for DeletionQueueCounter id={{false [298]}}
//
//	A pair of monotonic counters used to track the latest contract marked for deletion
//	and the latest deleted contract in queue.
func MakeDeletionQueueCounterStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Contracts", "DeletionQueueCounter")
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

// Make a storage key for MigrationInProgress id={{false [299]}}
//
//	A migration can span across multiple blocks. This storage defines a cursor to track the
//	progress of the migration, enabling us to resume from the last completed position.
func MakeMigrationInProgressStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Contracts", "MigrationInProgress")
}
func GetMigrationInProgress(state state.State, bhash types.Hash) (ret []byte, isSome bool, err error) {
	key, err := MakeMigrationInProgressStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetMigrationInProgressLatest(state state.State) (ret []byte, isSome bool, err error) {
	key, err := MakeMigrationInProgressStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
