package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func commandHelp() {
	var commands = getCliCommands()
	fmt.Printf("Welcome to the Pokedex!\n Usage:\n\n")
	for c := range commands {
		fmt.Printf("%v: %v\n", commands[c].name, commands[c].description)
	}
}

func exitHelp() {
	fmt.Println("Bye")
	os.Exit(1)
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display the help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Quit the program",
			callback:    exitHelp,
		},
	}
}

func main() {
	const prefixMsg = "pokedex >"
	var commands = getCliCommands()
	for {
		var userChoice = ""
		fmt.Printf("%v", prefixMsg)
		fmt.Scanf("%v", &userChoice)
		if command, ok := commands[userChoice]; ok {
			command.callback()
		} else {
			fmt.Println(errors.New("Invalid Option. Choices are :"))
			for c := range commands {
				fmt.Println(c)
			}
		}
	}

}
