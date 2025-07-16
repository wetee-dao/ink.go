package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type ContractCallBox struct {
	PackageName  string
	Name         string
	Funcs        []Func
	Constructors []Constructor
}

type Func struct {
	Selector   string
	FuncName   string
	ArgStr     string
	ArgTypeStr string
	Return     string
	IsMut      bool
}

type Constructor struct {
	Selector   string
	FuncName   string
	ArgStr     string
	ArgTypeStr string
	Return     string
}

func callGen(callData ContractCallBox) []byte {
	t := template.Must(template.New("call").Funcs(template.FuncMap{
		"CamelCase": func(v string) string {
			vsplit := strings.Split(v, "::")
			for i := range vsplit {
				vsplit[i] = UnderscoreToCamelCase(vsplit[i])
			}

			return strings.Join(vsplit, "")
		},
		"IsResult": func(v string) bool {
			return strings.HasPrefix(v, "util.Result[")
		},
	}).Parse(callTemp))
	var result bytes.Buffer
	err := t.Execute(&result, callData)
	fmt.Println(err)

	return result.Bytes()
}

var callTemp = `package {{.PackageName}}
import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	chain "github.com/wetee-dao/ink.go"
	"github.com/wetee-dao/ink.go/util"
)

{{ range .Constructors }}
func Deploy{{$.Name}}With{{CamelCase .FuncName}}({{.ArgTypeStr}} __ink_params chain.DeployParams) (*types.H160, error) {
	return __ink_params.Client.DeployContract(
		__ink_params.Code, __ink_params.Signer, types.NewU128(*big.NewInt(0)),
		util.InkContractInput{
			Selector: "{{.Selector}}",
			Args: []any{ {{.ArgStr}} },
		},
		__ink_params.Salt,
	)
}
{{ end }}

func Init{{.Name}}Contract(client *chain.ChainClient, address string) (*{{.Name}}, error) {
	contractAddress, err := util.HexToH160(address)
	if err != nil {
		return nil, err
	}
	return &{{.Name}}{
		ChainClient: client,
		Address:     contractAddress,
	}, nil
}

type {{.Name}} struct {
	ChainClient *chain.ChainClient
	Address     types.H160
}

func (c *{{.Name}}) Client() *chain.ChainClient {
	return c.ChainClient
}

func (c *{{.Name}}) ContractAddress() types.H160 {
	return c.Address
}

{{ range .Funcs }}
func (c *{{$.Name}}) {{if .IsMut}}DryRun{{else}}Query{{end}}{{CamelCase .FuncName}}(
	{{.ArgTypeStr}} __ink_params chain.DryRunParams,
) (*{{.Return}}, *chain.DryRunReturnGas, error) {
 	if c.ChainClient.Debug {
		fmt.Println()
		util.LogWithPurple("[ DryRun   method ]", "{{.FuncName}}")
	}
	v, gas, err := chain.DryRunInk[{{.Return}}](
		c,
		__ink_params.Origin,
		__ink_params.PayAmount,
		__ink_params.GasLimit,
		__ink_params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "{{.Selector}}",
			Args:     []any{ {{.ArgStr}} },
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	{{- if IsResult .Return}}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}
	{{end}}
	return v, gas, nil
}
{{if .IsMut}}
func (c *{{$.Name}}) Exec{{CamelCase .FuncName}}(
	{{.ArgTypeStr}} __ink_params chain.ExecParams,
) error {
 	_param := chain.DefaultParamWithOrigin(__ink_params.Signer.AccountID())
	_param.PayAmount = __ink_params.PayAmount
	_, gas, err := c.DryRun{{CamelCase .FuncName}}({{.ArgStr}}_param)
	if err != nil {
		return err
	}
	return chain.CallInk(
		c,
		gas.GasRequired,
		gas.StorageDeposit,
		util.InkContractInput{
			Selector: "{{.Selector}}",
			Args:     []any{ {{.ArgStr}} },
		},
		__ink_params,
	)
}

func (c *{{$.Name}}) CallOf{{CamelCase .FuncName}}(
	{{.ArgTypeStr}} __ink_params chain.DryRunParams,
) (*types.Call, error) {
	_, gas, err := c.DryRun{{CamelCase .FuncName}}({{.ArgStr}}__ink_params)
	if err != nil {
		return nil,err
	}
	return chain.CallOfTransaction(
		c,
		__ink_params.PayAmount,
		gas.GasRequired,
		gas.StorageDeposit,
		util.InkContractInput{
			Selector: "{{.Selector}}",
			Args:     []any{ {{.ArgStr}} },
		},
	)
}
{{ end }}
{{ end }}
`
