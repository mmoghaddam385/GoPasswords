package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

type command struct {
	helpText []string
	execute  func(...string)
}

// list of command names
var commands map[string]command

// I can't initialize the comands map to have this because there would be circular references
func initCommands() {
	commands = make(map[string]command)
	commands["help"] = command{[]string{"Show help message",
		"Optional argument: command name"}, showHelpCommand}

	commands["quit"] = command{[]string{"Quits the program"}, nil}
	commands["exit"] = command{[]string{"Quits the program"}, nil}

	commands["search"] = command{[]string{"Search for a record",
		"Required argument: query string",
		"This command returns the names of any record whose ",
		"site name contains the query string"}, searchCommand}

	commands["list"] = command{[]string{"List all records"}, listCommand}

	commands["view"] = command{[]string{"View the contents of a record",
		"Required argument: record sitename",
		"This command displays the attributes of the record in plain text",
		"The fields are masked after enter is pressed and the command exits"}, viewRecordCommand}

	commands["delete"] = command{[]string{"Delete a record permanently",
		"Required argument: record sitename"}, deleteRecordCommand}

	commands["rm"] = command{[]string{"See 'delete'"}, deleteRecordCommand}

	commands["create"] = command{[]string{"Create and save a new record",
		"No required arguments; you will be prompted for values"}, createRecordCommand}

	commands["new"] = command{[]string{"See 'create'"}, createRecordCommand}

	commands["update"] = command{[]string{"Update a record",
		"Required argument: record sitename"}, updateRecordCommand}

	commands["change-password"] = command{[]string{"Change Password",
		"No required arguments; you will be prompted for values"}, changePasswordCommand}
}

func showHelpCommand(args ...string) {
	// if arguments are supplied, display help text only for those commands
	if len(args) > 0 {
		for _, arg := range args {
			cmd := commands[arg]

			fmt.Printf("%v:\n", arg)
			if len(cmd.helpText) > 0 {
				printHelpText(cmd.helpText)
			} else {
				printHelpText([]string{"Command not found"})
			}

			fmt.Println()
		}
	} else { // no args? just print em all!
		for name, cmd := range commands {
			fmt.Printf("%v:\n", name)
			printHelpText(cmd.helpText)
			fmt.Println()
		}
	}
}

func printHelpText(helpText []string) {
	for _, textLine := range helpText {
		fmt.Printf("\t\t\t%v\n", textLine)
	}
}

func searchCommand(args ...string) {
	fmt.Println("Not yet implemented")
}

func listCommand(args ...string) {
	if len(recordMap) == 0 {
		fmt.Println("No records, make a new one with 'create'")
	} else {
		fmt.Println()
		for plainText := range recordMap {
			fmt.Println(plainText)
		}
		fmt.Println()
	}
}

func viewRecordCommand(args ...string) {
	if len(args) == 0 {
		fmt.Println("You must pass a record site name as a parameter")
		return
	}

	siteName := strings.Join(args, " ")

	if !recordExists(siteName) {
		fmt.Printf("'%v' isn't in the system, try 'create'\n", siteName)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	r := getRecord(siteName)

	fmt.Printf("site name: %v\n", r.sitename)
	fmt.Printf("user name: %v\n", r.username)
	fmt.Printf("password:  %v\t\t", r.password)

	scanner.Scan()

	//move cursor to beginning of the line, then up one line, then print
	fmt.Printf("\r\033[1Apassword: <REDACTED>                               \n")
}

func deleteRecordCommand(args ...string) {
	siteName := strings.Join(args, " ")

	if recordExists(siteName) {
		fmt.Print("Are you sure? (y/n) ")
		var answer string
		fmt.Scan(&answer)

		if strings.ToLower(answer) == "y" {
			deleteRecord(siteName)
			fmt.Println("Record deleted")
		} else {
			fmt.Println("Delete aborted")
		}
	} else {
		fmt.Printf("'%v' isn't in the system, try 'create'\n", siteName)
	}
}

func createRecordCommand(args ...string) {
	scanner := bufio.NewScanner(os.Stdin)
	var newRecord record

	fmt.Print("\nEnter site name: ")
	scanner.Scan()
	newRecord.sitename = scanner.Text()

	// disallow duplicate site names
	if recordExists(newRecord.sitename) {
		fmt.Printf("ERROR: A record already exists under '%v'\n", newRecord.sitename)
		return
	}

	fmt.Print("Enter user name: ")
	scanner.Scan()
	newRecord.username = scanner.Text()

	fmt.Print("Enter password: ")
	passwordBytes, _ := terminal.ReadPassword(0)
	newRecord.password = string(passwordBytes)
	fmt.Println()

	createNewRecord(newRecord)
}

func updateRecordCommand(args ...string) {

}

func changePasswordCommand(args ...string) {
	fmt.Println("Not yet implemented")
}
