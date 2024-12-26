package polkadotxcm

import types "github.com/wetee-dao/go-sdk/pallet/types"

func MakeSendCall(dest0 types.VersionedLocation, message1 types.VersionedXcm) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsSend:         true,
			AsSendDest0:    &dest0,
			AsSendMessage1: &message1,
		},
	}
}

// Teleport some assets from the local chain to some destination chain.
//
// **This function is deprecated: Use `limited_teleport_assets` instead.**
//
// Fee payment on the destination side is made from the asset in the `assets` vector of
// index `fee_asset_item`. The weight limit for fees is not provided and thus is unlimited,
// with all fees taken as needed from the asset.
//
//   - `origin`: Must be capable of withdrawing the `assets` and executing XCM.
//   - `dest`: Destination context for the assets. Will typically be `[Parent,
//     Parachain(..)]` to send from parachain to parachain, or `[Parachain(..)]` to send from
//     relay to parachain.
//   - `beneficiary`: A beneficiary location for the assets in the context of `dest`. Will
//     generally be an `AccountId32` value.
//   - `assets`: The assets to be withdrawn. This should include the assets used to pay the
//     fee on the `dest` chain.
//   - `fee_asset_item`: The index into `assets` of the item which should be used to pay
//     fees.
func MakeTeleportAssetsCall(dest0 types.VersionedLocation, beneficiary1 types.VersionedLocation, assets2 types.VersionedAssets, feeAssetItem3 uint32) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsTeleportAssets:              true,
			AsTeleportAssetsDest0:         &dest0,
			AsTeleportAssetsBeneficiary1:  &beneficiary1,
			AsTeleportAssetsAssets2:       &assets2,
			AsTeleportAssetsFeeAssetItem3: feeAssetItem3,
		},
	}
}

// Transfer some assets from the local chain to the destination chain through their local,
// destination or remote reserve.
//
// `assets` must have same reserve location and may not be teleportable to `dest`.
//   - `assets` have local reserve: transfer assets to sovereign account of destination
//     chain and forward a notification XCM to `dest` to mint and deposit reserve-based
//     assets to `beneficiary`.
//   - `assets` have destination reserve: burn local assets and forward a notification to
//     `dest` chain to withdraw the reserve assets from this chain's sovereign account and
//     deposit them to `beneficiary`.
//   - `assets` have remote reserve: burn local assets, forward XCM to reserve chain to move
//     reserves from this chain's SA to `dest` chain's SA, and forward another XCM to `dest`
//     to mint and deposit reserve-based assets to `beneficiary`.
//
// **This function is deprecated: Use `limited_reserve_transfer_assets` instead.**
//
// Fee payment on the destination side is made from the asset in the `assets` vector of
// index `fee_asset_item`. The weight limit for fees is not provided and thus is unlimited,
// with all fees taken as needed from the asset.
//
//   - `origin`: Must be capable of withdrawing the `assets` and executing XCM.
//   - `dest`: Destination context for the assets. Will typically be `[Parent,
//     Parachain(..)]` to send from parachain to parachain, or `[Parachain(..)]` to send from
//     relay to parachain.
//   - `beneficiary`: A beneficiary location for the assets in the context of `dest`. Will
//     generally be an `AccountId32` value.
//   - `assets`: The assets to be withdrawn. This should include the assets used to pay the
//     fee on the `dest` (and possibly reserve) chains.
//   - `fee_asset_item`: The index into `assets` of the item which should be used to pay
//     fees.
func MakeReserveTransferAssetsCall(dest0 types.VersionedLocation, beneficiary1 types.VersionedLocation, assets2 types.VersionedAssets, feeAssetItem3 uint32) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsReserveTransferAssets:              true,
			AsReserveTransferAssetsDest0:         &dest0,
			AsReserveTransferAssetsBeneficiary1:  &beneficiary1,
			AsReserveTransferAssetsAssets2:       &assets2,
			AsReserveTransferAssetsFeeAssetItem3: feeAssetItem3,
		},
	}
}

// Execute an XCM message from a local, signed, origin.
//
// An event is deposited indicating whether `msg` could be executed completely or only
// partially.
//
// No more than `max_weight` will be used in its attempted execution. If this is less than
// the maximum amount of weight that the message could take to be executed, then no
// execution attempt will be made.
func MakeExecuteCall(message0 types.VersionedXcm1, maxWeight1 types.Weight) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsExecute:           true,
			AsExecuteMessage0:   &message0,
			AsExecuteMaxWeight1: maxWeight1,
		},
	}
}

