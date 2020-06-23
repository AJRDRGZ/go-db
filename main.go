package main

import (
	"log"

	"github.com/AJRDRGZ/go-db/pkg/invoiceheader"
	"github.com/AJRDRGZ/go-db/pkg/invoiceitem"
	"github.com/AJRDRGZ/go-db/pkg/product"
	"github.com/AJRDRGZ/go-db/storage"
)

func main() {
	storage.NewMySQLDB()

	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

	storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	serviceHeader := invoiceheader.NewService(storageHeader)

	if err := serviceHeader.Migrate(); err != nil {
		log.Fatalf("header.Migrate: %v", err)
	}

	storageItem := storage.NewMySQLInvoiceItem(storage.Pool())
	serviceItem := invoiceitem.NewService(storageItem)

	if err := serviceItem.Migrate(); err != nil {
		log.Fatalf("item.Migrate: %v", err)
	}

}
