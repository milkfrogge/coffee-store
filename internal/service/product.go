package service

import (
	"context"
	"github.com/milkfrogge/coffee-store/internal/model"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]model.Product, error)
}
