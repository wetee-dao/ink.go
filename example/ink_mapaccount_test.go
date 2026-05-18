package example

import (
	"fmt"
	"testing"

	chain "github.com/wetee-dao/ink.go"
	"github.com/wetee-dao/ink.go/pallet/revive"
	"github.com/wetee-dao/ink.go/util"
)

func TestDssSubmitTx(t *testing.T) {
	client, err := chain.InitClient([]string{"wss://asset-hub-paseo.ibp.network"}, true)
	if err != nil {
		panic(err)
	}

	pk, err := chain.Sr25519PairFromSecret("0x43c0defe545b31e6f64ee9c17b35c6444a8c3527f0c68160203717c51367512d", uint16(42))
	if err != nil {
		util.LogWithPurple("Sr25519PairFromSecret", err)
		t.Fatal(err)
	}

	fmt.Println(pk.Address)

	runtimeCall := revive.MakeMapAccountCall()
	call, err := (runtimeCall).AsCall()
	if err != nil {
		util.LogWithPurple("MakeMapAccountCall", err)
		t.Fatal(err)
	}

	util.LogWithYellow("MakeMapAccount", pk.Address)
	err = client.SignAndSubmit(&pk, call, true, 0)
	if err != nil {
		t.Fatal(err)
	}
}
