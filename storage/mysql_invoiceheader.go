package storage

import (
	"database/sql"
	"fmt"

	"github.com/AJRDRGZ/go-db/pkg/invoiceheader"
)

const (
	mySQLMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
	mySQLCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES(?)`
)

// MySQLInvoiceHeader used for work with mySQL - invoiceHeader
type MySQLInvoiceHeader struct {
	db *sql.DB
}

// NewMySQLInvoiceHeader return a new pointer of MySQLInvoiceHeader
func NewMySQLInvoiceHeader(db *sql.DB) *MySQLInvoiceHeader {
	return &MySQLInvoiceHeader{db}
}

// Migrate implement the interface invoiceHeader.Storage
func (p *MySQLInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de invoiceHeader ejecutada correctamente")
	return nil
}

// CreateTx implement the interface invoiceHeader.Storage
func (p *MySQLInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(mySQLCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(m.Client)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	m.ID = uint(id)
	return nil
}
