package application

import (
	"context"
	"orders-center/internal/application/config"
	"orders-center/internal/application/connections"
	"orders-center/internal/application/grace"
	"orders-center/internal/application/start"
	cartrepo "orders-center/internal/domain/cart/repository"
	historyrepo "orders-center/internal/domain/history/repository"
	orderrepo "orders-center/internal/domain/order/repository"
	paymentrepo "orders-center/internal/domain/payment/repository"
	"orders-center/internal/infrastructure/db/pgxtx"

	cartsvc "orders-center/internal/domain/cart/service"
	historysvc "orders-center/internal/domain/history/service"
	ordersvc "orders-center/internal/domain/order/service"
	paymentsvc "orders-center/internal/domain/payment/service"
	orderfullrepo "orders-center/internal/service/orderfull/repository"
	orderfullsvc "orders-center/internal/service/orderfull/service"
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

	txManager := pgxtx.New(conns.DB)

	// init repositories
	cartRepo := cartrepo.NewRepo(conns.DB)
	historyRepo := historyrepo.NewRepo(conns.DB)
	orderRepo := orderrepo.NewRepo(conns.DB)
	paymentRepo := paymentrepo.NewRepo(conns.DB)

	orderFullRepo := orderfullrepo.NewRepo(conns.DB)

	// init services
	cartService := cartsvc.NewService(cartRepo, txManager)
	historyService := historysvc.NewService(historyRepo, txManager)
	orderService := ordersvc.NewService(orderRepo, txManager)
	paymentService := paymentsvc.NewService(paymentRepo, txManager)

	orderFullService := orderfullsvc.NewService(
		orderFullRepo,
		cartService,
		historyService,
		orderService,
		paymentService,
		txManager,
	)

	httpService := start.RunHTTP(cfg, errs)

	grace.KillThemSoftly(httpService).Shutdown(errs, conns)
}
