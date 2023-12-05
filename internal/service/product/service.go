package product

import (
	"github.com/milkfrogge/coffee-store/internal/repository"
	"log/slog"
)

var sortingType = map[int32]string{
	0: "name",
	1: "price",
	2: "created_at",
}

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
