package service

import "orders-center/internal/pkg/tx"

type Service struct {
	repo     HistoryRepository
	txManger tx.TransactionManager
}

func NewService(repo HistoryRepository, txManger tx.TransactionManager) *Service {
	return &Service{
		repo:     repo,
		txManger: txManger,
	}
}
