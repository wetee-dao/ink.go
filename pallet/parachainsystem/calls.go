package parachainsystem

import types "github.com/wetee-dao/go-sdk/pallet/types"

// Set the current validation data.
//
// This should be invoked exactly once per block. It will panic at the finalization
// phase if the call was not invoked.
//
// The dispatch origin for this call must be `Inherent`
//
// As a side effect, this function upgrades the current validation function
// if the appropriate time has come.
func MakeSetValidationDataCall(data0 types.ParachainInherentData) types.RuntimeCall {
	return types.RuntimeCall{
		IsParachainSystem: true,
		AsParachainSystemField0: &types.CumulusPalletParachainSystemPalletCall{
			IsSetValidationData:      true,
			AsSetValidationDataData0: data0,
		},
	}
}
func MakeSudoSendUpwardMessageCall(message0 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsParachainSystem: true,
		AsParachainSystemField0: &types.CumulusPalletParachainSystemPalletCall{
			IsSudoSendUpwardMessage:         true,
			AsSudoSendUpwardMessageMessage0: message0,
		},
	}
}

// Authorize an upgrade to a given `code_hash` for the runtime. The runtime can be supplied
// later.
//
// The `check_version` parameter sets a boolean flag for whether or not the runtime's spec
// version and name should be verified on upgrade. Since the authorization only has a hash,
// it cannot actually perform the verification.
//
// This call requires Root origin.
func MakeAuthorizeUpgradeCall(codeHash0 [32]byte, checkVersion1 bool) types.RuntimeCall {
	return types.RuntimeCall{
		IsParachainSystem: true,
		AsParachainSystemField0: &types.CumulusPalletParachainSystemPalletCall{
			IsAuthorizeUpgrade:              true,
			AsAuthorizeUpgradeCodeHash0:     codeHash0,
			AsAuthorizeUpgradeCheckVersion1: checkVersion1,
		},
	}
}

// Provide the preimage (runtime binary) `code` for an upgrade that has been authorized.
//
// If the authorization required a version check, this call will ensure the spec name
// remains unchanged and that the spec version has increased.
//
// Note that this function will not apply the new `code`, but only attempt to schedule the
// upgrade with the Relay Chain.
//
// All origins are allowed.
func MakeEnactAuthorizedUpgradeCall(code0 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsParachainSystem: true,
		AsParachainSystemField0: &types.CumulusPalletParachainSystemPalletCall{
			IsEnactAuthorizedUpgrade:      true,
			AsEnactAuthorizedUpgradeCode0: code0,
		},
	}
}
