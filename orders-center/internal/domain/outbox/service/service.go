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

	now := time.Now()
	outbox := &entity.Outbox{
		ID:            uuid.New(),
		AggregateID:   item.Order.ID,
		AggregateType: entity.AggregateTypeOrder,
		EventType:     entity.EventTypeOrderCreated,
		Payload:       data,
		CreatedAt:     &now,
		RetryCount:    0,
		Error:         "",
	}

	return s.repo.CreateTask(ctx, outbox)
}

func (s *Service) FetchOutboxMessagesLimit(ctx context.Context, limit int) ([]*entity.Outbox, error) {
	return s.repo.GetLimitedMessagesList(ctx, limit)
}

func (s *Service) UpdateBatch(ctx context.Context, list []*entity.Outbox) error {
	return s.repo.UpdateOutboxBatch(ctx, list)
}

func (s *Service) UpdateSingle(ctx context.Context, msg *entity.Outbox) error {
	return s.repo.UpdateOutboxSingle(ctx, msg)
}
