package service

import (
	"context"
	"orders-center/internal/domain/outbox/entity"
	"orders-center/internal/pkg/tx"
	full "orders-center/internal/service/orderfull/entity"
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

func (s *Service) CreateTask(ctx context.Context, item *full.OrderFull) error {
	return nil
}

func (s *Service) GetListTask(ctx context.Context, id string) ([]*entity.Outbox, error) {
	return nil, nil
}
