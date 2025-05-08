package connections

import (
	"context"
	"fmt"
	"orders-center/internal/application/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Connections struct {
	DB *pgxpool.Pool
}

func (c *Connections) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}

func New(ctx context.Context, cfg *config.Config) (*Connections, error) {

	db, err := pgxpool.New(ctx, cfg.Postgres.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	if err := db.Ping(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	return &Connections{
		DB: db,
	}, nil
}
