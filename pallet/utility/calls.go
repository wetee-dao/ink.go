package utility

import types "github.com/wetee-dao/go-sdk/pallet/types"

// Send a batch of dispatch calls.
//
// May be called from any origin except `None`.
//
//   - `calls`: The calls to be dispatched from the same origin. The number of call must not
//     exceed the constant: `batched_calls_limit` (available in constant metadata).
//
// If origin is root then the calls are dispatched without checking origin filter. (This
// includes bypassing `frame_system::Config::BaseCallFilter`).
//
// ## Complexity
// - O(C) where C is the number of calls to be batched.
//
// This will return `Ok` in all circumstances. To determine the success of the batch, an
// event is deposited. If a call failed and the batch was interrupted, then the
// `BatchInterrupted` event is deposited, along with the number of successful calls made
// and the error of the failed call. If all were successful, then the `BatchCompleted`
// event is deposited.
func MakeBatchCall(calls0 []types.RuntimeCall) types.RuntimeCall {
	return types.RuntimeCall{
		IsUtility: true,
		AsUtilityField0: &types.PalletUtilityPalletCall{
			IsBatch:       true,
			AsBatchCalls0: calls0,
		},
	}
}

// Send a call through an indexed pseudonym of the sender.
//
// Filter from origin are passed along. The call will be dispatched with an origin which
// use the same filter as the origin of this call.
//
// NOTE: If you need to ensure that any account-based filtering is not honored (i.e.
// because you expect `proxy` to have been used prior in the call stack and you do not want
// the call restrictions to apply to any sub-accounts), then use `as_multi_threshold_1`
// in the Multisig pallet instead.
//
// NOTE: Prior to version *12, this was called `as_limited_sub`.
//
// The dispatch origin for this call must be _Signed_.
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

// Send a batch of dispatch calls and atomically execute them.
// The whole transaction will rollback and fail if any of the calls failed.
//
// May be called from any origin except `None`.
//
//   - `calls`: The calls to be dispatched from the same origin. The number of call must not
//     exceed the constant: `batched_calls_limit` (available in constant metadata).
//
// If origin is root then the calls are dispatched without checking origin filter. (This
// includes bypassing `frame_system::Config::BaseCallFilter`).
//
// ## Complexity
// - O(C) where C is the number of calls to be batched.
func MakeBatchAllCall(calls0 []types.RuntimeCall) types.RuntimeCall {
	return types.RuntimeCall{
		IsUtility: true,
		AsUtilityField0: &types.PalletUtilityPalletCall{
			IsBatchAll:       true,
			AsBatchAllCalls0: calls0,
		},
	}
}

// Dispatches a function call with a provided origin.
//
// The dispatch origin for this call must be _Root_.
//
// ## Complexity
// - O(1).
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

// Send a batch of dispatch calls.
// Unlike `batch`, it allows errors and won't interrupt.
//
// May be called from any origin except `None`.
//
//   - `calls`: The calls to be dispatched from the same origin. The number of call must not
//     exceed the constant: `batched_calls_limit` (available in constant metadata).
//
// If origin is root then the calls are dispatch without checking origin filter. (This
// includes bypassing `frame_system::Config::BaseCallFilter`).
//
// ## Complexity
// - O(C) where C is the number of calls to be batched.
func MakeForceBatchCall(calls0 []types.RuntimeCall) types.RuntimeCall {
	return types.RuntimeCall{
		IsUtility: true,
		AsUtilityField0: &types.PalletUtilityPalletCall{
			IsForceBatch:       true,
			AsForceBatchCalls0: calls0,
		},
	}
}

// Dispatch a function call with a specified weight.
//
// This function does not check the weight of the call, and instead allows the
// Root origin to specify the weight of the call.
//
// The dispatch origin for this call must be _Root_.
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
