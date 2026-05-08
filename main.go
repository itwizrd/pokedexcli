package main

import (
	"time"
	"github.com/itwizrd/pokedexcli/internal/pokeapi"
	//"github.com/itwizrd/pokedexcli/internal/pokecache"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
    	pokeapiClient: client,
	}
	repl(cfg)
}