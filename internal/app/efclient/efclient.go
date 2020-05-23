package efclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Client ...
type Client struct {
	config     *Config
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
	ProductCategory   json.RawMessage   `json:"product_category"`
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

// ProductCategories ...
type ProductCategories = []ProductCategory

// ProductCategory ...
type ProductCategory struct {
	ID                        int               `json:"id"`
	Name                      string            `json:"name"`
	Description               string            `json:"description"`
	ParentProductCategory     json.RawMessage   `json:"parent_product_category"`
	CreatedAt                 string            `json:"created_at"`
	UpdatedAt                 string            `json:"updated_at"`
	Products                  []Product         `json:"products"`
	ChildrenProductCategories []ProductCategory `json:"children_product_categories"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewClient ...
func NewClient(config *Config) *Client {
	return &Client{
		BaseURL: config.BaseURL,
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

	if res.StatusCode != http.StatusOK {
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

// GetProducts ...
func (c *Client) GetProducts() (*Products, error) {

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

// GetProductCategories ...
func (c *Client) GetProductCategories() (*ProductCategories, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/product-categories", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	var res ProductCategories

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}
