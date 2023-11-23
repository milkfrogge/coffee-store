package product

import (
	"context"
	"github.com/milkfrogge/coffee-store/internal/converter"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) AddCountToProduct(ctx context.Context, request *desc.AddCountToProductRequest) (*emptypb.Empty, error) {
	const op = "Implementation.AddCountToProduct"
	i.log.Debug(op)

	err := i.ProductService.AddCountToProduct(ctx, converter.AddToProductToDTO(request))
	if err != nil {
		i.log.Error(err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil

}

func (i *Implementation) SubtractCountToProduct(ctx context.Context, request *desc.SubtractCountToProductRequest) (*emptypb.Empty, error) {
	const op = "Implementation.SubtractCountToProduct"
	i.log.Debug(op)

	err := i.ProductService.SubtractCountToProduct(ctx, converter.SubtractProductToDTO(request))
	if err != nil {
		i.log.Error(err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (i *Implementation) SubtractCountToManyProducts(ctx context.Context, request *desc.SubtractCountToManyProductsRequest) (*emptypb.Empty, error) {
	const op = "Implementation.SubtractCountToManyProducts"
	i.log.Debug(op)

	err := i.ProductService.SubtractCountToProducts(ctx, converter.SubtractManyProductsToDTO(request))
	if err != nil {
		i.log.Error(err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (i *Implementation) UpdateCategory(ctx context.Context, request *desc.UpdateCategoryRequest) (*emptypb.Empty, error) {
	const op = "Implementation.UpdateCategory"
	i.log.Debug(op)

	err := i.ProductService.UpdateCategory(ctx, converter.ProtoToCategory(request.Category))
	if err != nil {
		i.log.Error(err.Error())
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (i *Implementation) UpdateProduct(ctx context.Context, request *desc.UpdateProductRequest) (*emptypb.Empty, error) {
	const op = "Implementation.UpdateProduct"
	i.log.Debug(op)

	err := i.ProductService.UpdateProduct(ctx, converter.ProtoToProduct(request.Product))
	if err != nil {
		i.log.Error(err.Error())
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
