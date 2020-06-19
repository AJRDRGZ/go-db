package main

import (
	"log"

	"github.com/AJRDRGZ/go-db/pkg/product"
	"github.com/AJRDRGZ/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		ID:    90,
		Name:  "Curso testing",
		Price: 150,
	}
	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}
}
