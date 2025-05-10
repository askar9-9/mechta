package service

import (
	"context"
	"orders-center/internal/domain/cart/entity"
)

// Repository interfaces
type CartRepository interface {
	FindItemsByOrderID(ctx context.Context, orderID string) ([]*entity.OrderItem, error)
	AddItemsToOrder(ctx context.Context, items []*entity.OrderItem) error
}
