package main

import (
	"log"

	"github.com/AJRDRGZ/go-db/pkg/product"
	"github.com/AJRDRGZ/go-db/storage"
)

func main() {
	storage.NewMySQLDB()

	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		ID:    1,
		Name:  "Curso CSS",
		Price: 200,
	}
	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}
}
