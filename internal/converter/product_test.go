package converter

import (
	"github.com/milkfrogge/coffee-store/internal/model"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestCategoriesToProto(t *testing.T) {

	table := []struct {
		Name       string
		Categories []model.Category
	}{
		{
			Name: "good",
			Categories: []model.Category{
				{
					Id:   "0xdeadbeef",
					Name: "1",
				},
			},
		},
		{
			Name: "good",
			Categories: []model.Category{
				{
					Id:   "0xdeadbeef",
					Name: "1",
				},
			},
		},
	}

	for i := 0; i < len(table); i++ {

		proto := CategoriesToProto(table[i].Categories)

		require.Equal(t, &desc.GetAllCategoriesResponse{Category: []*desc.Category{&desc.Category{
			Id:   "0xdeadbeef",
			Name: "1",
		}}}, proto)

	}

}

func TestProductToProto(t *testing.T) {

	cat := model.Category{
		Id:   "0xdeadbeef",
		Name: "category",
	}

	tExp := time.Now()

	prod := model.Product{
		Id:            "0xdeadbeef",
		Name:          "prod",
		Description:   "descr",
		Price:         10,
		Count:         10,
		BaristaNeeded: false,
		KitchenNeeded: false,
		Category:      cat,
		Pics:          nil,
		CreatedAt:     tExp,
	}

	exp := &desc.Product{
		Id: "0xdeadbeef",
		Info: &desc.ProductInfo{
			Name:          "prod",
			Description:   "descr",
			Price:         10,
			Count:         10,
			BaristaNeeded: false,
			KitchenNeeded: false,
			Category: &desc.Category{
				Id:   "0xdeadbeef",
				Name: "category",
			},
			Pics: nil,
		},
		CreatedAt: timestamppb.New(tExp),
	}

	res := ProductToProto(prod)

	require.Equal(t, exp, res)

}

func TestCreateCategoryToDTO(t *testing.T) {

	in := desc.CreateCategoryRequest{Name: "test"}
	exp := model.CreateCategoryDTO{Name: "test"}

	res := CreateCategoryToDTO(&in)

	require.Equal(t, exp, res)

}

func TestProductCategoryToDTO(t *testing.T) {

	in := desc.CreateProductRequest{
		Name:          "test",
		Description:   "descr",
		Price:         10,
		Count:         10,
		BaristaNeeded: false,
		KitchenNeeded: false,
		CategoryId:    "0xdeadbeef",
		Pics:          []string{"test_pic.png"},
	}
	exp := model.CreateProductDTO{
		Name:          "test",
		Description:   "descr",
		Price:         10,
		Count:         10,
		CategoryId:    "0xdeadbeef",
		BaristaNeeded: false,
		KitchenNeeded: false,
		Pics:          []string{"test_pic.png"},
	}

	res := CreateProductToDTO(&in)

	require.Equal(t, exp, res)

}

func TestProductsToProto(t *testing.T) {

	tExp := time.Now()

	table := []struct {
		Name string
		In   []model.Product
		Exp  desc.GetAllProductsResponse
	}{
		{
			Name: "1",
			In: []model.Product{
				{
					Id:            "0xdeadbee1",
					Name:          "1",
					Description:   "desc",
					Price:         10,
					Count:         10,
					BaristaNeeded: false,
					KitchenNeeded: false,
					Category: model.Category{
						Id:   "0xdeadbeef",
						Name: "1",
					},
					Pics:      nil,
					CreatedAt: tExp,
				},
			},
			Exp: desc.GetAllProductsResponse{Product: []*desc.Product{&desc.Product{
				Id: "0xdeadbee1",
				Info: &desc.ProductInfo{
					Name:          "1",
					Description:   "desc",
					Price:         10,
					Count:         10,
					BaristaNeeded: false,
					KitchenNeeded: false,
					Category: &desc.Category{
						Id:   "0xdeadbeef",
						Name: "1",
					},
					Pics: nil,
				},
				CreatedAt: timestamppb.New(tExp),
			}}},
		},
	}

	for i := 0; i < len(table); i++ {
		res := ProductsToProto(table[i].In)

		require.Equal(t, &table[i].Exp, res)
	}

}
