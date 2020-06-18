package main

import (
	"fmt"
	"log"

	"github.com/AJRDRGZ/go-db/pkg/product"
	"github.com/AJRDRGZ/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:         "Curso de db con Go",
		Price:        70,
		Observations: "on fire",
	}
	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}

	fmt.Printf("%+v\n", m)
}
