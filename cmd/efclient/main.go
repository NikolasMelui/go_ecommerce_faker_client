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
	usersPlan := 10
	var createUsersWG sync.WaitGroup
	usersCount := c.CreateFakeUsers(&createUsersWG, usersPlan)
	log.Println(usersCount)

	// Create fake counterparties
	counterpatriesPlan := usersPlan * 2
	var createCounterpartiesWG sync.WaitGroup
	counterpartiesCount := c.CreateFakeCounterparties(&createCounterpartiesWG, counterpatriesPlan)
	log.Println(counterpartiesCount)

	// Create fake counterparty documents
	counterpartyDocumentsPlan := counterpatriesPlan * 2
	var createCounterpartyDocumentsWG sync.WaitGroup
	counterpartyDocumentsCount := c.CreateFakeCounterpartyDocuments(&createCounterpartyDocumentsWG, counterpartyDocumentsPlan)
	log.Println(counterpartyDocumentsCount)

	// Create first level fake product categories (PC)
	firstLvlPCPlan := 10
	var createFirstLvlPCWG sync.WaitGroup
	firstLvlPCFirstParentID := 0
	firstLvlPCLastParentID := 0
	firstLvlPCCount := c.CreateFakeProductCategories(&createFirstLvlPCWG, firstLvlPCPlan, firstLvlPCFirstParentID, firstLvlPCLastParentID)
	log.Println(firstLvlPCCount)

	// Create second level fake product categories (PC)
	secondLvlPCPlan := firstLvlPCPlan * 2
	secondLvlPCFirstParentID := 1
	secondLvlPCLastParentID := firstLvlPCPlan
	var createSecondLvlPCWG sync.WaitGroup
	secondLvlPCCount := c.CreateFakeProductCategories(&createSecondLvlPCWG, secondLvlPCPlan, secondLvlPCFirstParentID, secondLvlPCLastParentID)
	log.Println(secondLvlPCCount)

	// Create third level fake product categories (PC)
	thirdLvlPCPlan := secondLvlPCPlan * 4
	thirdLvlPCFirstParentID := firstLvlPCPlan + 1
	thirdLvlPCLastParentID := secondLvlPCPlan
	var createThirdLvlPCWG sync.WaitGroup
	thirdLvlPCCount := c.CreateFakeProductCategories(&createThirdLvlPCWG, thirdLvlPCPlan, thirdLvlPCFirstParentID, thirdLvlPCLastParentID)
	log.Println(thirdLvlPCCount)

	// Create fake labels
	labelsNeed := 20
	var createLabelsWG sync.WaitGroup
	labelsCount := c.CreateFakeLabels(&createLabelsWG, labelsNeed)
	log.Println(labelsCount)

	// Create fake products
	productsNeed := 100
	var createProductsWG sync.WaitGroup
	firstProductCategoryID := firstLvlPCPlan + secondLvlPCPlan + 1
	lastProductCategoryID := firstLvlPCPlan + secondLvlPCPlan + thirdLvlPCPlan
	productsCount := c.CreateFakeProducts(&createProductsWG, productsNeed, labelsNeed, firstProductCategoryID, lastProductCategoryID)
	log.Println(productsCount)

}
