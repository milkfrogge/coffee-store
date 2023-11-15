package product

import (
	"context"
	"github.com/milkfrogge/coffee-store/internal/converter"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetSingleProduct(ctx context.Context, request *desc.GetSingleProductRequest) (*desc.GetSingleProductResponse, error) {
	const op = "Implementation.GetSingleProduct"
	i.log.Debug(op)
	product, err := i.ProductService.GetSingleProduct(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &desc.GetSingleProductResponse{Product: converter.ProductToProto(product)}, nil
}

func (i *Implementation) GetAllProductsByCategory(ctx context.Context, request *desc.GetAllProductsByCategoryRequest) (*desc.GetAllProductsResponse, error) {
	const op = "Implementation.GetAllProductsByCategory"
	i.log.Debug(op)
	products, err := i.ProductService.GetAllProductsByCategory(ctx, request.Id)
	if err != nil {
		i.log.Error(err.Error())
		return &desc.GetAllProductsResponse{}, nil
	}

	return converter.ProductsToProto(products), nil

}

func (i *Implementation) GetAllCategories(ctx context.Context, empty *emptypb.Empty) (*desc.GetAllCategoriesResponse, error) {
	const op = "Implementation.GetAllCategories"
	i.log.Debug(op)

	categories, err := i.ProductService.GetAllCategories(ctx)
	if err != nil {
		i.log.Error(err.Error())
		return nil, err
	}

	return converter.CategoriesToProto(categories), nil
}

func (i *Implementation) GetAllProducts(ctx context.Context, _ *emptypb.Empty) (*desc.GetAllProductsResponse, error) {
	const op = "Implementation.GetAllProducts"
	i.log.Debug(op)
	products, err := i.ProductService.GetAllProducts(ctx)
	if err != nil {
		i.log.Error(err.Error())
		return &desc.GetAllProductsResponse{}, nil
	}

	return converter.ProductsToProto(products), nil

}
