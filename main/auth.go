package main

import (
	"GoPasswords/CryptoHelper"
	"bytes"
	"fmt"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

var decryptionKey []byte

const incorrectPasswordWaitTimeSeconds int = 5

// authenticate will attempt to authenticate the user until the correct password is entered
func authenticate() {
	passwordAttempt := getPasswordFromUser()

	hashedPasswordAttempt := CryptoHelper.HashPasswordForAuthentication(passwordAttempt, passwordSalt)

	if bytes.Equal(hashedPasswordAttempt, passwordHash) {
		//you're in!
		decryptionKey = CryptoHelper.HashPasswordForDecrypting(passwordAttempt, passwordSalt)
	} else {
		fmt.Println("Password is incorrect, try again.")
		forceDelay(incorrectPasswordWaitTimeSeconds)
		authenticate() // recursion woo
	}
}

func getPasswordFromUser() []byte {
	fmt.Print("Enter your password: ")

	password, _ := terminal.ReadPassword(0)

	fmt.Println() // go to a new line
	return password
}

// forceDelay makes the user wait before appempting to authenticate again
// no brute forcing on my watch! (unless they just kill the program and start it again :\ )
// (or run multiple instances D: )
func forceDelay(seconds int) {
	for n := seconds; n > 0; n-- {
		fmt.Printf("\rWait %v more second(s) before trying again...    ", n)

		time.Sleep(time.Second)
	}

	fmt.Println("\rWait 0 more second(s) before trying again...    ")
}
