package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// StrapiRes ...
// type StrapiRes struct {
// 	Code    int             `json:"code"`
// 	Message string          `json:"message"`
// 	Data    json.RawMessage `json:"data"`
// }

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	res, err := http.Get("http://localhost:1337/productss")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			nilError := errors.New(errRes.Message)
			log.Fatal(nilError)
		}
		newError := fmt.Errorf("Unknown error, status code: %d", res.StatusCode)
		log.Fatal(newError)
	}

	buffer, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(buffer))

	// var result StrapiResult
	// if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
	// 	log.Fatal(err)
	// }

	// var result StrapiRes

	// fmt.Print(&result)
}
