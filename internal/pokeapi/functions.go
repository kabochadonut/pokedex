package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"pokedex/internal/pokecache"
)

func GetLocationArea(url string, cache *pokecache.Cache) (LocationArea, error) {
	var locationArea LocationArea
	bytes, ok := cache.Get(url)
	if ok {
		err := json.Unmarshal(bytes, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}
	req.Header.Set("User-Agent", "pokedex-cli-app")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	cache.Add(url, bodyBytes)
	err = json.Unmarshal(bodyBytes, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}

func GetAreaInfo(url string, cache *pokecache.Cache) (AreaInfo, error) {
	var areaInfo AreaInfo
	bytes, ok := cache.Get(url)
	if ok {
		err := json.Unmarshal(bytes, &areaInfo)
		if err != nil {
			return AreaInfo{}, err
		}
		return areaInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaInfo{}, err
	}
	req.Header.Set("User-Agent", "pokedex-cli-app")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return AreaInfo{}, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return AreaInfo{}, err
	}
	cache.Add(url, bodyBytes)
	err = json.Unmarshal(bodyBytes, &areaInfo)
	if err != nil {
		return AreaInfo{}, err
	}
	return areaInfo, nil
}

func GetPokemon(url string, cache *pokecache.Cache) (Pokemon, error) {
	var pokemon Pokemon
	bytes, ok := cache.Get(url)
	if ok {
		err := json.Unmarshal(bytes, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	req.Header.Set("User-Agent", "pokedex-cli-app")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	cache.Add(url, bodyBytes)
	err = json.Unmarshal(bodyBytes, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
