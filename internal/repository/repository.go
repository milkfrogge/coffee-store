package repository

import (
	"context"
	"github.com/milkfrogge/coffee-store/internal/model"
)

//go:generate mockgen -source repository.go -destination mocks/product_repo.go

type ProductRepository interface {
	CreateCategory(ctx context.Context, category model.CreateCategoryDTO) (string, error)
	CreateProduct(ctx context.Context, product model.CreateProductDTO) (string, error)

	FindAllCategories(ctx context.Context) ([]model.Category, error)
	FindAllProducts(ctx context.Context) ([]model.Product, error)
	FindProductsByCategory(ctx context.Context, categoryId string) ([]model.Product, error)
	FindOneProduct(ctx context.Context, id string) (model.Product, error)

	UpdateCategory(ctx context.Context, category model.Category) error
	UpdateProduct(ctx context.Context, product model.Product) error
	UpdateCountOfProduct(ctx context.Context, product model.Product) error
	UpdateManyCountsOfProduct(ctx context.Context, products []model.Product) error

	DeleteCategory(ctx context.Context, categoryId string) error
	DeleteProduct(ctx context.Context, id string) error
}
