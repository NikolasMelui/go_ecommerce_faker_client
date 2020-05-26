package efclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"
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

	requestData := map[string]interface{}{
		"name":                    &productCategoryData.Name,
		"description":             &productCategoryData.Description,
		"parent_product_category": &productCategoryData.ParentProductCategory,
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

// CreateFakeProductCategories ...
func (c *Client) CreateFakeProductCategories(wg *sync.WaitGroup, count int, firstParentID int, lastParentID int) int {
	faker.Locale = locales.Ru
	ch := make(chan int, count)
	ch <- 0
	for i := 0; i < count; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond * 50)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			parentID := 0
			if firstParentID > 0 {
				rand.Seed(time.Now().UnixNano())
				parentID = rand.Intn(lastParentID-firstParentID+1) + firstParentID
			}
			fakeProductCategory := ProductCategoryData{
				Name:                  faker.Commerce().Department(),
				Description:           faker.Lorem().Sentence(10),
				ParentProductCategory: parentID,
			}
			log.Println(fakeProductCategory)
			_, err := c.CreateProductCategory(&fakeProductCategory)
			if err != nil {
				log.Print(fmt.Errorf("%v", err))
				// log.Fatal(err)
			} else {
				counter := <-ch
				ch <- counter + 1
			}
		}(wg)
	}
	wg.Wait()
	close(ch)
	return <-ch
}
