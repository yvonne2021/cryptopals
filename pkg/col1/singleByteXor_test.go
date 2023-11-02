package col1

import "testing"

// key is the byte used to xor the input
const keys = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func TestSingByteXor(t *testing.T) {
	for _, key := range keys {
		t.Run("single byte xor", func(t *testing.T) {
			got := singleByteXor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736", byte(key))
			want := "746865206b696420646f6e277420706c6179"

			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}
