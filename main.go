package main

import (
	"log"

	"github.com/AJRDRGZ/go-db/pkg/invoice"
	"github.com/AJRDRGZ/go-db/pkg/invoiceheader"
	"github.com/AJRDRGZ/go-db/pkg/invoiceitem"
	"github.com/AJRDRGZ/go-db/storage"
)

func main() {
	storage.NewMySQLDB()

	storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	storageItems := storage.NewMySQLInvoiceItem(storage.Pool())
	storageInvoice := storage.NewMySQLInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Alvaro Felipe",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 2},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}
}
