package parachainsystem

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for UnincludedSegment id={{false [188]}}
//
//	Latest included block descendants the runtime accepted. In other words, these are
//	ancestors of the currently executing block which have not been included in the observed
//	relay-chain state.
//
//	The segment length is limited by the capacity returned from the [`ConsensusHook`] configured
//	in the pallet.
func MakeUnincludedSegmentStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "UnincludedSegment")
}

var UnincludedSegmentResultDefaultBytes, _ = hex.DecodeString("00")

func GetUnincludedSegment(state state.State, bhash types.Hash) (ret []types1.Ancestor, err error) {
	key, err := MakeUnincludedSegmentStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(UnincludedSegmentResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetUnincludedSegmentLatest(state state.State) (ret []types1.Ancestor, err error) {
	key, err := MakeUnincludedSegmentStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(UnincludedSegmentResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for AggregatedUnincludedSegment id={{false [198]}}
//
//	Storage field that keeps track of bandwidth used by the unincluded segment along with the
//	latest HRMP watermark. Used for limiting the acceptance of new blocks with
//	respect to relay chain constraints.
func MakeAggregatedUnincludedSegmentStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "AggregatedUnincludedSegment")
}
func GetAggregatedUnincludedSegment(state state.State, bhash types.Hash) (ret types1.SegmentTracker, isSome bool, err error) {
	key, err := MakeAggregatedUnincludedSegmentStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetAggregatedUnincludedSegmentLatest(state state.State) (ret types1.SegmentTracker, isSome bool, err error) {
	key, err := MakeAggregatedUnincludedSegmentStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for PendingValidationCode id={{false [14]}}
//
//	In case of a scheduled upgrade, this storage field contains the validation code to be
//	applied.
//
//	As soon as the relay chain gives us the go-ahead signal, we will overwrite the
//	[`:code`][sp_core::storage::well_known_keys::CODE] which will result the next block process
//	with the new validation code. This concludes the upgrade process.
func MakePendingValidationCodeStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "PendingValidationCode")
}

var PendingValidationCodeResultDefaultBytes, _ = hex.DecodeString("00")

func GetPendingValidationCode(state state.State, bhash types.Hash) (ret []byte, err error) {
	key, err := MakePendingValidationCodeStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(PendingValidationCodeResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetPendingValidationCodeLatest(state state.State) (ret []byte, err error) {
	key, err := MakePendingValidationCodeStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(PendingValidationCodeResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for NewValidationCode id={{false [14]}}
//
//	Validation code that is set by the parachain and is to be communicated to collator and
//	consequently the relay-chain.
//
//	This will be cleared in `on_initialize` of each new block if no other pallet already set
//	the value.
func MakeNewValidationCodeStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "NewValidationCode")
}
func GetNewValidationCode(state state.State, bhash types.Hash) (ret []byte, isSome bool, err error) {
	key, err := MakeNewValidationCodeStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetNewValidationCodeLatest(state state.State) (ret []byte, isSome bool, err error) {
	key, err := MakeNewValidationCodeStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for ValidationData id={{false [200]}}
//
//	The [`PersistedValidationData`] set for this block.
//	This value is expected to be set only once per block and it's never stored
//	in the trie.
func MakeValidationDataStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "ValidationData")
}
func GetValidationData(state state.State, bhash types.Hash) (ret types1.PersistedValidationData, isSome bool, err error) {
	key, err := MakeValidationDataStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetValidationDataLatest(state state.State) (ret types1.PersistedValidationData, isSome bool, err error) {
	key, err := MakeValidationDataStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for DidSetValidationCode id={{false [8]}}
//
//	Were the validation data set to notify the relay chain?
func MakeDidSetValidationCodeStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "DidSetValidationCode")
}

var DidSetValidationCodeResultDefaultBytes, _ = hex.DecodeString("00")

func GetDidSetValidationCode(state state.State, bhash types.Hash) (ret bool, err error) {
	key, err := MakeDidSetValidationCodeStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(DidSetValidationCodeResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetDidSetValidationCodeLatest(state state.State) (ret bool, err error) {
	key, err := MakeDidSetValidationCodeStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(DidSetValidationCodeResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for LastRelayChainBlockNumber id={{false [4]}}
//
//	The relay chain block number associated with the last parachain block.
//
//	This is updated in `on_finalize`.
func MakeLastRelayChainBlockNumberStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "LastRelayChainBlockNumber")
}

var LastRelayChainBlockNumberResultDefaultBytes, _ = hex.DecodeString("00000000")

func GetLastRelayChainBlockNumber(state state.State, bhash types.Hash) (ret uint32, err error) {
	key, err := MakeLastRelayChainBlockNumberStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(LastRelayChainBlockNumberResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetLastRelayChainBlockNumberLatest(state state.State) (ret uint32, err error) {
	key, err := MakeLastRelayChainBlockNumberStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(LastRelayChainBlockNumberResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for UpgradeRestrictionSignal id={{false [202]}}
//
//	An option which indicates if the relay-chain restricts signalling a validation code upgrade.
//	In other words, if this is `Some` and [`NewValidationCode`] is `Some` then the produced
//	candidate will be invalid.
//
//	This storage item is a mirror of the corresponding value for the current parachain from the
//	relay-chain. This value is ephemeral which means it doesn't hit the storage. This value is
//	set after the inherent.
func MakeUpgradeRestrictionSignalStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "UpgradeRestrictionSignal")
}

var UpgradeRestrictionSignalResultDefaultBytes, _ = hex.DecodeString("00")

func GetUpgradeRestrictionSignal(state state.State, bhash types.Hash) (ret types1.OptionTUpgradeRestriction, err error) {
	key, err := MakeUpgradeRestrictionSignalStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(UpgradeRestrictionSignalResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetUpgradeRestrictionSignalLatest(state state.State) (ret types1.OptionTUpgradeRestriction, err error) {
	key, err := MakeUpgradeRestrictionSignalStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(UpgradeRestrictionSignalResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for UpgradeGoAhead id={{false [196]}}
//
//	Optional upgrade go-ahead signal from the relay-chain.
//
//	This storage item is a mirror of the corresponding value for the current parachain from the
//	relay-chain. This value is ephemeral which means it doesn't hit the storage. This value is
//	set after the inherent.
func MakeUpgradeGoAheadStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "UpgradeGoAhead")
}

var UpgradeGoAheadResultDefaultBytes, _ = hex.DecodeString("00")

func GetUpgradeGoAhead(state state.State, bhash types.Hash) (ret types1.OptionTUpgradeGoAhead, err error) {
	key, err := MakeUpgradeGoAheadStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(UpgradeGoAheadResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetUpgradeGoAheadLatest(state state.State) (ret types1.OptionTUpgradeGoAhead, err error) {
	key, err := MakeUpgradeGoAheadStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(UpgradeGoAheadResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for RelayStateProof id={{false [204]}}
//
//	The state proof for the last relay parent block.
//
//	This field is meant to be updated each block with the validation data inherent. Therefore,
//	before processing of the inherent, e.g. in `on_initialize` this data may be stale.
//
//	This data is also absent from the genesis.
func MakeRelayStateProofStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "RelayStateProof")
}
func GetRelayStateProof(state state.State, bhash types.Hash) (ret [][]byte, isSome bool, err error) {
	key, err := MakeRelayStateProofStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetRelayStateProofLatest(state state.State) (ret [][]byte, isSome bool, err error) {
	key, err := MakeRelayStateProofStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for RelevantMessagingState id={{false [206]}}
//
//	The snapshot of some state related to messaging relevant to the current parachain as per
//	the relay parent.
//
//	This field is meant to be updated each block with the validation data inherent. Therefore,
//	before processing of the inherent, e.g. in `on_initialize` this data may be stale.
//
//	This data is also absent from the genesis.
func MakeRelevantMessagingStateStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "RelevantMessagingState")
}
func GetRelevantMessagingState(state state.State, bhash types.Hash) (ret types1.MessagingStateSnapshot, isSome bool, err error) {
	key, err := MakeRelevantMessagingStateStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetRelevantMessagingStateLatest(state state.State) (ret types1.MessagingStateSnapshot, isSome bool, err error) {
	key, err := MakeRelevantMessagingStateStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for HostConfiguration id={{false [211]}}
//
//	The parachain host configuration that was obtained from the relay parent.
//
//	This field is meant to be updated each block with the validation data inherent. Therefore,
//	before processing of the inherent, e.g. in `on_initialize` this data may be stale.
//
//	This data is also absent from the genesis.
func MakeHostConfigurationStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "HostConfiguration")
}
func GetHostConfiguration(state state.State, bhash types.Hash) (ret types1.AbridgedHostConfiguration, isSome bool, err error) {
	key, err := MakeHostConfigurationStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetHostConfigurationLatest(state state.State) (ret types1.AbridgedHostConfiguration, isSome bool, err error) {
	key, err := MakeHostConfigurationStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for LastDmqMqcHead id={{false [213]}}
//
//	The last downward message queue chain head we have observed.
//
//	This value is loaded before and saved after processing inbound downward messages carried
//	by the system inherent.
func MakeLastDmqMqcHeadStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "LastDmqMqcHead")
}

var LastDmqMqcHeadResultDefaultBytes, _ = hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")

func GetLastDmqMqcHead(state state.State, bhash types.Hash) (ret [32]byte, err error) {
	key, err := MakeLastDmqMqcHeadStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(LastDmqMqcHeadResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetLastDmqMqcHeadLatest(state state.State) (ret [32]byte, err error) {
	key, err := MakeLastDmqMqcHeadStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(LastDmqMqcHeadResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for LastHrmpMqcHeads id={{false [214]}}
//
//	The message queue chain heads we have observed per each channel incoming channel.
//
//	This value is loaded before and saved after processing inbound downward messages carried
//	by the system inherent.
func MakeLastHrmpMqcHeadsStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "LastHrmpMqcHeads")
}

var LastHrmpMqcHeadsResultDefaultBytes, _ = hex.DecodeString("00")

func GetLastHrmpMqcHeads(state state.State, bhash types.Hash) (ret []types1.TupleOfUint32ByteArray32, err error) {
	key, err := MakeLastHrmpMqcHeadsStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(LastHrmpMqcHeadsResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetLastHrmpMqcHeadsLatest(state state.State) (ret []types1.TupleOfUint32ByteArray32, err error) {
	key, err := MakeLastHrmpMqcHeadsStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(LastHrmpMqcHeadsResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for ProcessedDownwardMessages id={{false [4]}}
//
//	Number of downward messages processed in a block.
//
//	This will be cleared in `on_initialize` of each new block.
func MakeProcessedDownwardMessagesStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "ProcessedDownwardMessages")
}

var ProcessedDownwardMessagesResultDefaultBytes, _ = hex.DecodeString("00000000")

func GetProcessedDownwardMessages(state state.State, bhash types.Hash) (ret uint32, err error) {
	key, err := MakeProcessedDownwardMessagesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ProcessedDownwardMessagesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetProcessedDownwardMessagesLatest(state state.State) (ret uint32, err error) {
	key, err := MakeProcessedDownwardMessagesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(ProcessedDownwardMessagesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for HrmpWatermark id={{false [4]}}
//
//	HRMP watermark that was set in a block.
//
//	This will be cleared in `on_initialize` of each new block.
func MakeHrmpWatermarkStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "HrmpWatermark")
}

var HrmpWatermarkResultDefaultBytes, _ = hex.DecodeString("00000000")

func GetHrmpWatermark(state state.State, bhash types.Hash) (ret uint32, err error) {
	key, err := MakeHrmpWatermarkStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(HrmpWatermarkResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetHrmpWatermarkLatest(state state.State) (ret uint32, err error) {
	key, err := MakeHrmpWatermarkStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(HrmpWatermarkResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for HrmpOutboundMessages id={{false [217]}}
//
//	HRMP messages that were sent in a block.
//
//	This will be cleared in `on_initialize` of each new block.
func MakeHrmpOutboundMessagesStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "HrmpOutboundMessages")
}

var HrmpOutboundMessagesResultDefaultBytes, _ = hex.DecodeString("00")

func GetHrmpOutboundMessages(state state.State, bhash types.Hash) (ret []types1.OutboundHrmpMessage, err error) {
	key, err := MakeHrmpOutboundMessagesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(HrmpOutboundMessagesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetHrmpOutboundMessagesLatest(state state.State) (ret []types1.OutboundHrmpMessage, err error) {
	key, err := MakeHrmpOutboundMessagesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(HrmpOutboundMessagesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for UpwardMessages id={{false [175]}}
//
//	Upward messages that were sent in a block.
//
//	This will be cleared in `on_initialize` of each new block.
func MakeUpwardMessagesStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "UpwardMessages")
}

var UpwardMessagesResultDefaultBytes, _ = hex.DecodeString("00")

func GetUpwardMessages(state state.State, bhash types.Hash) (ret [][]byte, err error) {
	key, err := MakeUpwardMessagesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(UpwardMessagesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetUpwardMessagesLatest(state state.State) (ret [][]byte, err error) {
	key, err := MakeUpwardMessagesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(UpwardMessagesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for PendingUpwardMessages id={{false [175]}}
//
//	Upward messages that are still pending and not yet send to the relay chain.
func MakePendingUpwardMessagesStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "PendingUpwardMessages")
}

var PendingUpwardMessagesResultDefaultBytes, _ = hex.DecodeString("00")

func GetPendingUpwardMessages(state state.State, bhash types.Hash) (ret [][]byte, err error) {
	key, err := MakePendingUpwardMessagesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(PendingUpwardMessagesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetPendingUpwardMessagesLatest(state state.State) (ret [][]byte, err error) {
	key, err := MakePendingUpwardMessagesStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(PendingUpwardMessagesResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for UpwardDeliveryFeeFactor id={{false [219]}}
//
//	The factor to multiply the base delivery fee by for UMP.
func MakeUpwardDeliveryFeeFactorStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "UpwardDeliveryFeeFactor")
}

var UpwardDeliveryFeeFactorResultDefaultBytes, _ = hex.DecodeString("000064a7b3b6e00d0000000000000000")

func GetUpwardDeliveryFeeFactor(state state.State, bhash types.Hash) (ret types.U128, err error) {
	key, err := MakeUpwardDeliveryFeeFactorStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(UpwardDeliveryFeeFactorResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetUpwardDeliveryFeeFactorLatest(state state.State) (ret types.U128, err error) {
	key, err := MakeUpwardDeliveryFeeFactorStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(UpwardDeliveryFeeFactorResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for AnnouncedHrmpMessagesPerCandidate id={{false [4]}}
//
//	The number of HRMP messages we observed in `on_initialize` and thus used that number for
//	announcing the weight of `on_initialize` and `on_finalize`.
func MakeAnnouncedHrmpMessagesPerCandidateStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "AnnouncedHrmpMessagesPerCandidate")
}

var AnnouncedHrmpMessagesPerCandidateResultDefaultBytes, _ = hex.DecodeString("00000000")

func GetAnnouncedHrmpMessagesPerCandidate(state state.State, bhash types.Hash) (ret uint32, err error) {
	key, err := MakeAnnouncedHrmpMessagesPerCandidateStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(AnnouncedHrmpMessagesPerCandidateResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetAnnouncedHrmpMessagesPerCandidateLatest(state state.State) (ret uint32, err error) {
	key, err := MakeAnnouncedHrmpMessagesPerCandidateStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(AnnouncedHrmpMessagesPerCandidateResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for ReservedXcmpWeightOverride id={{false [10]}}
//
//	The weight we reserve at the beginning of the block for processing XCMP messages. This
//	overrides the amount set in the Config trait.
func MakeReservedXcmpWeightOverrideStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "ReservedXcmpWeightOverride")
}
func GetReservedXcmpWeightOverride(state state.State, bhash types.Hash) (ret types1.Weight, isSome bool, err error) {
	key, err := MakeReservedXcmpWeightOverrideStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetReservedXcmpWeightOverrideLatest(state state.State) (ret types1.Weight, isSome bool, err error) {
	key, err := MakeReservedXcmpWeightOverrideStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for ReservedDmpWeightOverride id={{false [10]}}
//
//	The weight we reserve at the beginning of the block for processing DMP messages. This
//	overrides the amount set in the Config trait.
func MakeReservedDmpWeightOverrideStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "ReservedDmpWeightOverride")
}
func GetReservedDmpWeightOverride(state state.State, bhash types.Hash) (ret types1.Weight, isSome bool, err error) {
	key, err := MakeReservedDmpWeightOverrideStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetReservedDmpWeightOverrideLatest(state state.State) (ret types1.Weight, isSome bool, err error) {
	key, err := MakeReservedDmpWeightOverrideStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for CustomValidationHeadData id={{false [14]}}
//
//	A custom head data that should be returned as result of `validate_block`.
//
//	See `Pallet::set_custom_validation_head_data` for more information.
func MakeCustomValidationHeadDataStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "ParachainSystem", "CustomValidationHeadData")
}
func GetCustomValidationHeadData(state state.State, bhash types.Hash) (ret []byte, isSome bool, err error) {
	key, err := MakeCustomValidationHeadDataStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetCustomValidationHeadDataLatest(state state.State) (ret []byte, isSome bool, err error) {
	key, err := MakeCustomValidationHeadDataStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
