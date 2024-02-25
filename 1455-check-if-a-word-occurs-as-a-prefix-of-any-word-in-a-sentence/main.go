package main

import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	return withoutStr(sentence, searchWord)
	// return withoutStrHasPrefix(sentence, searchWord)
	// return strPackage(sentence, searchWord)
}

func strPackage(sentence string, searchWord string) int {
	for i, word := range strings.Split(sentence, " ") {
		if strings.HasPrefix(word, searchWord) {
			return i + 1
		}
	}
	return -1
}

func withoutStrHasPrefix(sentence string, searchWord string) int {
	sWordLen := len(searchWord)

	for i, word := range strings.Split(sentence, " ") {
		if len(word) < sWordLen {
			continue
		}

		if word[:sWordLen] == searchWord {
			return i + 1
		}
	}

	return -1
}

func withoutStr(sentence string, searchWord string) int {
	wordCount := 1 // from requirement: 1 <= sentence.length <= 100
	start := 0

	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' || i == len(sentence)-1 {
			end := i + 1

			if start+len(searchWord) < end {
				end = start + len(searchWord)
			}

			sub := sentence[start:end]
			fmt.Println(sub)

			if sub == searchWord {
				fmt.Println("found")
				return wordCount
			}

			wordCount++
			start = i + 1

			continue
		}
	}

	return -1
}
