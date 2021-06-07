package dto

import "github.com/dnbtr/fullcycle.hexagonal/application"

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

// Bind the Data received in the Request to the application's Product object
func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" { product.ID = p.ID }
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	// Check if the received product is valid
	_, err := product.IsValid()
	if err != nil { return &application.Product{}, err}

	return product, nil
}