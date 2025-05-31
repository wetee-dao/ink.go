package assets

import (
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Issue a new class of fungible assets from a public origin.
//
// This new asset class has no assets initially and its owner is the origin.
//
// The origin must conform to the configured `CreateOrigin` and have sufficient funds free.
//
// Funds of sender are reserved by `AssetDeposit`.
//
// Parameters:
// - `id`: The identifier of the new asset. This must not be currently in use to identify
// an existing asset. If [`NextAssetId`] is set, then this must be equal to it.
// - `admin`: The admin of this class of assets. The admin is the initial address of each
// member of the asset class's admin team.
// - `min_balance`: The minimum balance of this new asset that any single account must
// have. If an account's balance is reduced below this, then it collapses to zero.
//
// Emits `Created` event when successful.
//
// Weight: `O(1)`
func MakeCreateCall(id0 types.UCompact, admin1 types1.MultiAddress, minBalance2 types.U128) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsCreate:            true,
			AsCreateId0:         id0,
			AsCreateAdmin1:      admin1,
			AsCreateMinBalance2: minBalance2,
		},
	}
}

// Issue a new class of fungible assets from a privileged origin.
//
// This new asset class has no assets initially.
//
// The origin must conform to `ForceOrigin`.
//
// Unlike `create`, no funds are reserved.
//
// - `id`: The identifier of the new asset. This must not be currently in use to identify
// an existing asset. If [`NextAssetId`] is set, then this must be equal to it.
// - `owner`: The owner of this class of assets. The owner has full superuser permissions
// over this asset, but may later change and configure the permissions using
// `transfer_ownership` and `set_team`.
// - `min_balance`: The minimum balance of this new asset that any single account must
// have. If an account's balance is reduced below this, then it collapses to zero.
//
// Emits `ForceCreated` event when successful.
//
// Weight: `O(1)`
func MakeForceCreateCall(id0 types.UCompact, owner1 types1.MultiAddress, isSufficient2 bool, minBalance3 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsForceCreate:              true,
			AsForceCreateId0:           id0,
			AsForceCreateOwner1:        owner1,
			AsForceCreateIsSufficient2: isSufficient2,
			AsForceCreateMinBalance3:   minBalance3,
		},
	}
}

// Start the process of destroying a fungible asset class.
//
// `start_destroy` is the first in a series of extrinsics that should be called, to allow
// destruction of an asset class.
//
// The origin must conform to `ForceOrigin` or must be `Signed` by the asset's `owner`.
//
//   - `id`: The identifier of the asset to be destroyed. This must identify an existing
//     asset.
//
// It will fail with either [`Error::ContainsHolds`] or [`Error::ContainsFreezes`] if
// an account contains holds or freezes in place.
func MakeStartDestroyCall(id0 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsStartDestroy:    true,
			AsStartDestroyId0: id0,
		},
	}
}

// Destroy all accounts associated with a given asset.
//
// `destroy_accounts` should only be called after `start_destroy` has been called, and the
// asset is in a `Destroying` state.
//
// Due to weight restrictions, this function may need to be called multiple times to fully
// destroy all accounts. It will destroy `RemoveItemsLimit` accounts at a time.
//
//   - `id`: The identifier of the asset to be destroyed. This must identify an existing
//     asset.
//
// Each call emits the `Event::DestroyedAccounts` event.
func MakeDestroyAccountsCall(id0 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsDestroyAccounts:    true,
			AsDestroyAccountsId0: id0,
		},
	}
}

// Destroy all approvals associated with a given asset up to the max (T::RemoveItemsLimit).
//
// `destroy_approvals` should only be called after `start_destroy` has been called, and the
// asset is in a `Destroying` state.
//
// Due to weight restrictions, this function may need to be called multiple times to fully
// destroy all approvals. It will destroy `RemoveItemsLimit` approvals at a time.
//
//   - `id`: The identifier of the asset to be destroyed. This must identify an existing
//     asset.
//
// Each call emits the `Event::DestroyedApprovals` event.
func MakeDestroyApprovalsCall(id0 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsDestroyApprovals:    true,
			AsDestroyApprovalsId0: id0,
		},
	}
}