// Extoll that a particular destination can be communicated with through a particular
// version of XCM.
//
// - `origin`: Must be an origin specified by AdminOrigin.
// - `location`: The destination that is being described.
// - `xcm_version`: The latest version of XCM that `location` supports.
func MakeForceXcmVersionCall(location0 types.Location, version1 uint32) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsForceXcmVersion:          true,
			AsForceXcmVersionLocation0: &location0,
			AsForceXcmVersionVersion1:  version1,
		},
	}
}

// Set a safe XCM version (the version that XCM should be encoded with if the most recent
// version a destination can accept is unknown).
//
// - `origin`: Must be an origin specified by AdminOrigin.
// - `maybe_xcm_version`: The default XCM encoding version, or `None` to disable.
func MakeForceDefaultXcmVersionCall(maybeXcmVersion0 types.OptionTUint32) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsForceDefaultXcmVersion:                 true,
			AsForceDefaultXcmVersionMaybeXcmVersion0: &maybeXcmVersion0,
		},
	}
}

// Ask a location to notify us regarding their XCM version and any changes to it.
//
// - `origin`: Must be an origin specified by AdminOrigin.
// - `location`: The location to which we should subscribe for XCM version notifications.
func MakeForceSubscribeVersionNotifyCall(location0 types.VersionedLocation) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsForceSubscribeVersionNotify:          true,
			AsForceSubscribeVersionNotifyLocation0: &location0,
		},
	}
}

// Require that a particular destination should no longer notify us regarding any XCM
// version changes.
//
//   - `origin`: Must be an origin specified by AdminOrigin.
//   - `location`: The location to which we are currently subscribed for XCM version
//     notifications which we no longer desire.
func MakeForceUnsubscribeVersionNotifyCall(location0 types.VersionedLocation) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsForceUnsubscribeVersionNotify:          true,
			AsForceUnsubscribeVersionNotifyLocation0: &location0,
		},
	}
}

// Transfer some assets from the local chain to the destination chain through their local,
// destination or remote reserve.
//
// `assets` must have same reserve location and may not be teleportable to `dest`.
//   - `assets` have local reserve: transfer assets to sovereign account of destination
//     chain and forward a notification XCM to `dest` to mint and deposit reserve-based
//     assets to `beneficiary`.
//   - `assets` have destination reserve: burn local assets and forward a notification to
//     `dest` chain to withdraw the reserve assets from this chain's sovereign account and
//     deposit them to `beneficiary`.
//   - `assets` have remote reserve: burn local assets, forward XCM to reserve chain to move
//     reserves from this chain's SA to `dest` chain's SA, and forward another XCM to `dest`
//     to mint and deposit reserve-based assets to `beneficiary`.
//
// Fee payment on the destination side is made from the asset in the `assets` vector of
// index `fee_asset_item`, up to enough to pay for `weight_limit` of weight. If more weight
// is needed than `weight_limit`, then the operation will fail and the sent assets may be
// at risk.
//
//   - `origin`: Must be capable of withdrawing the `assets` and executing XCM.
//   - `dest`: Destination context for the assets. Will typically be `[Parent,
//     Parachain(..)]` to send from parachain to parachain, or `[Parachain(..)]` to send from
//     relay to parachain.
//   - `beneficiary`: A beneficiary location for the assets in the context of `dest`. Will
//     generally be an `AccountId32` value.
//   - `assets`: The assets to be withdrawn. This should include the assets used to pay the
//     fee on the `dest` (and possibly reserve) chains.
//   - `fee_asset_item`: The index into `assets` of the item which should be used to pay
//     fees.
//   - `weight_limit`: The remote-side weight limit, if any, for the XCM fee purchase.
func MakeLimitedReserveTransferAssetsCall(dest0 types.VersionedLocation, beneficiary1 types.VersionedLocation, assets2 types.VersionedAssets, feeAssetItem3 uint32, weightLimit4 types.WeightLimit) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsLimitedReserveTransferAssets:              true,
			AsLimitedReserveTransferAssetsDest0:         &dest0,
			AsLimitedReserveTransferAssetsBeneficiary1:  &beneficiary1,
			AsLimitedReserveTransferAssetsAssets2:       &assets2,
			AsLimitedReserveTransferAssetsFeeAssetItem3: feeAssetItem3,
			AsLimitedReserveTransferAssetsWeightLimit4:  weightLimit4,
		},
	}
}

