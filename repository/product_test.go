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

func TestProductRepository(t *testing.T) {
	suite.Run(t, new(productRepositorySuite))
}
