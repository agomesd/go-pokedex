package main

import (
	"bufio"
	"fmt"
	"github.com/agomesd/go-pokedex/consts"
	"github.com/agomesd/go-pokedex/poke-api"
	"github.com/agomesd/go-pokedex/utils"
	"os"
)

type config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

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

	for scanner.Scan() {
		userInput := scanner.Text()
		fmt.Print("Pokedex > ")
		cleaned := utils.CleanInput(userInput)
		command, ok := commands[cleaned[0]]
		if !ok {
			fmt.Println("Unknown command")
			return
		}

		command.callback(&config)

	}
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand) func(c *config) error {
	return func(c *config) error {
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

func commandMap(c *config) error {
	locationAreas, err := pokeapi.GetLocationAreas(c.Next)
	if err != nil {
		return err
	}
	pokeapi.PrintLocationAreas(locationAreas.Results)
	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous

	return nil
}

func commandMapBack(c *config) error {
	locationAreas, err := pokeapi.GetLocationAreas(c.Previous)
	if err != nil {
		return err
	}

	pokeapi.PrintLocationAreas(locationAreas.Results)

	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous

	return nil

}
