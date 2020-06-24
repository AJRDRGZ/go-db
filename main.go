package main

import (
	"fmt"
	"log"

	"github.com/AJRDRGZ/go-db/pkg/product"
	"github.com/AJRDRGZ/go-db/storage"
)

func main() {
	storage.NewMySQLDB()

	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:  "Curso de testing con Go",
		Price: 120,
	}
	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}

	fmt.Printf("%+v\n", m)
}
