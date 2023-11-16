package product

import (
	"github.com/milkfrogge/coffee-store/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateCreateProduct(t *testing.T) {
	table := []model.CreateProductDTO{
		{
			Name:        "first",
			Description: "lorem",
			Price:       1,
			Count:       1,
			CategoryId:  "0xdeadbeef",
			Pics:        nil,
		},
		{
			Name:        "second",
			Description: "lorem",
			Price:       1,
			Count:       1,
			CategoryId:  "0xdeadbeef",
			Pics:        []string{"image1"},
		},
	}
	for i := 0; i < len(table); i++ {
		err := ValidateCreateProduct(table[i])
		require.NoError(t, err)
	}
}

func TestValidateCreateProductError(t *testing.T) {
	table := []model.CreateProductDTO{
		{
			Name:        "first",
			Description: "lorem",
			Price:       1,
			Count:       0,
			CategoryId:  "0xdeadbeef",
			Pics:        nil,
		},
		{
			Name:        "second",
			Description: "lorem",
			Price:       0,
			Count:       1,
			CategoryId:  "0xdeadbeef",
			Pics:        []string{"image1"},
		},
		{
			Name:          "third",
			Description:   "lorem",
			Price:         1,
			Count:         1,
			BaristaNeeded: true,
			KitchenNeeded: true,
			CategoryId:    "0xdeadbeef",
			Pics:          []string{"image1"},
		},
	}

	expErrors := []error{ErrWrongCount, ErrWrongPrice, ErrBaristaAndKitchenNeeded}

	for i := 0; i < len(table); i++ {
		err := ValidateCreateProduct(table[i])
		require.EqualError(t, err, expErrors[i].Error())
	}
}

func TestValidateCreateCategory(t *testing.T) {

	table := []model.CreateCategoryDTO{
		{Name: "Coffee"},
		{Name: "Drinks"},
	}

	for i := 0; i < len(table); i++ {
		err := ValidateCreateCategory(table[i])
		require.NoError(t, err)
	}
}

func TestValidateCreateCategoryError(t *testing.T) {

	table := []model.CreateCategoryDTO{
		{Name: "1"},
	}

	expErrors := []error{
		ErrTooSmallCategoryName,
	}

	for i := 0; i < len(table); i++ {
		err := ValidateCreateCategory(table[i])
		require.EqualError(t, err, expErrors[i].Error())
	}
}
