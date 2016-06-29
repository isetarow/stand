package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Encrypt(data []byte, key string) []byte {
	block, _ := aes.NewCipher([]byte(key))

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return ciphertext
}

func Decrypt(encData []byte, key string) []byte {
	block, _ := aes.NewCipher([]byte(key))

	if len(encData) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encData[:aes.BlockSize]
	ciphertext := encData[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext
}
