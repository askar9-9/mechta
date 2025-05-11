package service

import (
	"context"
	cart "orders-center/internal/domain/cart/entity"
	history "orders-center/internal/domain/history/entity"
	order "orders-center/internal/domain/order/entity"
	payment "orders-center/internal/domain/payment/entity"
	"orders-center/internal/service/orderfull/entity"
)

type ENOService interface {
	CreateTask(ctx context.Context, item *entity.OrderFull) error
}

type CartService interface {
	GetItemsForOrder(ctx context.Context, id string) ([]*cart.OrderItem, error)
	AttachItemsToOrder(ctx context.Context, items []*cart.OrderItem) error
}

type HistoryService interface {
	LoadOrderHistory(ctx context.Context, id string) (*history.History, error)
	RecordOrderHistory(ctx context.Context, item *history.History) error
}

type OrderService interface {
	RegisterOrder(ctx context.Context, item *order.Order) error
	GetOrderDetails(ctx context.Context, id string) (*order.Order, error)
}

type PaymentService interface {
	InitializePayment(ctx context.Context, item *payment.OrderPayment) error
	GetPaymentInfo(ctx context.Context, id string) (*payment.OrderPayment, error)
}
