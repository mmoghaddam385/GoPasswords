package main

import "fmt"

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
		for plainText := range recordMap {
			fmt.Println(plainText)
		}
	}
}

func viewRecordCommand(args ...string) {
	fmt.Println("Not yet implemented")
}

func deleteRecordCommand(args ...string) {
	fmt.Println("Not yet implemented")
}

func createRecordCommand(args ...string) {
	fmt.Println("Not yet implemented")
}

func changePasswordCommand(args ...string) {
	fmt.Println("Not yet implemented")
}
