package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Product ...
type Product struct {
	ID                int               `json:"id"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Price             int               `json:"price"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	ProductCategory   int               `json:"product_category"`
	ProductProperties []ProductProperty `json:"property"`
	ProductImages     []ProductImage    `json:"images"`
	ProductOrders     []ProductOrder    `json:"orders"`
	ProductCarts      []ProductCart     `json:"carts"`
	ProductLables     []ProductLable    `json:"lables"`
}

// ProductProperty ...
type ProductProperty struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ProductImage ...
type ProductImage struct {
	ID int `json:"id"`
}

// ProductOrder ...
type ProductOrder struct {
	ID int `json:"id"`
}

// ProductCart ...
type ProductCart struct {
	ID int `json:"id"`
}

// ProductLable ...
type ProductLable struct {
	ID int `json:"id"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	res, err := http.Get("http://localhost:1337/products")

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

	var products []Product
	if err = json.NewDecoder(res.Body).Decode(&products); err != nil {
		log.Fatal(err)
	}

	for _, product := range products {
		fmt.Println("Product ID - ", product.ID)
		fmt.Println("Product Name - ", product.Name)
	}

}
