package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/dnbtr/fullcycle.hexagonal/adapters/db"
	"github.com/dnbtr/fullcycle.hexagonal/application"
	"github.com/stretchr/testify/require"
)

var TestDatabase *sql.DB

func setUp() {
	TestDatabase, _ = sql.Open("sqlite3", ":memory:")
	createTable(TestDatabase)
	createProduct(TestDatabase)
}

func createTable(db *sql.DB) {
	tableQuery := `CREATE TABLE products (
			"id" string,
			"name" string,
			"price" float,
			"status" string
			);`

	statement, err := db.Prepare(tableQuery)
	if err != nil {	log.Fatal(err.Error()) }

	statement.Exec()
}

func createProduct(db *sql.DB) {
	insertQuery := `INSERT INTO products VALUES("idTest","Test product",0,"disabled");`

	statement, err := db.Prepare(insertQuery)
	if err != nil {	log.Fatal(err.Error()) }

	statement.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer TestDatabase.Close()

	productDb := db.NewProductDb(TestDatabase)

	product, err := productDb.Get("idTest")

	require.Nil(t, err)
	require.Equal(t, "Test product", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer TestDatabase.Close()

	productDb := db.NewProductDb(TestDatabase)

	product := application.NewProduct()
	product.Name = "Test product"
	product.Price = 25

	// Testing create()
	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	// Testing update()
	product.Status = "enabled"
	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}
