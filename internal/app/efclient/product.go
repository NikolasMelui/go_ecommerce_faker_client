package efclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Products ...
type Products = []Product

// Product ...
type Product struct {
	ID                int             `json:"id"`
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	Price             int             `json:"price"`
	CreatedAt         string          `json:"created_at"`
	UpdatedAt         string          `json:"updated_at"`
	ProductCategory   json.RawMessage `json:"product_category"`
	ProductProperties []Property      `json:"property"`
	ProductImages     []Image         `json:"images"`
	ProductOrders     []Order         `json:"orders"`
	ProductCarts      []Cart          `json:"carts"`
	ProductLables     []Label         `json:"lables"`
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
