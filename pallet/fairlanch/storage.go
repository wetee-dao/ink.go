package fairlanch

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for BlockReward id={{false [338]}}
//
//	current block reward
func MakeBlockRewardStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "Fairlanch", "BlockReward")
}

var BlockRewardResultDefaultBytes, _ = hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")

func GetBlockReward(state state.State, bhash types.Hash) (ret types1.TupleOfU128U128, err error) {
	key, err := MakeBlockRewardStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(BlockRewardResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetBlockRewardLatest(state state.State) (ret types1.TupleOfU128U128, err error) {
	key, err := MakeBlockRewardStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(BlockRewardResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for Stakings
//
//	Staking
func MakeStakingsStorageKey(tupleOfByteArray32Uint6410 [32]byte, tupleOfByteArray32Uint6411 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfByteArray32Uint6410)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfByteArray32Uint6411)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Fairlanch", "Stakings", byteArgs...)
}
func GetStakings(state state.State, bhash types.Hash, tupleOfByteArray32Uint6410 [32]byte, tupleOfByteArray32Uint6411 uint64) (ret types1.Wstaking, isSome bool, err error) {
	key, err := MakeStakingsStorageKey(tupleOfByteArray32Uint6410, tupleOfByteArray32Uint6411)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetStakingsLatest(state state.State, tupleOfByteArray32Uint6410 [32]byte, tupleOfByteArray32Uint6411 uint64) (ret types1.Wstaking, isSome bool, err error) {
	key, err := MakeStakingsStorageKey(tupleOfByteArray32Uint6410, tupleOfByteArray32Uint6411)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for NextStakingRewards
//
//	next block reward
//	下一次奖励的区块高度
//	24小时执行一次奖励
func MakeNextStakingRewardsStorageKey(tupleOfUint64ByteArray320 uint64, tupleOfUint64ByteArray321 [32]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfUint64ByteArray320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfUint64ByteArray321)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Fairlanch", "NextStakingRewards", byteArgs...)
}
func GetNextStakingRewards(state state.State, bhash types.Hash, tupleOfUint64ByteArray320 uint64, tupleOfUint64ByteArray321 [32]byte) (ret []types1.TupleOfUint64U128, isSome bool, err error) {
	key, err := MakeNextStakingRewardsStorageKey(tupleOfUint64ByteArray320, tupleOfUint64ByteArray321)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetNextStakingRewardsLatest(state state.State, tupleOfUint64ByteArray320 uint64, tupleOfUint64ByteArray321 [32]byte) (ret []types1.TupleOfUint64U128, isSome bool, err error) {
	key, err := MakeNextStakingRewardsStorageKey(tupleOfUint64ByteArray320, tupleOfUint64ByteArray321)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for UserReward
func MakeUserRewardStorageKey(byteArray320 [32]byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byteArray320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Fairlanch", "UserReward", byteArgs...)
}

var UserRewardResultDefaultBytes, _ = hex.DecodeString("000000000000000000000000000000000000000000000000")

func GetUserReward(state state.State, bhash types.Hash, byteArray320 [32]byte) (ret types1.TupleOfUint64U128, err error) {
	key, err := MakeUserRewardStorageKey(byteArray320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(UserRewardResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetUserRewardLatest(state state.State, byteArray320 [32]byte) (ret types1.TupleOfUint64U128, err error) {
	key, err := MakeUserRewardStorageKey(byteArray320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(UserRewardResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for EpochRewardTotal
//
//	epoch_reward_total
//	奖励奖励池
func MakeEpochRewardTotalStorageKey(uint640 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Fairlanch", "EpochRewardTotal", byteArgs...)
}

var EpochRewardTotalResultDefaultBytes, _ = hex.DecodeString("00000000000000000000000000000000")

func GetEpochRewardTotal(state state.State, bhash types.Hash, uint640 uint64) (ret types.U128, err error) {
	key, err := MakeEpochRewardTotalStorageKey(uint640)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(EpochRewardTotalResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetEpochRewardTotalLatest(state state.State, uint640 uint64) (ret types.U128, err error) {
	key, err := MakeEpochRewardTotalStorageKey(uint640)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(EpochRewardTotalResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for Economics
//
//	economics
//	经济模型
//	0 => node mint reward
//	1 => tee mint reward
//	3 => app mint reward
//	WeAssetId => staking reward
func MakeEconomicsStorageKey(uint640 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Fairlanch", "Economics", byteArgs...)
}
func GetEconomics(state state.State, bhash types.Hash, uint640 uint64) (ret byte, isSome bool, err error) {
	key, err := MakeEconomicsStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetEconomicsLatest(state state.State, uint640 uint64) (ret byte, isSome bool, err error) {
	key, err := MakeEconomicsStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for Vtoken2token
//
//	vtoken transfer rate
//	vtoken 转换为 token 的比例
func MakeVtoken2tokenStorageKey(uint640 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "Fairlanch", "Vtoken2token", byteArgs...)
}
func GetVtoken2token(state state.State, bhash types.Hash, uint640 uint64) (ret types1.TupleOfUint64TupleOfU128U128, isSome bool, err error) {
	key, err := MakeVtoken2tokenStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetVtoken2tokenLatest(state state.State, uint640 uint64) (ret types1.TupleOfUint64TupleOfU128U128, isSome bool, err error) {
	key, err := MakeVtoken2tokenStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
