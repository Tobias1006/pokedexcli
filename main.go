package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var newScanner = bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		newScanner.Scan()
		var userPrompt = newScanner.Text()
		var command = cleanInput(userPrompt)
		cmd, exists := commandMap[command[0]]
		if exists == true {
			err := cmd.callback()
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

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Print("Welcome to the Pokedex! \n Usage: \n help: Displays a help message \n exit: Exit the Pokedex \n")
	return nil
}
