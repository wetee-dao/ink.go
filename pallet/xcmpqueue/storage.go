package xcmpqueue

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for InboundXcmpSuspended id={{false [392]}}
//
//	The suspended inbound XCMP channels. All others are not suspended.
//
//	This is a `StorageValue` instead of a `StorageMap` since we expect multiple reads per block
//	to different keys with a one byte payload. The access to `BoundedBTreeSet` will be cached
//	within the block and therefore only included once in the proof size.
//
//	NOTE: The PoV benchmarking cannot know this and will over-estimate, but the actual proof
//	will be smaller.
func MakeInboundXcmpSuspendedStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "XcmpQueue", "InboundXcmpSuspended")
}

var InboundXcmpSuspendedResultDefaultBytes, _ = hex.DecodeString("00")

func GetInboundXcmpSuspended(state state.State, bhash types.Hash) (ret []uint32, err error) {
	key, err := MakeInboundXcmpSuspendedStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(InboundXcmpSuspendedResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetInboundXcmpSuspendedLatest(state state.State) (ret []uint32, err error) {
	key, err := MakeInboundXcmpSuspendedStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(InboundXcmpSuspendedResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for OutboundXcmpStatus id={{false [395]}}
//
//	The non-empty XCMP channels in order of becoming non-empty, and the index of the first
//	and last outbound message. If the two indices are equal, then it indicates an empty
//	queue and there must be a non-`Ok` `OutboundStatus`. We assume queues grow no greater
//	than 65535 items. Queue indices for normal messages begin at one; zero is reserved in
//	case of the need to send a high-priority signal message this block.
//	The bool is true if there is a signal message waiting to be sent.
func MakeOutboundXcmpStatusStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "XcmpQueue", "OutboundXcmpStatus")
}

var OutboundXcmpStatusResultDefaultBytes, _ = hex.DecodeString("00")

func GetOutboundXcmpStatus(state state.State, bhash types.Hash) (ret []types1.OutboundChannelDetails, err error) {
	key, err := MakeOutboundXcmpStatusStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(OutboundXcmpStatusResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetOutboundXcmpStatusLatest(state state.State) (ret []types1.OutboundChannelDetails, err error) {
	key, err := MakeOutboundXcmpStatusStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(OutboundXcmpStatusResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for OutboundXcmpMessages
//
//	The messages outbound in a given XCMP channel.
func MakeOutboundXcmpMessagesStorageKey(tupleOfUint32Uint160 uint32, tupleOfUint32Uint161 uint16) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfUint32Uint160)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfUint32Uint161)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "XcmpQueue", "OutboundXcmpMessages", byteArgs...)
}

var OutboundXcmpMessagesResultDefaultBytes, _ = hex.DecodeString("00")

func GetOutboundXcmpMessages(state state.State, bhash types.Hash, tupleOfUint32Uint160 uint32, tupleOfUint32Uint161 uint16) (ret []byte, err error) {
	key, err := MakeOutboundXcmpMessagesStorageKey(tupleOfUint32Uint160, tupleOfUint32Uint161)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(OutboundXcmpMessagesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetOutboundXcmpMessagesLatest(state state.State, tupleOfUint32Uint160 uint32, tupleOfUint32Uint161 uint16) (ret []byte, err error) {
	key, err := MakeOutboundXcmpMessagesStorageKey(tupleOfUint32Uint160, tupleOfUint32Uint161)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(OutboundXcmpMessagesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for SignalMessages
//
//	Any signal messages waiting to be sent.
func MakeSignalMessagesStorageKey(uint320 uint32) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "XcmpQueue", "SignalMessages", byteArgs...)
}

var SignalMessagesResultDefaultBytes, _ = hex.DecodeString("00")

func GetSignalMessages(state state.State, bhash types.Hash, uint320 uint32) (ret []byte, err error) {
	key, err := MakeSignalMessagesStorageKey(uint320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(SignalMessagesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetSignalMessagesLatest(state state.State, uint320 uint32) (ret []byte, err error) {
	key, err := MakeSignalMessagesStorageKey(uint320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(SignalMessagesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for QueueConfig id={{false [401]}}
//
//	The configuration which controls the dynamics of the outbound queue.
func MakeQueueConfigStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "XcmpQueue", "QueueConfig")
}

var QueueConfigResultDefaultBytes, _ = hex.DecodeString("200000003000000008000000")

func GetQueueConfig(state state.State, bhash types.Hash) (ret types1.QueueConfigData, err error) {
	key, err := MakeQueueConfigStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(QueueConfigResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetQueueConfigLatest(state state.State) (ret types1.QueueConfigData, err error) {
	key, err := MakeQueueConfigStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(QueueConfigResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for QueueSuspended id={{false [8]}}
//
//	Whether or not the XCMP queue is suspended from executing incoming XCMs or not.
func MakeQueueSuspendedStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "XcmpQueue", "QueueSuspended")
}

var QueueSuspendedResultDefaultBytes, _ = hex.DecodeString("00")

func GetQueueSuspended(state state.State, bhash types.Hash) (ret bool, err error) {
	key, err := MakeQueueSuspendedStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(QueueSuspendedResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetQueueSuspendedLatest(state state.State) (ret bool, err error) {
	key, err := MakeQueueSuspendedStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(QueueSuspendedResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for DeliveryFeeFactor
//
//	The factor to multiply the base delivery fee by.
func MakeDeliveryFeeFactorStorageKey(uint320 uint32) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "XcmpQueue", "DeliveryFeeFactor", byteArgs...)
}

var DeliveryFeeFactorResultDefaultBytes, _ = hex.DecodeString("000064a7b3b6e00d0000000000000000")

func GetDeliveryFeeFactor(state state.State, bhash types.Hash, uint320 uint32) (ret types.U128, err error) {
	key, err := MakeDeliveryFeeFactorStorageKey(uint320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(DeliveryFeeFactorResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetDeliveryFeeFactorLatest(state state.State, uint320 uint32) (ret types.U128, err error) {
	key, err := MakeDeliveryFeeFactorStorageKey(uint320)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(DeliveryFeeFactorResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