// Complete destroying asset and unreserve currency.
//
// `finish_destroy` should only be called after `start_destroy` has been called, and the
// asset is in a `Destroying` state. All accounts or approvals should be destroyed before
// hand.
//
//   - `id`: The identifier of the asset to be destroyed. This must identify an existing
//     asset.
//
// Each successful call emits the `Event::Destroyed` event.
func MakeFinishDestroyCall(id0 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsFinishDestroy:    true,
			AsFinishDestroyId0: id0,
		},
	}
}

// Mint assets of a particular class.
//
// The origin must be Signed and the sender must be the Issuer of the asset `id`.
//
// - `id`: The identifier of the asset to have some amount minted.
// - `beneficiary`: The account to be credited with the minted assets.
// - `amount`: The amount of the asset to be minted.
//
// Emits `Issued` event when successful.
//
// Weight: `O(1)`
// Modes: Pre-existing balance of `beneficiary`; Account pre-existence of `beneficiary`.
func MakeMintCall(id0 types.UCompact, beneficiary1 types1.MultiAddress, amount2 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsMint:             true,
			AsMintId0:          id0,
			AsMintBeneficiary1: beneficiary1,
			AsMintAmount2:      amount2,
		},
	}
}

// Reduce the balance of `who` by as much as possible up to `amount` assets of `id`.
//
// Origin must be Signed and the sender should be the Manager of the asset `id`.
//
// Bails with `NoAccount` if the `who` is already dead.
//
// - `id`: The identifier of the asset to have some amount burned.
// - `who`: The account to be debited from.
// - `amount`: The maximum amount by which `who`'s balance should be reduced.
//
// Emits `Burned` with the actual amount burned. If this takes the balance to below the
// minimum for the asset, then the amount burned is increased to take it to zero.
//
// Weight: `O(1)`
// Modes: Post-existence of `who`; Pre & post Zombie-status of `who`.
func MakeBurnCall(id0 types.UCompact, who1 types1.MultiAddress, amount2 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsBurn:        true,
			AsBurnId0:     id0,
			AsBurnWho1:    who1,
			AsBurnAmount2: amount2,
		},
	}
}

// Move some assets from the sender account to another.
//
// Origin must be Signed.
//
// - `id`: The identifier of the asset to have some amount transferred.
// - `target`: The account to be credited.
// - `amount`: The amount by which the sender's balance of assets should be reduced and
// `target`'s balance increased. The amount actually transferred may be slightly greater in
// the case that the transfer would otherwise take the sender balance above zero but below
// the minimum balance. Must be greater than zero.
//
// Emits `Transferred` with the actual amount transferred. If this takes the source balance
// to below the minimum for the asset, then the amount transferred is increased to take it
// to zero.
//
// Weight: `O(1)`
// Modes: Pre-existence of `target`; Post-existence of sender; Account pre-existence of
// `target`.
func MakeTransferCall(id0 types.UCompact, target1 types1.MultiAddress, amount2 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsTransfer:        true,
			AsTransferId0:     id0,
			AsTransferTarget1: target1,
			AsTransferAmount2: amount2,
		},
	}
}

// Move some assets from the sender account to another, keeping the sender account alive.
//
// Origin must be Signed.
//
// - `id`: The identifier of the asset to have some amount transferred.
// - `target`: The account to be credited.
// - `amount`: The amount by which the sender's balance of assets should be reduced and
// `target`'s balance increased. The amount actually transferred may be slightly greater in
// the case that the transfer would otherwise take the sender balance above zero but below
// the minimum balance. Must be greater than zero.
//
// Emits `Transferred` with the actual amount transferred. If this takes the source balance
// to below the minimum for the asset, then the amount transferred is increased to take it
// to zero.
//
// Weight: `O(1)`
// Modes: Pre-existence of `target`; Post-existence of sender; Account pre-existence of
// `target`.
func MakeTransferKeepAliveCall(id0 types.UCompact, target1 types1.MultiAddress, amount2 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsTransferKeepAlive:        true,
			AsTransferKeepAliveId0:     id0,
			AsTransferKeepAliveTarget1: target1,
			AsTransferKeepAliveAmount2: amount2,
		},
	}
}

