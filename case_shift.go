package caseshift

import (
	"regexp"
	"strings"
)

func PascalCase(text string) string {
	subWords := GetSubWords(text)
	result := ""

	for _, word := range subWords {
		result += strings.Title(string(word))
	}
	return result
}

func GetSubWords(text string) []string {
	currentWord := ""
	words := []string{}

	for _, char := range text {
		if CharIsSeparator(string(char), text) {
			words = append(words, currentWord)
			currentWord = ""
			continue
		}
		currentWord += string(char)
	}
	words = append(words, currentWord)

	return words
}

func CharIsSeparator(char string, text string) bool {
	caps := regexp.MustCompile("[A-Z]")
	otherSeps := regexp.MustCompile("[-)_]")

	if strings.ToUpper(text) == text {
		return otherSeps.MatchString(char)
	}
	return caps.MatchString(char) || otherSeps.MatchString(char)
}
