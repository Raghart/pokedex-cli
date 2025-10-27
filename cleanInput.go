package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerText := strings.TrimSpace(strings.ToLower(text))
	return strings.Fields(lowerText)
}
