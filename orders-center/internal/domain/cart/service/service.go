package service

import (
	"orders-center/internal/pkg/tx"
)

type Service struct {
	repo      CartRepository
	txManager tx.TransactionManager
}

func NewService(repo CartRepository, txManager tx.TransactionManager) *Service {
	return &Service{
		repo:      repo,
		txManager: txManager,
	}
}

func (s *Service) FindItemsByOrderID() {
}
