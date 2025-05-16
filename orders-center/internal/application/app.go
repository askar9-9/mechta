package application

import (
	"context"
	"log"
	"orders-center/internal/application/config"
	"orders-center/internal/application/connections"
	"orders-center/internal/application/grace"
	"orders-center/internal/application/start"
	"orders-center/internal/delivery/client"
	cartrepo "orders-center/internal/domain/cart/repository"
	cartsvc "orders-center/internal/domain/cart/service"
	historyrepo "orders-center/internal/domain/history/repository"
	historysvc "orders-center/internal/domain/history/service"
	orderrepo "orders-center/internal/domain/order/repository"
	ordersvc "orders-center/internal/domain/order/service"
	outboxrepo "orders-center/internal/domain/outbox/repository"
	outboxsvc "orders-center/internal/domain/outbox/service"
	paymentrepo "orders-center/internal/domain/payment/repository"
	paymentsvc "orders-center/internal/domain/payment/service"
	"orders-center/internal/infrastructure/db/pgxtx"
	"orders-center/internal/service/cron/service"
	orderfullsvc "orders-center/internal/service/orderfull/service"
	"time"
)

func Run() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config init failed: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Postgres.ConnTimeout)
	defer cancel()

	conns, err := connections.New(ctx, cfg)
	if err != nil {
		log.Fatalf("connections init failed: %v", err)
	}

	connectionsService := grace.NewService("connections", func(ctx context.Context) error {
		conns.Close()
		return nil
	})

	txManager := pgxtx.New(conns.DB)

	// init repositories
	cartRepo := cartrepo.NewRepo(conns.DB)
	historyRepo := historyrepo.NewRepo(conns.DB)
	orderRepo := orderrepo.NewRepo(conns.DB)
	paymentRepo := paymentrepo.NewRepo(conns.DB)
	outboxRepo := outboxrepo.NewRepo(conns.DB)

	// init services
	cartService := cartsvc.NewService(cartRepo, txManager)
	historyService := historysvc.NewService(historyRepo, txManager)
	orderService := ordersvc.NewService(orderRepo, txManager)
	paymentService := paymentsvc.NewService(paymentRepo, txManager)
	outboxService := outboxsvc.NewService(outboxRepo, txManager)

	// combined business service
	orderFullService := orderfullsvc.NewService(
		cartService,
		historyService,
		orderService,
		paymentService,
		outboxService,
		txManager,
	)

	// HTTP server
	httpService := start.RunHTTP(cfg, orderFullService)

	cronService := service.NewWorkerPool(cfg)
	cronService.Start()

	cronGracefulService := grace.NewService("cron", func(ctx context.Context) error {
		cronService.Stop()
		return nil
	})

	oneCClient := client.NewOneCClient(cfg, conns.HTTPClient)

	orderEno1cService := start.RunOrderEnoService(
		cfg,
		oneCClient,
		outboxService,
		cronService,
		txManager,
	)

	grace.New(15*time.Second,
		httpService,
		orderEno1cService,
		cronGracefulService,
		connectionsService,
	).Wait()
}
