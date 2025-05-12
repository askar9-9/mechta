package repository

import (
	"context"
	"orders-center/internal/domain/history/entity"
	"orders-center/internal/infrastructure/db/pgxtx"

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

func (r *Repo) GetHistories(ctx context.Context, orderID string) ([]*entity.History, error) {
	q, ok := pgxtx.GetTx(ctx)
	if !ok {
		return nil, pgxtx.ErrNoTx
	}

	query := `
		SELECT type, type_id, old_value, value, date, user_id, order_id
		FROM history
		WHERE order_id = $1
	`

	rows, err := q.Query(ctx, query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []*entity.History
	for rows.Next() {
		var history entity.History
		err := rows.Scan(
			&history.Type,
			&history.TypeId,
			&history.OldValue,
			&history.Value,
			&history.Date,
			&history.UserID,
			&history.OrderID,
		)
		if err != nil {
			return nil, err
		}
		histories = append(histories, &history)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	if len(histories) == 0 {
		return nil, entity.ErrHistoryNotFound
	}

	return histories, nil
}

func (r *Repo) CreateHistories(ctx context.Context, histories []*entity.History) error {
	q, ok := pgxtx.GetTx(ctx)
	if !ok {
		return pgxtx.ErrNoTx
	}

	query := `
		INSERT INTO history (type, type_id, old_value, value, date, user_id, order_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	for _, history := range histories {
		_, err := q.Exec(ctx, query,
			history.Type,
			history.TypeId,
			history.OldValue,
			history.Value,
			history.Date,
			history.UserID,
			history.OrderID,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
