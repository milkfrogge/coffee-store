package converter

import (
	"github.com/milkfrogge/coffee-store/internal/model"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProductsToProto(products []model.Product) *desc.GetAllProductsResponse {
	r := desc.GetAllProductsResponse{Product: make([]*desc.Product, 0)}
	for _, product := range products {
		var updatedAt *timestamppb.Timestamp
		if product.UpdatedAt != nil {
			updatedAt = timestamppb.New(*product.UpdatedAt)
		}
		r.Product = append(r.Product, &desc.Product{
			Id: product.Id,
			Info: &desc.ProductInfo{
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Count:       product.Count,
				Category: &desc.Category{
					Id:   product.Category.Id,
					Name: product.Category.Name,
				},
				Pics: product.Pics,
			},
			CreatedAt: timestamppb.New(product.CreatedAt),
			UpdatedAt: updatedAt,
		})
	}
	return &r
}
