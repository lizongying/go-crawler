package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

type Mode string

const ECB Mode = "ECB"
const CBC Mode = "CBC"

// Aes key 16,24,32 AES-128，AES-192，AES-256
type Aes struct {
	key       []byte
	mode      Mode
	block     cipher.Block
	blockSize int
}

func NewAes(key []byte, mode Mode) (a *Aes, err error) {
	a = &Aes{
		key:  key,
		mode: mode,
	}
	a.block, err = aes.NewCipher(key)
	if err != nil {
		return
	}
	a.blockSize = a.block.BlockSize()

	return
}

func (a *Aes) Decrypt(s string) (decrypted []byte, err error) {
	if len(s) == 0 {
		err = errors.New("data size 0")
		return
	}

	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return
	}

	decrypted = make([]byte, len(data))

	if a.mode == ECB {
		blockMode := NewECBDecrypter(a.block, a.key[:a.blockSize])
		blockMode.CryptBlocks(decrypted, data)
	} else if a.mode == CBC {
		blockMode := cipher.NewCBCDecrypter(a.block, a.key[:a.blockSize])
		blockMode.CryptBlocks(decrypted, data)
	} else {
		err = errors.New("not support")
		return
	}

	decrypted = a.pkcs7UnPadding(decrypted)

	return
}

func (a *Aes) Encrypt(data []byte) (encryptedStr string, err error) {
	if len(data) == 0 {
		err = errors.New("data size 0")
		return
	}

	encryptBytes := a.pkcs7Padding(data)
	encrypted := make([]byte, len(encryptBytes))

	if a.mode == ECB {
		blockMode := NewECBEncrypter(a.block, a.key[:a.blockSize])
		blockMode.CryptBlocks(encrypted, encryptBytes)
	} else if a.mode == CBC {
		blockMode := cipher.NewCBCEncrypter(a.block, a.key[:a.blockSize])
		blockMode.CryptBlocks(encrypted, encryptBytes)
	} else {
		err = errors.New("not support")
		return
	}

	encryptedStr = base64.StdEncoding.EncodeToString(encrypted)

	return
}

func (a *Aes) pkcs7Padding(data []byte) []byte {
	padding := a.blockSize - len(data)%a.blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func (a *Aes) pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

type ECBDecrypter struct {
	block     cipher.Block
	blockSize int
	iv        []byte
}

func (e *ECBDecrypter) BlockSize() int {
	return e.blockSize
}

func (e *ECBDecrypter) CryptBlocks(dst, src []byte) {
	for bs, be := 0, e.blockSize; bs < len(src); bs, be = bs+e.blockSize, be+e.blockSize {
		e.block.Decrypt(dst[bs:be], src[bs:be])
	}

	return
}

func NewECBDecrypter(b cipher.Block, iv []byte) cipher.BlockMode {
	return &ECBDecrypter{
		block:     b,
		blockSize: b.BlockSize(),
		iv:        iv,
	}
}

type ECBEncrypter struct {
	block     cipher.Block
	blockSize int
	iv        []byte
}

func (e *ECBEncrypter) BlockSize() int {
	return e.blockSize
}

func (e *ECBEncrypter) CryptBlocks(dst, src []byte) {
	for bs, be := 0, e.blockSize; bs < len(src); bs, be = bs+e.blockSize, be+e.blockSize {
		e.block.Encrypt(dst[bs:be], src[bs:be])
	}

	return
}

func NewECBEncrypter(b cipher.Block, iv []byte) cipher.BlockMode {
	return &ECBEncrypter{
		block:     b,
		blockSize: b.BlockSize(),
		iv:        iv,
	}
}
