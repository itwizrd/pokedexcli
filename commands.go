package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
)


type cliCommand struct {
	name		string
	description	string
	callback	func() error
	config:		struct
}
type config struct {
	Next:		"next": "https://pokeapi.co/api/v2/location-area?limit=20&offset=20", //offset=x, if next offset+=20 if prev offset-=20
	Previous:	"previous": null,
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			config:		 null,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			config:		 null,
		},
		"map": {
			name:        "map",
			description: "Displays 20 map locations",
			callback:    map,
			config:		 Next,
		},
		"mapb": {
			name:		 "mapb",
			description: "Displays prev 20 map locations",
			callback:	 mapb,
			config:		 Previous,
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	command := getCommands()
	for _, c := range command {
		fmt.Printf("%s: %s\n",c.name,c.description)
	}
	return nil
}

type Location struct {
	name:	string
	URL:	string
}

func map() error {
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
	var loccations []Location

	err := json.Unmarshal(body, &locations)
	fmt.Printf("%s", body)
}