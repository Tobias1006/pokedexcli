package main

type config struct {
	commands map[string]cliCommand
	next     string
	previous string
}
