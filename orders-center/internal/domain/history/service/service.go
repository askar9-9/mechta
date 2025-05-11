package service

import (
	"context"
	"orders-center/internal/domain/history/entity"
	"orders-center/internal/pkg/tx"
)

type Service struct {
	repo     HistoryRepository
	txManger tx.TransactionManager
}

func NewService(repo HistoryRepository, txManger tx.TransactionManager) *Service {
	return &Service{
		repo:     repo,
		txManger: txManger,
	}
}

func (s *Service) LoadOrderHistory(ctx context.Context, id string) (*entity.History, error) {
	return nil, nil
}

func (s *Service) RecordOrderHistory(ctx context.Context, item *entity.History) error {
	return nil
}
