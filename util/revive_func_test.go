package util

import "testing"

func TestHexToH160(t *testing.T) {
	hex := "0xd64a86424a7d175d44b604589ec278cb55e2a149"
	_, err := HexToH160(hex)
	if err != nil {
		t.Error(err)
	}

	hex2 := "d64a86424a7d175d44b604589ec278cb55e2a149"
	_, err2 := HexToH160(hex2)
	if err2 != nil {
		t.Error(err2)
	}
}
