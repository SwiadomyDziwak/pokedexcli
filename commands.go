package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, c := range conf.commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
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
