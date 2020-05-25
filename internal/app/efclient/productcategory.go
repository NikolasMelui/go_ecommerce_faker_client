package efclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

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

// ProductCategoryData ...
type ProductCategoryData struct {
	Name                  string `json:"name"`
	Description           string `json:"description"`
	ParentProductCategory int    `json:"parent_product_category"`
}

// GetProductCategories ...
func (c *Client) GetProductCategories() (*ProductCategories, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/product-categories", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	var res ProductCategories

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
