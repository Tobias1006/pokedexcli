package main

import (
	"github.com/Tobias1006/pokedexcli/internal/pokeapi"
	"github.com/Tobias1006/pokedexcli/internal/pokecache"
)

type config struct {
	commands map[string]cliCommand
	cache    pokecache.Cache
	next     string
	previous string
	pokedex  map[string]pokeapi.PokemonData
}
