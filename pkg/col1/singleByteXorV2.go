package col1

import (
	"cryptopals/src/pkg/utils"
	"encoding/hex"
	"os"
	"strings"
)

/*
https://cryptopals.com/sets/1/challenges/3
plaintext -> bytes ->  xor with single byte key -> xoredBytes -> encoded hex(given hex string)

Reverse above steps to get a candidate plaintext.
1. decode the encoded hex string  to bytes
2. xor the bytes with a single byte key(try all possible bytes, 0000 0000 - 1111 1111, 0 - 255)
3. Score each candidate plaintexts using a character frequency weight map
4. Return the candidate plaintext with the highest score
*/

func singleByteXorV2(hexString string) ([]byte, int, error) {
	// step 1 decode hex to bytes
	rawBytes, _ := hex.DecodeString(hexString)

	var candidateBytes []byte
	maxScore := 0

	// step 2 xor bytes with a single byte key
	for key := 0; key < 256; key++ {
		xoredBytes := make([]byte, len(rawBytes))
		currentScore := 0
		for i, singleByte := range rawBytes {
			xoredByte := singleByte ^ byte(key)
			// step 3 score each candidate plaintexts using a character frequency weight map
			currentScore += utils.GetCharWeight(xoredByte)
			xoredBytes[i] = xoredByte
		}

		// step 4 Always keep the candidate plaintext with the highest score
		if currentScore > maxScore {
			candidateBytes = xoredBytes
			maxScore = currentScore
		}

		currentScore = 0
	}
	return candidateBytes, maxScore, nil
}

func findEncryptedString(filePath string) (string, int, error) {
	// Step 1. Read the file into bytes
	f, _ := os.ReadFile(filePath)

	// Step 2. Convert bytes into string and split string into lines
	lines := strings.Split(string(f), "\n")

	var candidateBytes []byte
	maxScore := 0

	// Step 3. For each line, find the candidate plaintext with the highest score
	for _, line := range lines {
		currentBytes, currentScore, _ := singleByteXorV2(line)

		// Step 4. Always keep the candidate plaintext with the highest score
		if currentScore > maxScore {
			maxScore = currentScore
			candidateBytes = currentBytes
		}
	}

	return string(candidateBytes), maxScore, nil
}
