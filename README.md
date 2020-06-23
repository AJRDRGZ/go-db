## Migrar tabla de productos

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

if err := serviceProduct.Migrate(); err != nil {
	log.Fatalf("product.Migrate: %v", err)
}
```

## Migrar tabla de invoiceheader

```go
storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

if err := serviceInvoiceHeader.Migrate(); err != nil {
	log.Fatalf("invoiceHeader.Migrate: %v", err)
}
```

## Migrar tabla de invoiceitem

```go
storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

if err := serviceInvoiceItem.Migrate(); err != nil {
	log.Fatalf("invoiceItem.Migrate: %v", err)
}
```

# Crear un producto

```go
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
```

# Consultar Productos

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

ms, err := serviceProduct.GetAll()
if err != nil {
	log.Fatalf("product.GetAll: %v", err)
}

fmt.Println(ms)
```

# Consultar un solo producto

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

m, err := serviceProduct.GetByID(3)
switch {
case errors.Is(err, sql.ErrNoRows):
	fmt.Println("no hay un producto con este id")
case err != nil:
	log.Fatalf("product.GetByID: %v", err)
default:
	fmt.Println(m)
}
```

# Actualizar un producto

```go
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
```

# Eliminar un producto

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

err := serviceProduct.Delete(3)
if err != nil {
	log.Fatalf("product.Delete: %v", err)
}
```

# Crear Factura (tx)

```go
storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
storageInvoice := storage.NewPsqlInvoice(
	storage.Pool(),
	storageHeader,
	storageItems,
)

m := &invoice.Model{
	Header: &invoiceheader.Model{
		Client: "Alexys",
	},
	Items: invoiceitem.Models{
		&invoiceitem.Model{ProductID: 4},
	},
}

serviceInvoice := invoice.NewService(storageInvoice)
if err := serviceInvoice.Create(m); err != nil {
	log.Fatalf("invoice.Create: %v", err)
}
```
