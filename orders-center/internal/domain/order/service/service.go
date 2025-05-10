package service

import "orders-center/internal/pkg/tx"

type Service struct {
	repo      OrderRepository
	txManager tx.TransactionManager
}

func NewService(repo OrderRepository, txManager tx.TransactionManager) *Service {
	return &Service{
		repo:      repo,
		txManager: txManager,
	}
}
