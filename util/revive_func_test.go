package util

import "testing"

func TestHexToH160(t *testing.T) {
	hex := "0x0c17c8bf3e4054632f59c2ea44a7efce60804642"
	_, err := HexToH160(hex)
	if err != nil {
		t.Error(err)
	}

	hex2 := "0c17c8bf3e4054632f59c2ea44a7efce60804642"
	_, err2 := HexToH160(hex2)
	if err2 != nil {
		t.Error(err2)
	}
}
