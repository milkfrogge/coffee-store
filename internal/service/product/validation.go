package product

import "github.com/milkfrogge/coffee-store/internal/model"

func ValidateCreateProduct(product model.CreateProductDTO) error {
	if product.Price <= 0 {
		return ErrWrongPrice
	}
	if product.Count <= 0 {
		return ErrWrongCount
	}

	return nil
}
