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

func (c *Client) ListPokemons(id string) (Pokemon, error) {
	url := baseURL + "/" + "location-area/" + id

	cachedUrl, ok := c.cache.Get(url)
	if ok {
		// fmt.Println("************** cache hit for", url)
		pokemonsResp := Pokemon{}
		err := json.Unmarshal(cachedUrl, &pokemonsResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonsResp, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	// fmt.Println("url:", url)
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, dat)

	pokemonsResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonsResp)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonsResp, nil
}

func (c *Client) GetPokemonInfo(name string) (PokemonInfo, error) {

	// https://pokeapi.co/api/v2/pokemon/{id or name}/
	url := baseURL + "/" + "pokemon/" + name

	cachedUrl, ok := c.cache.Get(url)
	if ok {
		// fmt.Println("************** cache hit for", url)
		pokemonsInfoResp := PokemonInfo{}
		err := json.Unmarshal(cachedUrl, &pokemonsInfoResp)
		if err != nil {
			return PokemonInfo{}, err
		}
		return pokemonsInfoResp, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	// fmt.Println("url:", url)
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}
	c.cache.Add(url, dat)

	pokemonsInfoResp := PokemonInfo{}
	err = json.Unmarshal(dat, &pokemonsInfoResp)
	if err != nil {
		return PokemonInfo{}, err
	}

	return pokemonsInfoResp, nil
}
