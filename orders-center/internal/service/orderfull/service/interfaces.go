package service

import (
	"context"
	cart "orders-center/internal/domain/cart/entity"
	history "orders-center/internal/domain/history/entity"
	order "orders-center/internal/domain/order/entity"
	payment "orders-center/internal/domain/payment/entity"
)

type OrderFullRepository interface {
}

type CartService interface {
	AddItemsToOrder(ctx context.Context, items []*cart.OrderItem) error
}

type HistoryService interface {
	CreateOrderHistory(ctx context.Context, item *history.History) error
}

type OrderService interface {
	CreateOrder(ctx context.Context, item *order.Order) error
}

type PaymentService interface {
	CreateOrderPayment(ctx context.Context, item *payment.OrderPayment) error
}
