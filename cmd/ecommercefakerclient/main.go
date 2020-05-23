package main

import (
	"fmt"
	"log"

	"github.com/nikolasMelui/go_ecommerce_faker_client/internal/app/ecommercefakerclient"
)

func main() {
	c := ecommercefakerclient.NewClient()

	products, err := c.GetProducts()
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range *products {
		fmt.Println("Product ID - ", product.ID)
		fmt.Println("Product Name - ", product.Name)
	}
}
