package weteedsecret

import types "github.com/wetee-dao/go-sdk/gen/types"

// 注册 dkg 节点
// register dkg node
func MakeRegisterNodeCall(pubkey0 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEDsecret: true,
		AsWeTEEDsecretField0: &types.WeteeDsecretPalletCall{
			IsRegisterNode:        true,
			AsRegisterNodePubkey0: pubkey0,
		},
	}
}

// 上传共识节点代码
// update consensus node code
func MakeUploadCodeCall(mrenclave0 []byte, mrsigner1 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEDsecret: true,
		AsWeTEEDsecretField0: &types.WeteeDsecretPalletCall{
			IsUploadCode:           true,
			AsUploadCodeMrenclave0: mrenclave0,
			AsUploadCodeMrsigner1:  mrsigner1,
		},
	}
}
