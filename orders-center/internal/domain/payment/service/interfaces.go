package service

import (
	"context"
	"orders-center/internal/domain/payment/entity"
)

// Repository interfaces
type PaymentRepository interface {
	GetOrderPaymentsByOrderID(ctx context.Context, orderID string) ([]*entity.OrderPayment, error)
	CreateOrderPayments(ctx context.Context, payments []*entity.OrderPayment) error
}
