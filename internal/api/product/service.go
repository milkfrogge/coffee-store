package product

import (
	"github.com/milkfrogge/coffee-store/internal/service"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
)

type Implementation struct {
	log *slog.Logger
	desc.UnimplementedProductV1Server
	service.ProductService
	tracer trace.TracerProvider
}

func NewImplementation(s service.ProductService, l *slog.Logger) *Implementation {

	tracerProvider := otel.GetTracerProvider()

	return &Implementation{
		ProductService: s,
		log:            l,
		tracer:         tracerProvider,
	}
}
