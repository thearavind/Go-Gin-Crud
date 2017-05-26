package handlers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// Takes in a string and returns AES Encrypted and base64 Encoded string
func encrypt(rawString []byte) (string, error) {
	/*TODO move this to env variable*/
	cipherBlock, err := aes.NewCipher([]byte("secret*#key#*for*#AES&encryption"))
	if err != nil {
		return "", err
	}
	encodedRawString := base64.StdEncoding.EncodeToString(rawString)
	cipherText := make([]byte, aes.BlockSize+len(encodedRawString))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(cipherBlock, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(encodedRawString))
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Takes in the encrypted string and returns the decrypted text
func decrypt(encryptedString string) (string, error) {
	cipherBlock, err := aes.NewCipher([]byte("secret*#key#*for*#AES&encryption"))
	if err != nil {
		return "", err
	}
	decodedString, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}
	if len(decodedString) < aes.BlockSize {
		return "", errors.New("cipher text is too short")
	}
	iv := decodedString[:aes.BlockSize]
	decodedString = decodedString[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(cipherBlock, iv)
	cfb.XORKeyStream(decodedString, decodedString)
	decryptedString, err := base64.StdEncoding.DecodeString(string(decodedString))
	if err != nil {
		return "", err
	}
	return string(decryptedString), nil
}
