package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/agomesd/go-pokedex/consts"
	"github.com/agomesd/go-pokedex/pokeapi"

	"github.com/agomesd/go-pokedex/utils"
)

type config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, location string) error
}

var caughtPokemon = make(map[string]pokeapi.PokemonInfo)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var commands = map[string]cliCommand{}
	config := config{
		Previous: consts.LocationAreasEndpoint,
		Next:     consts.LocationAreasEndpoint,
	}

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp(commands),
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the names of 20 location areas in the Pokemon world",
		callback:    commandMap,
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the names of the 20 previous location areas in the Pokemon world",
		callback:    commandMapBack,
	}

	commands["explore"] = cliCommand{
		name:        "explore",
		description: "It takes the name of a location area and lists all the Pokémon located there",
		callback:    commandExplore,
	}

	commands["catch"] = cliCommand{
		name:        "catch",
		description: "Allows the user to try and catch a pokemon",
		callback:    commandCatch,
	}

	commands["inspect"] = cliCommand{
		name:        "inspect",
		description: "Inspect pokedex to see if you have caught the pokemon",
		callback:    commandInspect,
	}

	for scanner.Scan() {
		userInput := scanner.Text()
		fmt.Print("Pokedex > ")
		cleaned := utils.CleanInput(userInput)
		command, ok := commands[cleaned[0]]
		var arg string
		if len(cleaned) == 2 {
			arg = cleaned[1]
		}

		if !ok {
			fmt.Println("Unknown command")
			return
		}

		command.callback(&config, arg)

	}
}

func commandExit(c *config, location string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand) func(c *config, location string) error {
	return func(c *config, location string) error {
		fmt.Println("")
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println("")
		for key, value := range commands {
			fmt.Printf("%s: %s\n", key, value.description)
		}
		return nil
	}
}

func commandMap(c *config, location string) error {
	if len(location) == 0 {
		fmt.Println("please enter location")
		return nil
	}
	locationAreas, err := pokeapi.GetLocationAreas(c.Next)
	if err != nil {
		return err
	}
	pokeapi.PrintLocationAreas(locationAreas.Results)
	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous

	return nil
}

func commandMapBack(c *config, location string) error {
	locationAreas, err := pokeapi.GetLocationAreas(c.Previous)
	if err != nil {
		return err
	}

	pokeapi.PrintLocationAreas(locationAreas.Results)

	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous

	return nil

}

func commandExplore(c *config, location string) error {
	if len(location) == 0 {
		fmt.Println("please enter location")
		return nil
	}
	fmt.Println("")
	fmt.Printf("Exploring %s...\n", location)
	locationAreasInfo, err := pokeapi.GetLocationAreaInfo(location)
	if err != nil {
		return err
	}
	pokeapi.PrintLocationAreasPokemon(locationAreasInfo)

	return nil
}

func commandCatch(c *config, pokemon string) error {
	if len(pokemon) == 0 {
		fmt.Println("enter a pokemon name")
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	pokemonInfo, err := pokeapi.GetPokemonInfo(pokemon)
	if err != nil {
		return err
	}

	roll := rand.Float64()
	caught := utils.TryCatchPokemon(roll, pokemonInfo.BaseExperience)
	if caught {
		fmt.Printf("%s was caught!\n", pokemon)
		caughtPokemon[pokemon] = pokemonInfo
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}

func commandInspect(c *config, pokemon string) error {
	caught, ok := caughtPokemon[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	utils.PrintInfo(caught)
	return nil
}
