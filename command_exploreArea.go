package main

import (
	"fmt"

	"github.com/Tobias1006/pokedexcli/internal/pokeapi"
)

func exploreArea(c *config, areaName ...string) error {
	pokemonData, err := pokeapi.GetPokemonInArea(areaName[0], c.cache)
	if err != nil {
		return err
	}
	for _, pokemonEncounter := range pokemonData.PokemonEncounters {
		fmt.Printf("%s\n", pokemonEncounter.Pokemon.Name)
	}
	return nil
}
