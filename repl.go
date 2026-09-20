package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	var lowerString = strings.ToLower(text)
	var stringSlice = strings.Fields(lowerString)
	return stringSlice
}

func repl(c *config) {
	var newScanner = bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		newScanner.Scan()
		var userPrompt = newScanner.Text()
		var command = cleanInput(userPrompt)
		cmd, exists := cmdMap[command[0]]
		if exists == true {
			err := cmd.callback(c)
			if err != nil {
				fmt.Printf("Error occured: %v\n", err)
				return
			}
			continue
		} else {
			fmt.Print("Unknown command\n")
			continue
		}
	}
}
