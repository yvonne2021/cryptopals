package col1

import (
	"encoding/hex"
)

func repeatingKeyXor(s string, key string) string {
	// convert string to bytes
	bytes := []byte(s)
	keyBytes := []byte(key)

	// xor bytes
	bytesResult := make([]byte, len(bytes))
	for i := range bytes {
		bytesResult[i] = bytes[i] ^ keyBytes[i%len(keyBytes)]
	}

	// encode bytes to hex
	return hex.EncodeToString(bytesResult)
}

//func breakRepeatingKeyXor(bytes []byte) string {
//	// xor bytes
//	bytesResult := make([]byte, len(bytes))
//	for i := range bytes {
//		bytesResult[i] = bytes[i] ^ byte(0x05)
//	}
//
//	// encode bytes to hex
//	return hex.EncodeToString(bytesResult)
//}

func getHammingDistance(s1 string, s2 string) int {
	//TODO: check if s1 and s2 have the same length

	// convert string to bytes
	bytes1 := []byte(s1)
	bytes2 := []byte(s2)

	// xor bytes
	bytesResult := make([]byte, len(bytes1))
	for i := range bytes1 {
		bytesResult[i] = bytes1[i] ^ bytes2[i]
	}

	// count number of 1s in binary representation
	// and that will be the number of different bits
	count := 0
	for _, b := range bytesResult {
		// The following uses Brian Kernighan's algorithm to efficiently count the number of set bits in an integer.
		// It's more efficient than simply checking each bit individually, especially for larger integers or strings,
		// because it skips over the bits that are 0 and only operates on the ones that are 1.
		for b != 0 {
			// bitwise AND (&) operator
			// a & b = 1; only if a = b = 1 else = 0
			// e.g. 0011 & 0101 = 0001
			b &= b - 1
			count++
		}
	}
	return count
}
