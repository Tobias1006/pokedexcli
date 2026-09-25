package main

import "fmt"

func commandHelp(c *config, params ...any) error {
	fmt.Print("Welcome to the Pokedex! \n")
	for _, method := range c.commands {
		fmt.Printf("%s: %s \n", method.name, method.description)
	}
	return nil
}
