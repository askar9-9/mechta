package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"orders-center/internal/domain/payment/entity"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetOrderPaymentByOrderID(ctx context.Context, orderID string) (*entity.OrderPayment, error) {
	return nil, nil
}

func (r *Repo) CreateOrderPayment(ctx context.Context, payment *entity.OrderPayment) error {
	return nil
}
