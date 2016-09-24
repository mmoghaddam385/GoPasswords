package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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
