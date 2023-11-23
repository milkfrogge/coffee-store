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
			Name:          product.Name,
			Description:   product.Description,
			Price:         product.Price,
			Count:         product.Count,
			BaristaNeeded: product.BaristaNeeded,
			KitchenNeeded: product.KitchenNeeded,
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
	out.BaristaNeeded = r.BaristaNeeded
	out.KitchenNeeded = r.KitchenNeeded
	out.CategoryId = r.CategoryId
	out.Description = r.Description
	return out
}

func AddToProductToDTO(r *desc.AddCountToProductRequest) model.CountToProductDTO {
	return model.CountToProductDTO{
		Id:    r.Product.Id,
		Count: r.Product.Count,
	}
}

func SubtractProductToDTO(r *desc.SubtractCountToProductRequest) model.CountToProductDTO {
	return model.CountToProductDTO{
		Id:    r.Product.Id,
		Count: r.Product.Count,
	}
}

func SubtractManyProductsToDTO(r *desc.SubtractCountToManyProductsRequest) []model.CountToProductDTO {

	out := make([]model.CountToProductDTO, len(r.Products))

	for i := 0; i < len(r.Products); i++ {
		out[i] = model.CountToProductDTO{
			Id:    r.Products[i].Id,
			Count: r.Products[i].Count,
		}
	}

	return out

}

func ProtoToProduct(product *desc.Product) model.Product {
	return model.Product{
		Id:            product.Id,
		Name:          product.Info.Name,
		Description:   product.Info.Description,
		Price:         product.Info.Price,
		Count:         product.Info.Price,
		BaristaNeeded: product.Info.BaristaNeeded,
		KitchenNeeded: product.Info.KitchenNeeded,
		Category: model.Category{
			Id:   product.Info.Category.Id,
			Name: product.Info.Category.Name,
		},
		Pics:      product.Info.Pics,
		CreatedAt: product.CreatedAt.AsTime(),
	}
}

func ProtoToCategory(category *desc.Category) model.Category {
	return model.Category{
		Id:   category.Id,
		Name: category.Name,
	}
}