// Move some assets from one account to another.
//
// Origin must be Signed and the sender should be the Admin of the asset `id`.
//
// - `id`: The identifier of the asset to have some amount transferred.
// - `source`: The account to be debited.
// - `dest`: The account to be credited.
// - `amount`: The amount by which the `source`'s balance of assets should be reduced and
// `dest`'s balance increased. The amount actually transferred may be slightly greater in
// the case that the transfer would otherwise take the `source` balance above zero but
// below the minimum balance. Must be greater than zero.
//
// Emits `Transferred` with the actual amount transferred. If this takes the source balance
// to below the minimum for the asset, then the amount transferred is increased to take it
// to zero.
//
// Weight: `O(1)`
// Modes: Pre-existence of `dest`; Post-existence of `source`; Account pre-existence of
// `dest`.
func MakeForceTransferCall(id0 types.UCompact, source1 types1.MultiAddress, dest2 types1.MultiAddress, amount3 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsForceTransfer:        true,
			AsForceTransferId0:     id0,
			AsForceTransferSource1: source1,
			AsForceTransferDest2:   dest2,
			AsForceTransferAmount3: amount3,
		},
	}
}

// Disallow further unprivileged transfers of an asset `id` from an account `who`. `who`
// must already exist as an entry in `Account`s of the asset. If you want to freeze an
// account that does not have an entry, use `touch_other` first.
//
// Origin must be Signed and the sender should be the Freezer of the asset `id`.
//
// - `id`: The identifier of the asset to be frozen.
// - `who`: The account to be frozen.
//
// Emits `Frozen`.
//
// Weight: `O(1)`
func MakeFreezeCall(id0 types.UCompact, who1 types1.MultiAddress) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsFreeze:     true,
			AsFreezeId0:  id0,
			AsFreezeWho1: who1,
		},
	}
}

// Allow unprivileged transfers to and from an account again.
//
// Origin must be Signed and the sender should be the Admin of the asset `id`.
//
// - `id`: The identifier of the asset to be frozen.
// - `who`: The account to be unfrozen.
//
// Emits `Thawed`.
//
// Weight: `O(1)`
func MakeThawCall(id0 types.UCompact, who1 types1.MultiAddress) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsThaw:     true,
			AsThawId0:  id0,
			AsThawWho1: who1,
		},
	}
}

// Disallow further unprivileged transfers for the asset class.
//
// Origin must be Signed and the sender should be the Freezer of the asset `id`.
//
// - `id`: The identifier of the asset to be frozen.
//
// Emits `Frozen`.
//
// Weight: `O(1)`
func MakeFreezeAssetCall(id0 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsFreezeAsset:    true,
			AsFreezeAssetId0: id0,
		},
	}
}

// Allow unprivileged transfers for the asset again.
//
// Origin must be Signed and the sender should be the Admin of the asset `id`.
//
// - `id`: The identifier of the asset to be thawed.
//
// Emits `Thawed`.
//
// Weight: `O(1)`
func MakeThawAssetCall(id0 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsThawAsset:    true,
			AsThawAssetId0: id0,
		},
	}
}

// Change the Owner of an asset.
//
// Origin must be Signed and the sender should be the Owner of the asset `id`.
//
// - `id`: The identifier of the asset.
// - `owner`: The new Owner of this asset.
//
// Emits `OwnerChanged`.
//
// Weight: `O(1)`
func MakeTransferOwnershipCall(id0 types.UCompact, owner1 types1.MultiAddress) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsTransferOwnership:       true,
			AsTransferOwnershipId0:    id0,
			AsTransferOwnershipOwner1: owner1,
		},
	}
}

// Change the Issuer, Admin and Freezer of an asset.
//
// Origin must be Signed and the sender should be the Owner of the asset `id`.
//
// - `id`: The identifier of the asset to be frozen.
// - `issuer`: The new Issuer of this asset.
// - `admin`: The new Admin of this asset.
// - `freezer`: The new Freezer of this asset.
//
// Emits `TeamChanged`.
//
// Weight: `O(1)`
func MakeSetTeamCall(id0 types.UCompact, issuer1 types1.MultiAddress, admin2 types1.MultiAddress, freezer3 types1.MultiAddress) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsSetTeam:         true,
			AsSetTeamId0:      id0,
			AsSetTeamIssuer1:  issuer1,
			AsSetTeamAdmin2:   admin2,
			AsSetTeamFreezer3: freezer3,
		},
	}
}

