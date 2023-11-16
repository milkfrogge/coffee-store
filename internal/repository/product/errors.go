package product

import "errors"

var (
	ErrNoSuchCategory = errors.New("no category with requested id")
	ErrNoSuchProduct  = errors.New("no product with requested id")
	ErrDbIsDown       = errors.New("db is down")
)
