package util

import (
	"encoding/json"
	"fmt"
)

func LogWithRed(tag string, a ...any) {
	b := make([]any, 0, len(a)+1)
	b = append(b, tag+": ")
	b = append(b, a...)
	fmt.Println(b...)
}

func PrintJson(v any) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println("------------------ json --------------------------")
	fmt.Println(string(b))
	fmt.Println("------------------ json --------------------------")
}
