package service

import (
	"context"
	"orders-center/internal/pkg/tx"
	"orders-center/internal/service/order_eno_1c/entity"
)

type Service struct {
	repo      OutboxRepository
	txManager tx.TransactionManager
}

func NewService(repo OutboxRepository, txManager tx.TransactionManager) *Service {
	return &Service{
		repo:      repo,
		txManager: txManager,
	}
}

func (s *Service) CreateTask(ctx context.Context, orderID string) error {
	return nil
}

func (s *Service) GetListTask(ctx context.Context, id string) ([]*entity.Outbox, error) {
	return nil, nil
}
