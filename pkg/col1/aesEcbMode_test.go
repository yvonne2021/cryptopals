package col1

import (
	"os"
	"testing"
)

func TestAesEcbMode(t *testing.T) {
	t.Run("aes ecb mode", func(t *testing.T) {
		encryptedBytes, _ := os.ReadFile("7.txt")
		got := aesInEcbMode(encryptedBytes, []byte("YELLOW SUBMARINE"))
		want := "746865206b696420646f6e277420706c6179"

		if string(got) != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
