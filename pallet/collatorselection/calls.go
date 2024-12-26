package collatorselection

import (
	types1 "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types "github.com/wetee-dao/go-sdk/pallet/types"
)

// Set the list of invulnerable (fixed) collators. These collators must do some
// preparation, namely to have registered session keys.
//
// The call will remove any accounts that have not registered keys from the set. That is,
// it is non-atomic; the caller accepts all `AccountId`s passed in `new` _individually_ as
// acceptable Invulnerables, and is not proposing a _set_ of new Invulnerables.
//
// This call does not maintain mutual exclusivity of `Invulnerables` and `Candidates`. It
// is recommended to use a batch of `add_invulnerable` and `remove_invulnerable` instead. A
// `batch_all` can also be used to enforce atomicity. If any candidates are included in
// `new`, they should be removed with `remove_invulnerable_candidate` after execution.
//
// Must be called by the `UpdateOrigin`.
func MakeSetInvulnerablesCall(new0 [][32]byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsCollatorSelection: true,
		AsCollatorSelectionField0: &types.PalletCollatorSelectionPalletCall{
			IsSetInvulnerables:     true,
			AsSetInvulnerablesNew0: new0,
		},
	}
}

// Set the ideal number of non-invulnerable collators. If lowering this number, then the
// number of running collators could be higher than this figure. Aside from that edge case,
// there should be no other way to have more candidates than the desired number.
//
// The origin for this call must be the `UpdateOrigin`.
func MakeSetDesiredCandidatesCall(max0 uint32) types.RuntimeCall {
	return types.RuntimeCall{
		IsCollatorSelection: true,
		AsCollatorSelectionField0: &types.PalletCollatorSelectionPalletCall{
			IsSetDesiredCandidates:     true,
			AsSetDesiredCandidatesMax0: max0,
		},
	}
}

// Set the candidacy bond amount.
//
// If the candidacy bond is increased by this call, all current candidates which have a
// deposit lower than the new bond will be kicked from the list and get their deposits
// back.
//
// The origin for this call must be the `UpdateOrigin`.
func MakeSetCandidacyBondCall(bond0 types1.U128) types.RuntimeCall {
	return types.RuntimeCall{
		IsCollatorSelection: true,
		AsCollatorSelectionField0: &types.PalletCollatorSelectionPalletCall{
			IsSetCandidacyBond:      true,
			AsSetCandidacyBondBond0: bond0,
		},
	}
}

// Register this account as a collator candidate. The account must (a) already have
// registered session keys and (b) be able to reserve the `CandidacyBond`.
//
// This call is not available to `Invulnerable` collators.
func MakeRegisterAsCandidateCall() types.RuntimeCall {
	return types.RuntimeCall{
		IsCollatorSelection: true,
		AsCollatorSelectionField0: &types.PalletCollatorSelectionPalletCall{
			IsRegisterAsCandidate: true,
		},
	}
}

// Deregister `origin` as a collator candidate. Note that the collator can only leave on
// session change. The `CandidacyBond` will be unreserved immediately.
//
// This call will fail if the total number of candidates would drop below
// `MinEligibleCollators`.
func MakeLeaveIntentCall() types.RuntimeCall {
	return types.RuntimeCall{
		IsCollatorSelection: true,
		AsCollatorSelectionField0: &types.PalletCollatorSelectionPalletCall{
			IsLeaveIntent: true,
		},
	}
}

// Add a new account `who` to the list of `Invulnerables` collators. `who` must have
// registered session keys. If `who` is a candidate, they will be removed.
//
// The origin for this call must be the `UpdateOrigin`.
func MakeAddInvulnerableCall(who0 [32]byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsCollatorSelection: true,
		AsCollatorSelectionField0: &types.PalletCollatorSelectionPalletCall{
			IsAddInvulnerable:     true,
			AsAddInvulnerableWho0: who0,
		},
	}
}

// Remove an account `who` from the list of `Invulnerables` collators. `Invulnerables` must
// be sorted.
//
// The origin for this call must be the `UpdateOrigin`.
func MakeRemoveInvulnerableCall(who0 [32]byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsCollatorSelection: true,
		AsCollatorSelectionField0: &types.PalletCollatorSelectionPalletCall{
			IsRemoveInvulnerable:     true,
			AsRemoveInvulnerableWho0: who0,
		},
	}
}

// Update the candidacy bond of collator candidate `origin` to a new amount `new_deposit`.
//
// Setting a `new_deposit` that is lower than the current deposit while `origin` is
// occupying a top-`DesiredCandidates` slot is not allowed.
//
// This call will fail if `origin` is not a collator candidate, the updated bond is lower
// than the minimum candidacy bond, and/or the amount cannot be reserved.
func MakeUpdateBondCall(newDeposit0 types1.U128) types.RuntimeCall {
	return types.RuntimeCall{
		IsCollatorSelection: true,
		AsCollatorSelectionField0: &types.PalletCollatorSelectionPalletCall{
			IsUpdateBond:            true,
			AsUpdateBondNewDeposit0: newDeposit0,
		},
	}
}

// The caller `origin` replaces a candidate `target` in the collator candidate list by
// reserving `deposit`. The amount `deposit` reserved by the caller must be greater than
// the existing bond of the target it is trying to replace.
//
// This call will fail if the caller is already a collator candidate or invulnerable, the
// caller does not have registered session keys, the target is not a collator candidate,
// and/or the `deposit` amount cannot be reserved.
func MakeTakeCandidateSlotCall(deposit0 types1.U128, target1 [32]byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsCollatorSelection: true,
		AsCollatorSelectionField0: &types.PalletCollatorSelectionPalletCall{
			IsTakeCandidateSlot:         true,
			AsTakeCandidateSlotDeposit0: deposit0,
			AsTakeCandidateSlotTarget1:  target1,
		},
	}
}
