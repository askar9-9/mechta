package service

import (
	"context"
	"github.com/google/uuid"
	"orders-center/internal/domain/order/entity"
)

// Repository interfaces
type OrderRepository interface {
	GetOrderByID(ctx context.Context, id uuid.UUID) (*entity.Order, error)
	CreateOrder(ctx context.Context, order *entity.Order) error
}
