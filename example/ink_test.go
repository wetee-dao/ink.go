package example_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	chain "github.com/wetee-dao/ink.go"

	"github.com/wetee-dao/ink.go/example/contracts/cloud"
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
	contractAddress, err := util.HexToH160(contractAddressStr)
	if err != nil {
		util.LogWithPurple("HexToH160", err)
		return
	}

	// init contract
	contract := cloud.Cloud{
		ChainClient: chainClient,
		Address:     contractAddress,
	}

	// query data
	data, _, err := contract.QueryUserPods(
		chain.DefaultParamWithOragin(p.AccountID()),
	)
	if err != nil {
		fmt.Println(err)
	}
	initLen := len(*data)

	_, gas, err := contract.DryRunCreateUserPod(
		chain.DefaultParamWithOragin(p.AccountID()),
	)
	if err != nil {
		fmt.Println(err)
	}

	err = contract.CallCreateUserPod(
		chain.CallParams{
			Signer:              &p,
			PayAmount:           types.NewU128(*big.NewInt(0)),
			GasLimit:            gas.GasRequired,
			StorageDepositLimit: gas.StorageDeposit,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(err)

	// query data
	data2, _, err := contract.QueryUserPods(
		chain.DefaultParamWithOragin(p.AccountID()),
	)
	if err != nil {
		fmt.Println(err)
	}
	initLen2 := len(*data2)

	if initLen2-initLen != 1 {
		t.Fatal("create pod failed")
	}
}
