package weteebridge

import (
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

func MakeInkCallbackCall(clusterId0 uint64, callId1 types.U128, args2 []types1.InkArg, value3 types.U128, error4 types1.OptionTByteSlice) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeTEEBridge: true,
		AsWeTEEBridgeField0: &types1.WeteeTeeBridgePalletCall{
			IsInkCallback:           true,
			AsInkCallbackClusterId0: clusterId0,
			AsInkCallbackCallId1:    callId1,
			AsInkCallbackArgs2:      args2,
			AsInkCallbackValue3:     value3,
			AsInkCallbackError4:     error4,
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
func MakeDeleteCallCall(clusterId0 uint64, callId1 types.U128) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeTEEBridge: true,
		AsWeTEEBridgeField0: &types1.WeteeTeeBridgePalletCall{
			IsDeleteCall:           true,
			AsDeleteCallClusterId0: clusterId0,
			AsDeleteCallCallId1:    callId1,
		},
	}
}
