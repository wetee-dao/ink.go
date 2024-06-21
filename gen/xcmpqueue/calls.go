package xcmpqueue

import types "github.com/wetee-dao/go-sdk/gen/types"

// Suspends all XCM executions for the XCMP queue, regardless of the sender's origin.
//
// - `origin`: Must pass `ControllerOrigin`.
func MakeSuspendXcmExecutionCall() types.RuntimeCall {
	return types.RuntimeCall{
		IsXcmpQueue: true,
		AsXcmpQueueField0: &types.CumulusPalletXcmpQueuePalletCall{
			IsSuspendXcmExecution: true,
		},
	}
}

// Resumes all XCM executions for the XCMP queue.
//
// Note that this function doesn't change the status of the in/out bound channels.
//
// - `origin`: Must pass `ControllerOrigin`.
func MakeResumeXcmExecutionCall() types.RuntimeCall {
	return types.RuntimeCall{
		IsXcmpQueue: true,
		AsXcmpQueueField0: &types.CumulusPalletXcmpQueuePalletCall{
			IsResumeXcmExecution: true,
		},
	}
}

// Overwrites the number of pages which must be in the queue for the other side to be
// told to suspend their sending.
//
// - `origin`: Must pass `Root`.
// - `new`: Desired value for `QueueConfigData.suspend_value`
func MakeUpdateSuspendThresholdCall(new0 uint32) types.RuntimeCall {
	return types.RuntimeCall{
		IsXcmpQueue: true,
		AsXcmpQueueField0: &types.CumulusPalletXcmpQueuePalletCall{
			IsUpdateSuspendThreshold:     true,
			AsUpdateSuspendThresholdNew0: new0,
		},
	}
}

// Overwrites the number of pages which must be in the queue after which we drop any
// further messages from the channel.
//
// - `origin`: Must pass `Root`.
// - `new`: Desired value for `QueueConfigData.drop_threshold`
func MakeUpdateDropThresholdCall(new0 uint32) types.RuntimeCall {
	return types.RuntimeCall{
		IsXcmpQueue: true,
		AsXcmpQueueField0: &types.CumulusPalletXcmpQueuePalletCall{
			IsUpdateDropThreshold:     true,
			AsUpdateDropThresholdNew0: new0,
		},
	}
}

// Overwrites the number of pages which the queue must be reduced to before it signals
// that message sending may recommence after it has been suspended.
//
// - `origin`: Must pass `Root`.
// - `new`: Desired value for `QueueConfigData.resume_threshold`
func MakeUpdateResumeThresholdCall(new0 uint32) types.RuntimeCall {
	return types.RuntimeCall{
		IsXcmpQueue: true,
		AsXcmpQueueField0: &types.CumulusPalletXcmpQueuePalletCall{
			IsUpdateResumeThreshold:     true,
			AsUpdateResumeThresholdNew0: new0,
		},
	}
}
