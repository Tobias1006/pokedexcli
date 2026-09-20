package main

func main() {
	var c config
	c.commands = cmdMap
	c.next = "https://pokeapi.co/api/v2/location-area/"
	repl(&c)
}
