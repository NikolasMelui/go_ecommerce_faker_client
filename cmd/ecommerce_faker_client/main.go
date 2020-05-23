package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

// BaseURLV1 ...
var BaseURLV1 = "http://localhost:1337"

// Client ...
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// Products ...
type Products = []Product

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
	c := NewClient()

	products, err := c.getProducts()
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range *products {
		fmt.Println("Product ID - ", product.ID)
		fmt.Println("Product Name - ", product.Name)
	}
}

// NewClient ...
func NewClient() *Client {
	return &Client{
		BaseURL: BaseURLV1,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK {
		var errRes errorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}
		return fmt.Errorf("Unknown error, status code: %d", res.StatusCode)
	}

	result := v

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return err
	}

	return nil
}

func (c *Client) getProducts() (*Products, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/products", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	var res Products

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}
