package service

import (
	"context"
	"orders-center/internal/service/order_eno_1c/entity"
)

// Repository intefaces
type OrderENO1CRepository interface {
	CreateTask(ctx context.Context, item *entity.Outbox) error
	GetListTask(ctx context.Context, id string) ([]*entity.Outbox, error)
}
