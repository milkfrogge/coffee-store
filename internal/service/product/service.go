package product

import (
	"github.com/milkfrogge/coffee-store/internal/repository"
	"log/slog"
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
