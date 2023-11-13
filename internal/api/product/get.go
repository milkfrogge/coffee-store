package product

import (
	"context"
	"github.com/milkfrogge/coffee-store/internal/converter"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
)

func (i *Implementation) GetAllProducts(ctx context.Context, _ *desc.GetAllProductsRequest) (*desc.GetAllProductsResponse, error) {
	const op = "Implementation.GetAllProducts"
	i.log.Info(op)
	products, err := i.ProductService.GetAll(ctx)
	if err != nil {
		i.log.Error(err.Error())
		return &desc.GetAllProductsResponse{}, nil
	}

	return converter.ProductsToProto(products), nil

}
