package CryptoHelper

import (
	"crypto/rand"
	"crypto/sha256"
	"io"
)

const hashRepititions int = 10
const saltLength int = 32

// HashPasswordForAuthentication hashes a password with a given salt for authentication
func HashPasswordForAuthentication(password []byte, salt []byte) []byte {
	return hashPassword(password, salt, hashRepititions)
}

// HashPasswordForDecrypting hashes a password with a given salt to be used as a decryption key
func HashPasswordForDecrypting(password []byte, salt []byte) []byte {
	return hashPassword(password, salt, hashRepititions/2)
}

// hashPassword hashes a password with a given salt and a given number of iterations
func hashPassword(password []byte, salt []byte, iterations int) []byte {
	hasher := sha256.New()

	var hashed = password

	//hash our password with salt 'iterations' times for added security
	for n := 0; n < iterations; n++ {
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
