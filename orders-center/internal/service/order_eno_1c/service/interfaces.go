package service

import (
	"context"
	"orders-center/internal/domain/outbox/entity"
	"orders-center/internal/service/cron/service"
)

type OutboxService interface {
	FetchOutboxMessagesLimit(ctx context.Context, limited int) ([]*entity.Outbox, error)
	UpdateBatch(ctx context.Context, list []*entity.Outbox) error
	UpdateSingle(ctx context.Context, msg *entity.Outbox) error
}

type CronService interface {
	Submit(ctx context.Context, job service.Job) error
}

type OneCClient interface {
	SendMessage(ctx context.Context, msg *entity.Outbox) error
}
