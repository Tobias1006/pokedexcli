package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/Tobias1006/pokedexcli/internal/pokeapi"
)

func catchPokemon(c *config, params ...any) error {
	pokemon, ok := params[0].(string)
	if ok {
		var catchModifier int
		info, err := pokeapi.GetPokemonInfo(pokemon)
		if err != nil {
			return fmt.Errorf("Couldn't load pokemon data")
		}
		if int((info.BaseExperience/635)*100) <= 1 {
			catchModifier = 1
		} else {
			catchModifier = int((info.BaseExperience / 635) * 100)
		}
		catchChance := int(rand.IntN(100) / catchModifier)
		fmt.Printf("Throwing a Pokeball at %s... \n", pokemon)
		if catchChance < 50 {
			fmt.Printf("%s escaped! \n", pokemon)
		} else {
			fmt.Printf("%s  was caught! \n", pokemon)
			c.pokedex[pokemon] = info
		}
		return nil
	} else if !ok {
		return fmt.Errorf("No pokemon found \n")
	}
	return nil
}
