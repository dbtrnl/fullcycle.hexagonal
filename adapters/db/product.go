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

	// Working with prepared statements to avoid SQL Injection
	statement, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id=?")
	if err != nil {	return nil, err	}

	err = statement.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {	return nil, err	}

	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	p.db.QueryRow("select id from products where id=?", product.GetId()).Scan(&rows)

	if rows == 0 {
		_, err := p.create(product)
		if err != nil { return nil, err }
	} else {
		_, err := p.update(product)
		if err != nil { return nil, err }
	}

	return product, nil
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	statement, err := p.db.Prepare(`INSERT INTO products(id, name, price, status) values(?,?,?,?);`)
	if err != nil {	return nil, err	}

	_, err = statement.Exec(
		product.GetId(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	if err != nil { return nil, err }

	err = statement.Close()
	if err != nil { return nil, err }

	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.db.Exec("update products set name=?, price=?, status=? where id=?",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId())
	if err != nil {	return nil, err	}

	return product, nil
}
