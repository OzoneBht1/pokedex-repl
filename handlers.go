package main

import (
	"fmt"
	"os"
)

func commandHelp() {
	var commands = getCliCommands()
	fmt.Printf("Welcome to the Pokedex!\n Usage:\n\n")
	for c := range commands {
		fmt.Printf("%v: %v\n", commands[c].name, commands[c].description)
	}
}

func exitHelp() {
	fmt.Fprintln(os.Stdout, []any{"Bye"}...)
	os.Exit(1)
}
