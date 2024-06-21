package weteeasset

import (
	types1 "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types "github.com/wetee-dao/go-sdk/gen/types"
)

// create dao asset.
// 创建 WETEE 资产
func MakeCreateAssetCall(daoId0 uint64, metadata1 types.DaoAssetMeta, amount2 types1.U128, initDaoAsset3 types1.U128) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEAsset: true,
		AsWeTEEAssetField0: &types.WeteeAssetsPalletCall{
			IsCreateAsset:              true,
			AsCreateAssetDaoId0:        daoId0,
			AsCreateAssetMetadata1:     metadata1,
			AsCreateAssetAmount2:       amount2,
			AsCreateAssetInitDaoAsset3: initDaoAsset3,
		},
	}
}

// 设置加入WETEE所需要的最小抵押
func MakeSetExistenialDepositCall(daoId0 uint64, existenialDeposit1 types1.U128) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEAsset: true,
		AsWeTEEAssetField0: &types.WeteeAssetsPalletCall{
			IsSetExistenialDeposit:                   true,
			AsSetExistenialDepositDaoId0:             daoId0,
			AsSetExistenialDepositExistenialDeposit1: existenialDeposit1,
		},
	}
}

// You should have created the asset first.
// 设置资产元数据
func MakeSetMetadataCall(daoId0 uint64, metadata1 types.DaoAssetMeta) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEAsset: true,
		AsWeTEEAssetField0: &types.WeteeAssetsPalletCall{
			IsSetMetadata:          true,
			AsSetMetadataDaoId0:    daoId0,
			AsSetMetadataMetadata1: metadata1,
		},
	}
}

// Users destroy their own assets.
// 销毁资产
func MakeBurnCall(daoId0 uint64, amount1 types1.U128) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEAsset: true,
		AsWeTEEAssetField0: &types.WeteeAssetsPalletCall{
			IsBurn:        true,
			AsBurnDaoId0:  daoId0,
			AsBurnAmount1: amount1,
		},
	}
}

// This function transfers the given amount from the source to the destination.
//
// # Arguments
//
// * `amount` - The amount to transfer
// * `source` - The source account
// * `destination` - The destination account
// 转移资产
func MakeTransferCall(dest0 types.MultiAddress, daoId1 uint64, amount2 types1.UCompact) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEAsset: true,
		AsWeTEEAssetField0: &types.WeteeAssetsPalletCall{
			IsTransfer:        true,
			AsTransferDest0:   dest0,
			AsTransferDaoId1:  daoId1,
			AsTransferAmount2: amount2,
		},
	}
}

// 成为会员
func MakeJoinCall(daoId0 uint64, shareExpect1 uint32, existenialDeposit2 types1.UCompact) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEAsset: true,
		AsWeTEEAssetField0: &types.WeteeAssetsPalletCall{
			IsJoin:                   true,
			AsJoinDaoId0:             daoId0,
			AsJoinShareExpect1:       shareExpect1,
			AsJoinExistenialDeposit2: existenialDeposit2,
		},
	}
}
