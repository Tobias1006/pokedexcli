package main

import (
	"fmt"

	"github.com/Tobias1006/pokedexcli/internal/pokeapi"
)

func commandMapB(c *config, params ...string) error {
	if c.previous == "" {
		fmt.Print("you're on the first page \n")
		return nil
	} else {
		locationData, err := pokeapi.GetLocationAreas(c.previous, c.cache)
		if err != nil {
			return err
		}
		c.next = locationData.Next
		c.previous = locationData.Previous
		for _, location := range locationData.Results {
			fmt.Printf("%s\n", location.Name)
		}
		return nil
	}
}
