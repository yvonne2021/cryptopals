package col1

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

// convert hex to base64
func hexToBase64(source string) string {
	// decode hex to bytes
	bytes, _ := hex.DecodeString(source)

	// encode bytes to base64
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func hexXor(hexSourceA, hexSourceB string) (string, error) {
	if len(hexSourceA) != len(hexSourceB) {
		return "", errors.New("hex strings must be of equal length")
	}

	// decode hex to bytes
	bytesA, _ := hex.DecodeString(hexSourceA)
	bytesB, _ := hex.DecodeString(hexSourceB)

	// xor bytes
	bytesResult := make([]byte, len(bytesA))
	for i := range bytesA {
		bytesResult[i] = bytesA[i] ^ bytesB[i]
	}

	// encode bytes to hex
	return hex.EncodeToString(bytesResult), nil
}

func singleByteXor(hexSource string, key byte) string {
	// decode hex to bytes
	bytes, _ := hex.DecodeString(hexSource)

	// xor bytes
	bytesResult := make([]byte, len(bytes))
	for i := range bytes {
		bytesResult[i] = bytes[i] ^ key
	}

	// encode to hex
	return string(bytesResult)
}
