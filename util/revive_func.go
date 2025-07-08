package util

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

// Get selector of contract function
func FuncToSelector(f string) [4]byte {
	if strings.HasPrefix(f, "0x") {
		f = f[2:]
		hex, err := hex.DecodeString(f)
		if err != nil {
			fmt.Println("Error DecodeString hash:", err)
			return [4]byte{}
		}
		return [4]byte(hex[:4])
	}

	inputBytes := []byte(f)

	// Create a new BLAKE2s hash with default settings to produce a 32-byte hash
	hash, err := blake2b.New256(nil)
	if err != nil {
		fmt.Println("Error creating BLAKE2s hash:", err)
		return [4]byte{}
	}

	// Write the data to the hash
	_, err = hash.Write(inputBytes)
	if err != nil {
		fmt.Println("Error writing to hash:", err)
		return [4]byte{}
	}

	// Sum the hash
	hashSum := hash.Sum(nil)

	return [4]byte(hashSum[:4])
}

// Convert a h160 hex string to a byte array
func HexToH160(hexString string) (types.H160, error) {
	// Remove the "0x" prefix if it exists
	if strings.HasPrefix(hexString, "0x") {
		hexString = hexString[2:]
	}

	// Decode the hexadecimal string to a byte slice
	src := []byte(hexString)
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return [20]byte{}, err
	}

	// Copy the byte slice to a fixed-size byte array
	var byteArray [20]byte
	copy(byteArray[:20], dst[:])

	return byteArray, nil
}

func H160FromPublicKey(bytes []byte) (types.H160, error) {
	if len(bytes) < 20 {
		return types.H160{}, errors.New("invalid byte array length")
	}

	if IsEthDerived(bytes) {
		var byteArray [20]byte
		copy(byteArray[:20], bytes[:])

		return byteArray, nil
	}

	account_hash := Keccak256Hash(bytes)

	var byteArray [20]byte
	copy(byteArray[:20], account_hash[12:])

	return byteArray, nil
}

func Keccak256Hash(data []byte) []byte {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	return hasher.Sum(nil)
}

var eth = []byte{0xEE, 0xEE, 0xEE, 0xEE, 0xEE, 0xEE, 0xEE, 0xEE, 0xEE, 0xEE, 0xEE, 0xEE}

func IsEthDerived(account_bytes []byte) bool {
	return hex.EncodeToString(account_bytes[:20]) == hex.EncodeToString(eth)
}
