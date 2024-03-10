package utility

import types "github.com/wetee-dao/go-sdk/gen/types"

// See [`Pallet::batch`].
func MakeBatchCall(calls0 []types.RuntimeCall) types.RuntimeCall {
	return types.RuntimeCall{
		IsUtility: true,
		AsUtilityField0: &types.PalletUtilityPalletCall{
			IsBatch:       true,
			AsBatchCalls0: calls0,
		},
	}
}

// See [`Pallet::as_derivative`].
func MakeAsDerivativeCall(index0 uint16, call1 types.RuntimeCall) types.RuntimeCall {
	return types.RuntimeCall{
		IsUtility: true,
		AsUtilityField0: &types.PalletUtilityPalletCall{
			IsAsDerivative:       true,
			AsAsDerivativeIndex0: index0,
			AsAsDerivativeCall1:  &call1,
		},
	}
}

// See [`Pallet::batch_all`].
func MakeBatchAllCall(calls0 []types.RuntimeCall) types.RuntimeCall {
	return types.RuntimeCall{
		IsUtility: true,
		AsUtilityField0: &types.PalletUtilityPalletCall{
			IsBatchAll:       true,
			AsBatchAllCalls0: calls0,
		},
	}
}

// See [`Pallet::dispatch_as`].
func MakeDispatchAsCall(asOrigin0 types.OriginCaller, call1 types.RuntimeCall) types.RuntimeCall {
	return types.RuntimeCall{
		IsUtility: true,
		AsUtilityField0: &types.PalletUtilityPalletCall{
			IsDispatchAs:          true,
			AsDispatchAsAsOrigin0: &asOrigin0,
			AsDispatchAsCall1:     &call1,
		},
	}
}

// See [`Pallet::force_batch`].
func MakeForceBatchCall(calls0 []types.RuntimeCall) types.RuntimeCall {
	return types.RuntimeCall{
		IsUtility: true,
		AsUtilityField0: &types.PalletUtilityPalletCall{
			IsForceBatch:       true,
			AsForceBatchCalls0: calls0,
		},
	}
}

// See [`Pallet::with_weight`].
func MakeWithWeightCall(call0 types.RuntimeCall, weight1 types.Weight) types.RuntimeCall {
	return types.RuntimeCall{
		IsUtility: true,
		AsUtilityField0: &types.PalletUtilityPalletCall{
			IsWithWeight:        true,
			AsWithWeightCall0:   &call0,
			AsWithWeightWeight1: weight1,
		},
	}
}
