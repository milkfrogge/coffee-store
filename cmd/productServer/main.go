package main

import (
	"fmt"
	productApi "github.com/milkfrogge/coffee-store/internal/api/product"
	"github.com/milkfrogge/coffee-store/internal/config"
	productRepo "github.com/milkfrogge/coffee-store/internal/repository/product/postgres"
	productService "github.com/milkfrogge/coffee-store/internal/service/product"
	"github.com/milkfrogge/coffee-store/pkg/interceptors"
	"github.com/milkfrogge/coffee-store/pkg/jaeger"
	loggerClient "github.com/milkfrogge/coffee-store/pkg/logger"
	desc "github.com/milkfrogge/coffee-store/pkg/product_v1"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"io"
	"log/slog"
	"net"
	"os"
	"time"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout), &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Info("Init config")
	err := config.Load()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	cfg, err := config.NewProductConfig()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info("Init grayLogClient")
	grayLogClient, err := loggerClient.NewGrayLogLogger(fmt.Sprintf("%s:%s", cfg.GraylogHost, cfg.GraylogPort))
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info("Combine logger with graylog")
	logger = slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, grayLogClient), &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Info("Init tracer")
	tracer := jaeger.NewTracer(fmt.Sprintf("http://%s:%s/api/traces", cfg.JaegerHost, cfg.JaegerPort))
	provider, err := tracer.NewTracerProvider("ProductService")
	if err != nil {
		logger.Error(err.Error())
		return
	}
	otel.SetTracerProvider(provider)

	logger.Info("Init repo")
	repository, err := productRepo.NewProductPostgresRepository(cfg.GetPostgresDsn(), logger)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	logger.Info("Init service")
	s := productService.NewService(repository, logger)

	logger.Info("Init implementation of server")
	sImpl := productApi.NewImplementation(s, logger)

	logger.Info("Register implementation of server")
	srv := grpc.NewServer(grpc.Creds(insecure.NewCredentials()), grpc.ChainUnaryInterceptor(interceptors.TracingUnaryInterceptor(logger), interceptors.TimeoutUnaryInterceptor(time.Second*100)))

	reflection.Register(srv)

	desc.RegisterProductV1Server(srv, sImpl)

	logger.Info("Starting server")

	list, err := net.Listen("tcp", cfg.Address())
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = srv.Serve(list)
	if err != nil {
		logger.Error(err.Error())
		return
	}

}
