package main

import (
	"fmt"
)

func inspectPokemon(c *config, params ...any) error {
	pokemon, ok := params[0].(string)
	if ok {
		entry, inPokedex := c.pokedex[pokemon]
		if inPokedex {
			baseStats := make(map[string]int)
			types := make([]string, 0)
			for _, stat := range entry.Stats {
				baseStats[stat.Stat.Name] = stat.BaseStat
			}
			for _, typ := range entry.Types {
				types = append(types, typ.Type.Name)
			}
			fmt.Printf("Name: %s\nHeight: %v\nWeight: %v\nStats:\n", entry.Name, baseStats["height"], baseStats["weight"])
			fmt.Printf("  -hp: %v\n  -attack: %v\n  -defense: %v\n", baseStats["hp"], baseStats["attack"], baseStats["defense"])
			fmt.Printf("  -special-attack: %v\n  -special-defense: %v\n  -speed: %v\n", baseStats["special-attack"], baseStats["special-defense"], baseStats["speed"])
			fmt.Printf("Types:\n")
			for _, typ := range types {
				fmt.Printf(" - %s \n", typ)
			}
		} else {
			fmt.Printf("You have not yet caught %s, so no data is available \n", pokemon)
		}
	} else if !ok {
		return fmt.Errorf("pokemon not found \n")
	}
	return nil
}
