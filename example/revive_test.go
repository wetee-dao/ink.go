package example_test

import (
	"fmt"
	"math/big"
	"os"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	chain "github.com/wetee-dao/ink.go"

	"github.com/wetee-dao/ink.go/example/contracts/dao"
	"github.com/wetee-dao/ink.go/util"
)

var UTIL = big.NewInt(1_000_000_000_000)

func ExampleRevive() {
	chainClient, err := chain.ClientInit("ws://127.0.0.1:9944", false)
	if err != nil {
		panic(err)
	}

	// 初始化私钥
	var testSecretSeed = "0x167d9a020688544ea246b056799d6a771e97c9da057e4d0b87024537f99177bc"
	p, err := chain.Ed25519PairFromSecret(testSecretSeed, 42)
	if err != nil {
		panic(err)
	}

	// 获取合约地址
	contractAddressStr := "0x1547E25E7fe95a931E96907C70529d57D2438aD1"
	contractAddress, err := util.HexToH160(contractAddressStr)
	if err != nil {
		util.LogWithPurple("HexToH160", err)
		return
	}

	// init contract
	abiRaw, err := os.ReadFile("./contracts/dao.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	contract, err := chain.NewRevive(
		chainClient,
		contractAddress,
		abiRaw,
	)
	if err != nil {
		util.LogWithPurple("NewRevive", err)
		return
	}

	// query Member::list
	memberList, err := chain.QueryInk[[]types.H160](
		contract,
		util.NewAccountID(p.PublicKey),
		util.InkContractInput{
			Selector: "Member::list",
			Args:     []any{},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	// query Gov::track return type is dao.Track
	_, err = chain.QueryInk[util.Option[dao.Track]](
		contract,
		util.NewAccountID(p.PublicKey),
		util.InkContractInput{
			Selector: "Gov::track",
			Args: []any{
				types.NewU16(0),
			},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	// dry run contract
	result, err := chain.DryRunInk(
		contract,
		util.NewAccountID(p.PublicKey),
		types.NewU128(*big.NewInt(0)),
		util.NewNone[types.Weight](),
		util.NewNone[types.U128](),
		util.InkContractInput{
			Selector: "Erc20::enable_transfer",
			Args:     []any{},
		},
	)
	if err == nil {
		fmt.Println(result)
	}

	// err = chainClient.MapReviveAccount(&p)
	// if err != nil {
	// 	panic(err)
	// }

	err = chain.CallInk(
		contract,
		&p,
		types.NewU128(*big.NewInt(0)),
		types.Weight{
			RefTime:   types.NewUCompact(big.NewInt(1_100_000_000)),
			ProofSize: types.NewUCompact(big.NewInt(100_000)),
		},
		types.NewU128(*big.NewInt(110_000_000_000)),
		util.InkContractInput{
			Selector: "Member::public_join",
			Args:     []any{},
		},
	)

	isIn := false
	for _, m := range *memberList {
		if m.Hex() == p.H160Address().Hex() {
			isIn = true
		}
	}

	if isIn {
		if err == nil {
			fmt.Println("error")
		}
	} else {
		if err != nil {
			fmt.Println(err)
		}
	}

	// Output:
}
