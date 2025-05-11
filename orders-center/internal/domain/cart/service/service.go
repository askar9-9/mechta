package service

import (
	"context"
	"orders-center/internal/domain/cart/entity"
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

func (s *Service) GetItemsForOrder(ctx context.Context, orderID string) ([]*entity.OrderItem, error) {
	return s.repo.FindItemsByOrderID(ctx, orderID)
}

func (s *Service) AttachItemsToOrder(ctx context.Context, items []*entity.OrderItem) error {
	return s.repo.AddItemsToOrder(ctx, items)
}
