package fairlanch

import (
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

func MakeVStakingCall(vassertId0 uint64, vamount1 types.U128) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsFairlanch: true,
		AsFairlanchField0: &types1.WeteeFairlanchPalletCall{
			IsVStaking:           true,
			AsVStakingVassertId0: vassertId0,
			AsVStakingVamount1:   vamount1,
		},
	}
}
func MakeVUnstakingCall(vassertId0 uint64, amount1 types.U128) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsFairlanch: true,
		AsFairlanchField0: &types1.WeteeFairlanchPalletCall{
			IsVUnstaking:           true,
			AsVUnstakingVassertId0: vassertId0,
			AsVUnstakingAmount1:    amount1,
		},
	}
}
