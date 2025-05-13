package service

import (
	"context"
	"github.com/google/uuid"
	"orders-center/internal/domain/payment/entity"
	"orders-center/internal/pkg/tx"
)

type Service struct {
	repo      PaymentRepository
	txManager tx.TransactionManager
}

func NewService(repo PaymentRepository, txManager tx.TransactionManager) *Service {
	return &Service{
		repo:      repo,
		txManager: txManager,
	}
}

func (s *Service) InitializePayment(ctx context.Context, items []*entity.OrderPayment) error {
	return s.repo.CreateOrderPayments(ctx, items)
}

func (s *Service) GetPaymentInfo(ctx context.Context, orderID uuid.UUID) ([]*entity.OrderPayment, error) {
	return s.repo.GetOrderPaymentsByOrderID(ctx, orderID)
}