// Teleport some assets from the local chain to some destination chain.
//
// Fee payment on the destination side is made from the asset in the `assets` vector of
// index `fee_asset_item`, up to enough to pay for `weight_limit` of weight. If more weight
// is needed than `weight_limit`, then the operation will fail and the sent assets may be
// at risk.
//
//   - `origin`: Must be capable of withdrawing the `assets` and executing XCM.
//   - `dest`: Destination context for the assets. Will typically be `[Parent,
//     Parachain(..)]` to send from parachain to parachain, or `[Parachain(..)]` to send from
//     relay to parachain.
//   - `beneficiary`: A beneficiary location for the assets in the context of `dest`. Will
//     generally be an `AccountId32` value.
//   - `assets`: The assets to be withdrawn. This should include the assets used to pay the
//     fee on the `dest` chain.
//   - `fee_asset_item`: The index into `assets` of the item which should be used to pay
//     fees.
//   - `weight_limit`: The remote-side weight limit, if any, for the XCM fee purchase.
func MakeLimitedTeleportAssetsCall(dest0 types.VersionedLocation, beneficiary1 types.VersionedLocation, assets2 types.VersionedAssets, feeAssetItem3 uint32, weightLimit4 types.WeightLimit) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsLimitedTeleportAssets:              true,
			AsLimitedTeleportAssetsDest0:         &dest0,
			AsLimitedTeleportAssetsBeneficiary1:  &beneficiary1,
			AsLimitedTeleportAssetsAssets2:       &assets2,
			AsLimitedTeleportAssetsFeeAssetItem3: feeAssetItem3,
			AsLimitedTeleportAssetsWeightLimit4:  weightLimit4,
		},
	}
}

// Set or unset the global suspension state of the XCM executor.
//
// - `origin`: Must be an origin specified by AdminOrigin.
// - `suspended`: `true` to suspend, `false` to resume.
func MakeForceSuspensionCall(suspended0 bool) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsForceSuspension:           true,
			AsForceSuspensionSuspended0: suspended0,
		},
	}
}

// Transfer some assets from the local chain to the destination chain through their local,
// destination or remote reserve, or through teleports.
//
// Fee payment on the destination side is made from the asset in the `assets` vector of
// index `fee_asset_item` (hence referred to as `fees`), up to enough to pay for
// `weight_limit` of weight. If more weight is needed than `weight_limit`, then the
// operation will fail and the sent assets may be at risk.
//
// `assets` (excluding `fees`) must have same reserve location or otherwise be teleportable
// to `dest`, no limitations imposed on `fees`.
//   - for local reserve: transfer assets to sovereign account of destination chain and
//     forward a notification XCM to `dest` to mint and deposit reserve-based assets to
//     `beneficiary`.
//   - for destination reserve: burn local assets and forward a notification to `dest` chain
//     to withdraw the reserve assets from this chain's sovereign account and deposit them
//     to `beneficiary`.
//   - for remote reserve: burn local assets, forward XCM to reserve chain to move reserves
//     from this chain's SA to `dest` chain's SA, and forward another XCM to `dest` to mint
//     and deposit reserve-based assets to `beneficiary`.
//   - for teleports: burn local assets and forward XCM to `dest` chain to mint/teleport
//     assets and deposit them to `beneficiary`.
//
//   - `origin`: Must be capable of withdrawing the `assets` and executing XCM.
//   - `dest`: Destination context for the assets. Will typically be `X2(Parent,
//     Parachain(..))` to send from parachain to parachain, or `X1(Parachain(..))` to send
//     from relay to parachain.
//   - `beneficiary`: A beneficiary location for the assets in the context of `dest`. Will
//     generally be an `AccountId32` value.
//   - `assets`: The assets to be withdrawn. This should include the assets used to pay the
//     fee on the `dest` (and possibly reserve) chains.
//   - `fee_asset_item`: The index into `assets` of the item which should be used to pay
//     fees.
//   - `weight_limit`: The remote-side weight limit, if any, for the XCM fee purchase.
func MakeTransferAssetsCall(dest0 types.VersionedLocation, beneficiary1 types.VersionedLocation, assets2 types.VersionedAssets, feeAssetItem3 uint32, weightLimit4 types.WeightLimit) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsTransferAssets:              true,
			AsTransferAssetsDest0:         &dest0,
			AsTransferAssetsBeneficiary1:  &beneficiary1,
			AsTransferAssetsAssets2:       &assets2,
			AsTransferAssetsFeeAssetItem3: feeAssetItem3,
			AsTransferAssetsWeightLimit4:  weightLimit4,
		},
	}
}

// Claims assets trapped on this pallet because of leftover assets during XCM execution.
//
// - `origin`: Anyone can call this extrinsic.
// - `assets`: The exact assets that were trapped. Use the version to specify what version
// was the latest when they were trapped.
// - `beneficiary`: The location/account where the claimed assets will be deposited.
func MakeClaimAssetsCall(assets0 types.VersionedAssets, beneficiary1 types.VersionedLocation) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsClaimAssets:             true,
			AsClaimAssetsAssets0:      &assets0,
			AsClaimAssetsBeneficiary1: &beneficiary1,
		},
	}
}

