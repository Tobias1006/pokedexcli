package main

import (
	"fmt"
	"os"
)

func commandExit(c *config, params ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
