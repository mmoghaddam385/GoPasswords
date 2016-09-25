package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kardianos/osext"
)

/*

Format of settings file:

Line 1: filepath of data folder

*/

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

// loadPasswordFile will load the settings file that contains the location of the data
// if the file doesn't exist, then we create it and prompt the user for the data location
func loadSettingsFile() {
	settingsFile, err := ioutil.ReadFile(getSettingsFileName())

	//defer our error handling function
	defer func(err error) {
		recover()

		if err != nil {
			// file doesn't exist yet, lets make it! then try loading it again
			if os.IsNotExist(err) {
				makeSettingsFile()
				loadSettingsFile()
			} else {
				panic("Error loading settings file: " + err.Error())
			}
		}
	}(err)

	if err == nil {
		dataDirectory = string(settingsFile)
	}
}

func makeSettingsFile() {
	fmt.Println("I can't find your settings file! Let's create a new one!")

	changeDataDirectory("Where is your data saved? (Or where do you want it to be saved)")
}

func changeDataDirectory(msg string) {
	//this will simply open for writing if the file already exists
	settingsFile, err := os.Create(getSettingsFileName())

	if err != nil {
		panic("Error creating settings file (" + getSettingsFileName() + "): " + err.Error())
	}

	var dataLocation string
	fmt.Println(msg)
	fmt.Scanln(&dataLocation)

	//append a slash to the end if there isn't one already
	if dataLocation[len(dataLocation)-1] != '/' {
		dataLocation += "/"
	}

	dataDirectory = dataLocation

	settingsFile.WriteString(dataLocation)
	settingsFile.Close()

}

func forceChangeDataDirectory(newDir string) {
	//this will simply open for writing if the file already exists
	settingsFile, err := os.Create(getSettingsFileName())

	if err != nil {
		panic("Error creating settings file (" + getSettingsFileName() + "): " + err.Error())
	}

	dataDirectory = newDir

	settingsFile.WriteString(newDir)
	settingsFile.Close()
}
