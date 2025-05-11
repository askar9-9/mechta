package service

import (
	"context"
	"orders-center/internal/pkg/tx"
	"orders-center/internal/service/order_eno_1c/entity"
	full "orders-center/internal/service/orderfull/entity"
)

type Service struct {
	repo      OrderENO1CRepository
	txManager tx.TransactionManager
}

func NewService(repo OrderENO1CRepository, txManager tx.TransactionManager) *Service {
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
