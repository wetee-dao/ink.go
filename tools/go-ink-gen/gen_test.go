package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/wetee-dao/ink.go/util"
)

func TestGen(t *testing.T) {

	data, err := os.ReadFile("../../ink/dao.json")
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
