package repository

import (
	"ecommerce/config"
	"ecommerce/model"
	"ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

type productRepositorySuite struct {
	// needed to use suite functionalities from testify
	suite.Suite
	// the functionalities we nned to test
	repository ProductRepository

	// helper for db
	cleanupExecutor utils.TruncateTableExecutor
}

func (suite *productRepositorySuite) SetupSuite() {
	// this function runs once before all tests in the suite
	db := config.ConnectDB()
	repository := InitializeProductRepository(db)
	suite.repository = repository

	suite.cleanupExecutor = utils.InitTruncateTableExecutor(db)
}

func (suite *productRepositorySuite) TearDownSuite() {
	suite.cleanupExecutor.TruncateTable([]string{"product"})
}

func (suite *productRepositorySuite) TestCreateProduct() {

	data := model.Product{
		Sku:   "sku1",
		Name:  "name1",
		Price: 1,
	}

	product, err := suite.repository.Create(data.Sku, data.Name, data.Price)
	suite.NoError(err, "no error when create a new product")
	suite.Equal(data.Sku, product.Sku)
	suite.Equal(data.Name, product.Name)
	suite.Equal(data.Price, product.Price)
	suite.NotEmpty(product.ID.Value())
}

func (suite *productRepositorySuite) TestCreateProductError() {
	data := model.Product{Price: 0}

	product, err := suite.repository.Create(data.Sku, data.Name, data.Price)
	suite.Nil(product)
	suite.NotNil(err)
	suite.ErrorContains(err, "sku no valid")
}

func (suite *productRepositorySuite) TestGetProduct() {
	expected := model.Product{
		Sku:   "sku2",
		Name:  "name2",
		Price: 2.5,
	}
	created, err := suite.repository.Create(expected.Sku, expected.Name, expected.Price)
	suite.NoError(err, "no error when get a product")
	suite.NotNil(created)

	product, err := suite.repository.Get(created.ID)
	suite.NoError(err)
	suite.Equal(expected.Sku, product.Sku)
	suite.Equal(expected.Name, product.Name)
	suite.Equal(expected.Price, product.Price)
}

func (suite *productRepositorySuite) TestUpdateProduct() {
	created, err := suite.repository.Create("sku-test", "name-test", 0)
	suite.NoError(err)
	suite.NotNil(created)

	expected := model.Product{
		ID:    created.ID,
		Sku:   "sku3",
		Name:  "name3",
		Price: 3.9,
	}

	updated, err := suite.repository.Update(&expected)
	suite.NoError(err, "no error when update a product")
	suite.Equal(expected.ID, updated.ID)
	suite.Equal(expected.Sku, updated.Sku)
	suite.Equal(expected.Name, updated.Name)
	suite.Equal(expected.Price, updated.Price)
}

func (suite *productRepositorySuite) TestDeleteProduct() {
	created, err := suite.repository.Create("sku-test", "name-test", 6)
	suite.NoError(err)
	suite.NotNil(created)

	err = suite.repository.Delete(created.ID)
	suite.NoError(err, "no error when delete a product")
}

func TestProductRepository(t *testing.T) {
	suite.Run(t, new(productRepositorySuite))
}
