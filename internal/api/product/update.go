package product

import (
	"context"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateProduct(ctx context.Context, request *desc.UpdateProductRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (i *Implementation) AddCountToProduct(ctx context.Context, request *desc.AddCountToProductRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (i *Implementation) UpdateCategory(ctx context.Context, request *desc.UpdateCategoryRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
