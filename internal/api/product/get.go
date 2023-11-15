package product

import (
	"context"
	"github.com/milkfrogge/coffee-store/internal/converter"
	"github.com/milkfrogge/coffee-store/pkg/jaeger"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetSingleProduct(ctx context.Context, request *desc.GetSingleProductRequest) (*desc.GetSingleProductResponse, error) {
	const op = "Implementation.GetSingleProduct"
	i.log.Debug(op)

	ctx, err := jaeger.ExtractMetaFromGRPC(ctx)
	if err != nil {
		i.log.Error(err.Error())
	}
	ctx, span := i.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	product, err := i.ProductService.GetSingleProduct(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &desc.GetSingleProductResponse{Product: converter.ProductToProto(product)}, nil
}

func (i *Implementation) GetAllProductsByCategory(ctx context.Context, request *desc.GetAllProductsByCategoryRequest) (*desc.GetAllProductsResponse, error) {
	const op = "Implementation.GetAllProductsByCategory"
	i.log.Debug(op)

	ctx, err := jaeger.ExtractMetaFromGRPC(ctx)
	if err != nil {
		i.log.Error(err.Error())
	}
	ctx, span := i.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	products, err := i.ProductService.GetAllProductsByCategory(ctx, request.Id)
	if err != nil {
		i.log.Error(err.Error())
		return &desc.GetAllProductsResponse{}, nil
	}

	span.AddEvent("convert to protobuf")

	return converter.ProductsToProto(products), nil

}

func (i *Implementation) GetAllCategories(ctx context.Context, _ *emptypb.Empty) (*desc.GetAllCategoriesResponse, error) {
	const op = "Implementation.GetAllCategories"
	i.log.Debug(op)

	ctx, err := jaeger.ExtractMetaFromGRPC(ctx)
	if err != nil {
		i.log.Error(err.Error())
	}
	ctx, span := i.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	categories, err := i.ProductService.GetAllCategories(ctx)
	if err != nil {
		i.log.Error(err.Error())
		return nil, err
	}

	span.AddEvent("convert to protobuf")

	return converter.CategoriesToProto(categories), nil
}

func (i *Implementation) GetAllProducts(ctx context.Context, _ *emptypb.Empty) (*desc.GetAllProductsResponse, error) {
	const op = "Implementation.GetAllProducts"
	i.log.Debug(op)

	ctx, err := jaeger.ExtractMetaFromGRPC(ctx)
	if err != nil {
		i.log.Error(err.Error())
	}
	ctx, span := i.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	products, err := i.ProductService.GetAllProducts(ctx)
	if err != nil {
		i.log.Error(err.Error())
		return &desc.GetAllProductsResponse{}, nil
	}

	span.AddEvent("convert to protobuf")

	return converter.ProductsToProto(products), nil

}
