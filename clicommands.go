package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...any) error
}

var cmdMap = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}, "help": {
		name:        "help",
		description: "Helps with any questions",
		callback:    commandHelp,
	}, "map": {
		name:        "map",
		description: "Prints the next 20 map locations",
		callback:    commandMap,
	}, "mapb": {
		name:        "mapb",
		description: "Prints the previous 20 map locations",
		callback:    commandMapB,
	}, "explore": {
		name:        "explore",
		description: "Explores the given location",
		callback:    exploreArea,
	}, "catch": {
		name:        "catch",
		description: "Tries to catch a pokemon",
		callback:    catchPokemon,
	}, "inspect": {
		name:        "inspect",
		description: "Gives information about previously caught",
		callback:    inspectPokemon,
	}, "pokedex": {
		name:        "pokedex",
		description: "Lists all the pokemon in your pokedex",
		callback:    showPokedex,
	},
}
