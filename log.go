package client

import (
	"fmt"
)

func LogWithRed(tag string, a ...interface{}) {
	b := make([]interface{}, 0, len(a)+1)
	b = append(b, tag+":")
	b = append(b, a...)
	fmt.Println(b...)
}
