// cmd/iamsad/dog.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Dog struct (empty)
type Dog struct{}

// SubBreedsResponse struct
type SubBreedsResponse struct {
	Status    string   `json:"status"`
	SubBreeds []string `json:"message"`
}

// BreedsResponse struct
type BreedsResponse struct {
	Breeds map[string][]string `json:"message"`
}

// Breeds function
func (Dog) Breeds() []string {
	response, err := http.Get("https://dog.ceo/api/breeds/list/all")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject BreedsResponse
	json.Unmarshal(responseData, &responseObject)

	keys := make([]string, len(responseObject.Breeds))

	i := 0
	for k := range responseObject.Breeds {
		keys[i] = k
		i++
	}

	return keys
}

// SubBreeds function
func (Dog) SubBreeds(breed string) []string {
	response, err := http.Get("https://dog.ceo/api/breed/" + breed + "/list")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject SubBreedsResponse
	json.Unmarshal(responseData, &responseObject)

	return responseObject.SubBreeds
}