package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
	"unicode"

	"github.com/wetee-dao/go-sdk/util"
)

// Revive module
type ReviveGen struct {
	Abi        *util.InkAbi
	TypeMap    map[int]util.AbiSubType
	TypeResult map[int]string
}

func NewReviveGen(abiRaw []byte) (*ReviveGen, error) {
	abi, err := util.InitAbi(abiRaw)
	if err != nil {
		return nil, err
	}

	var typeMap = map[int]util.AbiSubType{}
	for _, t := range abi.Types {
		typeMap[t.Id] = t.Type
	}

	return &ReviveGen{
		Abi:        abi,
		TypeMap:    typeMap,
		TypeResult: map[int]string{},
	}, nil
}

func (r *ReviveGen) SaveTypes() {
	for i, t := range r.Abi.Spec.Messages {
		// if t.Label != "Erc20::balance_of" {
		// 	continue
		// }

		util.LogWithYellow("--------------------------------------------------" + t.Label)
		for _, arg := range r.Abi.Spec.Messages[i].Args {
			fmt.Print("arg: " + arg.Label + " ")
			r.RecursionTypes(arg.Type.Type, arg.Type.DisplayName[len(arg.Type.DisplayName)-1], 1)
			fmt.Println("")
		}

		r.RecursionTypes(t.ReturnType.Type, t.ReturnType.DisplayName[len(t.ReturnType.DisplayName)-1], 1)
	}

	name := r.Abi.Contract.Name

	var data = "package " + name + "\n"
	data += "import (\n"
	data += "  \"github.com/wetee-dao/go-sdk/util\"\n"
	data += "  \"github.com/centrifuge/go-substrate-rpc-client/v4/types\"\n"
	data += "  \"github.com/centrifuge/go-substrate-rpc-client/v4/scale\"\n"
	data += ")\n"

	var keys []int
	for k := range r.TypeResult {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, t := range keys {
		data += r.TypeResult[t]
	}

	if err := os.MkdirAll("./"+name, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	formattedData, err := format.Source([]byte(data))
	if err != nil {
		log.Fatal(err)
	}

	formattedData, err = formatAndCleanCode(formattedData)
	if err != nil {
		log.Fatalf("Error formatting and cleaning code: %v", err)
	}

	err = os.WriteFile("./"+name+"/types.go", formattedData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *ReviveGen) RecursionTypes(ty int, name string, level int) []string {
	path := r.TypeMap[ty].Path
	fields := [][]string{}
	def := r.TypeMap[ty].Def
	curtype := r.GetType(r.TypeMap[ty])

	name = UnderscoreToCamelCase(name)
	cpath := ""
	if len(path) > 0 {
		cpath = " [" + path[len(path)-1] + "]"
	}

	var typeStr = curtype
	if typeStr == "Primitive" {
		typeStr += " " + fmt.Sprint(*def.Primitive)
	}
	fmt.Println(getLevelSpace(level) + name + cpath + " (" + typeStr + ")")

	returnType := ""
	returnTraits := ""
	if def.Composite != nil {
		for _, v := range def.Composite.Fields {
			f := r.RecursionTypes(v.Type, v.Name, level+1)
			fields = append(fields, f)
		}
	} else if def.Variant != nil {
		enums := [][][]string{}
		for _, v := range def.Variant.Variants {
			fmt.Println(getLevelSpace(level+1) + v.Name)
			enumItem := [][]string{}
			for i := 0; i < len(v.Fields); i++ {
				subfield := v.Fields[i]
				f := r.RecursionTypes(subfield.Type, subfield.Name, level+2)
				enumItem = append(enumItem, f)
			}
			enums = append(enums, enumItem)
		}

		var enumName = name
		if len(path) > 0 {
			enumName = path[len(path)-1]
		}
		returnTraits = r.EnumGen(ty, enumName, def.Variant.Variants, enums)
	} else if def.Sequence != nil {
		f := r.RecursionTypes(def.Sequence.Type, "", level+1)
		f[0] = name
		f[1] = "[]" + f[1]
		return f
	} else if def.Array != nil {
		f := r.RecursionTypes(def.Array.Type, "Length "+fmt.Sprint(def.Array.Len), level+1)
		f[0] = name
		f[1] = "[" + fmt.Sprint(def.Array.Len) + "]" + f[1]
		return f
	} else if def.Tuple != nil {
		for i, v := range *def.Tuple {
			f := r.RecursionTypes(v, "F"+fmt.Sprint(i), level+1)
			fields = append(fields, f)
		}
	} else if def.Range != nil {
		util.LogWithRed(getLevelSpace(level) + "Range")
	} else if def.Compact != nil {
		util.LogWithRed(getLevelSpace(level) + "Compact")
	} else if def.BitSequence != nil {
		util.LogWithRed(getLevelSpace(level) + "BitSequence")
	} else if def.Primitive != nil {
		returnType = primitiveMapping[fmt.Sprint(*def.Primitive)]
	} else {
		util.LogWithRed(getLevelSpace(level) + "unknown")
	}

	typeName := ""
	if len(path) > 0 {
		typeName = path[len(path)-1]
	}
	if typeName == "" && curtype == "Tuple" {
		typeName = "Tuple_" + fmt.Sprint(ty)
	}

	if !r.CheckTypeIsSkip(ty, typeName) {
		var typeStr = ""
		if len(fields) > 1 {
			typeStr += ("type " + typeName + " struct {  // " + curtype + "\n")
			for _, v := range fields {
				typeStr += ("  " + v[0] + "   " + v[1] + "\n")
			}
			typeStr += ("}" + "\n")
		} else if len(fields) == 1 && fields[0][0] == "" {
			typeStr += ("type " + typeName + " = " + fields[0][1] + "  // " + curtype + "\n")
		} else if len(fields) == 1 && fields[0][0] != "" {
			typeStr += ("type " + typeName + " struct {  // " + curtype + "\n")
			for _, v := range fields {
				typeStr += ("  " + v[0] + "   " + v[1] + "\n")
			}
			typeStr += ("}" + "\n")
		}
		r.TypeResult[ty] = typeStr
	}

	if returnType == "" && len(path) > 0 {
		returnType = path[len(path)-1]
	}

	return []string{
		name, addTypePrefix(returnType) + returnTraits, curtype,
	}
}

// Get type of AbiSubType
func (r *ReviveGen) GetType(data util.AbiSubType) string {
	def := data.Def
	if def.Composite != nil {
		return "Composite"
	} else if def.Variant != nil {
		return "Variant|Enum"
	} else if def.Sequence != nil {
		return "Sequence|Vec"
	} else if def.Array != nil {
		return "Array"
	} else if def.Tuple != nil {
		return "Tuple"
	} else if def.Range != nil {
		return "Range"
	} else if def.Compact != nil {
		return "Compact"
	} else if def.BitSequence != nil {
		return "BitSequence"
	} else if def.Primitive != nil {
		return "Primitive"
	} else {
		return "Unknown"
	}
}

func (r *ReviveGen) CheckTypeIsSkip(ty int, name string) bool {
	name = strings.ReplaceAll(name, " ", "")
	_, ok := r.TypeResult[ty]
	if ok {
		return true
	}

	_, skip := skipTypes[name]

	return skip
}

func addTypePrefix(ty string) string {
	t, ok := typePrefix[ty]
	if ok {
		ty = t + "." + ty
	}
	return ty
}

// Get level space string
func getLevelSpace(level int) string {
	str := ""
	for i := 0; i < level-1; i++ {
		str += "  "
	}
	return str + "--"
}

// Convert underscore to camel case
func UnderscoreToCamelCase(s string) string {
	parts := strings.Split(s, "_")
	var result string
	for _, part := range parts {
		if len(part) > 0 {
			runes := []rune(part)
			runes[0] = unicode.ToUpper(runes[0])
			result += string(runes)
		}
	}
	return result
}

func formatAndCleanCode(src []byte) ([]byte, error) {
	// Use goimports to format the code, remove unused imports, and add missing ones.
	cmd := exec.Command("goimports")
	cmd.Stdin = bytes.NewReader(src)
	var out bytes.Buffer
	cmd.Stdout = &out
	var errBuf bytes.Buffer
	cmd.Stderr = &errBuf

	err := cmd.Run()
	if err != nil {
		errMsg := errBuf.String()
		if errMsg != "" {
			return nil, fmt.Errorf("goimports error: %s", errMsg)
		}
		return nil, fmt.Errorf("goimports failed: %w", err)
	}

	return out.Bytes(), nil
}

// ink type to go type
var primitiveMapping = map[string]string{
	"bool":   "bool",
	"i8":     "int8",
	"i16":    "int16",
	"i32":    "int32",
	"i64":    "int64",
	"i128":   "big.Int",
	"u8":     "byte",
	"u16":    "uint16",
	"u32":    "uint32",
	"u64":    "uint64",
	"u128":   "big.Int",
	"String": "string",
	"&str":   "string",
}

var skipTypes = map[string]bool{
	"U256":      true,
	"H256":      true,
	"H160":      true,
	"LangError": true,
}

var typePrefix = map[string]string{
	"U256":   "types",
	"H256":   "types",
	"H160":   "types",
	"Result": "util",
	"Option": "util",
}
