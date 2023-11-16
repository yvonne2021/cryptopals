package col1

import "crypto/aes"

func aesInEcbMode(encryptedBytes []byte, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)

	sourceText := make([]byte, len(encryptedBytes))
	for i := 0; i < len(encryptedBytes); i += cipher.BlockSize() {
		cipher.Decrypt(sourceText[i:i+cipher.BlockSize()], encryptedBytes[i:i+cipher.BlockSize()])
	}
	return sourceText
}
