package product

import "context"

func (s *Service) DeleteProduct(ctx context.Context, id string) error {
	const op = "Product.Service.DeleteProduct"
	s.log.Debug(op)

	err := s.repo.DeleteProduct(ctx, id)

	return err
}

func (s *Service) DeleteCategory(ctx context.Context, id string) error {
	const op = "Product.Service.DeleteCategory"
	s.log.Debug(op)

	err := s.repo.DeleteCategory(ctx, id)

	return err
}
