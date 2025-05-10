package service

import (
	"context"
	"orders-center/internal/pkg/tx"
	"orders-center/internal/service/orderfull/entity"
)

type Service struct {
	repo      OrderFullRepository
	cartSvc   CartService
	histSvc   HistoryService
	orderSvc  OrderService
	paySvc    PaymentService
	txManager tx.TransactionManager
}

func NewService(
	repo OrderFullRepository,
	cart CartService,
	history HistoryService,
	order OrderService,
	payment PaymentService,
	txManager tx.TransactionManager,
) *Service {
	return &Service{
		repo:      repo,
		cartSvc:   cart,
		histSvc:   history,
		orderSvc:  order,
		paySvc:    payment,
		txManager: txManager,
	}
}

func (s *Service) CreateOrderFull(ctx context.Context, orderFull *entity.OrderFull) error {
	return s.txManager.Do(ctx, func(ctx context.Context) error {
		if err := s.orderSvc.CreateOrder(ctx, orderFull.Order); err != nil {
			return err
		}

		if err := s.cartSvc.AddItemsToOrder(ctx, orderFull.Items); err != nil {
			return err
		}

		if err := s.paySvc.CreateOrderPayment(ctx, orderFull.Payment); err != nil {
			return err
		}

		if err := s.histSvc.CreateOrderHistory(ctx, orderFull.History); err != nil {
			return err
		}

		return nil
	})
}
