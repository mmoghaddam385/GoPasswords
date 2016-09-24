package CryptoHelper

import (
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

const hashRepititions int = 10
const saltLength int = 32

// HashPassword hashes a password with a given salt
func HashPassword(password []byte, salt []byte) []byte {
	fmt.Printf("SHA256 available? %v\n", crypto.SHA256.Available())
	hasher := sha256.New()

	var hashed = password

	//hash our password with salt hashRepititions times for added security
	for n := 0; n < hashRepititions; n++ {
		hasher.Write(append(hashed[:], salt[:]...))
		hashed = hasher.Sum(nil)
	}

	return hashed
}

// GenerateRandomSalt generates a random salt for use with password hashing
func GenerateRandomSalt() []byte {
	salt := make([]byte, saltLength)
	io.ReadFull(rand.Reader, salt)

	return salt
}
