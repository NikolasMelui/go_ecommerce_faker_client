package efclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Products ...
type Products = []Product

// Product ...
type Product struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Price           int             `json:"price"`
	Orders          []Order         `json:"orders"`
	Carts           []Cart          `json:"carts"`
	Properties      []Property      `json:"property"`
	Lables          []Label         `json:"lables"`
	Images          []Image         `json:"images"`
	ProductCategory json.RawMessage `json:"product_category"`
	CreatedAt       string          `json:"created_at"`
	UpdatedAt       string          `json:"updated_at"`
}

// Property ...
type Property struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Image ...
type Image struct {
	ID int `json:"id"`
}

// Order ...
type Order struct {
	ID int `json:"id"`
}

// Cart ...
type Cart struct {
	ID int `json:"id"`
}

// Label ...
type Label struct {
	ID int `json:"id"`
}

// GetProducts ...
func (c *Client) GetProducts() (*Products, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/products", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	var res Products

	if err := c.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateProduct ...
func (c *Client) CreateProduct() (*Product, error) {

	requestBody, err := json.Marshal(map[string]string{
		"name":             "FirstPostedProduct",
		"description":      "Description of the first posted product",
		"price":            "1000",
		"product_category": "1",
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(requestBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/products", c.BaseURL), bytes.NewBuffer(requestBody))
	// req, err := http.NewRequest("POST", fmt.Sprintf("%s/products", c.BaseURL), bytes.NewBuffer(r))
	if err != nil {
		return nil, err
	}

	var res Product

	if err := c.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
