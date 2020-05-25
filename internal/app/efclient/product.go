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
	Price           float64         `json:"price"`
	Orders          []Order         `json:"orders"`
	Carts           []Cart          `json:"carts"`
	Properties      []Property      `json:"property"`
	Labels          []Label         `json:"labels"`
	Images          []Image         `json:"images"`
	ProductCategory json.RawMessage `json:"product_category"`
	CreatedAt       string          `json:"created_at"`
	UpdatedAt       string          `json:"updated_at"`
}

// ProductData ...
type ProductData struct {
	Name            string
	Description     string
	Price           float32
	Labels          []int
	ProductCategory int
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
func (c *Client) CreateProduct(productData *ProductData) (*Product, error) {

	requestData := map[string]interface{}{
		"name":             &productData.Name,
		"description":      &productData.Description,
		"price":            &productData.Price,
		"labels":           &productData.Labels,
		"product_category": &productData.ProductCategory,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/products", c.BaseURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	var res Product

	if err := c.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
