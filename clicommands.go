package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
	},
}
