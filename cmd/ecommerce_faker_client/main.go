package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
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

	io.Copy(os.Stdout, res.Body)

	// var result StrapiRes

	// if err = json.NewDecoder(res.Body).Decode(&result); err == nil {
	// 	log.Fatal(err)
	// }

	// fmt.Print(&result)
}
