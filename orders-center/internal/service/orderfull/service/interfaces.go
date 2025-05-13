package service

import (
	"context"
	"github.com/google/uuid"
	cart "orders-center/internal/domain/cart/entity"
	history "orders-center/internal/domain/history/entity"
	order "orders-center/internal/domain/order/entity"
	payment "orders-center/internal/domain/payment/entity"
	full "orders-center/internal/service/orderfull/entity"
)

type OutboxService interface {
	CreateOrderFullTask(ctx context.Context, item *full.OrderFull) error
}

type CartService interface {
	GetItemsForOrder(ctx context.Context, orderID uuid.UUID) ([]*cart.OrderItem, error)
	AttachItemsToOrder(ctx context.Context, items []*cart.OrderItem) error
}

type HistoryService interface {
	LoadOrderHistory(ctx context.Context, orderID uuid.UUID) ([]*history.History, error)
	RecordOrderHistory(ctx context.Context, item []*history.History) error
}

type OrderService interface {
	RegisterOrder(ctx context.Context, item *order.Order) error
	GetOrderDetails(ctx context.Context, orderID uuid.UUID) (*order.Order, error)
}

type PaymentService interface {
	InitializePayment(ctx context.Context, items []*payment.OrderPayment) error
	GetPaymentInfo(ctx context.Context, orderID uuid.UUID) ([]*payment.OrderPayment, error)
}
