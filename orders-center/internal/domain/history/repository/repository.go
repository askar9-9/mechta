package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"orders-center/internal/domain/history/entity"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetHistory(ctx context.Context, id string) (*entity.History, error) {
	return nil, nil
}

func (r *Repo) CreateHistory(ctx context.Context, history *entity.History) error {
	return nil
}
