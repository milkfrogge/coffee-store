package product

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/milkfrogge/coffee-store/internal/utils"
	"log/slog"
	"time"
)

type PostgresRepository struct {
	conn *pgx.Conn
	log  *slog.Logger
}

func NewProductPostgresRepository(dsn string, log *slog.Logger) (*PostgresRepository, error) {

	ctx := context.Background()

	a, err := utils.WithAttempts(func() (any, error) {
		ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		conn, err := pgx.Connect(ctxTimeout, dsn)
		if err != nil {
			log.Error(fmt.Sprintf("Can`t establish connect to db: %s", err.Error()))
			return nil, err
		}

		return &PostgresRepository{
			conn: conn,
			log:  log,
		}, nil
	}, 5, log)

	if err != nil {
		return nil, err
	}

	return a.(*PostgresRepository), nil

}
