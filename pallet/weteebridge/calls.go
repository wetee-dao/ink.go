package weteebridge

import types "github.com/wetee-dao/go-sdk/pallet/types"

// 注册 dkg 节点
// register dkg node
func MakeTestCall(pubkey0 [32]byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEBridge: true,
		AsWeTEEBridgeField0: &types.WeteeTeeBridgePalletCall{
			IsTest:        true,
			AsTestPubkey0: pubkey0,
		},
	}
}
