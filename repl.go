package main

import (
	"strings"
)

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
