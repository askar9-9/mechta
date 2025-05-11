package service

import (
	"context"
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

func (s *Service) InitializePayment(ctx context.Context, item *entity.OrderPayment) error {
	return nil
}

func (s *Service) GetPaymentInfo(ctx context.Context, id string) (*entity.OrderPayment, error) {
	return nil, nil
}
