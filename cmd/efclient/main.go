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
	"syreclabs.com/go/faker/locales"
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

	// Config the fakers locale
	faker.Locale = locales.Ru

	// Create 10 users
	var userWG sync.WaitGroup
	for i := 0; i < 10; i++ {
		userWG.Add(1)
		time.Sleep(time.Millisecond * 50)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fakeUser := efclient.UserData{
				Username: faker.Internet().UserName(),
				Email:    faker.Internet().Email(),
				Password: "password",
				Phone:    faker.PhoneNumber().CellPhone(),
			}
			log.Println(fakeUser)
			_, err := c.CreateUser(&fakeUser)
			if err != nil {
				log.Print(fmt.Errorf("%v", err))
				// log.Fatal(err)
			}
		}(&userWG)
	}
	userWG.Wait()

	// Create 20 first level product categories
	var firstLevelProductCategoryWG sync.WaitGroup
	for i := 0; i < 30; i++ {
		firstLevelProductCategoryWG.Add(1)
		time.Sleep(time.Millisecond * 50)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fakeProductCategory := efclient.ProductCategoryData{
				Name:        faker.Commerce().Department(),
				Description: faker.Lorem().Sentence(10),
			}
			log.Println(fakeProductCategory)
			_, err := c.CreateProductCategory(&fakeProductCategory)
			if err != nil {
				log.Print(fmt.Errorf("%v", err))
				// log.Fatal(err)
			}
		}(&firstLevelProductCategoryWG)
	}
	firstLevelProductCategoryWG.Wait()

	// Create 40 second level product categories
	var secondLevelProductCategoryWG sync.WaitGroup
	for i := 0; i < 60; i++ {
		secondLevelProductCategoryWG.Add(1)
		time.Sleep(time.Millisecond * 50)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			minParentProductCategory := 1
			maxParentProductCategory := 30
			fakeParentProductCategory := rand.Intn(maxParentProductCategory-minParentProductCategory+1) + minParentProductCategory
			fakeProductCategory := efclient.ProductCategoryData{
				Name:                  faker.Commerce().Department(),
				Description:           faker.Lorem().Sentence(10),
				ParentProductCategory: fakeParentProductCategory,
			}
			log.Println(fakeProductCategory)
			_, err := c.CreateProductCategory(&fakeProductCategory)
			if err != nil {
				log.Print(fmt.Errorf("%v", err))
				// log.Fatal(err)
			}
		}(&secondLevelProductCategoryWG)
	}
	secondLevelProductCategoryWG.Wait()

	// Create 80 third level product categories
	var thirdLevelProductCategoryWG sync.WaitGroup
	for i := 0; i < 90; i++ {
		thirdLevelProductCategoryWG.Add(1)
		time.Sleep(time.Millisecond * 50)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			minParentProductCategory := 31
			maxParentProductCategory := 60
			fakeParentProductCategory := rand.Intn(maxParentProductCategory-minParentProductCategory+1) + minParentProductCategory
			fakeProductCategory := efclient.ProductCategoryData{
				Name:                  faker.Commerce().Department(),
				Description:           faker.Lorem().Sentence(10),
				ParentProductCategory: fakeParentProductCategory,
			}
			log.Println(fakeProductCategory)
			_, err := c.CreateProductCategory(&fakeProductCategory)
			if err != nil {
				log.Print(fmt.Errorf("%v", err))
				// log.Fatal(err)
			}
		}(&thirdLevelProductCategoryWG)
	}
	// Create 20 fake labels
	var labelWG sync.WaitGroup
	for i := 0; i < 30; i++ {
		labelWG.Add(1)
		time.Sleep(time.Millisecond * 50)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fakeLabel := efclient.LabelData{
				Name:        faker.Commerce().Color(),
				Description: faker.Lorem().Sentence(10),
			}
			log.Println(fakeLabel)
			_, err := c.CreateLabel(&fakeLabel)
			if err != nil {
				log.Print(fmt.Errorf("%v", err))
				// log.Fatal(err)
			}
		}(&labelWG)
	}
	thirdLevelProductCategoryWG.Wait()
	labelWG.Wait()

	// Create 500 fake products
	var productWG sync.WaitGroup
	for i := 0; i < 500; i++ {
		productWG.Add(1)
		time.Sleep(time.Millisecond * 50)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			maxLabels := 30
			fakeLabels := []int{rand.Intn(maxLabels-1+1) + 1, rand.Intn(maxLabels-1+1) + 1}
			minProductCategory := 40
			maxProductCategory := 90
			fakeProductCategory := rand.Intn(maxProductCategory-minProductCategory+1) + minProductCategory
			fakeProduct := efclient.ProductData{
				Name:            faker.Commerce().ProductName(),
				Description:     faker.Lorem().Sentence(20),
				Price:           faker.Commerce().Price(),
				Labels:          fakeLabels,
				ProductCategory: fakeProductCategory,
			}
			log.Println(fakeProduct)
			_, err := c.CreateProduct(&fakeProduct)
			if err != nil {
				log.Print(fmt.Errorf("%v", err))
				// log.Fatal(err)
			}
		}(&productWG)
	}
	productWG.Wait()
}
