package product

import (
	"github.com/milkfrogge/coffee-store/internal/model"
	"unicode/utf8"
)

const MinimumCategoryNameLength = 4

func ValidateCreateProduct(product model.CreateProductDTO) error {
	if product.Price <= 0 {
		return ErrWrongPrice
	}
	if product.Count <= 0 {
		return ErrWrongCount
	}
	if product.BaristaNeeded == product.KitchenNeeded && product.BaristaNeeded {
		return ErrBaristaAndKitchenNeeded
	}

	return nil
}

func ValidateCreateCategory(category model.CreateCategoryDTO) error {

	if utf8.RuneCountInString(category.Name) < MinimumCategoryNameLength {
		return ErrTooSmallCategoryName
	}

	return nil
}
