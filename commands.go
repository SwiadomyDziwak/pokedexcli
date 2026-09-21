package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, c := range loadCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	//fmt.Println("exit: Exit the Pokedex")
	//fmt.Println("help: Displays a help message")
	return nil
}

func loadCommands() map[string]cliCommand {
	commands := make(map[string]cliCommand)

	commands["exit"] = cliCommand{
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	}
	commands["help"] = cliCommand{
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	}
	return commands
}
