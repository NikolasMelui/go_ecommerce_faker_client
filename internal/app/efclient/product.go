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

// LabelData ...
type LabelData struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ProductCategoryData ...
type ProductCategoryData struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	ParentProductCategory int    `json:"parent_product_category"`
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

// CreateProductCategory ...
func (c *Client) CreateProductCategory(productCategoryData *ProductCategoryData) (*ProductCategory, error) {

	emptyProductParentCategory := 0
	var checkedProductParentCategory *int
	if &productCategoryData.ParentProductCategory != nil {
		checkedProductParentCategory = &productCategoryData.ParentProductCategory
	} else {
		checkedProductParentCategory = &emptyProductParentCategory
	}

	requestData := map[string]interface{}{
		"name":                    &productCategoryData.Name,
		"description":             &productCategoryData.Description,
		"product_parent_category": checkedProductParentCategory,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/product-categories", c.BaseURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	var res ProductCategory
	if err := c.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateLabel ...
func (c *Client) CreateLabel(labelData *LabelData) (*Label, error) {

	requestData := map[string]interface{}{
		"name":        &labelData.Name,
		"description": &labelData.Description,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/product-labels", c.BaseURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	var res Label
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
