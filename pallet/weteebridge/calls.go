package weteebridge

import (
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

func MakeInkCallbackCall(callId0 types.U128, data1 []byte, value2 types.U128) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeTEEBridge: true,
		AsWeTEEBridgeField0: &types1.WeteeTeeBridgePalletCall{
			IsInkCallback:        true,
			AsInkCallbackCallId0: callId0,
			AsInkCallbackData1:   data1,
			AsInkCallbackValue2:  value2,
		},
	}
}
