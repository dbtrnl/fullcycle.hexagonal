package application_test

import (
	"testing"

	"github.com/dnbtr/fullcycle.hexagonal/application"
	mock_application "github.com/dnbtr/fullcycle.hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	testController := gomock.NewController(t)
	defer testController.Finish()

	product := mock_application.NewMockProductInterface(testController)
	persistence := mock_application.NewMockProductPersistenceInterface(testController)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Get("teste")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	testController := gomock.NewController(t)
	defer testController.Finish()

	product := mock_application.NewMockProductInterface(testController)
	persistence := mock_application.NewMockProductPersistenceInterface(testController)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Create("teste", 29)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
