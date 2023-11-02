package col1

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	t.Run("convert hex to base64", func(t *testing.T) {
		got := hexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
		want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestHexXor(t *testing.T) {
	t.Run("xor two equal length hex strings", func(t *testing.T) {
		got, _ := hexXor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
		want := "746865206b696420646f6e277420706c6179"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
