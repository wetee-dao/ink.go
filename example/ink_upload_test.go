package example

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	chain "github.com/wetee-dao/ink.go"
	"github.com/wetee-dao/ink.go/util"
)

func TestUpload(t *testing.T) {
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

	data, err := os.ReadFile("./pod.polkavm")
	if err != nil {
		util.LogWithPurple("read file error", err)
		t.Fatal(err)
	}
	res, err := chainClient.UploadInkCode(data, &p)
	if err != nil {
		util.LogWithPurple("UploadInkCode", err)
		t.Fatal(err)
	}

	fmt.Println(res.Hex())
}

func TestInitWithCode(t *testing.T) {
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

	data, err := os.ReadFile("./pod.polkavm")
	if err != nil {
		util.LogWithPurple("read file error", err)
		t.Fatal(err)
	}

	bytes := make([]byte, 32)
	_, err = rand.Read(bytes)
	if err != nil {
		t.Fatal(err)
	}
	randomBytes := [32]byte{}
	copy(randomBytes[:], bytes)

	res, err := chainClient.DeployContract(
		util.InkCode{Upload: &data}, &p, types.NewU128(*big.NewInt(0)),
		util.InkContractInput{
			Selector: "new",
			Args: []any{
				types.NewU64(1000),
				p.H160Address(),
			},
		},
		util.NewSome(randomBytes),
	)
	if err != nil {
		util.LogWithPurple("DeployContract", err)
		t.Fatal(err)
	}

	fmt.Println(res.Hex())
}
