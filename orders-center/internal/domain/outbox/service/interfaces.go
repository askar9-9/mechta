package service

import (
	"context"
	"orders-center/internal/service/order_eno_1c/entity"
)

type OutboxRepository interface {
	CreateTask(ctx context.Context, item *entity.Outbox) error
	GetLimitedTaskList(ctx context.Context, limited int) ([]*entity.Outbox, error)
}
