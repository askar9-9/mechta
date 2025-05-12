package repository

import (
	"context"
	"errors"
	"orders-center/internal/domain/order/entity"
	"orders-center/internal/infrastructure/db/pgxtx"

	"github.com/jackc/pgx/v5"
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

func (r *Repo) GetOrderByID(ctx context.Context, id string) (*entity.Order, error) {
	q, ok := pgxtx.GetTx(ctx)
	if !ok {
		return nil, pgxtx.ErrNoTx
	}

	query := `
		SELECT id, type, status, city, subdivision, price, platform, general_id, order_number, executor, created_at, updated_at
		FROM orders
		WHERE id = $1
	`

	row := q.QueryRow(ctx, query, id)

	var order entity.Order
	err := row.Scan(
		&order.ID,
		&order.Type,
		&order.Status,
		&order.City,
		&order.Subdivision,
		&order.Price,
		&order.Platform,
		&order.GeneralID,
		&order.OrderNumber,
		&order.Executor,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, entity.ErrOrderNotFound
		}
		return nil, err
	}

	return &order, nil
}

func (r *Repo) CreateOrder(ctx context.Context, order *entity.Order) error {
	q, ok := pgxtx.GetTx(ctx)
	if !ok {
		return pgxtx.ErrNoTx
	}

	query := `
		INSERT INTO orders (id, type, status, city, subdivision, price, platform, general_id, order_number, executor, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err := q.Exec(ctx, query,
		order.ID,
		order.Type,
		order.Status,
		order.City,
		order.Subdivision,
		order.Price,
		order.Platform,
		order.GeneralID,
		order.OrderNumber,
		order.Executor,
		order.CreatedAt,
		order.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
