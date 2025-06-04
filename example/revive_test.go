package example_test

import (
	"fmt"
	"math/big"
	"os"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	chain "github.com/wetee-dao/go-sdk"
	"github.com/wetee-dao/go-sdk/ink/dao"
	"github.com/wetee-dao/go-sdk/module"
	"github.com/wetee-dao/go-sdk/util"
)

func ExampleRevive() {
	client, err := chain.ClientInit("ws://127.0.0.1:9944", true)
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile("../ink/dao.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	contract := "0x5578537B6c44A654DdF461a0948FE50bb2E5F76C"
	contractAddress, err := util.HexToH160(contract)
	if err != nil {
		util.LogWithPurple("HexToH160", err)
		return
	}

	revice, err := module.NewRevive(
		client,
		contractAddress,
		data,
	)
	if err != nil {
		util.LogWithPurple("NewRevive", err)
		return
	}

	_, err = client.BalanceOfH160(contract)
	if err != nil {
		util.LogWithPurple("BalanceOfH160", err)
		return
	}

	mlist := make([]types.H160, 0, 10)
	err = revice.QueryInk(
		util.NewAccountID(signature.TestKeyringPairAlice.PublicKey),
		types.NewU128(*big.NewInt(0)),
		util.NewNone[types.Weight](),
		util.NewNone[types.U128](),
		util.InkContractInput{
			Selector: util.FuncToSelector("Member::list"),
			Args:     []any{},
		},
		&mlist,
	)
	if err != nil {
		fmt.Println(err)
	}

	track := util.NewSome(dao.Track{})
	err = revice.QueryInk(
		util.NewAccountID(signature.TestKeyringPairAlice.PublicKey),
		types.NewU128(*big.NewInt(0)),
		util.NewNone[types.Weight](),
		util.NewNone[types.U128](),
		util.InkContractInput{
			Selector: util.FuncToSelector("Gov::track"),
			Args: []any{
				types.NewU16(0),
			},
		},
		&track,
	)
	if err != nil {
		fmt.Println(err)
	}

	// v, err := revice.DryRunInk(
	// 	util.NewAccountID(signature.TestKeyringPairAlice.PublicKey),
	// 	types.NewU128(*big.NewInt(0)),
	// 	util.NewNone[types.Weight](),
	// 	util.NewNone[types.U128](),
	// 	contract,
	// 	util.InkContractInput{
	// 		Selector: util.FuncToSelector("Member::levae"),
	// 		Args:     []any{},
	// 	},
	// )
	// fmt.Println(err)
	// if err == nil {
	// 	fmt.Println(v)
	// }

	// Output:
}
