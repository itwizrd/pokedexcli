package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

func mapf(cfg *config) error {
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