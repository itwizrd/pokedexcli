package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"net/http"
	"io"
)


type cliCommand struct {
	name		string
	description	string
	callback	func() error
	config
}

type config struct {
	URL		string
	action	string
}

//"next":"https://pokeapi.co/api/v2/location-area?limit=20&offset=20", //offset=x, if next offset+=20 if prev offset-=20
// "previous":null

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
			callback:    pokemap,
			config:		 config{
				"https://pokeapi.co/api/v2/location-area?limit=20&offset=20",
				"Next",
			},
		},
		"mapb": {
			name:		 "mapb",
			description: "Displays prev 20 map locations",
			callback:	 mapb,
			config:		 config{
				"https://pokeapi.co/api/v2/location-area?limit=20&offset=20",
				"Previous",
			},
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n",c.name,c.description)
	}
	return nil
}

type Location struct {
	name	string
	URL		string
}

func pokemap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area?offset=0&limit=20")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	var locations []Location

	err = json.Unmarshal(body, &locations)
	fmt.Printf("%s", body)
	return nil
}

func mapb() error {
	return nil
}