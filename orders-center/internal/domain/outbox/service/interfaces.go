package service

import (
	"context"
	"orders-center/internal/domain/outbox/entity"
)

type OutboxRepository interface {
	CreateTask(ctx context.Context, item *entity.Outbox) error
	GetLimitedMessagesList(ctx context.Context, limited int) ([]*entity.Outbox, error)
	UpdateOutboxBatch(ctx context.Context, items []*entity.Outbox) error
	UpdateOutboxSingle(ctx context.Context, item *entity.Outbox) error
}
