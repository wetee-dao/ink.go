package weteebridge

import (
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

func MakeInkCallbackCall(clusterId0 uint64, callId1 types.U128, args2 []byte, value3 types.U128) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeTEEBridge: true,
		AsWeTEEBridgeField0: &types1.WeteeTeeBridgePalletCall{
			IsInkCallback:           true,
			AsInkCallbackClusterId0: clusterId0,
			AsInkCallbackCallId1:    callId1,
			AsInkCallbackArgs2:      args2,
			AsInkCallbackValue3:     value3,
		},
	}
}
func MakeSetTeeApiCall(workId0 types1.WorkId, meta1 types1.ApiMeta) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeTEEBridge: true,
		AsWeTEEBridgeField0: &types1.WeteeTeeBridgePalletCall{
			IsSetTeeApi:        true,
			AsSetTeeApiWorkId0: workId0,
			AsSetTeeApiMeta1:   meta1,
		},
	}
}
