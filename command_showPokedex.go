package main

import "fmt"

func showPokedex(c *config, params ...any) error {
	for _, entry := range c.pokedex {
		fmt.Printf("%s \n", entry.Name)
	}
	return nil
}
