package asset

import (
	types1 "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types "github.com/wetee-dao/go-sdk/pallet/types"
)

// create we asset.
// 创建 WETEE 资产
func MakeCreateAssetCall(metadata0 types.AssetMeta, initAmount1 types1.U128) types.RuntimeCall {
	return types.RuntimeCall{
		IsAsset: true,
		AsAssetField0: &types.WeteeAssetsPalletCall{
			IsCreateAsset:            true,
			AsCreateAssetMetadata0:   metadata0,
			AsCreateAssetInitAmount1: initAmount1,
		},
	}
}

// You should have created the asset first.
// 设置资产元数据
func MakeSetMetadataCall(assetId0 uint64, metadata1 types.AssetMeta) types.RuntimeCall {
	return types.RuntimeCall{
		IsAsset: true,
		AsAssetField0: &types.WeteeAssetsPalletCall{
			IsSetMetadata:          true,
			AsSetMetadataAssetId0:  assetId0,
			AsSetMetadataMetadata1: metadata1,
		},
	}
}

// Users destroy their own assets.
// 销毁资产
func MakeBurnCall(assetId0 uint64, amount1 types1.U128) types.RuntimeCall {
	return types.RuntimeCall{
		IsAsset: true,
		AsAssetField0: &types.WeteeAssetsPalletCall{
			IsBurn:         true,
			AsBurnAssetId0: assetId0,
			AsBurnAmount1:  amount1,
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
func MakeTransferCall(dest0 types.MultiAddress, assetId1 uint64, amount2 types1.UCompact) types.RuntimeCall {
	return types.RuntimeCall{
		IsAsset: true,
		AsAssetField0: &types.WeteeAssetsPalletCall{
			IsTransfer:         true,
			AsTransferDest0:    dest0,
			AsTransferAssetId1: assetId1,
			AsTransferAmount2:  amount2,
		},
	}
}