// Set the metadata for an asset.
//
// Origin must be Signed and the sender should be the Owner of the asset `id`.
//
// Funds of sender are reserved according to the formula:
// `MetadataDepositBase + MetadataDepositPerByte * (name.len + symbol.len)` taking into
// account any already reserved funds.
//
// - `id`: The identifier of the asset to update.
// - `name`: The user friendly name of this asset. Limited in length by `StringLimit`.
// - `symbol`: The exchange symbol for this asset. Limited in length by `StringLimit`.
// - `decimals`: The number of decimals this asset uses to represent one unit.
//
// Emits `MetadataSet`.
//
// Weight: `O(1)`
func MakeSetMetadataCall(id0 types.UCompact, name1 []byte, symbol2 []byte, decimals3 byte) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsSetMetadata:          true,
			AsSetMetadataId0:       id0,
			AsSetMetadataName1:     name1,
			AsSetMetadataSymbol2:   symbol2,
			AsSetMetadataDecimals3: decimals3,
		},
	}
}

// Clear the metadata for an asset.
//
// Origin must be Signed and the sender should be the Owner of the asset `id`.
//
// Any deposit is freed for the asset owner.
//
// - `id`: The identifier of the asset to clear.
//
// Emits `MetadataCleared`.
//
// Weight: `O(1)`
func MakeClearMetadataCall(id0 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsClearMetadata:    true,
			AsClearMetadataId0: id0,
		},
	}
}

// Force the metadata for an asset to some value.
//
// Origin must be ForceOrigin.
//
// Any deposit is left alone.
//
// - `id`: The identifier of the asset to update.
// - `name`: The user friendly name of this asset. Limited in length by `StringLimit`.
// - `symbol`: The exchange symbol for this asset. Limited in length by `StringLimit`.
// - `decimals`: The number of decimals this asset uses to represent one unit.
//
// Emits `MetadataSet`.
//
// Weight: `O(N + S)` where N and S are the length of the name and symbol respectively.
func MakeForceSetMetadataCall(id0 types.UCompact, name1 []byte, symbol2 []byte, decimals3 byte, isFrozen4 bool) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsForceSetMetadata:          true,
			AsForceSetMetadataId0:       id0,
			AsForceSetMetadataName1:     name1,
			AsForceSetMetadataSymbol2:   symbol2,
			AsForceSetMetadataDecimals3: decimals3,
			AsForceSetMetadataIsFrozen4: isFrozen4,
		},
	}
}

// Clear the metadata for an asset.
//
// Origin must be ForceOrigin.
//
// Any deposit is returned.
//
// - `id`: The identifier of the asset to clear.
//
// Emits `MetadataCleared`.
//
// Weight: `O(1)`
func MakeForceClearMetadataCall(id0 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsForceClearMetadata:    true,
			AsForceClearMetadataId0: id0,
		},
	}
}

// Alter the attributes of a given asset.
//
// Origin must be `ForceOrigin`.
//
// - `id`: The identifier of the asset.
// - `owner`: The new Owner of this asset.
// - `issuer`: The new Issuer of this asset.
// - `admin`: The new Admin of this asset.
// - `freezer`: The new Freezer of this asset.
// - `min_balance`: The minimum balance of this new asset that any single account must
// have. If an account's balance is reduced below this, then it collapses to zero.
// - `is_sufficient`: Whether a non-zero balance of this asset is deposit of sufficient
// value to account for the state bloat associated with its balance storage. If set to
// `true`, then non-zero balances may be stored without a `consumer` reference (and thus
// an ED in the Balances pallet or whatever else is used to control user-account state
// growth).
// - `is_frozen`: Whether this asset class is frozen except for permissioned/admin
// instructions.
//
// Emits `AssetStatusChanged` with the identity of the asset.
//
// Weight: `O(1)`
func MakeForceAssetStatusCall(id0 types.UCompact, owner1 types1.MultiAddress, issuer2 types1.MultiAddress, admin3 types1.MultiAddress, freezer4 types1.MultiAddress, minBalance5 types.UCompact, isSufficient6 bool, isFrozen7 bool) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsForceAssetStatus:              true,
			AsForceAssetStatusId0:           id0,
			AsForceAssetStatusOwner1:        owner1,
			AsForceAssetStatusIssuer2:       issuer2,
			AsForceAssetStatusAdmin3:        admin3,
			AsForceAssetStatusFreezer4:      freezer4,
			AsForceAssetStatusMinBalance5:   minBalance5,
			AsForceAssetStatusIsSufficient6: isSufficient6,
			AsForceAssetStatusIsFrozen7:     isFrozen7,
		},
	}
}

