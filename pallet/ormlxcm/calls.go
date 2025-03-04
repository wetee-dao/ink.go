package ormlxcm

import types "github.com/wetee-dao/go-sdk/pallet/types"

// Send an XCM message as parachain sovereign.
func MakeSendAsSovereignCall(dest0 types.VersionedLocation, message1 types.VersionedXcm) types.RuntimeCall {
	return types.RuntimeCall{
		IsOrmlXcm: true,
		AsOrmlXcmField0: &types.OrmlXcmModuleCall{
			IsSendAsSovereign:         true,
			AsSendAsSovereignDest0:    &dest0,
			AsSendAsSovereignMessage1: &message1,
		},
	}
}
