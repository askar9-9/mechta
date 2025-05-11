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
	return nil
}

func (s *Service) GetOrderDetails(ctx context.Context, id string) (*entity.Order, error) {
	return nil, nil
}
