package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"

	chain "github.com/wetee-dao/go-sdk"
	gtypes "github.com/wetee-dao/go-sdk/gen/types"
)

func main() {
	client, err := chain.ClientInit("ws://xiaobai.asyou.me:30002", true)
	if err != nil {
		panic(err)
	}

	var testSecretSeed = "0x167d9a020688544ea246b056799d6a771e97c9da057e4d0b87024537f99177bc"
	p, err := signature.KeyringPairFromSecret(testSecretSeed, 42)
	if err != nil {
		panic(err)
	}

	worker := &chain.Worker{
		Client: client,
		Signer: &p,
	}

	name := "XXXXX"
	domain := "wetee.test"
	ipstrs := strings.Split("127.0.0.1", ".")
	if len(ipstrs) != 4 {
		printErrorStack(err)
		return
	}
	iparr := []uint8{}
	for _, ipstr := range ipstrs {
		i, err := strconv.Atoi(ipstr)
		if err != nil {
			printErrorStack(err)
			return
		}
		iparr = append(iparr, uint8(i))
	}

	// iparr
	err = worker.ClusterRegister(name, []gtypes.Ip{
		{
			Ipv4: gtypes.OptionTUint32{
				IsNone: true,
			},

			Ipv6: gtypes.OptionTU128{
				IsNone: true,
			},
			Domain: gtypes.OptionTByteSlice{
				IsSome:       true,
				AsSomeField0: []byte(domain),
			},
		},
	}, uint32(30000), uint8(1), true)

	printErrorStack(err)
}

func printErrorStack(err error) {
	if err != nil {
		pc, file, line, ok := runtime.Caller(1)
		if !ok {
			fmt.Printf("runtime.Caller error\n")
			return
		}
		funcName := runtime.FuncForPC(pc).Name()
		fmt.Printf("Error: %s\nFile: %s\nLine: %d\nFunction: %s\n",
			err.Error(), file, line, funcName)
	}
}
