package main

import (
	"database/sql"

	db2 "github.com/cassioglay/arquitetura-hexagonal/adapters/db"
	"github.com/cassioglay/arquitetura-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbApter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbApter)
	product, _ := productService.Create("Product Exemplo", 10)

	productService.Enable(product)

}
