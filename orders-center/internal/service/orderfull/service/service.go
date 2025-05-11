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
	enoSvc    ENOService
	txManager tx.TransactionManager
}

func NewService(
	cart CartService,
	history HistoryService,
	order OrderService,
	payment PaymentService,
	eno ENOService,
	txManager tx.TransactionManager,
) *Service {
	return &Service{
		cartSvc:   cart,
		histSvc:   history,
		orderSvc:  order,
		paySvc:    payment,
		enoSvc:    eno,
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

		if err := s.enoSvc.CreateTask(ctx, orderFull); err != nil {
			return err
		}
		return nil
	})
}
