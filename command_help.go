package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n",c.name,c.description)
	}
	fmt.Printf("\n")
	return nil
}