// Approve an amount of asset for transfer by a delegated third-party account.
//
// Origin must be Signed.
//
// Ensures that `ApprovalDeposit` worth of `Currency` is reserved from signing account
// for the purpose of holding the approval. If some non-zero amount of assets is already
// approved from signing account to `delegate`, then it is topped up or unreserved to
// meet the right value.
//
// NOTE: The signing account does not need to own `amount` of assets at the point of
// making this call.
//
// - `id`: The identifier of the asset.
// - `delegate`: The account to delegate permission to transfer asset.
// - `amount`: The amount of asset that may be transferred by `delegate`. If there is
// already an approval in place, then this acts additively.
//
// Emits `ApprovedTransfer` on success.
//
// Weight: `O(1)`
func MakeApproveTransferCall(id0 types.UCompact, delegate1 types1.MultiAddress, amount2 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsApproveTransfer:          true,
			AsApproveTransferId0:       id0,
			AsApproveTransferDelegate1: delegate1,
			AsApproveTransferAmount2:   amount2,
		},
	}
}

// Cancel all of some asset approved for delegated transfer by a third-party account.
//
// Origin must be Signed and there must be an approval in place between signer and
// `delegate`.
//
// Unreserves any deposit previously reserved by `approve_transfer` for the approval.
//
// - `id`: The identifier of the asset.
// - `delegate`: The account delegated permission to transfer asset.
//
// Emits `ApprovalCancelled` on success.
//
// Weight: `O(1)`
func MakeCancelApprovalCall(id0 types.UCompact, delegate1 types1.MultiAddress) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsCancelApproval:          true,
			AsCancelApprovalId0:       id0,
			AsCancelApprovalDelegate1: delegate1,
		},
	}
}

// Cancel all of some asset approved for delegated transfer by a third-party account.
//
// Origin must be either ForceOrigin or Signed origin with the signer being the Admin
// account of the asset `id`.
//
// Unreserves any deposit previously reserved by `approve_transfer` for the approval.
//
// - `id`: The identifier of the asset.
// - `delegate`: The account delegated permission to transfer asset.
//
// Emits `ApprovalCancelled` on success.
//
// Weight: `O(1)`
func MakeForceCancelApprovalCall(id0 types.UCompact, owner1 types1.MultiAddress, delegate2 types1.MultiAddress) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsForceCancelApproval:          true,
			AsForceCancelApprovalId0:       id0,
			AsForceCancelApprovalOwner1:    owner1,
			AsForceCancelApprovalDelegate2: delegate2,
		},
	}
}

// Transfer some asset balance from a previously delegated account to some third-party
// account.
//
// Origin must be Signed and there must be an approval in place by the `owner` to the
// signer.
//
// If the entire amount approved for transfer is transferred, then any deposit previously
// reserved by `approve_transfer` is unreserved.
//
// - `id`: The identifier of the asset.
// - `owner`: The account which previously approved for a transfer of at least `amount` and
// from which the asset balance will be withdrawn.
// - `destination`: The account to which the asset balance of `amount` will be transferred.
// - `amount`: The amount of assets to transfer.
//
// Emits `TransferredApproved` on success.
//
// Weight: `O(1)`
func MakeTransferApprovedCall(id0 types.UCompact, owner1 types1.MultiAddress, destination2 types1.MultiAddress, amount3 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsTransferApproved:             true,
			AsTransferApprovedId0:          id0,
			AsTransferApprovedOwner1:       owner1,
			AsTransferApprovedDestination2: destination2,
			AsTransferApprovedAmount3:      amount3,
		},
	}
}

// Create an asset account for non-provider assets.
//
// A deposit will be taken from the signer account.
//
//   - `origin`: Must be Signed; the signer account must have sufficient funds for a deposit
//     to be taken.
//   - `id`: The identifier of the asset for the account to be created.
//
// Emits `Touched` event when successful.
func MakeTouchCall(id0 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsTouch:    true,
			AsTouchId0: id0,
		},
	}
}

