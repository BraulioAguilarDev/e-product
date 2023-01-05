package usecases

import (
	"ecommerce/model"
	"ecommerce/repository"

	"github.com/google/uuid"
)

type productUsecase struct {
	productRepository repository.ProductRepository
}

type ProductUsecase interface {
	Create(sku, name string, price float64) (*model.Product, error)
	Update(product *model.Product) (*model.Product, error)
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*model.Product, error)
}

func InitializeProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepository: repo,
	}
}

func (p *productUsecase) Create(sku, name string, price float64) (*model.Product, error) {
	return p.productRepository.Create(sku, name, price)
}

func (p *productUsecase) Update(product *model.Product) (*model.Product, error) {
	return p.productRepository.Update(product)
}

func (p *productUsecase) Delete(id uuid.UUID) error {
	return p.productRepository.Delete(id)
}

func (p *productUsecase) Get(id uuid.UUID) (*model.Product, error) {
	return p.productRepository.Get(id)
}
