package main

import (
	"fmt"
)

func mapf(cfg *config) error {
	res, err := cfg.pokeapiClient.ListLocations(cfg.nextURL)
	if err != nil {
		return err
	}
	cfg.nextURL = res.Next
	cfg.prevURL = res.Previous
	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func mapb(cfg *config) error {
	if cfg.prevURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := cfg.pokeapiClient.ListLocations(cfg.prevURL)
	if err != nil {
		return err
	}
	cfg.nextURL = res.Next
	cfg.prevURL = res.Previous
	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil	
}