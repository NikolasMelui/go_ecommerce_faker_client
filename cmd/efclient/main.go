package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/nikolasMelui/go_ecommerce_faker_client/internal/app/efclient"
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
	products, err := c.GetProducts()
	if err != nil {
		log.Fatal(err)
	}
	for _, product := range *products {
		fmt.Println("Product ID - ", product.ID)
		fmt.Println("Product Name - ", product.Name)
		fmt.Println("ProductCategory ID - ", string(product.ProductCategory))
	}

	// Get product-categories
	productCategories, err := c.GetProductCategories()
	if err != nil {
		log.Fatal(err)
	}

	for _, productCategory := range *productCategories {
		fmt.Println("ProductCategory ID - ", productCategory.ID)
		fmt.Println("ProductCategory Name - ", productCategory.Name)
	}
}
