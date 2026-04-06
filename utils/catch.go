package utils

import (
	"fmt"

	"github.com/agomesd/go-pokedex/pokeapi"
)

func TryCatchPokemon(roll float64, baseExp int) bool {

	normalized := float64(baseExp) / float64(300)
	catchChance := 1.0 - float64(normalized)
	if catchChance > roll {
		return true
	} else {
		return false
	}
}

func PrintInfo(pokemonInfo pokeapi.PokemonInfo) {
	fmt.Println("")
	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemonInfo.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}

}
