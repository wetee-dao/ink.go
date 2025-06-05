package example_test

import (
	"fmt"
	"math/big"
	"os"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	chain "github.com/wetee-dao/go-sdk"

	"github.com/wetee-dao/go-sdk/example/contracts/dao"
	"github.com/wetee-dao/go-sdk/util"
)

func ExampleRevive() {
	chainClient, err := chain.ClientInit("ws://127.0.0.1:9944", true)
	if err != nil {
		panic(err)
	}

	abiRaw, err := os.ReadFile("./contracts/dao.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	contractStr := "0x5578537B6c44A654DdF461a0948FE50bb2E5F76C"
	contractAddress, err := util.HexToH160(contractStr)
	if err != nil {
		util.LogWithPurple("HexToH160", err)
		return
	}

	// init contract
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
	_, err = chain.QueryInk[[]types.H160](
		contract,
		util.NewAccountID(signature.TestKeyringPairAlice.PublicKey),
		types.NewU128(*big.NewInt(0)),
		util.NewNone[types.Weight](),
		util.NewNone[types.U128](),
		util.InkContractInput{
			Selector: util.FuncToSelector("Member::list"),
			Args:     []any{},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	// query Gov::track return type is dao.Track
	_, err = chain.QueryInk[dao.Track](
		contract,
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
	)
	if err != nil {
		fmt.Println(err)
	}

	// dry run contract
	result, err := chain.DryRunInk(
		contract,
		util.NewAccountID(signature.TestKeyringPairAlice.PublicKey),
		types.NewU128(*big.NewInt(0)),
		util.NewNone[types.Weight](),
		util.NewNone[types.U128](),
		util.InkContractInput{
			Selector: util.FuncToSelector("Member::levae"),
			Args:     []any{},
		},
	)
	if err == nil {
		fmt.Println(result)
	}

	var testSecretSeed = "0x167d9a020688544ea246b056799d6a771e97c9da057e4d0b87024537f99177bc"
	p, err := chain.Ed25519PairFromSecret(testSecretSeed, 42)
	if err != nil {
		panic(err)
	}

	err = chain.CallInk(
		contract,
		&p,
		types.NewU128(*big.NewInt(0)),
		types.Weight{
			ProofSize: types.NewUCompact(big.NewInt(1_000_000_000)),
			RefTime:   types.NewUCompact(big.NewInt(1_00_000)),
		},
		types.NewU128(*UTIL),
		util.InkContractInput{
			Selector: util.FuncToSelector("Member::levae"),
			Args:     []any{},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
}

var UTIL = big.NewInt(1_000_000_000_000)
