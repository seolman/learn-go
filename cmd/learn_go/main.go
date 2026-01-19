package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	wordMap := make(map[string]struct{})
	for _, msg := range messages {
		words := strings.Split(strings.ToLower(msg), " ")
		for _, word := range words {
			if word != "" {
				if _, ok := wordMap[word]; !ok {
					wordMap[word] = struct{}{}
				}
			}
		}
	}

	return len(wordMap)
}
