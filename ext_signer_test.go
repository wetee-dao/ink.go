package ink

import (
	"fmt"
	"testing"
)

func TestEd25519(t *testing.T) {
	var testSecretSeed = "0x167d9a020688544ea246b056799d6a771e97c9da057e4d0b87024537f99177bc"
	p, err := Ed25519PairFromSecret(testSecretSeed, 42)
	if err != nil {
		panic(err)
	}

	var body = []byte("<Bytes>hello<Bytes>")
	sig, err := p.Sign(body)
	if err != nil {
		panic(err)
	}
	if !p.Verify(body, sig) {
		panic("verify failed")
	}

	PrintBytes(p.PublicKey)
	PrintBytes(body)
	PrintBytes(sig)
}

func TestEr25519(t *testing.T) {
	// p := signature.TestKeyringPairAlice
	var testSecretSeed = "0x167d9a020688544ea246b056799d6a771e97c9da057e4d0b87024537f99177bc"
	p, err := Sr25519PairFromSecret(testSecretSeed, 42)
	if err != nil {
		panic(err)
	}

	var body = []byte("<Bytes>hello<Bytes>")

	sig, err := p.Sign(body)
	if err != nil {
		panic(err)
	}

	if !p.Verify(body, sig) {
		panic("verify failed")
	}

	PrintBytes(p.PublicKey)
	PrintBytes(body)
	PrintBytes(sig)
}

func PrintBytes(b []byte) {
	for i := range b {
		fmt.Print(fmt.Sprint(b[i]) + ",")
	}
	fmt.Println()
}
