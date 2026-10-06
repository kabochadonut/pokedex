package main

import (
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	commands map[string]cliCommand
	map_urls MapUrl
	cache    pokecache.Cache
	pokedex  map[string]pokeapi.Pokemon
}

type MapUrl struct {
	base string
	next string
	prev string
}
