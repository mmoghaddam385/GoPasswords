package CryptoHelper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

const nonceSize int = 16

// DecryptString will decrypt the given string with the given key and iv
// and return the result as a string
func DecryptString(cipherText string, key []byte, iv []byte) string {
	cipherTextAsBytes, _ := base64.URLEncoding.DecodeString(cipherText)
	return string(Decrypt(cipherTextAsBytes, key, iv))
}

// Decrypt will decrypt the given bytes with the given key and iv
// and return the result as a byte slice
func Decrypt(cipherText []byte, key []byte, iv []byte) []byte {
	plainText, err := getCipher(key).Open(nil, iv, cipherText, nil)
	if err != nil {
		panic("Error decrypting data: " + err.Error())
	}

	return plainText
}

// EncryptString encrypts a string with the given key and iv
// and returns the result as a byte slice
func EncryptString(plainText string, key []byte, iv []byte) []byte {
	plainTextAsBytes, _ := base64.URLEncoding.DecodeString(plainText)
	return Encrypt(plainTextAsBytes, key, iv)
}

// Encrypt will encrypt the given bytes with the given key and iv
// and return the result as a byte slice
func Encrypt(plainText []byte, key []byte, iv []byte) []byte {
	return getCipher(key).Seal(nil, iv, plainText, nil)
}

func getCipher(key []byte) cipher.AEAD {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic("Error getting AES cipher!!!! " + err.Error())
	}

	aesgcmCipher, err := cipher.NewGCM(block)
	if err != nil {
		panic("Error getting GCM mode AES Cipher!!!! " + err.Error())
	}

	return aesgcmCipher
}
