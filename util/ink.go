package util

import (
	"encoding/json"
)

type InkAbi struct {
	Spec    Spec      `json:"spec"`
	Types   []AbiType `json:"types"`
	Version int       `json:"version,omitempty"`
}

type Spec struct {
	Constructors []Message           `json:"constructors"`
	Docs         []any               `json:"docs"`
	LangError    TypeWithDisplayName `json:"lang_error"`
	Events       []SpecEvent         `json:"events"`
	Messages     []Message           `json:"messages"`
}

type SpecEvent struct {
	Args  []EventArg `json:"args"`
	Docs  []string   `json:"docs"`
	Label string     `json:"label"`
}
type EventArg struct {
	Docs    []string            `json:"docs"`
	Indexed bool                `json:"indexed"`
	Payable bool                `json:"payable"`
	Type    TypeWithDisplayName `json:"type"`
}

type TypeWithDisplayName struct {
	DisplayName []string `json:"displayName"`
	Type        int      `json:"type"`
}

type Message struct {
	Args       []MessageArg        `json:"args"`
	Docs       []string            `json:"docs"`
	Label      string              `json:"label"`
	Payable    bool                `json:"payable"`
	ReturnType TypeWithDisplayName `json:"returnType"`
	Selector   string              `json:"selector"`
}

type MessageArg struct {
	Label string              `json:"label"`
	Type  TypeWithDisplayName `json:"type"`
}

type AbiType struct {
	Id   int        `json:"id"`
	Type AbiSubType `json:"type"`
}

type AbiSubType struct {
	Path []string `json:"path,omitempty"`
	Def  TypeDef  `json:"def"`
	Docs []string `json:"docs,omitempty"`
}

type TypeDef struct {
	Composite   *TypeDefComposite   `json:"composite,omitempty"`
	Variant     *TypeDefVariant     `json:"variant,omitempty"`
	Sequence    *TypeDefSequence    `json:"sequence,omitempty"`
	Array       *TypeDefArray       `json:"array,omitempty"`
	Tuple       *TypeDefTuple       `json:"tuple,omitempty"`
	Primitive   *TypeDefPrimitive   `json:"primitive,omitempty"`
	Compact     *TypeDefCompact     `json:"compact,omitempty"`
	BitSequence *TypeDefBitSequence `json:"bitSequence,omitempty"`
	Range       *TypeDefRange       `json:"range,omitempty"`
}

type TypeDefComposite struct {
	Fields []SubField `json:"fields"`
}

type SubField struct {
	Name     string   `json:"name,omitempty"`
	Type     int      `json:"type"`
	TypeName string   `json:"typeName"`
	Docs     []string `json:"docs,omitempty"`
}

type TypeDefRange struct {
	Start     int  `json:"start"`
	End       int  `json:"end"`
	Inclusive bool `json:"inclusive"`
}

type TypeDefVariant struct {
	Variants []SubVariant `json:"variants"`
}

type SubVariant struct {
	Name   string     `json:"name"`
	Fields []SubField `json:"fields"`
	Index  int        `json:"index"`
	Docs   []string   `json:"docs,omitempty"`
}

type TypeDefSequence struct {
	Type int `json:"type"`
}

type TypeDefArray struct {
	Len  int `json:"len"`
	Type int `json:"type"`
}

type TypeDefTuple []int

type TypeDefPrimitive string

type TypeDefCompact struct {
	Type int `json:"type"`
}

type TypeDefBitSequence struct {
	BitStoreType int `json:"bitStoreType"`
	BitOrderType int `json:"bitOrderType"`
}

func InitAbi(raw []byte) (*InkAbi, error) {
	var abi InkAbi
	err := json.Unmarshal(raw, &abi)
	if err != nil {
		return nil, err
	}
	return &abi, nil
}
