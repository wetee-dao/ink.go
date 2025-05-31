package util

import (
	"fmt"
)

func LogWithRed(tag string, a ...any) {
	b := make([]any, 0, len(a)+1)
	b = append(b, tag+": ")
	b = append(b, a...)
	fmt.Println(b...)
}
