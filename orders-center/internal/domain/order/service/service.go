package service

import (
	"context"
	"orders-center/internal/domain/order/entity"
	"orders-center/internal/pkg/tx"
)

type Service struct {
	repo      OrderRepository
	txManager tx.TransactionManager
}

func NewService(repo OrderRepository, txManager tx.TransactionManager) *Service {
	return &Service{
		repo:      repo,
		txManager: txManager,
	}
}

func (s *Service) RegisterOrder(ctx context.Context, item *entity.Order) error {
	if item == nil {
		return entity.ErrOrderRequired
	}

	if err := item.Validate(); err != nil {
		return err
	}

	return s.repo.CreateOrder(ctx, item)
}

func (s *Service) GetOrderDetails(ctx context.Context, orderID string) (*entity.Order, error) {
	if orderID == "" {
		return nil, entity.ErrOrderIDRequired
	}
	return s.repo.GetOrderByID(ctx, orderID)
}
