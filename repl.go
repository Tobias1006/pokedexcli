package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var lowerString = strings.ToLower(text)
	var stringSlice = strings.Fields(lowerString)
	return stringSlice
}
