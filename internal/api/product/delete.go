package product

import (
	"context"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteProduct(ctx context.Context, request *desc.DeleteCategoryRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (i *Implementation) DeleteCategory(ctx context.Context, request *desc.DeleteCategoryRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
