package product

import (
	"context"
	"github.com/milkfrogge/coffee-store/internal/model"
	"github.com/milkfrogge/coffee-store/internal/repository"
	"log/slog"
	"time"
)

type Service struct {
	repo repository.ProductRepository
	log  *slog.Logger
}

func NewService(repo repository.ProductRepository, log *slog.Logger) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]model.Product, error) {
	const op = "Product.Service.GetAll"
	s.log.Info(op)
	out := make([]model.Product, 0)
	out = append(out, model.Product{
		Id:          "zxc",
		Name:        "zxc",
		Description: "zxc",
		Price:       2,
		Count:       23,
		Category: model.Category{
			Id:   "asd",
			Name: "cat",
		},
		Pics:      nil,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	})
	return out, nil
}
