package main

import (
	"log"

	"github.com/AJRDRGZ/go-db/pkg/invoiceitem"
	"github.com/AJRDRGZ/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// if err := serviceProduct.Migrate(); err != nil {
	// 	log.Fatalf("product.Migrate: %v", err)
	// }

	// storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	// serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

	// if err := serviceInvoiceHeader.Migrate(); err != nil {
	// 	log.Fatalf("invoiceHeader.Migrate: %v", err)
	// }

	storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceItem.Migrate: %v", err)
	}
}
