package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(*config) error
}

type config struct {
	Next		*string
	Previous	*string
}


func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
func repl() {
	cfg := &config{}
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		c, exists := getCommands()[input[0]]
		if exists {
			err := c.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 map locations",
			callback:    mapf,
		},
		"mapb": {
			name:		 "mapb",
			description: "Displays prev 20 map locations",
			callback:	 mapb,
		},
	}
}
