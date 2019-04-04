// cmd/iamsad/cat.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Cat struct (empty)
type Cat struct{}

// RandomCatPhotoResponse struct
type RandomCatPhotoResponse []struct {
	Photo string `json:"url"`
}

// Photo function
func (Cat) Photo(breed string) string {
	var response *http.Response
	var err error

	response, err = http.Get("https://api.thecatapi.com/v1/images/search?limit=1")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject RandomCatPhotoResponse
	json.Unmarshal(responseData, &responseObject)

	return responseObject[0].Photo
}
