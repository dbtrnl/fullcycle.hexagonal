package main

import (
	"database/sql"

	db2 "github.com/dnbtr/fullcycle.hexagonal/adapters/db"
	"github.com/dnbtr/fullcycle.hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Produto Exemplo", 666)

	productService.Enable(product)
}