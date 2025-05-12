package service

import (
	"context"
	"github.com/google/uuid"
	"orders-center/internal/domain/cart/entity"
)

// Repository interfaces
type CartRepository interface {
	FindItemsByOrderID(ctx context.Context, orderID uuid.UUID) ([]*entity.OrderItem, error)
	AddItemsToOrder(ctx context.Context, items []*entity.OrderItem) error
}
