package cli_test

import (
	"fmt"
	"testing"

	"github.com/dnbtr/fullcycle.hexagonal/adapters/cli"
	mock_application "github.com/dnbtr/fullcycle.hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	testController := gomock.NewController(t)
	defer testController.Finish()

	// Defining the product used in this test
	productId := "TestID"
	productName := "Produto Teste CLI"
	productPrice := 25.99
	productStatus := "enabled"

	// Defining the Product Mock
	productMock := mock_application.NewMockProductInterface(testController)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	// Defining the Service Mock
	service := mock_application.NewMockProductServiceInterface(testController)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	// Using gomock.Any() as the value does not matter
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	// First test - CREATE
	resultExpected := fmt.Sprintf("Product ID:%s, Name:%s, Price:%f, Status:%s was created",
		productId, productName, productPrice, productStatus)
	result, err := cli.Run(service, "create", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	// Second test - ENABLE
	// Only the productId is used, other parameters are discarded
	resultExpected = fmt.Sprintf("Product %s has been enabled.", productName)
	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	// Third test - DISABLE
	// Only the productId is used, other parameters are discarded
	resultExpected = fmt.Sprintf("Product %s has been disabled.", productName)
	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	// Fourth test - GET
	resultExpected = fmt.Sprintf("Product ID:%s\nName: %s\nPrice: %f\nStatus: %s",
		productId, productName, productPrice, productStatus)
	result, err = cli.Run(service, "get", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
