package example_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	chain "github.com/wetee-dao/ink.go"

	"github.com/wetee-dao/ink.go/tools/go-ink-gen/cloud"
	"github.com/wetee-dao/ink.go/util"
)

func TestInk(t *testing.T) {
	chainClient, err := chain.ClientInit("ws://127.0.0.1:9944", true)
	if err != nil {
		panic(err)
	}

	// 初始化私钥
	p, err := chain.Sr25519PairFromSecret("//Alice", 42)
	if err != nil {
		util.LogWithPurple("Sr25519PairFromSecret", err)
		t.Fatal(err)
	}

	// 获取合约地址
	contractAddressStr := "0xF39250328320705cdE6B478c6bc425239014a6D4"
	// init contract
	contract, err := cloud.InitCloudContract(chainClient, contractAddressStr)
	if err != nil {
		util.LogWithPurple("HexToH160", err)
		return
	}

	// query data
	data, _, err := contract.QueryUserPods(
		chain.DefaultParamWithOrigin(p.AccountID()),
	)
	if err != nil {
		fmt.Println(err)
	}
	initLen := len(*data)

	err = contract.CallCreateUserPod(
		chain.CallParams{
			Signer:    &p,
			PayAmount: types.NewU128(*big.NewInt(0)),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// query data
	data2, _, err := contract.QueryUserPods(
		chain.DefaultParamWithOrigin(p.AccountID()),
	)
	if err != nil {
		fmt.Println(err)
	}
	initLen2 := len(*data2)

	if initLen2-initLen != 1 {
		t.Fatal("create pod failed")
	}
}
