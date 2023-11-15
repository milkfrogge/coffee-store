package product

import (
	"github.com/milkfrogge/coffee-store/internal/repository"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
)

type Service struct {
	repo   repository.ProductRepository
	log    *slog.Logger
	tracer trace.TracerProvider
}

func NewService(repo repository.ProductRepository, log *slog.Logger) *Service {

	tp := otel.GetTracerProvider()

	return &Service{
		repo:   repo,
		log:    log,
		tracer: tp,
	}
}