// Transfer assets from the local chain to the destination chain using explicit transfer
// types for assets and fees.
//
// `assets` must have same reserve location or may be teleportable to `dest`. Caller must
// provide the `assets_transfer_type` to be used for `assets`:
//   - `TransferType::LocalReserve`: transfer assets to sovereign account of destination
//     chain and forward a notification XCM to `dest` to mint and deposit reserve-based
//     assets to `beneficiary`.
//   - `TransferType::DestinationReserve`: burn local assets and forward a notification to
//     `dest` chain to withdraw the reserve assets from this chain's sovereign account and
//     deposit them to `beneficiary`.
//   - `TransferType::RemoteReserve(reserve)`: burn local assets, forward XCM to `reserve`
//     chain to move reserves from this chain's SA to `dest` chain's SA, and forward another
//     XCM to `dest` to mint and deposit reserve-based assets to `beneficiary`. Typically
//     the remote `reserve` is Asset Hub.
//   - `TransferType::Teleport`: burn local assets and forward XCM to `dest` chain to
//     mint/teleport assets and deposit them to `beneficiary`.
//
// On the destination chain, as well as any intermediary hops, `BuyExecution` is used to
// buy execution using transferred `assets` identified by `remote_fees_id`.
// Make sure enough of the specified `remote_fees_id` asset is included in the given list
// of `assets`. `remote_fees_id` should be enough to pay for `weight_limit`. If more weight
// is needed than `weight_limit`, then the operation will fail and the sent assets may be
// at risk.
//
// `remote_fees_id` may use different transfer type than rest of `assets` and can be
// specified through `fees_transfer_type`.
//
// The caller needs to specify what should happen to the transferred assets once they reach
// the `dest` chain. This is done through the `custom_xcm_on_dest` parameter, which
// contains the instructions to execute on `dest` as a final step.
//
//	This is usually as simple as:
//	`Xcm(vec![DepositAsset { assets: Wild(AllCounted(assets.len())), beneficiary }])`,
//	but could be something more exotic like sending the `assets` even further.
//
//   - `origin`: Must be capable of withdrawing the `assets` and executing XCM.
//   - `dest`: Destination context for the assets. Will typically be `[Parent,
//     Parachain(..)]` to send from parachain to parachain, or `[Parachain(..)]` to send from
//     relay to parachain, or `(parents: 2, (GlobalConsensus(..), ..))` to send from
//     parachain across a bridge to another ecosystem destination.
//   - `assets`: The assets to be withdrawn. This should include the assets used to pay the
//     fee on the `dest` (and possibly reserve) chains.
//   - `assets_transfer_type`: The XCM `TransferType` used to transfer the `assets`.
//   - `remote_fees_id`: One of the included `assets` to be be used to pay fees.
//   - `fees_transfer_type`: The XCM `TransferType` used to transfer the `fees` assets.
//   - `custom_xcm_on_dest`: The XCM to be executed on `dest` chain as the last step of the
//     transfer, which also determines what happens to the assets on the destination chain.
//   - `weight_limit`: The remote-side weight limit, if any, for the XCM fee purchase.
func MakeTransferAssetsUsingTypeAndThenCall(dest0 types.VersionedLocation, assets1 types.VersionedAssets, assetsTransferType2 types.TransferType, remoteFeesId3 types.VersionedAssetId, feesTransferType4 types.TransferType, customXcmOnDest5 types.VersionedXcm, weightLimit6 types.WeightLimit) types.RuntimeCall {
	return types.RuntimeCall{
		IsPolkadotXcm: true,
		AsPolkadotXcmField0: &types.PalletXcmPalletCall{
			IsTransferAssetsUsingTypeAndThen:                    true,
			AsTransferAssetsUsingTypeAndThenDest0:               &dest0,
			AsTransferAssetsUsingTypeAndThenAssets1:             &assets1,
			AsTransferAssetsUsingTypeAndThenAssetsTransferType2: &assetsTransferType2,
			AsTransferAssetsUsingTypeAndThenRemoteFeesId3:       &remoteFeesId3,
			AsTransferAssetsUsingTypeAndThenFeesTransferType4:   &feesTransferType4,
			AsTransferAssetsUsingTypeAndThenCustomXcmOnDest5:    &customXcmOnDest5,
			AsTransferAssetsUsingTypeAndThenWeightLimit6:        weightLimit6,
		},
	}
}
