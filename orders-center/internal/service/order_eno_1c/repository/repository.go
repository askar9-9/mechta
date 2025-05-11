package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"orders-center/internal/service/order_eno_1c/entity"
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

func (r *Repo) GetListTask(ctx context.Context, id string) ([]*entity.Outbox, error) {
	return nil, nil
}
