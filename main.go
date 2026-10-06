package main

import (
	"fmt"
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokecache"
	"time"
)

func main() {
	config := config{
		commands: Commands(),
		map_urls: MapUrl{
			next: "https://pokeapi.co/api/v2/location-area/",
			prev: "",
		},
		cache:   pokecache.NewCache(5 * time.Second),
		pokedex: make(map[string]pokeapi.Pokemon),
	}
	err := repl(&config)
	if err != nil {
		fmt.Println(err)
	}
}
