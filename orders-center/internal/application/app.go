package application

import (
	"context"
	"orders-center/internal/application/config"
	"orders-center/internal/application/connections"
	"orders-center/internal/application/grace"
	"orders-center/internal/application/start"
)

func Run() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Postgres.ConnTimeout)
	defer cancel()

	conns, err := connections.New(ctx, cfg)
	if err != nil {
		panic(err)
	}

	errs := make(chan error, 1)

	httpService := start.RunHTTP(cfg, errs)

	grace.KillThemSoftly(httpService).Shutdown(errs, conns)
}
