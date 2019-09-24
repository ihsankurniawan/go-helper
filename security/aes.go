package security

import (
	"io"
	"fmt"
	"errors"
	"crypto/aes"
	"crypto/rand"
	"crypto/cipher"
	"encoding/hex"
	"encoding/base64"
)

func AesEncrypt(plainText string, secretKey string) (string, error) {
	// key must be 32 bytes or 32 character
	key := []byte(secretKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	text := []byte(plainText)
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return fmt.Sprintf("%0x\n", ciphertext), nil
}

func AesDecrypt(cipherText string, secretKey string) (string, error) {
	// key must be 32 bytes or 32 character
	key := []byte(secretKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	text, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return "", err
	}

	return string(data), nil
}