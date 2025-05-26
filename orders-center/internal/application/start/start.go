package start

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"orders-center/internal/application/config"
	"orders-center/internal/application/grace"
	v1 "orders-center/internal/delivery/api/v1"
	"orders-center/internal/delivery/client"
	"orders-center/internal/pkg/tx"
	ordereno1c "orders-center/internal/service/order_eno_1c/service"
)

func RunHTTP(cfg *config.Config, svc v1.OrderService) grace.Service {
	server := &http.Server{
		Addr:    cfg.HTTP.Addr,
		Handler: v1.Routes(svc),
	}

	go func() {
		fmt.Printf("HTTP server starting on %s\n", cfg.HTTP.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("HTTP server error: %v\n", err)
		}
	}()

	return grace.NewService("http", func(ctx context.Context) error {
		fmt.Println("Shutting down HTTP server...")
		return server.Shutdown(ctx)
	})
}

func RunOrderEnoService(
	cfg *config.Config,
	oneCClient *client.OneCClient,
	outboxSvc ordereno1c.OutboxService,
	cronSvc ordereno1c.CronService,
	txManager tx.TransactionManager,
) grace.Service {
	svc := ordereno1c.NewOrderEno1cService(
		cfg,
		oneCClient,
		outboxSvc,
		cronSvc,
		txManager,
	)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic in OrderEno1cService: %v\n", r)
			}
		}()

		fmt.Println("OrderEno1cService started")
		if err := svc.Start(ctx); err != nil && !errors.Is(err, context.Canceled) {
			fmt.Printf("OrderEno1cService error: %v\n", err)
		}
		fmt.Println("OrderEno1cService exited")
	}()

	return grace.NewService(svc.Name(), func(shutdownCtx context.Context) error {
		fmt.Println("Shutting down OrderEno1cService...")
		cancel()

		select {
		case <-shutdownCtx.Done():
			return shutdownCtx.Err()
		case <-ctx.Done():
			return nil
		}
	})
}
