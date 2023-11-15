package converter

import (
	"github.com/milkfrogge/coffee-store/internal/model"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProductsToProto(products []model.Product) *desc.GetAllProductsResponse {
	r := desc.GetAllProductsResponse{Product: make([]*desc.Product, 0)}
	for _, product := range products {
		r.Product = append(r.Product, ProductToProto(product))
	}
	return &r
}

func ProductToProto(product model.Product) *desc.Product {
	return &desc.Product{
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
	}
}

func CategoriesToProto(categories []model.Category) *desc.GetAllCategoriesResponse {
	r := desc.GetAllCategoriesResponse{Category: make([]*desc.Category, 0)}
	for _, category := range categories {
		r.Category = append(r.Category, &desc.Category{
			Id:   category.Id,
			Name: category.Name,
		})
	}
	return &r
}

func CreateCategoryToDTO(r *desc.CreateCategoryRequest) model.CreateCategoryDTO {
	out := model.CreateCategoryDTO{}
	out.Name = r.Name
	return out
}

func CreateProductToDTO(r *desc.CreateProductRequest) model.CreateProductDTO {
	out := model.CreateProductDTO{}
	out.Name = r.Name
	out.Pics = r.Pics
	out.Count = r.Count
	out.Price = r.Price
	out.CategoryId = r.CategoryId
	out.Description = r.Description
	return out
}
