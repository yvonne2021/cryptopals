package col1

import (
	"testing"
)

func TestRepeatingKeyXor(t *testing.T) {
	t.Run("repeating key xor", func(t *testing.T) {
		got := repeatingKeyXor("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal", "ICE")
		want := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

//func TestBreakRepeatingKeyXor(t *testing.T) {
//	t.Run("break repeating key xor", func(t *testing.T) {
//		bytes, _ := os.ReadFile("6.txt")
//		got := breakRepeatingKeyXor(bytes)
//		want := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
//
//		if got != want {
//			t.Errorf("got %q want %q", got, want)
//		}
//	})
//}

func TestGetHammingDistance(t *testing.T) {
	t.Run("get haming distance", func(t *testing.T) {
		got := getHammingDistance("this is a test", "wokka wokka!!!")
		want := 37

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
