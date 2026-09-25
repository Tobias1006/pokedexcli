package pokeapi

import (
	"encoding/json"
	"fmt"
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
		if res.StatusCode > 299 {
			return emptyLocData, fmt.Errorf("Unsuccessful request.")
		}
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
	var emptyLocPokData locationPokemonData
	var locPokemonData locationPokemonData

	var fullUrl = "https://pokeapi.co/api/v2/location-area/" + area
	cachedVal, ok := cache.Get("pok", fullUrl)
	if ok {
		data := cachedVal
		if err := json.Unmarshal(data, &locPokemonData); err != nil {
			return emptyLocPokData, err
		}
	} else {
		res, err := http.Get(fullUrl)
		if res.StatusCode > 299 {
			return emptyLocPokData, fmt.Errorf("Unsuccessful request.")
		}
		if err != nil {
			return emptyLocPokData, err
		}
		defer res.Body.Close()

		byteSlice, err := io.ReadAll(res.Body)
		if err := json.Unmarshal(byteSlice, &locPokemonData); err != nil {
			return emptyLocPokData, err
		}
		cache.Add("loc", fullUrl, byteSlice)
	}
	return locPokemonData, nil
}

func GetPokemonInfo(pokemon string) (PokemonData, error) {
	var emptyPokData PokemonData
	var pokemonData PokemonData

	var fullUrl = "https://pokeapi.co/api/v2/pokemon/" + pokemon
	res, err := http.Get(fullUrl)
	if res.StatusCode > 299 {
		return emptyPokData, fmt.Errorf("Unsuccessful request.")
	}
	if err != nil {
		return emptyPokData, err
	}
	defer res.Body.Close()
	byteSlice, err := io.ReadAll(res.Body)
	if err := json.Unmarshal(byteSlice, &pokemonData); err != nil {
		return emptyPokData, err
	}
	return pokemonData, nil
}
