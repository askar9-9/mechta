package service

import (
	"context"
	"orders-center/internal/domain/order/entity"
)

// Repository interfaces
type OrderRepository interface {
	GetOrderByID(ctx context.Context, id string) (*entity.Order, error)
	CreateOrder(ctx context.Context, order *entity.Order) error
}
