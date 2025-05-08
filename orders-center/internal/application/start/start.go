package start

import (
	"context"
	"fmt"
	"net/http"
	"orders-center/internal/application/config"
	"orders-center/internal/application/grace"
	v1 "orders-center/internal/delivery/api/v1"
	"time"
)

func RunHTTP(cfg *config.Config, errs chan<- error) grace.Service {
	server := &http.Server{
		Addr:    cfg.HTTP.Addr,
		Handler: v1.Routes(),
	}

	// Start the server in a goroutine
	go func() {
		fmt.Printf("HTTP server starting on %s\n", cfg.HTTP.Addr)
		if err := server.ListenAndServe(); err != nil {
			// Only report errors that are not due to the server closing
			if err != http.ErrServerClosed {
				errs <- fmt.Errorf("failed to start HTTP server: %w", err)
			}
		}
	}()

	return grace.NewService("http", func() {
		fmt.Println("Initiating graceful shutdown of HTTP server...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			errs <- fmt.Errorf("HTTP server graceful shutdown failed: %w", err)
		} else {
			fmt.Println("HTTP server gracefully stopped")
		}
	})
}
