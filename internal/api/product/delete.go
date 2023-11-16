package product

import (
	"context"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteProduct(ctx context.Context, request *desc.DeleteCategoryRequest) (*emptypb.Empty, error) {
	const op = "Implementation.DeleteCategory"
	i.log.Debug(op)

	return &emptypb.Empty{}, i.ProductService.DeleteProduct(ctx, request.Id)

}

func (i *Implementation) DeleteCategory(ctx context.Context, request *desc.DeleteCategoryRequest) (*emptypb.Empty, error) {
	const op = "Implementation.DeleteCategory"
	i.log.Debug(op)

	return &emptypb.Empty{}, i.ProductService.DeleteCategory(ctx, request.Id)

}
