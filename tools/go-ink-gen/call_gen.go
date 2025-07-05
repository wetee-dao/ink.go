package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type ContractCallBox struct {
	PackageName string
	Name        string
	Funcs       []Func
}

type Func struct {
	FuncName   string
	ArgStr     string
	ArgTypeStr string
	Return     string
	IsMut      bool
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

func Init{{.Name}}Contract(client *chain.ChainClient, address string) (chain.Ink, error) {
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
	{{.ArgTypeStr}} params chain.DryRunCallParams,
) (*{{.Return}}, *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRunInk[{{.Return}}](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "{{.FuncName}}",
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
func (c *{{$.Name}}) Call{{CamelCase .FuncName}}(
	{{.ArgTypeStr}} __ink_params chain.CallParams,
) error {
 	_param := chain.DefaultParamWithOrigin(__ink_params.Signer.AccountID())
	_param.PayAmount = __ink_params.PayAmount
	_, gas, err := c.DryRun{{CamelCase .FuncName}}({{.ArgStr}}_param)
	if err != nil {
		return err
	}
	return chain.CallInk(
		c,
		__ink_params.Signer,
		__ink_params.PayAmount,
		gas.GasRequired,
		gas.StorageDeposit,
		util.InkContractInput{
			Selector: "{{.FuncName}}",
			Args:     []any{ {{.ArgStr}} },
		},
	)
}

func (c *{{$.Name}}) TxCallOf{{CamelCase .FuncName}}(
	{{.ArgTypeStr}} __ink_params chain.CallParams,
) (*types.Call, error) {
 	_param := chain.DefaultParamWithOrigin(__ink_params.Signer.AccountID())
	_param.PayAmount = __ink_params.PayAmount
	_, gas, err := c.DryRun{{CamelCase .FuncName}}({{.ArgStr}}_param)
	if err != nil {
		return nil,err
	}
	return chain.TxCall(
		c,
		__ink_params.Signer,
		__ink_params.PayAmount,
		gas.GasRequired,
		gas.StorageDeposit,
		util.InkContractInput{
			Selector: "{{.FuncName}}",
			Args:     []any{ {{.ArgStr}} },
		},
	)
}
{{ end }}
{{ end }}
`
