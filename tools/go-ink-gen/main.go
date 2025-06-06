package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wetee-dao/ink.go/util"
)

func main() {
	json := flag.String("json", "", "contract ABI json file")
	flag.Parse() // 解析命令行参数

	data, err := os.ReadFile("./" + *json)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	revice, err := NewReviveGen(
		data,
	)
	if err != nil {
		util.LogWithPurple("NewRevive", err)
		return
	}

	revice.SaveTypes()
}
