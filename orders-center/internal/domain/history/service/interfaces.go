package service

import (
	"context"
	"github.com/google/uuid"
	"orders-center/internal/domain/history/entity"
)

type HistoryRepository interface {
	GetHistories(ctx context.Context, orderID uuid.UUID) ([]*entity.History, error)
	CreateHistories(ctx context.Context, history []*entity.History) error
}
