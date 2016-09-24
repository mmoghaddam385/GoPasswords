package main

import (
	"GoPasswords/CryptoHelper"
	"fmt"

	"github.com/kardianos/osext"
	"golang.org/x/crypto/ssh/terminal"
)

var settingsFileName string
var dataDirectory string

// getSettingsFileName returns the file path of the executable + "/settings"
func getSettingsFileName() string {
	if settingsFileName == "" {
		var err error
		settingsFileName, err = osext.ExecutableFolder()

		settingsFileName += "/gp_settings"

		if err != nil {
			panic("Error getting executable path: " + err.Error())
		}
	}

	return settingsFileName
}

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
}

func loadPasswordFile() {

}
