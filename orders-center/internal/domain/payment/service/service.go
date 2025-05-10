package service

import "orders-center/internal/pkg/tx"

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
