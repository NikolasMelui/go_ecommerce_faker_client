package efclient

import (
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
