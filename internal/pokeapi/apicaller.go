package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetLocationAreas(url string) (locationData, error) {
	var emptyLocData locationData

	res, err := http.Get(url)
	if err != nil {
		return emptyLocData, err
	}
	defer res.Body.Close()

	var locationData locationData

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationData); err != nil {
		return emptyLocData, err
	}

	return locationData, err
}
