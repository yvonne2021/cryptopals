package col1

import "encoding/hex"

/*
requirement:
https://cryptopals.com/sets/1/challenges/3
plaintext -> bytes ->  xor with single byte key -> bytes -> encoded hex(given hex string)
*/

func singByteXor(hexSource string, key byte) string {
	// decode hex to bytes, as we always operate on raw bytes
	bytes, _ := hex.DecodeString(hexSource)

	// xor bytes
	bytesResult := make([]byte, len(bytes))
	for i := range bytes {
		bytesResult[i] = bytes[i] ^ key
	}
	return string(bytesResult)
}
