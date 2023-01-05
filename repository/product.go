package repository

import (
	"context"
	"database/sql"
	"ecommerce/model"

	"github.com/google/uuid"
)

type productRepository struct {
	db *model.Queries
}

type ProductRepository interface {
	Create(sku, name string, price float64) (*model.Product, error)
	Update(product *model.Product) (*model.Product, error)
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*model.Product, error)
}

func InitializeProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: model.New(db),
	}
}

func (p *productRepository) Create(sku, name string, price float64) (*model.Product, error) {
	product, err := p.db.CreateProduct(context.Background(), model.CreateProductParams{
		Sku:   sku,
		Name:  name,
		Price: price,
	})
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productRepository) Update(product *model.Product) (*model.Product, error) {
	updated, err := p.db.UpdateProduct(context.Background(), model.UpdateProductParams{
		Sku:   product.Sku,
		Name:  product.Name,
		Price: product.Price,
		ID:    product.ID,
	})
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (p *productRepository) Delete(id uuid.UUID) error {
	return p.db.DeleteProduct(context.Background(), id)
}

func (p *productRepository) Get(id uuid.UUID) (*model.Product, error) {
	product, err := p.db.GetProduct(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
