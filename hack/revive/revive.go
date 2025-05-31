package main

import (
	"fmt"
	"math/big"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	chain "github.com/wetee-dao/go-sdk"
	"github.com/wetee-dao/go-sdk/module"
	"github.com/wetee-dao/go-sdk/util"
)

func main() {
	client, err := chain.ClientInit("ws://127.0.0.1:9944", true)
	if err != nil {
		panic(err)
	}

	revice := module.Revive{
		Client: client,
	}

	b, err := revice.BalanceOfH160("0x2B9c0Cc310CAaFcf5E7c9A09cf0dC582053DAbAF")
	if err != nil {
		util.LogWithRed("getBy32", err)
		return
	}
	fmt.Println(b)

	mlist := [][20]byte{}
	err = revice.QueryInk(
		util.NewAccountID(signature.TestKeyringPairAlice.PublicKey),
		types.NewU128(*big.NewInt(0)),
		util.NewNone[types.Weight](),
		util.NewNone[types.U128](),
		"0x2B9c0Cc310CAaFcf5E7c9A09cf0dC582053DAbAF",
		util.InkContractInput{
			Selector: util.FuncToSelector("Member::list"),
			Args:     []any{},
		},
		&mlist,
	)
	if err == nil {
		fmt.Println(mlist)
	}
}
