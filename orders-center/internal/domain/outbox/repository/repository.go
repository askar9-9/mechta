package repository

import (
	"context"
	"orders-center/internal/domain/outbox/entity"
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

func (r *Repo) CreateTask(ctx context.Context, item *entity.Outbox) error {
	q, ok := pgxtx.GetTx(ctx)
	if !ok {
		return pgxtx.ErrNoTx
	}

	query := `
		INSERT INTO outbox_messages (
			id, aggregate_id, aggregate_type, event_type, payload, created_at, processed_at, retry_count, error
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		)
	`
	_, err := q.Exec(ctx, query, item.ID, item.AggregateID, item.AggregateType, item.EventType, item.Payload, item.CreatedAt, item.ProcessedAt, item.RetryCount, item.Error)
	if err != nil {
		return err
	}

	return nil
}
func (r *Repo) GetLimitedTaskList(ctx context.Context, limit int) ([]*entity.Outbox, error) {
	return nil, nil
}
