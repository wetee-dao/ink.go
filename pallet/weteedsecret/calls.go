package weteedsecret

import types "github.com/wetee-dao/go-sdk/pallet/types"

// 注册 dkg 节点
// register dkg node
func MakeRegisterNodeCall(sender0 [32]byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEDsecret: true,
		AsWeTEEDsecretField0: &types.WeteeDsecretPalletCall{
			IsRegisterNode:        true,
			AsRegisterNodeSender0: sender0,
		},
	}
}

// 上传共识节点代码
// update consensus node code
func MakeUploadCodeCall(signature0 []byte, signer1 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEDsecret: true,
		AsWeTEEDsecretField0: &types.WeteeDsecretPalletCall{
			IsUploadCode:           true,
			AsUploadCodeSignature0: signature0,
			AsUploadCodeSigner1:    signer1,
		},
	}
}

// 上传共识节点代码
// update consensus node code
func MakeUploadClusterProofCall(cid0 uint64, report1 []byte, pubs2 [][32]byte, sigs3 []types.MultiSignature) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEDsecret: true,
		AsWeTEEDsecretField0: &types.WeteeDsecretPalletCall{
			IsUploadClusterProof:        true,
			AsUploadClusterProofCid0:    cid0,
			AsUploadClusterProofReport1: report1,
			AsUploadClusterProofPubs2:   pubs2,
			AsUploadClusterProofSigs3:   sigs3,
		},
	}
}

// 上传 devloper，report hash 启动应用
func MakeWorkLaunchCall(work0 types.WorkId, report1 types.OptionTByteSlice, deployKey2 [32]byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEDsecret: true,
		AsWeTEEDsecretField0: &types.WeteeDsecretPalletCall{
			IsWorkLaunch:           true,
			AsWorkLaunchWork0:      work0,
			AsWorkLaunchReport1:    report1,
			AsWorkLaunchDeployKey2: deployKey2,
		},
	}
}
