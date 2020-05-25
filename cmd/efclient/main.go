package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/nikolasMelui/go_ecommerce_faker_client/internal/app/efclient"
	"syreclabs.com/go/faker"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/efclient.toml", "path to the config file")
}

func main() {

	flag.Parse()

	config := efclient.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	c := efclient.NewClient(config)

	// Get products
	chProducts := make(chan *efclient.Products)
	go func() {
		products, err := c.GetProducts()
		if err != nil {
			log.Fatal(err)
		}
		chProducts <- products
	}()
	products := <-chProducts
	for _, product := range *products {
		fmt.Println("Product ID - ", product.ID)
		fmt.Println("Product Name - ", product.Name)
		fmt.Println("ProductCategory ID - ", string(product.ProductCategory))
	}

	// Get product-categories
	chProductCategories := make(chan *efclient.ProductCategories)
	go func() {
		productCategories, err := c.GetProductCategories()
		if err != nil {
			log.Fatal(err)
		}
		chProductCategories <- productCategories
	}()
	productCategories := <-chProductCategories
	for _, productCategory := range *productCategories {
		fmt.Println("ProductCategory ID - ", productCategory.ID)
		fmt.Println("ProductCategory Name - ", productCategory.Name)
	}

	var productWG sync.WaitGroup
	// Create 50 fake products
	for i := 0; i < 500; i++ {
		productWG.Add(1)
		time.Sleep(time.Millisecond * 50)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			fakeLabels := []int{rand.Intn(5-1+1) + 1, rand.Intn(5-1+1) + 1}
			fakeProduct := efclient.ProductData{
				Name:            faker.Commerce().ProductName(),
				Description:     faker.Lorem().Sentence(20),
				Price:           faker.Commerce().Price(),
				Labels:          fakeLabels,
				ProductCategory: rand.Intn(6-3) + 3,
			}
			_, err := c.CreateProduct(&fakeProduct)
			if err != nil {
				log.Fatal(err)
			}
		}(&productWG)
	}
	productWG.Wait()
}
