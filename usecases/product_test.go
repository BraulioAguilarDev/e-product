package usecases

import (
	"ecommerce/mocks"
	"ecommerce/model"
	"testing"

	"github.com/google/uuid"
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

func (suite *productUsecaseSuite) TestUpdateProduct() {
	id, _ := uuid.NewRandom()
	product := &model.Product{
		ID:    id,
		Sku:   "sk2",
		Name:  "name2",
		Price: 1,
	}
	suite.repository.On("Update", product).Return(product, nil)

	updated, err := suite.usecase.Update(product)
	suite.NoError(err)
	suite.NotNil(updated)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *productUsecaseSuite) TestDeleteProduct() {
	id, err := uuid.NewRandom()
	suite.NoError(err)
	suite.repository.On("Delete", id).Return(nil)

	err = suite.usecase.Delete(id)
	suite.NoError(err)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *productUsecaseSuite) TestGetProduct() {
	id, _ := uuid.NewRandom()

	product := &model.Product{
		ID:    id,
		Sku:   "sk3",
		Name:  "name3",
		Price: 3,
	}

	suite.repository.On("Get", id).Return(product, nil)

	result, err := suite.repository.Get(id)
	suite.NoError(err)
	suite.Equal(product.Name, result.Name)
	suite.Equal(product.Sku, result.Sku)
	suite.Equal(product.Price, result.Price)
	suite.repository.AssertExpectations(suite.T())
}

func TestProductUsecase(t *testing.T) {
	suite.Run(t, new(productUsecaseSuite))
}
