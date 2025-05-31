package util

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/blake2b"
)

// Get selector of contract function
func FuncToSelector(f string) [4]byte {
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
func HexToH160(hexString string) (H160, error) {
	// Remove the "0x" prefix if it exists
	if hexString[:2] == "0x" {
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
