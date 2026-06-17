package main

import (
	"strings"
)

func cleanInput(text string) []string {
	// split user's input into "words" based on whitespace
	// should also lowercase the input and trim any leading or trailing whitespace. For example:
	// hello world -> ["hello", "world"]
	trimmed := strings.TrimSpace(strings.ToLower(text))
	words := strings.Fields(trimmed)
	return words
}

