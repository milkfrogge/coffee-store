package product

import (
	"context"
	"github.com/milkfrogge/coffee-store/internal/model"
	"log/slog"
)

func (s *Service) UpdateProduct(ctx context.Context, product model.Product) error {
	err := s.repo.UpdateProduct(ctx, product)

	if err != nil {
		return err
	}

	return err
}

func (s *Service) AddCountToProduct(ctx context.Context, dto model.CountToProductDTO) error {

	product, err := s.repo.FindOneProduct(ctx, dto.Id)
	if err != nil {
		return err
	}

	err = s.repo.UpdateCountOfProduct(ctx, product.Id, product.Count+dto.Count)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SubtractCountToProduct(ctx context.Context, dto model.CountToProductDTO) error {
	product, err := s.repo.FindOneProduct(ctx, dto.Id)
	if err != nil {
		return err
	}

	err = s.repo.UpdateCountOfProduct(ctx, product.Id, product.Count-dto.Count)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SubtractCountToProducts(ctx context.Context, dto []model.CountToProductDTO) error {

	in := make(map[string]uint64)

	for i := 0; i < len(dto); i++ {
		product, err := s.repo.FindOneProduct(ctx, dto[i].Id)
		if err != nil {
			slog.Error(err.Error())
			return err
		}

		in[dto[i].Id] = product.Count - dto[i].Count
	}

	err := s.repo.UpdateManyCountsOfProduct(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateCategory(ctx context.Context, category model.Category) error {
	err := s.repo.UpdateCategory(ctx, category)
	if err != nil {
		return err
	}
	return err
}
