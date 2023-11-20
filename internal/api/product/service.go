package product

import (
	"github.com/milkfrogge/coffee-store/internal/service"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
	"log/slog"
)

type Implementation struct {
	log *slog.Logger
	desc.UnimplementedProductV1Server
	service.ProductService
}

func NewImplementation(s service.ProductService, l *slog.Logger) *Implementation {

	return &Implementation{
		ProductService: s,
		log:            l,
	}
}
