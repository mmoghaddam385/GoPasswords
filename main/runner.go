package main

import (
	"GoPasswords/CryptoHelper"
	"fmt"

	"golang.org/x/crypto/ssh/terminal"
)

//TODO: Ensure that the program is writting to stdout so that carriage returns work properly

/*******************************************************************************************/
/**                                                                                       **/
/**                                       Main                                            **/
/**                                                                                       **/
/*******************************************************************************************/

func main() {

	fmt.Println("enter a password!")

	password, _ := terminal.ReadPassword(0)

	fmt.Printf("\nYou entered this password: %v\n", string(password))

	hash := CryptoHelper.HashPassword(password, nil)

	fmt.Printf("Hashed: %v\n", hash)

	loadSettingsFile()
	loadPasswordFile()
}
