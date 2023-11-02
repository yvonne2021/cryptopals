package col1

import (
	"fmt"
	"testing"
)

func TestSingleXorCipherV2(t *testing.T) {
	t.Run("single byte xor", func(t *testing.T) {
		b, i, _ := singleByteXorV2("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
		fmt.Println("Plaintext: ", string(b))
		fmt.Println("score: ", i)
	})
}

func TestFindEncryptedString(t *testing.T) {
	t.Run("single byte xor", func(t *testing.T) {
		s, i, _ := findEncryptedString("4.txt")
		fmt.Println("Plaintext: ", s)
		fmt.Println("score: ", i)
	})
}
