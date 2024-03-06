package main

import (
	"fmt"
	"math/big"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/wetee-dao/go-sdk/gen/balances"

	chain "github.com/wetee-dao/go-sdk"
	gtypes "github.com/wetee-dao/go-sdk/gen/types"
)

func main() {
	client, err := chain.ClientInit("ws://192.168.111.105:30001")
	if err != nil {
		panic(err)
	}

	// 1 unit of transfer
	bal, ok := new(big.Int).SetString("5000000000000000000", 10)
	if !ok {
		panic(fmt.Errorf("failed to convert balance"))
	}

	minter, _ := types.NewMultiAddressFromHexAccountID("0x8eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a48")
	minterWrap := gtypes.MultiAddress{
		IsId:       true,
		AsIdField0: minter.AsID,
	}

	c := balances.MakeTransferCall(minterWrap, types.NewUCompact(bal))
	err = client.SignAndSubmit(&signature.TestKeyringPairAlice, c, false)
	if err != nil {
		panic(err)
	}
}
