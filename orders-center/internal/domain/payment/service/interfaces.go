package service

import (
	"context"
	"github.com/google/uuid"
	"orders-center/internal/domain/payment/entity"
)

// Repository interfaces
type PaymentRepository interface {
	GetOrderPaymentsByOrderID(ctx context.Context, orderID uuid.UUID) ([]*entity.OrderPayment, error)
	CreateOrderPayments(ctx context.Context, payments []*entity.OrderPayment) error
}
