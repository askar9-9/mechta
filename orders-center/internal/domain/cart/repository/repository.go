package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"orders-center/internal/domain/cart/entity"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) FindItemsByOrderID(ctx context.Context, orderID string) ([]*entity.OrderItem, error) {
	return nil, nil
}

func (r *Repo) AddItemsToOrder(ctx context.Context, items []*entity.OrderItem) error {
	return nil
}
