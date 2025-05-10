package service

import (
	"context"
	"orders-center/internal/domain/history/entity"
)

type HistoryRepository interface {
	GetHistory(ctx context.Context, id string) (*entity.History, error)
	CreateHistory(ctx context.Context, history *entity.History) error
}
