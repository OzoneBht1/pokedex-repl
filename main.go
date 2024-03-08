package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/OzoneBht1/pokedex-repl/pokeapi"
)

const baseUrl string = "https://pokeapi.co/api/v2/location-area/"

type cliCommand struct {
	name        string
	description string
	callback    func()
}

var config = pokeapi.Config{
	Next:     nil,
	Previous: nil,
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
		"map": {
			name:        "map",
			description: "Fetch the next 20 pokemon locations",
			callback:    handleMapCallback,
		},
		"mapb": {
			name:        "mapb",
			description: "Fetch the previous 20 pokemon locations",
			callback:    handleMapbCallback,
		},
	}
}

func handleMapCallback() {
	url := baseUrl
	if config.Next != nil {
		url = *config.Next
	}
	result, err := pokeapi.HandleFetch(url)
	if err != nil {
		log.Fatalf("Error during map:%v", err)
	}
	printResult(result)
	config.Previous = result.Previous
	config.Next = result.Next
}

func handleMapbCallback() {
	if config.Previous == nil {
		fmt.Println("Invalid. Cannot call previous anymore. You need to call map first.")
		return
	}
	result, err := pokeapi.HandleFetch(*config.Previous)
	if err != nil {
		log.Fatalf("Error during map:%v", err)
	}
	printResult(result)
	config.Previous = result.Previous
	config.Next = result.Next
}

func printResult(data pokeapi.ApiResponse) {
	for idx, mapInfo := range data.Results {
		fmt.Printf("%d: %v\n", idx+1, mapInfo.Name)
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
