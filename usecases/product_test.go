package usecases

import (
	"ecommerce/mocks"
	"ecommerce/model"
	"testing"

	"github.com/stretchr/testify/suite"
)

type productUsecaseSuite struct {
	suite.Suite
	repository *mocks.ProductRepository

	usecase ProductUsecase
}

func (suite *productUsecaseSuite) SetupTest() {
	repository := new(mocks.ProductRepository)
	// inject the repository to usecase
	usecase := InitializeProductUsecase(repository)

	suite.repository = repository
	suite.usecase = usecase
}

func (suite *productUsecaseSuite) TestCreateProduct() {
	product := model.Product{
		Sku:   "sku1",
		Name:  "name1",
		Price: 1,
	}

	// We need to specify what function we are going to call inside usecase
	suite.repository.On("Create", product.Sku, product.Name, product.Price).Return(&product, nil)

	// the real operation that we will test
	created, err := suite.usecase.Create(product.Sku, product.Name, product.Price)

	// assertions to make sure that all works correctly
	suite.NoError(err)
	suite.Equal(product.Sku, created.Sku)
	suite.Equal(product.Name, created.Name)
	suite.Equal(product.Price, created.Price)
	suite.repository.AssertExpectations(suite.T())
}

func TestProductUsecase(t *testing.T) {
	suite.Run(t, new(productUsecaseSuite))
}
