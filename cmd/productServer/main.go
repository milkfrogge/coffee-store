package main

import (
	"fmt"
	"github.com/milkfrogge/coffee-store/internal/config"
	"github.com/milkfrogge/coffee-store/internal/repository/product"
	"log/slog"
	"os"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

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

	repository, err := product.NewProductPostgresRepository(cfg.GetPostgresDsn(), logger)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	fmt.Println(repository)

}
