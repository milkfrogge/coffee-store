package product

import (
	"context"
	"github.com/milkfrogge/coffee-store/internal/converter"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
)

func (i *Implementation) CreateProduct(ctx context.Context, request *desc.CreateProductRequest) (*desc.CreateProductResponse, error) {
	const op = "Implementation.CreateProduct"
	i.log.Debug(op)

	id, err := i.ProductService.CreateProduct(ctx, converter.CreateProductToDTO(request))
	if err != nil {
		return nil, err
	}
	return &desc.CreateProductResponse{Id: id}, nil
}

func (i *Implementation) CreateCategory(ctx context.Context, request *desc.CreateCategoryRequest) (*desc.CreateCategoryResponse, error) {
	const op = "Implementation.CreateCategory"
	i.log.Debug(op)

	id, err := i.ProductService.CreateCategory(ctx, converter.CreateCategoryToDTO(request))
	if err != nil {
		return nil, err
	}

	return &desc.CreateCategoryResponse{Id: id}, nil
}
