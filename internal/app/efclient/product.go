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

// Products ...
type Products = []Product

// Product ...
type Product struct {
	ID              int             `json:"id"`
	Title           string          `json:"title"`
	Description     string          `json:"description"`
	Price           float32         `json:"price"`
	Orders          json.RawMessage `json:"orders"`
	Properties      []Property      `json:"property"`
	ProductLabels   []Label         `json:"product_labels"`
	Images          []Image         `json:"images"`
	ProductCategory json.RawMessage `json:"product_category"`
	CreatedAt       string          `json:"created_at"`
	UpdatedAt       string          `json:"updated_at"`
}

// ProductData ...
type ProductData struct {
	Title           string
	Description     string
	Price           float32
	ProductLabels   []int
	ProductCategory int
}

// Property ...
type Property struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
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
		"title":            &productData.Title,
		"description":      &productData.Description,
		"price":            &productData.Price,
		"product_labels":   &productData.ProductLabels,
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

// CreateFakeProducts ...
func (c *Client) CreateFakeProducts(wg *sync.WaitGroup, count int, productLabelsCount int, firstProductCategoryID int, lastProductCategoryID int) int {
	faker.Locale = locales.Ru
	ch := make(chan int, count)
	ch <- 0
	for i := 0; i < count; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond * 50)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			fakeProductLabels := []int{rand.Intn(productLabelsCount-1+1) + 1, rand.Intn(productLabelsCount-1+1) + 1}
			fakeProductCategoryID := rand.Intn(lastProductCategoryID-firstProductCategoryID+1) + firstProductCategoryID
			fakeProduct := ProductData{
				Title:           faker.Commerce().ProductName(),
				Description:     faker.Lorem().Sentence(20),
				Price:           faker.Commerce().Price(),
				ProductLabels:   fakeProductLabels,
				ProductCategory: fakeProductCategoryID,
			}
			log.Println(fakeProduct)
			_, err := c.CreateProduct(&fakeProduct)
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
