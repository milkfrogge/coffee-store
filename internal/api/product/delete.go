package product

import (
	"context"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteProduct(ctx context.Context, request *desc.DeleteCategoryRequest) (*emptypb.Empty, error) {
	const op = "Implementation.DeleteCategory"
	i.log.Debug(op)
	err := i.ProductService.DeleteProduct(ctx, request.Id)
	if err != nil {
		i.log.Error(err.Error())
	}
	return &emptypb.Empty{}, err

}

func (i *Implementation) DeleteCategory(ctx context.Context, request *desc.DeleteCategoryRequest) (*emptypb.Empty, error) {
	const op = "Implementation.DeleteCategory"
	i.log.Debug(op)
	err := i.ProductService.DeleteCategory(ctx, request.Id)
	if err != nil {
		i.log.Error(err.Error())
	}
	return &emptypb.Empty{}, err

}
