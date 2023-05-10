package utils

import (
	"testing"
)

func TestAes(t *testing.T) {
	aes, _ := NewAes([]byte("crzjmwlcmgylxtyl"), ECB)
	str, _ := aes.Encrypt([]byte("丰"))
	t.Log("Encrypt", str)

	bs, _ := aes.Decrypt(str)
	t.Log("Decrypt", string(bs))
}
