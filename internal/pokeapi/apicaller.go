package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Tobias1006/pokedexcli/internal/pokecache"
)

func GetLocationAreas(url string, cache pokecache.Cache) (locationData, error) {
	var emptyLocData locationData
	var locationData locationData

	cachedVal, ok := cache.Get("loc", url)
	if ok {
		data := cachedVal
		if err := json.Unmarshal(data, &locationData); err != nil {
			return emptyLocData, err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return emptyLocData, err
		}
		defer res.Body.Close()

		byteSlice, err := io.ReadAll(res.Body)
		if err := json.Unmarshal(byteSlice, &locationData); err != nil {
			return emptyLocData, err
		}
		cache.Add("loc", url, byteSlice)
	}
	return locationData, nil
}

func GetPokemonInArea(area string, cache pokecache.Cache) (locationPokemonData, error) {
	var emptyPokData locationPokemonData
	var pokemonData locationPokemonData

	var fullUrl = "https://pokeapi.co/api/v2/location-area/" + area
	cachedVal, ok := cache.Get("pok", fullUrl)
	if ok {
		data := cachedVal
		if err := json.Unmarshal(data, &pokemonData); err != nil {
			return emptyPokData, err
		}
	} else {
		res, err := http.Get(fullUrl)
		if err != nil {
			return emptyPokData, err
		}
		defer res.Body.Close()

		byteSlice, err := io.ReadAll(res.Body)
		if err := json.Unmarshal(byteSlice, &pokemonData); err != nil {
			return emptyPokData, err
		}
		cache.Add("loc", fullUrl, byteSlice)
	}
	return pokemonData, nil
}