// Return the deposit (if any) of an asset account or a consumer reference (if any) of an
// account.
//
// The origin must be Signed.
//
//   - `id`: The identifier of the asset for which the caller would like the deposit
//     refunded.
//   - `allow_burn`: If `true` then assets may be destroyed in order to complete the refund.
//
// It will fail with either [`Error::ContainsHolds`] or [`Error::ContainsFreezes`] if
// the asset account contains holds or freezes in place.
//
// Emits `Refunded` event when successful.
func MakeRefundCall(id0 types.UCompact, allowBurn1 bool) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsRefund:           true,
			AsRefundId0:        id0,
			AsRefundAllowBurn1: allowBurn1,
		},
	}
}

// Sets the minimum balance of an asset.
//
// Only works if there aren't any accounts that are holding the asset or if
// the new value of `min_balance` is less than the old one.
//
// Origin must be Signed and the sender has to be the Owner of the
// asset `id`.
//
// - `id`: The identifier of the asset.
// - `min_balance`: The new value of `min_balance`.
//
// Emits `AssetMinBalanceChanged` event when successful.
func MakeSetMinBalanceCall(id0 types.UCompact, minBalance1 types.U128) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsSetMinBalance:            true,
			AsSetMinBalanceId0:         id0,
			AsSetMinBalanceMinBalance1: minBalance1,
		},
	}
}

// Create an asset account for `who`.
//
// A deposit will be taken from the signer account.
//
//   - `origin`: Must be Signed by `Freezer` or `Admin` of the asset `id`; the signer account
//     must have sufficient funds for a deposit to be taken.
//   - `id`: The identifier of the asset for the account to be created.
//   - `who`: The account to be created.
//
// Emits `Touched` event when successful.
func MakeTouchOtherCall(id0 types.UCompact, who1 types1.MultiAddress) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsTouchOther:     true,
			AsTouchOtherId0:  id0,
			AsTouchOtherWho1: who1,
		},
	}
}

// Return the deposit (if any) of a target asset account. Useful if you are the depositor.
//
// The origin must be Signed and either the account owner, depositor, or asset `Admin`. In
// order to burn a non-zero balance of the asset, the caller must be the account and should
// use `refund`.
//
// - `id`: The identifier of the asset for the account holding a deposit.
// - `who`: The account to refund.
//
// It will fail with either [`Error::ContainsHolds`] or [`Error::ContainsFreezes`] if
// the asset account contains holds or freezes in place.
//
// Emits `Refunded` event when successful.
func MakeRefundOtherCall(id0 types.UCompact, who1 types1.MultiAddress) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsRefundOther:     true,
			AsRefundOtherId0:  id0,
			AsRefundOtherWho1: who1,
		},
	}
}

// Disallow further unprivileged transfers of an asset `id` to and from an account `who`.
//
// Origin must be Signed and the sender should be the Freezer of the asset `id`.
//
// - `id`: The identifier of the account's asset.
// - `who`: The account to be unblocked.
//
// Emits `Blocked`.
//
// Weight: `O(1)`
func MakeBlockCall(id0 types.UCompact, who1 types1.MultiAddress) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsBlock:     true,
			AsBlockId0:  id0,
			AsBlockWho1: who1,
		},
	}
}

// Transfer the entire transferable balance from the caller asset account.
//
// NOTE: This function only attempts to transfer _transferable_ balances. This means that
// any held, frozen, or minimum balance (when `keep_alive` is `true`), will not be
// transferred by this function. To ensure that this function results in a killed account,
// you might need to prepare the account by removing any reference counters, storage
// deposits, etc...
//
// The dispatch origin of this call must be Signed.
//
//   - `id`: The identifier of the asset for the account holding a deposit.
//   - `dest`: The recipient of the transfer.
//   - `keep_alive`: A boolean to determine if the `transfer_all` operation should send all
//     of the funds the asset account has, causing the sender asset account to be killed
//     (false), or transfer everything except at least the minimum balance, which will
//     guarantee to keep the sender asset account alive (true).
func MakeTransferAllCall(id0 types.UCompact, dest1 types1.MultiAddress, keepAlive2 bool) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsAssets: true,
		AsAssetsField0: &types1.PalletAssetsPalletCall{
			IsTransferAll:           true,
			AsTransferAllId0:        id0,
			AsTransferAllDest1:      dest1,
			AsTransferAllKeepAlive2: keepAlive2,
		},
	}
}
