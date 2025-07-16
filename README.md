# Golang client for ink! v6 contract
> Base on github.com/centrifuge/go-substrate-rpc-client

## Installation
### 1. Install the code generation tool for ink! contracts
used to generate contract call code through ABI files.
```
go install github.com/wetee-dao/ink.go/tools/go-ink-gen@latest
```

### 2. Install the ink.go package
```
go get github.com/wetee-dao/ink.go
```

## Code generation
Through go-ink-gen, we can directly generate smart contract code in Golang from ABI files.
For example
```
go-ink-gen -json xxx.json
```
The files will be divided into `types.go` and `calls.go` in the directory named after the contract name.

In `types.go`. All types, including Error, are automatically converted into golang structs.
For example, Complete example types.go](https://github.com/wetee-dao/ink.go/blob/main/example/contracts/dao/types.go)
```go
type AccountId = [32]byte
type Error struct {
	NotEnoughBalance        *bool
	MustCallByMainContract  *bool
}
```

In the calls.go file, it contains all the Query, DryRun, and Call functions.
For example, Complete example calls.go](https://github.com/wetee-dao/ink.go/blob/main/example/contracts/dao/calls.go)
```go
func (c *Dao) QueryMemberList(
	params chain.DryRunParams,
) ([]types.H160, *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[[]types.H160](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::list",
			Args:     []any{},
		},
	)
	if err != nil && !errors.Is(err, chain.ContractReverted) {
		return *v, nil, err
	}
	return *v, gas, nil
}

```

## Query contract data
After the complete code is generated, users can quickly complete the contract invocation.
For example
```go
// Step1: connect to chain
chainClient, err := chain.ClientInit("ws://127.0.0.1:9944", false)
if err != nil {
    panic(err)
}


// Step2: initialize the account
var testSecretSeed = "0x167d9a020688544ea246b056799d6a771e97c9da057e4d0b87024537f99177bc"
p, err := chain.Ed25519PairFromSecret(testSecretSeed, 42)
if err != nil {
    panic(err)
}

// Step3: initialize contract 
contract,_ := dao.InitDaoContract(chainClient,"0x1547E25E7fe95a931E96907C70529d57D2438aD1")

// Step4: query contract data
_, _, err = contract.QueryMemberList(
    chain.DryRunParams{
        Origin:              util.NewAccountID(p.PublicKey),
        PayAmount:           types.NewU128(*big.NewInt(0)),
        GasLimit:            util.NewNone[types.Weight](),
        StorageDepositLimit: util.NewNone[types.U128](),
    },
)
if err != nil {
    fmt.Println(err)
}
```

## Dry-run contract 
Similar queries can also complete the DryRun steps in 4 steps.
For example
```go
// Step1: connect to chain
chainClient, err := chain.ClientInit("ws://127.0.0.1:9944", false)
if err != nil {
    panic(err)
}


// Step2: initialize the account
var testSecretSeed = "0x167d9a020688544ea246b056799d6a771e97c9da057e4d0b87024537f99177bc"
p, err := chain.Ed25519PairFromSecret(testSecretSeed, 42)
if err != nil {
    panic(err)
}

// Step3: initialize contract 
contract,_ := dao.InitDaoContract(chainClient,"0x1547E25E7fe95a931E96907C70529d57D2438aD1")

// Step4: dry-run contract
result, gas, err := contract.DryRunMemberPublicJoin(
	chain.DefaultParamWithOrigin(p.AccountID()),
)
if err == nil {
	fmt.Println(result.E)
} else {
	fmt.Println(err)
	return
}
```

## Call contract
```go
// Step1: connect to chain
chainClient, err := chain.ClientInit("ws://127.0.0.1:9944", false)
if err != nil {
    panic(err)
}

// Step2: initialize the account
var testSecretSeed = "0x167d9a020688544ea246b056799d6a771e97c9da057e4d0b87024537f99177bc"
p, err := chain.Ed25519PairFromSecret(testSecretSeed, 42)
if err != nil {
    panic(err)
}

// Step3: initialize contract 
contract,_ := dao.InitDaoContract(chainClient,"0x1547E25E7fe95a931E96907C70529d57D2438aD1")

// Step4: call contract
err = contract.ExecMemberPublicJoin(
    chain.ExecParams{
        Signer:    &p,
        PayAmount: types.NewU128(*big.NewInt(0)),
    },
)
if err != nil {
	fmt.Println(err)
}
```