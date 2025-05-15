package connections

import (
	"context"
	"fmt"
	"net/http"
	"orders-center/internal/application/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Connections struct {
	DB         *pgxpool.Pool
	HTTPClient *http.Client
}

func (c *Connections) Close() {
	if c.DB != nil {
		c.DB.Close()
	}

	if c.HTTPClient != nil {
		if transport, ok := c.HTTPClient.Transport.(*http.Transport); ok {
			transport.CloseIdleConnections()
		}
	}
}

func New(ctx context.Context, cfg *config.Config) (*Connections, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    cfg.OneC.MaxIdleConnections,
			MaxConnsPerHost: cfg.OneC.MaxConnsPerHost,
			IdleConnTimeout: time.Duration(cfg.OneC.IdleConnTimeout) * time.Millisecond,
		},
		Timeout: cfg.OneC.Timeout * time.Millisecond,
	}

	db, err := pgxpool.New(ctx, cfg.Postgres.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	if err := db.Ping(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	return &Connections{
		DB:         db,
		HTTPClient: httpClient,
	}, nil
}
