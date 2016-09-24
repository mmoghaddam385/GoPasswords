package main

import (
	"GoPasswords/CryptoHelper"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

		fmt.Printf("exePath: %v\n", settingsFileName)

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

// loadPasswordFile will load the settings file that contains the location of the data
// if the file doesn't exist, then we create it and prompt the user for the data location
func loadSettingsFile() {
	settingsFile, err := ioutil.ReadFile(getSettingsFileName())

	//defer our error handling function
	defer func(err error) {
		recover()

		if err != nil {
			// file doesn't exist yet, lets make it! then try loading it again
			if strings.Contains(err.Error(), "no such file") {
				makeSettingsFile()
				loadSettingsFile()
			} else {
				panic("Error loading settings file: " + err.Error())
			}
		}
	}(err)

	if err == nil {
		fmt.Printf("This was in the file: %v\n", string(settingsFile))
	}
}

func makeSettingsFile() {
	fmt.Println("I can't find your settings file! Let's create a new one!")

	file, err := os.Create(getSettingsFileName())

	if err != nil {
		panic("Error creating settings file: " + err.Error())
	}

	var dataLocation string
	fmt.Println("Where is your data saved? (Or where do you want it to be saved)")
	fmt.Scanln(&dataLocation)

	//append a slash to the end if there isn't one already
	if dataLocation[len(dataLocation)-1] != '/' {
		dataLocation += "/"
	}

	file.WriteString(dataLocation)
	file.Close()
}

func loadPasswordFile() {

}
