package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/wetee-dao/go-sdk/util"
)

// BaseT,
// TupleT(u8, u16),
// StructT{ f1: i32, f2: i32 },
// StrT(Vec<u8>),

// type Test struct {
// 	BaseT  *bool
// 	TupleT *struct {
// 		F0 byte
// 		F1 uint16
// 	}
// 	StructT *struct {
// 		F1 int32
// 		F2 int32
// 	}
// 	StrT *[]byte
// }

type EnumBox struct {
	Name  string
	Items []EnumItem
}

type EnumItem struct {
	Name   string
	Type   string //Inline Base Tuple Struct
	Fields []EnumItemField
	Index  int
}

type EnumItemField struct {
	Name string
	Type string
}

// Enum generator
func (r *ReviveGen) EnumGen(ty int, name string, items []util.SubVariant, subs [][][]string) string {
	var traits = ""
	if r.CheckTypeIsSkip(ty, name) {
		return traits
	}

	// 为Result和Option类型添加泛型
	if name == "Result" || name == "Option" {
		if name == "Result" {
			return "[" + subs[0][0][1] + "," + subs[1][0][1] + "]"
		} else if name == "Option" {
			return "[" + subs[1][0][1] + "]"
		}
	}

	typeStr := ("type " + name + " struct { // Enum" + "\n")
	tempItems := make([]EnumItem, 0, len(items))
	for i, v := range items {
		typeStr += ("  " + v.Name)
		if len(v.Fields) == 0 { // enum base type
			typeStr += (" *bool // " + fmt.Sprint(v.Index) + "\n")
			tempItems = append(tempItems, EnumItem{
				Name:   v.Name,
				Type:   "Base",
				Fields: []EnumItemField{},
				Index:  v.Index,
			})
			continue
		} else if len(v.Fields) == 1 && subs[i][0][2] == "Primitive" { // inline type
			typeStr += (" *" + subs[i][0][1] + " // " + fmt.Sprint(v.Index) + "\n")
			tempItems = append(tempItems, EnumItem{
				Name:   v.Name,
				Type:   "Inline",
				Fields: []EnumItemField{},
				Index:  v.Index,
			})
		} else { // multiple fields
			typeStr += (" *struct{ // " + fmt.Sprint(v.Index) + "\n")
			itemType := "Struct"
			itemFields := make([]EnumItemField, 0, len(v.Fields))
			for j, subfield := range v.Fields {
				typeName := UnderscoreToCamelCase(subfield.Name)
				if subfield.Name == "" { // Tuple
					typeName = "F" + fmt.Sprint(j)
					itemType = "Tuple"
				}
				typeStr += ("    " + typeName + " " + subs[i][j][1] + "\n")
				itemFields = append(itemFields, EnumItemField{
					Name: typeName,
					Type: subs[i][j][1],
				})
			}
			typeStr += ("  }\n")
			tempItems = append(tempItems, EnumItem{
				Name:   v.Name,
				Type:   itemType,
				Fields: itemFields,
				Index:  v.Index,
			})
		}
	}
	typeStr += ("}\n")

	p := EnumBox{
		Name:  name,
		Items: tempItems,
	}
	t := template.Must(template.New("scale").Parse(enumScaleTemp))

	var result bytes.Buffer
	t.Execute(&result, p)

	r.TypeResult[ty] = typeStr + result.String()

	return ""
}

var enumScaleTemp = `func (ty {{.Name}}) Encode(encoder scale.Encoder) (err error) {
{{range $outerIndex, $outerItem := .Items}}
	{{if eq .Type "Base"}}
	if ty.{{.Name}} != nil {
		err = encoder.PushByte({{.Index}})
		if err != nil {
			return err
		}
		return nil
	}
	{{end}}
	{{if eq .Type "Tuple"}}
	if ty.{{.Name}} != nil {
		err = encoder.PushByte({{.Index}})
		if err != nil {
			return err
		}
		{{range .Fields}}
		err = encoder.Encode(ty.{{$outerItem.Name}}.{{.Name}})
		if err != nil {
			return err
		}
		{{end}}
		return nil
	}
	{{end}}
	{{if eq .Type "Struct"}}
	if ty.{{.Name}} != nil {
		err = encoder.PushByte({{.Index}})
		if err != nil {
			return err
		}
		{{range .Fields}}
		err = encoder.Encode(ty.{{$outerItem.Name}}.{{.Name}})
		if err != nil {
			return err
		}
		{{end}}
		return nil
	}
	{{end}}
	{{if eq .Type "Inline"}}
	if ty.{{.Name}} != nil {
		err = encoder.PushByte({{.Index}})
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.{{.Name}})
		if err != nil {
			return err
		}
		return nil
	}
	{{end}}
{{end}}
	return fmt.Errorf("unrecognized enum")
}

func (ty *{{.Name}}) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
{{range $outerIndex, $outerItem := .Items}}
	{{if eq .Type "Base"}}
	case {{.Index}}:
		t := true
		ty.{{.Name}} = &t
		return
	{{end}}
	{{if eq .Type "Tuple"}}
	case {{.Index}}:
		ty.{{.Name}} = &struct {
			{{range .Fields}}
			{{.Name}} {{.Type}}
			{{end}}
		}{}

		{{range .Fields}}
		err = decoder.Decode(&ty.{{$outerItem.Name}}.{{.Name}})
		if err != nil {
			return err
		}
		{{end}}
		
		return
	{{end}}
	{{if eq .Type "Struct"}}
	case {{.Index}}:
		ty.{{.Name}} = &struct {
			{{range .Fields}}
			{{.Name}} {{.Type}}
			{{end}}
		}{}

		{{range .Fields}}
		err = decoder.Decode(&ty.{{$outerItem.Name}}.{{.Name}})
		if err != nil {
			return err
		}
		{{end}}

		return
	{{end}}
	{{if eq .Type "Inline"}}
	case {{.Index}}:
		err = decoder.Decode(ty.{{.Name}})
		if err != nil {
			return err
		}
		return
	{{end}}
{{end}}
	default:
		return fmt.Errorf("unrecognized enum")
	}
}
`
