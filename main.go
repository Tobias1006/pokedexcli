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
		var cleanedPrompt = cleanInput(userPrompt)
		fmt.Printf("Your command was: %s \n", cleanedPrompt[0])
	}
}
