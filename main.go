package main

import (
	"github.com/Tobias1006/pokedexcli/internal/pokecache"
)

func main() {
	var config config
	config.commands = cmdMap
	config.cache = *pokecache.NewCache(5000)
	config.next = "https://pokeapi.co/api/v2/location-area/"
	repl(&config)
}
