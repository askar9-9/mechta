package service

import (
	"context"
	"orders-center/internal/pkg/tx"
	"orders-center/internal/service/orderfull/entity"
)

type Service struct {
	cartSvc   CartService
	histSvc   HistoryService
	orderSvc  OrderService
	paySvc    PaymentService
	outboxSvc OutboxService
	txManager tx.TransactionManager
}

func NewService(
	cart CartService,
	history HistoryService,
	order OrderService,
	payment PaymentService,
	outboxSvc OutboxService,
	txManager tx.TransactionManager,
) *Service {
	return &Service{
		cartSvc:   cart,
		histSvc:   history,
		orderSvc:  order,
		paySvc:    payment,
		outboxSvc: outboxSvc,
		txManager: txManager,
	}
}

func (s *Service) CreateOrderFull(ctx context.Context, orderFull *entity.OrderFull) error {
	return s.txManager.Do(ctx, func(ctx context.Context) error {
		if err := s.orderSvc.RegisterOrder(ctx, orderFull.Order); err != nil {
			return err
		}

		if err := s.cartSvc.AttachItemsToOrder(ctx, orderFull.Items); err != nil {
			return err
		}

		if err := s.paySvc.InitializePayment(ctx, orderFull.Payment); err != nil {
			return err
		}

		if err := s.histSvc.RecordOrderHistory(ctx, orderFull.History); err != nil {
			return err
		}

		if err := s.outboxSvc.CreateTask(ctx, orderFull); err != nil {
			return err
		}
		return nil
	})
}
