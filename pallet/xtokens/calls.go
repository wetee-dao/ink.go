package xtokens

import (
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Transfer native currencies.
//
// `dest_weight_limit` is the weight for XCM execution on the dest
// chain, and it would be charged from the transferred assets. If set
// below requirements, the execution may fail and assets wouldn't be
// received.
//
// It's a no-op if any error on local XCM execution or message sending.
// Note sending assets out per se doesn't guarantee they would be
// received. Receiving depends on if the XCM message could be delivered
// by the network, and if the receiving chain would handle
// messages correctly.
func MakeTransferCall(currencyId0 uint64, amount1 types.U128, dest2 types1.VersionedLocation, destWeightLimit3 types1.WeightLimit) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsXTokens: true,
		AsXTokensField0: &types1.OrmlXtokensModuleCall{
			IsTransfer:                 true,
			AsTransferCurrencyId0:      currencyId0,
			AsTransferAmount1:          amount1,
			AsTransferDest2:            &dest2,
			AsTransferDestWeightLimit3: destWeightLimit3,
		},
	}
}

// Transfer `Asset`.
//
// `dest_weight_limit` is the weight for XCM execution on the dest
// chain, and it would be charged from the transferred assets. If set
// below requirements, the execution may fail and assets wouldn't be
// received.
//
// It's a no-op if any error on local XCM execution or message sending.
// Note sending assets out per se doesn't guarantee they would be
// received. Receiving depends on if the XCM message could be delivered
// by the network, and if the receiving chain would handle
// messages correctly.
func MakeTransferMultiassetCall(asset0 types1.VersionedAsset, dest1 types1.VersionedLocation, destWeightLimit2 types1.WeightLimit) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsXTokens: true,
		AsXTokensField0: &types1.OrmlXtokensModuleCall{
			IsTransferMultiasset:                 true,
			AsTransferMultiassetAsset0:           &asset0,
			AsTransferMultiassetDest1:            &dest1,
			AsTransferMultiassetDestWeightLimit2: destWeightLimit2,
		},
	}
}

// Transfer native currencies specifying the fee and amount as
// separate.
//
// `dest_weight_limit` is the weight for XCM execution on the dest
// chain, and it would be charged from the transferred assets. If set
// below requirements, the execution may fail and assets wouldn't be
// received.
//
// `fee` is the amount to be spent to pay for execution in destination
// chain. Both fee and amount will be subtracted form the callers
// balance.
//
// If `fee` is not high enough to cover for the execution costs in the
// destination chain, then the assets will be trapped in the
// destination chain
//
// It's a no-op if any error on local XCM execution or message sending.
// Note sending assets out per se doesn't guarantee they would be
// received. Receiving depends on if the XCM message could be delivered
// by the network, and if the receiving chain would handle
// messages correctly.
func MakeTransferWithFeeCall(currencyId0 uint64, amount1 types.U128, fee2 types.U128, dest3 types1.VersionedLocation, destWeightLimit4 types1.WeightLimit) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsXTokens: true,
		AsXTokensField0: &types1.OrmlXtokensModuleCall{
			IsTransferWithFee:                 true,
			AsTransferWithFeeCurrencyId0:      currencyId0,
			AsTransferWithFeeAmount1:          amount1,
			AsTransferWithFeeFee2:             fee2,
			AsTransferWithFeeDest3:            &dest3,
			AsTransferWithFeeDestWeightLimit4: destWeightLimit4,
		},
	}
}

// Transfer `Asset` specifying the fee and amount as separate.
//
// `dest_weight_limit` is the weight for XCM execution on the dest
// chain, and it would be charged from the transferred assets. If set
// below requirements, the execution may fail and assets wouldn't be
// received.
//
// `fee` is the Asset to be spent to pay for execution in
// destination chain. Both fee and amount will be subtracted form the
// callers balance For now we only accept fee and asset having the same
// `Location` id.
//
// If `fee` is not high enough to cover for the execution costs in the
// destination chain, then the assets will be trapped in the
// destination chain
//
// It's a no-op if any error on local XCM execution or message sending.
// Note sending assets out per se doesn't guarantee they would be
// received. Receiving depends on if the XCM message could be delivered
// by the network, and if the receiving chain would handle
// messages correctly.
func MakeTransferMultiassetWithFeeCall(asset0 types1.VersionedAsset, fee1 types1.VersionedAsset, dest2 types1.VersionedLocation, destWeightLimit3 types1.WeightLimit) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsXTokens: true,
		AsXTokensField0: &types1.OrmlXtokensModuleCall{
			IsTransferMultiassetWithFee:                 true,
			AsTransferMultiassetWithFeeAsset0:           &asset0,
			AsTransferMultiassetWithFeeFee1:             &fee1,
			AsTransferMultiassetWithFeeDest2:            &dest2,
			AsTransferMultiassetWithFeeDestWeightLimit3: destWeightLimit3,
		},
	}
}

// Transfer several currencies specifying the item to be used as fee
//
// `dest_weight_limit` is the weight for XCM execution on the dest
// chain, and it would be charged from the transferred assets. If set
// below requirements, the execution may fail and assets wouldn't be
// received.
//
// `fee_item` is index of the currencies tuple that we want to use for
// payment
//
// It's a no-op if any error on local XCM execution or message sending.
// Note sending assets out per se doesn't guarantee they would be
// received. Receiving depends on if the XCM message could be delivered
// by the network, and if the receiving chain would handle
// messages correctly.
func MakeTransferMulticurrenciesCall(currencies0 []types1.TupleOfUint64U128, feeItem1 uint32, dest2 types1.VersionedLocation, destWeightLimit3 types1.WeightLimit) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsXTokens: true,
		AsXTokensField0: &types1.OrmlXtokensModuleCall{
			IsTransferMulticurrencies:                 true,
			AsTransferMulticurrenciesCurrencies0:      currencies0,
			AsTransferMulticurrenciesFeeItem1:         feeItem1,
			AsTransferMulticurrenciesDest2:            &dest2,
			AsTransferMulticurrenciesDestWeightLimit3: destWeightLimit3,
		},
	}
}

// Transfer several `Asset` specifying the item to be used as fee
//
// `dest_weight_limit` is the weight for XCM execution on the dest
// chain, and it would be charged from the transferred assets. If set
// below requirements, the execution may fail and assets wouldn't be
// received.
//
// `fee_item` is index of the Assets that we want to use for
// payment
//
// It's a no-op if any error on local XCM execution or message sending.
// Note sending assets out per se doesn't guarantee they would be
// received. Receiving depends on if the XCM message could be delivered
// by the network, and if the receiving chain would handle
// messages correctly.
func MakeTransferMultiassetsCall(assets0 types1.VersionedAssets, feeItem1 uint32, dest2 types1.VersionedLocation, destWeightLimit3 types1.WeightLimit) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsXTokens: true,
		AsXTokensField0: &types1.OrmlXtokensModuleCall{
			IsTransferMultiassets:                 true,
			AsTransferMultiassetsAssets0:          &assets0,
			AsTransferMultiassetsFeeItem1:         feeItem1,
			AsTransferMultiassetsDest2:            &dest2,
			AsTransferMultiassetsDestWeightLimit3: destWeightLimit3,
		},
	}
}
