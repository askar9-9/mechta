package v1

import (
	"context"
	full "orders-center/internal/service/orderfull/entity"
)

type OrderService interface {
	CreateOrderFull(ctx context.Context, item *full.OrderFull) error
}
