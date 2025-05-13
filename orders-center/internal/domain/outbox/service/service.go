package service

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"orders-center/internal/domain/outbox/entity"
	"orders-center/internal/pkg/tx"
	full "orders-center/internal/service/orderfull/entity"
	"time"
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

func (s *Service) CreateOrderFullTask(ctx context.Context, item *full.OrderFull) error {
	data, err := json.Marshal(&item)
	if err != nil {
		return err
	}

	outbox := &entity.Outbox{
		ID:            uuid.New(),
		AggregateID:   item.Order.ID,
		AggregateType: entity.AggregateTypeOrder,
		EventType:     entity.EventTypeOrderCreated,
		Payload:       data,
		CreatedAt:     time.Now(),
		SyncAt:        time.Time{},
		ProcessedAt:   time.Time{},
		RetryCount:    0,
		Error:         "",
	}

	return s.repo.CreateTask(ctx, outbox)
}

func (s *Service) GetListTask(ctx context.Context, id string) ([]*entity.Outbox, error) {
	return nil, nil
}
