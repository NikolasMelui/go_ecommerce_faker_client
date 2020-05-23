package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// StrapiRes ...
type StrapiRes struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func main() {
	res, err := http.Get("http://localhost:1337/")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		buffer, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		result := string(buffer)
		fmt.Print(result)
	}

	// var result StrapiRes

	// if err = json.NewDecoder(res.Body).Decode(&result); err == nil {
	// 	log.Fatal(err)
	// }

	// fmt.Print(&result)
}
