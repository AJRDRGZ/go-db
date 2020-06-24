package main

import (
	"fmt"
	"log"

	"github.com/AJRDRGZ/go-db/pkg/product"
	"github.com/AJRDRGZ/go-db/storage"
)

func main() {
	driver := storage.Postgres

	storage.New(driver)

	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}

	serviceProduct := product.NewService(myStorage)

	m, err := serviceProduct.GetByID(4)
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(m)
}
