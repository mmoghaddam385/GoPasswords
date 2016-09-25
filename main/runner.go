package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//TODO: Ensure that the program is writting to stdout so that carriage returns work properly

/*******************************************************************************************/
/**                                                                                       **/
/**                                       Main                                            **/
/**                                                                                       **/
/*******************************************************************************************/

func main() {
	loadSettingsFile()
	loadPasswordFile()

	authenticate()

	loadRecords()

	startMainLoop()
}

// startMainLoop starts the main logic loop that runs the program
func startMainLoop() {
	initCommands()

	scanner := bufio.NewScanner(os.Stdin)

	for getInput(scanner) {
	}
}

// getInput will get input and execute the command
// it will return true if we should keep going, false if we should quit
func getInput(scanner *bufio.Scanner) bool {
	fmt.Print("Enter command (try 'help'): ")

	scanner.Scan()
	input := scanner.Text()

	splitInput := strings.Split(input, " ")
	commandName := splitInput[0]
	commandArgs := splitInput[1:]

	cmd := commands[commandName]

	if cmd.execute != nil {
		cmd.execute(commandArgs[:]...)
	} else if strings.Compare(commandName, "quit") == 0 || strings.Compare(commandName, "exit") == 0 {
		return false
	} else {
		fmt.Printf("Command '%v' unrecognized, try 'help'\n", commandName)
	}

	return true
}
