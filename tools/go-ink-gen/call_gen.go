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
) ({{.Return}}, *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[{{.Return}}](
		c,
		params.Origin,
		params.Amount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "{{.FuncName}}",
			Args:     []any{ {{.ArgStr}} },
		},
	)
	{{if IsResult .Return}}
	if v.IsErr {
		return *v, nil, errors.New("Contract Reverted: " + v.E.Error())
	}
	{{end}}
	return *v, gas, err
}
{{if .IsMut}}
func (c *{{$.Name}}) Call{{CamelCase .FuncName}}(
	{{.ArgTypeStr}} params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.Amount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "{{.FuncName}}",
			Args:     []any{ {{.ArgStr}} },
		},
	)
	return err
}
{{ end }}
{{ end }}
`
