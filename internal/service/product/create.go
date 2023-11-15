package product

import (
	"context"
	"github.com/milkfrogge/coffee-store/internal/model"
)

func (s *Service) CreateCategory(ctx context.Context, category model.CreateCategoryDTO) (string, error) {
	const op = "Product.Service.CreateCategory"
	s.log.Debug(op)

	ctx, span := s.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	id, err := s.repo.CreateCategory(ctx, category)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *Service) CreateProduct(ctx context.Context, product model.CreateProductDTO) (string, error) {
	const op = "Product.Service.CreateProduct"
	s.log.Debug(op)

	ctx, span := s.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	err := ValidateCreateProduct(product)
	if err != nil {
		return "", err
	}

	id, err := s.repo.CreateProduct(ctx, product)
	if err != nil {
		return "", err
	}

	return id, nil
}
