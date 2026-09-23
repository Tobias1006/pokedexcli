package main

import "github.com/Tobias1006/pokedexcli/internal/pokecache"

func main() {
	var c config
	c.commands = cmdMap
	c.cache = *pokecache.NewCache(5000)
	c.next = "https://pokeapi.co/api/v2/location-area/"
	repl(&c)
}
