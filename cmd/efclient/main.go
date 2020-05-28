package main

import (
	"flag"
	"log"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/nikolasMelui/go_ecommerce_faker_client/internal/app/efclient"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/efclient.toml", "path to the config file")
}

func main() {

	// Parse flags and create Config
	flag.Parse()
	config := efclient.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// Create Client
	c := efclient.NewClient(config)

	// Create fake users
	var createFakeUsersWG sync.WaitGroup
	fakeUsersCount := c.CreateFakeUsers(&createFakeUsersWG, 10)
	log.Println(fakeUsersCount)

	// Create fake counterparties
	var createFakeCounterpartiesWG sync.WaitGroup
	fakeCounterpartiesCount := c.CreateFakeCounterparties(&createFakeCounterpartiesWG, 20)
	log.Println(fakeCounterpartiesCount)

	// Create first level fake product categories
	var createFirstLvlFakeProductCategoryWG sync.WaitGroup
	firstLvlFakeProductCategoriesCount := c.CreateFakeProductCategories(&createFirstLvlFakeProductCategoryWG, 10, 0, 0)
	log.Println(firstLvlFakeProductCategoriesCount)

	// Create second level fake product categories
	var createSecondLvlFakeProductCategoryWG sync.WaitGroup
	secondLvlFakeProductCategoriesCount := c.CreateFakeProductCategories(&createSecondLvlFakeProductCategoryWG, 20, 1, 10)
	log.Println(secondLvlFakeProductCategoriesCount)

	// Create third level fake product categories
	var createThirdLvlFakeProductCategoryWG sync.WaitGroup
	thirdLvlFakeProductCategoriesCount := c.CreateFakeProductCategories(&createThirdLvlFakeProductCategoryWG, 80, 11, 30)
	log.Println(thirdLvlFakeProductCategoriesCount)

	// Create fake labels
	var createFakeLabelsWG sync.WaitGroup
	fakeLabelsCount := c.CreateFakeLabels(&createFakeLabelsWG, 20)
	log.Println(fakeLabelsCount)

	// Create fake products
	var createFakeProductsWG sync.WaitGroup
	fakeProductsCount := c.CreateFakeProducts(&createFakeProductsWG, 100, 20, 21, 90)
	log.Println(fakeProductsCount)

}
