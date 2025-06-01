package main

import (
	"fmt"
	"math/big"
	"os"

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

	data, err := os.ReadFile("./DAO.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	contract := "0xFF9C30963C48949325E93e1D1BF02429f35382bA"
	contractAddress, err := util.HexToH160(contract)
	if err != nil {
		util.LogWithRed("HexToH160", err)
		return
	}

	revice, err := module.NewRevive(
		client,
		contractAddress,
		data,
	)
	if err != nil {
		util.LogWithRed("NewRevive", err)
		return
	}

	b, err := client.BalanceOfH160(contract)
	if err != nil {
		util.LogWithRed("BalanceOfH160", err)
		return
	}
	fmt.Println(b)
	fmt.Println("")

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
	if err == nil {
		for _, m := range mlist {
			fmt.Println(m.Hex())
		}
		fmt.Println("")
	} else {
		fmt.Println("-----------------")
		fmt.Println(err)
	}

	track := util.NewSome(types.NewU16(10))
	err = revice.QueryInk(
		util.NewAccountID(signature.TestKeyringPairAlice.PublicKey),
		types.NewU128(*big.NewInt(0)),
		util.NewNone[types.Weight](),
		util.NewNone[types.U128](),
		util.InkContractInput{
			Selector: util.FuncToSelector("Gov::defalut_track"),
			Args:     []any{},
		},
		&track,
	)
	if err == nil {
		fmt.Println(track)
		fmt.Println("")
	} else {
		fmt.Println("-----------------")
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
}
