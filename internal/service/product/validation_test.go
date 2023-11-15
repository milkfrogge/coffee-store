package product

import (
	"github.com/milkfrogge/coffee-store/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateCreateProduct(t *testing.T) {
	table := []model.CreateProductDTO{
		model.CreateProductDTO{
			Name:        "first",
			Description: "lorem",
			Price:       1,
			Count:       1,
			CategoryId:  "0xdeadbeef",
			Pics:        nil,
		},
		model.CreateProductDTO{
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
		model.CreateProductDTO{
			Name:        "first",
			Description: "lorem",
			Price:       1,
			Count:       0,
			CategoryId:  "0xdeadbeef",
			Pics:        nil,
		},
		model.CreateProductDTO{
			Name:        "second",
			Description: "lorem",
			Price:       0,
			Count:       1,
			CategoryId:  "0xdeadbeef",
			Pics:        []string{"image1"},
		},
	}

	expErrors := []error{ErrWrongCount, ErrWrongPrice}

	for i := 0; i < len(table); i++ {
		err := ValidateCreateProduct(table[i])
		require.EqualError(t, err, expErrors[i].Error())
	}
}
