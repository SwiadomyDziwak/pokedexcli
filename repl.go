package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startREPL(conf *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		if len(userInput) == 0 {
			continue
		}
		cleanedInput := cleanInput(userInput)
		command := cleanedInput[0]

		function, err := loadCommands()[command]
		if err != true {
			fmt.Println("Unknown command")
			continue
		}
		function.callback(conf)
	}
}

func cleanInput(text string) []string {
	cleaned := strings.TrimSpace(text)
	cleaned = strings.ToLower(cleaned)
	var cleanedWords []string
	for _, word := range strings.Split(cleaned, " ") {
		if word == "" {
			continue
		}
		cleanedWord := strings.TrimSpace(word)
		cleanedWords = append(cleanedWords, cleanedWord)
	}
	return cleanedWords
}
