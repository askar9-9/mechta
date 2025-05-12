package repository

import (
	"context"
	"orders-center/internal/domain/outbox/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) CreateTask(ctx context.Context, item *entity.Outbox) error {
	return nil
}
func (r *Repo) GetLimitedTaskList(ctx context.Context, limit int) ([]*entity.Outbox, error) {
	return nil, nil
}
