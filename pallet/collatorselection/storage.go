package collatorselection

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for Invulnerables id={{false [377]}}
//
//	The invulnerable, permissioned collators. This list must be sorted.
func MakeInvulnerablesStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "CollatorSelection", "Invulnerables")
}

var InvulnerablesResultDefaultBytes, _ = hex.DecodeString("00")

func GetInvulnerables(state state.State, bhash types.Hash) (ret [][32]byte, err error) {
	key, err := MakeInvulnerablesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(InvulnerablesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetInvulnerablesLatest(state state.State) (ret [][32]byte, err error) {
	key, err := MakeInvulnerablesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(InvulnerablesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for CandidateList id={{false [378]}}
//
//	The (community, limited) collation candidates. `Candidates` and `Invulnerables` should be
//	mutually exclusive.
//
//	This list is sorted in ascending order by deposit and when the deposits are equal, the least
//	recently updated is considered greater.
func MakeCandidateListStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "CollatorSelection", "CandidateList")
}

var CandidateListResultDefaultBytes, _ = hex.DecodeString("00")

func GetCandidateList(state state.State, bhash types.Hash) (ret []types1.CandidateInfo, err error) {
	key, err := MakeCandidateListStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CandidateListResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetCandidateListLatest(state state.State) (ret []types1.CandidateInfo, err error) {
	key, err := MakeCandidateListStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CandidateListResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for LastAuthoredBlock
//
//	Last block authored by collator.
func MakeLastAuthoredBlockStorageKey(byteArray320 [32]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteArray320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "CollatorSelection", "LastAuthoredBlock", byteArgs...)
}

var LastAuthoredBlockResultDefaultBytes, _ = hex.DecodeString("0000000000000000")

func GetLastAuthoredBlock(state state.State, bhash types.Hash, byteArray320 [32]byte) (ret uint64, err error) {
	key, err := MakeLastAuthoredBlockStorageKey(byteArray320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(LastAuthoredBlockResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetLastAuthoredBlockLatest(state state.State, byteArray320 [32]byte) (ret uint64, err error) {
	key, err := MakeLastAuthoredBlockStorageKey(byteArray320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(LastAuthoredBlockResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for DesiredCandidates id={{false [4]}}
//
//	Desired number of candidates.
//
//	This should ideally always be less than [`Config::MaxCandidates`] for weights to be correct.
func MakeDesiredCandidatesStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "CollatorSelection", "DesiredCandidates")
}

var DesiredCandidatesResultDefaultBytes, _ = hex.DecodeString("00000000")

func GetDesiredCandidates(state state.State, bhash types.Hash) (ret uint32, err error) {
	key, err := MakeDesiredCandidatesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(DesiredCandidatesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetDesiredCandidatesLatest(state state.State) (ret uint32, err error) {
	key, err := MakeDesiredCandidatesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(DesiredCandidatesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for CandidacyBond id={{false [6]}}
//
//	Fixed amount to deposit to become a collator.
//
//	When a collator calls `leave_intent` they immediately receive the deposit back.
func MakeCandidacyBondStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "CollatorSelection", "CandidacyBond")
}

var CandidacyBondResultDefaultBytes, _ = hex.DecodeString("00000000000000000000000000000000")

func GetCandidacyBond(state state.State, bhash types.Hash) (ret types.U128, err error) {
	key, err := MakeCandidacyBondStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CandidacyBondResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetCandidacyBondLatest(state state.State) (ret types.U128, err error) {
	key, err := MakeCandidacyBondStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CandidacyBondResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
