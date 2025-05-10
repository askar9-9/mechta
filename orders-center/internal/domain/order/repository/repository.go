package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"orders-center/internal/domain/order/entity"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetOrderByID(ctx context.Context, id string) (*entity.Order, error) {
	return nil, nil
}

func (r *Repo) CreateOrder(ctx context.Context, order *entity.Order) error {
	return nil
}
