package main

import (
	"GoPasswords/CryptoHelper"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

/*

Format of password file:

Line 1: password hash (base64 encoded)
Line 2: password salt (base64 encoded)

*/

var passwordHash []byte
var passwordSalt []byte

const passwordFileName string = "auth"

func loadPasswordFile() {
	passwordFileContents, err := ioutil.ReadFile(dataDirectory + passwordFileName)

	//defer our error handling function
	defer func(err error) {
		recover()

		if err != nil {
			if os.IsNotExist(err) {
				resolveMissingPasswordFile()
				loadPasswordFile()
			} else {
				panic("Error loading password file: " + err.Error())
			}
		}
	}(err)

	if err == nil {
		split := strings.Split(string(passwordFileContents), "\n")

		passwordHash, _ = base64.URLEncoding.DecodeString(split[0])
		passwordSalt, _ = base64.URLEncoding.DecodeString(split[1])
	}

}

// resolveMissingPasswordFile will either make a new password file or
// change the data directory to point to an already existing data folder
func resolveMissingPasswordFile() {
	fmt.Printf("I can't load your password file from '%v'.\n", dataDirectory)

	var isValidInput = false
	var input string

	//keep polling for input until a valid response is given
	for !isValidInput { // no while loop boo :(
		fmt.Print("Would you like to create a [N]ew password file or [C]hange the data directory? ")
		fmt.Scan(&input)

		strings.ToLower(input)

		if input == "n" || input == "c" {
			isValidInput = true
		}
	}

	switch input {
	case "n":
		createNewPasswordFile()
	case "c":
		changeDataDirectory("Where is your data saved?")
	default:
		panic("How the hell did you input '" + input + "' and get away with it??")
	}

}

func createNewPasswordFile() {
	fmt.Println("Ok, I'll make you a new password file.")

	rawPasswordPtr := getNewPassword()
	salt := CryptoHelper.GenerateRandomSalt()

	hashedPasswordRaw := CryptoHelper.HashPasswordForAuthentication(rawPasswordPtr, salt)

	hashedPassword := base64.URLEncoding.EncodeToString(hashedPasswordRaw)
	encodedSalt := base64.URLEncoding.EncodeToString(salt)

	authFile, err := os.Create(dataDirectory + passwordFileName)

	if err != nil {
		panic("Error creating password file: " + err.Error())
	}

	authFile.WriteString(hashedPassword + "\n")
	authFile.WriteString(encodedSalt)

	authFile.Close()
}

// getNewPassword gets a new password from the user.
//  the password is confirmed by multiple entries
func getNewPassword() []byte {
	var isValidPassword = false
	var password []byte

	for !isValidPassword {
		fmt.Print("Enter your password: ")
		password, _ = terminal.ReadPassword(0)

		fmt.Print("\nConfirm your password: ")
		passwordConfirm, _ := terminal.ReadPassword(0)

		fmt.Print("\n")

		if bytes.Equal(password, passwordConfirm) {
			isValidPassword = true
		} else {
			fmt.Println("Passwords don't match, try again")
		}
	}

	return password
}
