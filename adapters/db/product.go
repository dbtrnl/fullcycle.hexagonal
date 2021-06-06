package db

import (
	"database/sql"

	"github.com/dnbtr/fullcycle.hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	statement, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id=?")
	if err != nil {
		return nil, err
	}

	err = statement.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
