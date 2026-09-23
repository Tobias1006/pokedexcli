package main

import (
	"fmt"

	"github.com/Tobias1006/pokedexcli/internal/pokeapi"
)

func commandMap(c *config, params ...string) error {
	locationData, err := pokeapi.GetLocationAreas(c.next, c.cache)
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
