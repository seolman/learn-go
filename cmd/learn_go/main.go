package main

import (
	"strings"
)

func removeProfanity(message *string) {
	replacements := map[string]string{
		"fubb":  "****",
		"shiz":  "****",
		"witch": "*****",
	}

	currentMessage := *message

	for badWord, replacement := range replacements {
		currentMessage = strings.ReplaceAll(currentMessage, badWord, replacement)
	}

	*message = currentMessage
}
