package product

import "errors"

var (
	ErrWrongPrice = errors.New("price cannot be <=0")
	ErrWrongCount = errors.New("count of product cannot be <=0")
)
