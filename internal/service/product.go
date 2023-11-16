package service

import (
	"context"
	"github.com/milkfrogge/coffee-store/internal/model"
)

type ProductService interface {
	CreateCategory(ctx context.Context, category model.CreateCategoryDTO) (string, error)
	CreateProduct(ctx context.Context, product model.CreateProductDTO) (string, error)

	GetAllCategories(ctx context.Context) ([]model.Category, error)
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	GetAllProductsByCategory(ctx context.Context, categoryId string) ([]model.Product, error)
	GetSingleProduct(ctx context.Context, id string) (model.Product, error)

	UpdateProduct(ctx context.Context, product model.Product) error
	AddCountToProduct(ctx context.Context, dto model.CountToProductDTO) error
	SubtractCountToProduct(ctx context.Context, dto model.CountToProductDTO) error
	SubtractCountToProducts(ctx context.Context, dto []model.CountToProductDTO) error
	UpdateCategory(ctx context.Context, category model.Category) error

	DeleteProduct(ctx context.Context, id string) error
	DeleteCategory(ctx context.Context, id string) error
}

//func (i *Implementation) SubtractCountToProduct(ctx context.Context, request *desc.SubtractCountToProductRequest) (*emptypb.Empty, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (i *Implementation) SubtractCountToManyProducts(context.Context, *desc.SubtractCountToManyProductsRequest) (*emptypb.Empty, error) {
//	//TODO implement me
//	panic("implement me")
//}
