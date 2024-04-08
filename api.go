package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var defaultUrl = "https://pokeapi.co/api/v2/location-area/"

type apiResponse struct {
	Next     *string
	Previous *string
	Results  []location
}

type location struct {
	Name string
	Url  string
}

func getLocations(url *string) (*apiResponse, error) {
	if url == nil {
		url = &defaultUrl
	}

	// make request
	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	// read body
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	if res.StatusCode > 299 {
		err := fmt.Errorf("response faild with status code %d and body %s", res.StatusCode, body)
		log.Fatal(err)
		return nil, err
	}

	// parse JSON
	apiRes := apiResponse{}
	if err := json.Unmarshal(body, &apiRes); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return &apiRes, nil
}
