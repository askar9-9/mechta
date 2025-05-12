package service

import (
	"context"
	"github.com/google/uuid"
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

func (s *Service) LoadOrderHistory(ctx context.Context, orderID uuid.UUID) ([]*entity.History, error) {
	if orderID == uuid.Nil {
		return nil, entity.ErrOrderIDRequired
	}

	return s.repo.GetHistories(ctx, orderID)
}

func (s *Service) RecordOrderHistory(ctx context.Context, items []*entity.History) error {
	if len(items) == 0 {
		return entity.ErrHistoryItemsRequired
	}

	for _, item := range items {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return s.repo.CreateHistories(ctx, items)
}
