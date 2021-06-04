package application_test

import (
	"testing"

	"github.com/dnbtr/fullcycle.hexagonal/application"
	"github.com/stretchr/testify/require"
)

/*
	This convention is followed by GoLand IDE (appears in autocomplete)
	Testing 'Enable' func from 'Product' struct
	If it was a different package, it would be called 'TestApplicationProduct_Enable'
*/
func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED

	// If price > 10, must not return error
	product.Price = 10
	err := product.Enable()
	require.Nil(t, err)

	// If price = 0, must return error
	product.Price = 0
	err = product.Enable()
	require.Equal(t, "Price must be greater than zero", err.Error())
}
