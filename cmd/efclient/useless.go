package main

// Get products
// chProducts := make(chan *efclient.Products)
// go func() {
// 	products, err := c.GetProducts()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	chProducts <- products
// }()
// products := <-chProducts
// for _, product := range *products {
// 	log.Println("Product ID - ", product.ID)
// 	log.Println("Product Name - ", product.Name)
// 	log.Println("ProductCategory ID - ", string(product.ProductCategory))
// }

// Get product-categories
// chProductCategories := make(chan *efclient.ProductCategories)
// go func() {
// 	productCategories, err := c.GetProductCategories()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	chProductCategories <- productCategories
// }()
// productCategories := <-chProductCategories
// for _, productCategory := range *productCategories {
// 	log.Println("ProductCategory ID - ", productCategory.ID)
// 	log.Println("ProductCategory Name - ", productCategory.Name)
// }
