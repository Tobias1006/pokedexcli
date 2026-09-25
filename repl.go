package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Tobias1006/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	var lowerString = strings.ToLower(text)
	var stringSlice = strings.Fields(lowerString)
	return stringSlice
}

func repl(c *config) {
	c.pokedex = make(map[string]pokeapi.PokemonData)
	var newScanner = bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		newScanner.Scan()
		var userPrompt = newScanner.Text()
		var command = cleanInput(userPrompt)
		cmd, exists := cmdMap[command[0]]
		if exists == true {
			if len(command) > 1 {
				areaOrPokemon := command[1]
				err := cmd.callback(c, areaOrPokemon)
				if err != nil {
					fmt.Printf("Error occured: %v\n", err)
					continue
				}
			} else {
				err := cmd.callback(c)
				if err != nil {
					fmt.Printf("Error occured: %v\n", err)
					continue
				}
			}
		} else {
			fmt.Print("Unknown command\n")
			continue
		}
	}
}
