package grace

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Service interface {
	Name() string
	Stop(ctx context.Context) error
}

type ShutdownManager struct {
	services []Service
	timeout  time.Duration
}

func New(timeout time.Duration, services ...Service) *ShutdownManager {
	return &ShutdownManager{
		services: services,
		timeout:  timeout,
	}
}

func (g *ShutdownManager) Wait() {
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	// Wait for a signal
	<-ctx.Done()
	log.Println("Инициирована процедура graceful shutdown...")

	// Context for shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), g.timeout)
	defer cancel()

	// Create a WaitGroup to wait for all services to stop
	var wg sync.WaitGroup
	errCh := make(chan error, len(g.services))

	for _, svc := range g.services {
		if svc == nil {
			continue
		}

		wg.Add(1)
		go func(s Service) {
			defer wg.Done()
			log.Printf("Остановка сервиса: %s", s.Name())

			// Create a new context with timeout for each service
			serviceCtx, serviceCancel := context.WithTimeout(shutdownCtx, g.timeout/2)
			defer serviceCancel()

			if err := s.Stop(serviceCtx); err != nil {
				if !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
					log.Printf("Ошибка при остановке сервиса %s: %v", s.Name(), err)
					errCh <- err
				}
			} else {
				log.Printf("Сервис %s успешно остановлен", s.Name())
			}
		}(svc)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(errCh)

	// Check for errors
	exitCode := 0
	for err := range errCh {
		if err != nil {
			exitCode = 1
			log.Printf("Ошибка при graceful shutdown: %v", err)
		}
	}

	log.Println("Все сервисы остановлены. До свидания!")
	os.Exit(exitCode)
}

func NewService(name string, stopFn func(ctx context.Context) error) Service {
	return &serviceImpl{
		name: name,
		stop: stopFn,
	}
}

type serviceImpl struct {
	name string
	stop func(ctx context.Context) error
}

func (s *serviceImpl) Stop(ctx context.Context) error {
	return s.stop(ctx)
}

func (s *serviceImpl) Name() string {
	return s.name
}
