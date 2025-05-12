package service

import (
	"context"
	"orders-center/internal/domain/history/entity"
)

type HistoryRepository interface {
	GetHistories(ctx context.Context, orderID string) ([]*entity.History, error)
	CreateHistories(ctx context.Context, history []*entity.History) error
}
