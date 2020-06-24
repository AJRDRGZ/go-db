package storage

import (
	"database/sql"
	"fmt"

	"github.com/AJRDRGZ/go-db/pkg/invoiceitem"
)

const (
	mySQLMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
	mySQLCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES(?, ?)`
)

// MySQLInvoiceItem used for work with mySQL - invoiceItem
type MySQLInvoiceItem struct {
	db *sql.DB
}

// NewMySQLInvoiceItem return a new pointer of MySQLInvoiceItem
func NewMySQLInvoiceItem(db *sql.DB) *MySQLInvoiceItem {
	return &MySQLInvoiceItem{db}
}

// Migrate implement the interface invoiceitem.Storage
func (p *MySQLInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de invoiceItem ejecutada correctamente")
	return nil
}

// CreateTx implement the interface invoiceitem.Storage
func (p *MySQLInvoiceItem) CreateTx(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(mySQLCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {
		result, err := stmt.Exec(headerID, item.ProductID)
		if err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		item.ID = uint(id)
	}

	return nil
}
