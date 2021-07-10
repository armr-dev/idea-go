package idea

import "testing"

func TestIdea(t *testing.T) {
	t.Run("Test Encrypt", func(t *testing.T) {
		var plainText = []byte("abcdefgh")
		var key = []byte("abcdefghijklmnop")
		var expected  = []byte("num sei")

		cypheredText := EncryptBlock(plainText, key)

		if string(expected) != string(cypheredText) {
			t.Error(cypheredText)
		}
	})
}