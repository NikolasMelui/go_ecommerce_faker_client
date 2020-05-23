package main

import (
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

func main() {
	res, err := http.Get("http://localhost:1337/products")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {

		buffer, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatal(err)
		}

		// var result json.RawMessage
		// var result string

		// if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		// 	log.Fatal(err)
		// }

		fmt.Print(string(buffer))
	}

	// var result StrapiRes

	// if err = json.NewDecoder(res.Body).Decode(&result); err == nil {
	// 	log.Fatal(err)
	// }

	// fmt.Print(&result)
}
