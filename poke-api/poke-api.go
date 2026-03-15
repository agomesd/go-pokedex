package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetLocationAreasResponse struct {
	Count    int64                    `json:"count"`
	Next     string                   `json:"next"`
	Previous string                   `json:"previous"`
	Results  []GetLocationAreasResult `json:"results"`
}

type GetLocationAreasResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
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

func PrintLocationAreas(locationAreasResults []GetLocationAreasResult) {
	fmt.Println("")
	for _, locationArea := range locationAreasResults {
		fmt.Println(locationArea.Name)
	}
}
