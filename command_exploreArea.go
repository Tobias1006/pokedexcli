package main

import (
	"fmt"

	"github.com/Tobias1006/pokedexcli/internal/pokeapi"
)

func exploreArea(c *config, params ...any) error {
	value, ok := params[0].(string)
	if ok {
		areaName := value
		pokemonData, err := pokeapi.GetPokemonInArea(areaName, c.cache)
		if err != nil {
			return err
		}
		for _, pokemonEncounter := range pokemonData.PokemonEncounters {
			fmt.Printf("%s\n", pokemonEncounter.Pokemon.Name)
		}
		return nil
	} else {
		return fmt.Errorf("No area named.")
	}
}
