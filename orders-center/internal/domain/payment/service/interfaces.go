package service

import (
	"context"
	"orders-center/internal/domain/payment/entity"
)

// Repository interfaces
type PaymentRepository interface {
	GetOrderPaymentByOrderID(ctx context.Context, orderID string) (*entity.OrderPayment, error)
	CreateOrderPayment(ctx context.Context, payment *entity.OrderPayment) error
}
