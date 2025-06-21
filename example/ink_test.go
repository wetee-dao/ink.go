package example_test

import (
	"fmt"
	"math/big"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	chain "github.com/wetee-dao/ink.go"

	"github.com/wetee-dao/ink.go/example/contracts/dao"
	"github.com/wetee-dao/ink.go/util"
)

func ExampleInk() {
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
	contract := dao.Dao{
		ChainClient: chainClient,
		Address:     contractAddress,
	}

	// query Member::list
	_, _, err = contract.QueryMemberList(
		chain.DryRunCallParams{
			Origin:              util.NewAccountID(p.PublicKey),
			Amount:              types.NewU128(*big.NewInt(0)),
			GasLimit:            util.NewNone[types.Weight](),
			StorageDepositLimit: util.NewNone[types.U128](),
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	// query Gov::track return type is dao.Track
	_, _, err = contract.QueryGovTrack(
		0,
		chain.DryRunCallParams{
			Origin:              util.NewAccountID(p.PublicKey),
			Amount:              types.NewU128(*big.NewInt(0)),
			GasLimit:            util.NewNone[types.Weight](),
			StorageDepositLimit: util.NewNone[types.U128](),
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	// dry run contract
	result, _, err := contract.DryRunErc20EnableTransfer(
		chain.DryRunCallParams{
			Origin:              util.NewAccountID(p.PublicKey),
			Amount:              types.NewU128(*big.NewInt(0)),
			GasLimit:            util.NewNone[types.Weight](),
			StorageDepositLimit: util.NewNone[types.U128](),
		},
	)
	if err == nil {
		fmt.Println(result.E)
	}

	err = contract.CallErc20EnableTransfer(
		chain.CallParams{
			Signer: &p,
			Amount: types.NewU128(*big.NewInt(0)),
			GasLimit: types.Weight{
				RefTime:   types.NewUCompact(big.NewInt(1_100_000_000)),
				ProofSize: types.NewUCompact(big.NewInt(100_000)),
			},
			StorageDepositLimit: types.NewU128(*big.NewInt(110_000_000_000)),
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
}
