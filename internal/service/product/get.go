package product

import (
	"context"
	"github.com/milkfrogge/coffee-store/internal/model"
)

func (s *Service) GetAllCategories(ctx context.Context) ([]model.Category, error) {
	const op = "Product.Service.GetAllCategories"
	s.log.Debug(op)

	ctx, span := s.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	categories, err := s.repo.FindAllCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *Service) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	const op = "Product.Service.GetAllProducts"
	s.log.Debug(op)

	ctx, span := s.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	products, err := s.repo.FindAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *Service) GetAllProductsByCategory(ctx context.Context, categoryId string) ([]model.Product, error) {
	const op = "Product.Service.GetAllProductsByCategory"
	s.log.Debug(op)

	ctx, span := s.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	products, err := s.repo.FindProductsByCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *Service) GetSingleProduct(ctx context.Context, id string) (model.Product, error) {
	const op = "Product.Service.GetSingleProduct"
	s.log.Debug(op)

	ctx, span := s.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	product, err := s.repo.FindOneProduct(ctx, id)
	if err != nil {
		return model.Product{}, err
	}
	span.AddEvent("returning")
	return product, nil
}
