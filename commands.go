package main

import (
	"encoding/json"
	"fmt"
	"os"
	"net/http"
	"io"
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
		},
		"mapb": {
			name:		 "mapb",
			description: "Displays prev 20 map locations",
			callback:	 mapb,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n",c.name,c.description)
	}
	fmt.Printf("\n")
	return nil
}

type Locations struct {
	Count    int		`json:"count"`
	Next     *string	`json:"next"`
	Previous *string	`json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func pokemap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	return mapHelper(url, cfg)
}

func mapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := *cfg.Previous
	return mapHelper(url, cfg)
}

func mapHelper(url string, cfg *config) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}
	locations := Locations{}

	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}
	cfg.Next = locations.Next
	cfg.Previous = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}