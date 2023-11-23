package interceptors

import (
	"context"
	"github.com/milkfrogge/coffee-store/pkg/jaeger"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"log/slog"
)

func TracingUnaryInterceptor(log *slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx, err = jaeger.ExtractMetaFromGRPC(ctx)
		if err != nil {
			log.Warn(err.Error())

		}

		tracer := otel.GetTracerProvider()

		ctx, span := tracer.Tracer(info.FullMethod).Start(ctx, info.FullMethod)
		defer span.End()

		res, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}
