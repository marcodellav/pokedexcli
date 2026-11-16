package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cachedUrl, ok := c.cache.Get(url)
	if ok {
		// fmt.Println("************** cache hit for", url)
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(cachedUrl, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	c.cache.Add(url, dat)

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) ListPokemons(id string) (RespShallowPokemons, error) {
	url := baseURL + "/" + "location-area/" + id

	cachedUrl, ok := c.cache.Get(url)
	if ok {
		// fmt.Println("************** cache hit for", url)
		pokemonsResp := RespShallowPokemons{}
		err := json.Unmarshal(cachedUrl, &pokemonsResp)
		if err != nil {
			return RespShallowPokemons{}, err
		}
		return pokemonsResp, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	// fmt.Println("url:", url)
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	c.cache.Add(url, dat)

	pokemonsResp := RespShallowPokemons{}
	err = json.Unmarshal(dat, &pokemonsResp)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	return pokemonsResp, nil
}

func (c *Client) GetPokemonInfo(name string) (RespShallowPokemonsInfo, error) {

	// https://pokeapi.co/api/v2/pokemon/{id or name}/
	url := baseURL + "/" + "pokemon/" + name

	cachedUrl, ok := c.cache.Get(url)
	if ok {
		// fmt.Println("************** cache hit for", url)
		pokemonsInfoResp := RespShallowPokemonsInfo{}
		err := json.Unmarshal(cachedUrl, &pokemonsInfoResp)
		if err != nil {
			return RespShallowPokemonsInfo{}, err
		}
		return pokemonsInfoResp, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemonsInfo{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemonsInfo{}, err
	}
	// fmt.Println("url:", url)
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemonsInfo{}, err
	}
	c.cache.Add(url, dat)

	pokemonsInfoResp := RespShallowPokemonsInfo{}
	err = json.Unmarshal(dat, &pokemonsInfoResp)
	if err != nil {
		return RespShallowPokemonsInfo{}, err
	}

	return pokemonsInfoResp, nil
}
