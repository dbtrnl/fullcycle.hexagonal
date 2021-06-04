package application_test

import (
	"fmt"
	"testing"

	"github.com/dnbtr/fullcycle.hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

/*
	This convention is followed by GoLand IDE (appears in autocomplete)
	Testing 'Enable' func from 'Product' struct
	If it was a different package, it would be called 'TestApplicationProduct_Enable'
*/
func TestProduct_Enable(t *testing.T) {
	fmt.Printf("Running TestProduct_Enable test suite... ")
	product := application.Product{}
	product.Name = "TestProduct_Enable name"
	product.Status = application.DISABLED

	// If price > 10, must not return error
	product.Price = 10
	err := product.Enable()
	require.Nil(t, err)

	// If price = 0, must return error
	product.Price = 0
	err = product.Enable()
	require.Equal(t, "Price must be greater than zero to enable product", err.Error())
	fmt.Printf("OK\n\n")
}
func TestProduct_Disable(t *testing.T) {
	fmt.Printf("Running TestProduct_Disable test suite... ")
	product := application.Product{}
	product.Name = "TestProduct_Disable name"
	product.Status = application.ENABLED

	// If price is 0, must be disabled
	product.Price = 0
	err := product.Disable()
	require.Nil(t, err)

	// If price != must not be disabled
	product.Price = 10
	err = product.Disable()
	require.Equal(t, "Price must be zero to disable product", err.Error())
	fmt.Printf("OK\n\n")
}

func TestProduct_IsValid(t *testing.T) {
	fmt.Printf("Running TestProduct_IsValid test suite... ")
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "TestProduct_IsValid name"
	product.Status = application.DISABLED
	product.Price = 50

	// Valid product test
	_, err := product.IsValid()
	require.Nil(t, err)

	// Invalid status test
	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "The status must be 'enabled' or 'disabled", err.Error())

	// Enabled status test
	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	// Negative price test
	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "The price must be equal or greater than zero", err.Error())
	fmt.Printf("OK\n\n")
}

func TestProduct_GetID(t *testing.T) {
	fmt.Printf("Running TestProduct_GetID test suite... ")
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "TestProduct_GetID Name"
	product.Status = application.DISABLED
	product.Price = 0

	// Testing GetID
	id := product.GetId()
	require.Equal(t, id, product.ID)
	fmt.Printf("OK\n\n")
}
func TestProduct_GetName(t *testing.T) {
	fmt.Printf("Running TestProduct_GetName test suite... ")
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "TestProduct_GetName Name"
	product.Status = application.DISABLED
	product.Price = 0

	// Testing GetName
	name := product.GetName()
	require.Equal(t, name, product.Name)
	fmt.Printf("OK\n\n")
}
func TestProduct_GetStatus(t *testing.T) {
	fmt.Printf("Running TestProduct_GetStatus test suite... ")
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "TestProduct_GetStatus Name"
	product.Status = application.DISABLED
	product.Price = 0

	// Testing GetStatus
	status := product.GetStatus()
	require.Equal(t, status, product.Status)
	fmt.Printf("OK\n\n")

}
func TestProduct_GetPrice(t *testing.T) {
	fmt.Printf("Running TestProduct_GetPrice test suite... ")
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "TestProduct_GetPrice Name"
	product.Status = application.DISABLED
	product.Price = 842

	// Testing GetPrice
	price := product.GetPrice()
	require.NotEqual(t, price, product.Price)
	fmt.Printf("OK\n\n")
}
