package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/agomesd/go-pokedex/consts"
)

type GetLocationAreasResponse struct {
	Count    int64         `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []PokeAPIItem `json:"results"`
}

type LocationAreaInfo struct {
	ID                   int64  `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int64  `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod PokeAPIItem `json:"encounter_method"`
		VersionDetails  []struct {
			Rate    int64       `json:"rate"`
			Version PokeAPIItem `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location PokeAPIItem `json:"location"`
	Names    []struct {
		Name     string      `json:"name"`
		Language PokeAPIItem `json:"language"`
	} `json:"names"`
	PokemonEncounter []struct {
		Pokemon        PokeAPIItem `json:"pokemon"`
		VersionDetails []struct {
			Version          PokeAPIItem `json:"version"`
			MaxChance        int64       `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int64                    `json:"min_level"`
				MaxLevel        int64                    `json:"max_level"`
				ConditionValues []map[string]interface{} `json:"condition_values"`
				Chance          int64                    `json:"chance"`
				Method          PokeAPIItem              `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type PokeAPIItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonInfo struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool        `json:"is_hidden"`
		Slot     int         `json:"slot"`
		Ability  PokeAPIItem `json:"ability"`
	} `json:"abilitites"`
	Forms        []PokeAPIItem `json:"forms"`
	GameIndicies []struct {
		GameIndex int         `json:"game_index"`
		Version   PokeAPIItem `json:"version"`
	} `json:"game_indices"`
}

func GetLocationAreas(url string) (GetLocationAreasResponse, error) {
	req, err := http.Get(url)
	if err != nil {
		return GetLocationAreasResponse{}, err
	}

	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)

	json := GetLocationAreasResponse{}

	decoder.Decode(&json)

	return json, nil
}

func GetLocationAreaInfo(location string) (LocationAreaInfo, error) {
	fullURL := consts.LocationAreasEndpoint + location
	req, err := http.Get(fullURL)
	if err != nil {
		return LocationAreaInfo{}, err
	}

	defer req.Body.Close()

	locationAreaInfo := LocationAreaInfo{}

	decoder := json.NewDecoder(req.Body)

	decoder.Decode(&locationAreaInfo)

	return locationAreaInfo, nil
}

func GetPokemonInfo(pokemon string) (PokemonInfo, error) {
	fullURL := consts.PokemonInfoEndpoint + pokemon
	req, err := http.Get(fullURL)
	if err != nil {
		return PokemonInfo{}, err
	}

	defer req.Body.Close()

	pokemonInfo := PokemonInfo{}

	decoder := json.NewDecoder(req.Body)

	decoder.Decode(&pokemonInfo)

	return pokemonInfo, nil
}

func PrintLocationAreas(locationAreasResults []PokeAPIItem) {
	fmt.Println("")
	for _, locationArea := range locationAreasResults {
		fmt.Println(locationArea.Name)
	}
}

func PrintLocationAreasPokemon(locationAreaInfo LocationAreaInfo) {
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationAreaInfo.PokemonEncounter {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
}
