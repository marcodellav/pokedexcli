package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(c *config) error {

	type LocationAreas struct {
		Count    int     `json:"count"`
		Next     *string `json:"next"`
		Previous *string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}

	url := c.URL
	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	locAreas := LocationAreas{}
	err = json.Unmarshal(body, &locAreas)
	c.URL = locAreas.Next
	c.Previous = locAreas.Previous

	if err != nil {
		fmt.Println(err)
	}
	for i := range locAreas.Results {
		fmt.Println(locAreas.Results[i].Name)
	}

	// fmt.Println("url:", url)

	return err

}
