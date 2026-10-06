package main

import (
	"bufio"
	//	"encoding/json"
	"errors"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokecache"
	"strings"
)

func repl(config *config) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); ok {
			words := cleanInput(scanner.Text())
			if len(words) < 1 {
				continue
			}
			arg := ""
			if len(words) > 1 {
				arg = words[1]
			}
			err := runCommand(words[0], arg, config)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func Commands() map[string]cliCommand {
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
			description: "Lists next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area>",
			description: "Lists pokemon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "catch a pokemon, add it to your pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "inspect the stats of a pokemon in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list all of the pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}
}

func commandExit(config *config, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, _ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for name, command := range config.commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}

func commandMap(config *config, _ string) error {
	if config.map_urls.next == "" {
		fmt.Println("You're on the last page")
		return nil
	}
	l, err := getLocationArea(config.map_urls.next, &config.cache)
	if err != nil {
		return err
	}
	for _, loc := range l.Results {
		fmt.Println(loc.Name)
	}
	config.map_urls.prev = config.map_urls.next
	if l.Next == nil {
		config.map_urls.next = ""
	} else {
		config.map_urls.next = *l.Next
	}
	return nil
}

func commandMapb(config *config, _ string) error {
	if config.map_urls.prev == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	l, err := getLocationArea(config.map_urls.prev, &config.cache)
	if err != nil {
		return err
	}
	for _, loc := range l.Results {
		fmt.Println(loc.Name)
	}
	config.map_urls.next = config.map_urls.prev
	if l.Prev == nil {
		config.map_urls.prev = ""
	} else {
		config.map_urls.prev = *l.Prev
	}
	return nil
}

func commandExplore(config *config, arg string) error {
	if arg == "" {
		return errors.New("Usage: explore <area>")
	}
	fmt.Printf("Exploring %s...\n", arg)
	l, err := pokeapi.GetAreaInfo("https://pokeapi.co/api/v2/location-area/"+arg, &config.cache)
	if err != nil {
		return err
	}
	fmt.Println("Found pokemon:")
	for _, loc := range l.PokemonEncounters {
		fmt.Printf("- %s\n", loc.Pokemon.Name)
	}
	return nil
}

func commandCatch(config *config, arg string) error {
	if arg == "" {
		return errors.New("Usage: catch <pokemon>")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	pokemon, err := pokeapi.GetPokemon("https://pokeapi.co/api/v2/pokemon/"+arg, &config.cache)
	if err != nil {
		return err
	}
	config.pokedex[pokemon.Name] = pokemon
	fmt.Printf("Caught %s!\n", pokemon.Name)
	return nil
}

func getLocationArea(url string, cache *pokecache.Cache) (pokeapi.LocationArea, error) {
	l, err := pokeapi.GetLocationArea(url, cache)
	if err != nil {
		return pokeapi.LocationArea{}, err
	}
	return l, nil
}

func commandInspect(config *config, arg string) error {
	pokemon, ok := config.pokedex[arg]
	if ok {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
		return nil
	}
	return errors.New("You have not caught that pokemon")
}

func commandPokedex(config *config, _ string) error {
	for name := range config.pokedex {
		fmt.Printf("  - %s\n", name)
	}
	return nil
}

func runCommand(command string, arg string, config *config) error {
	commands := config.commands
	if c, ok := commands[command]; ok {
		return c.callback(config, arg)
	} else {
		return errors.New("Unknown command")
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return strings.Fields(text)
}
