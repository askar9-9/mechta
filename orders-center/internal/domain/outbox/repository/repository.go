package repository

import (
	"context"
	"orders-center/internal/domain/outbox/entity"
	"orders-center/internal/infrastructure/db/pgxtx"

	"errors"
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

func (r *Repo) GetLimitedMessagesList(ctx context.Context, limit int) ([]*entity.Outbox, error) {
	q, ok := pgxtx.GetTx(ctx)
	if !ok {
		return nil, pgxtx.ErrNoTx
	}

	query := `
		SELECT id, aggregate_id, aggregate_type, event_type, payload, created_at, processed_at, retry_count, error
		FROM outbox_messages
		WHERE processed_at IS NULL
		  AND (sync_at IS NULL OR sync_at < NOW() - INTERVAL '10 minutes')
		ORDER BY created_at ASC, sync_at ASC
		LIMIT $1
		FOR UPDATE SKIP LOCKED
	`

	rows, err := q.Query(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*entity.Outbox
	for rows.Next() {
		var msg entity.Outbox
		err := rows.Scan(
			&msg.ID,
			&msg.AggregateID,
			&msg.AggregateType,
			&msg.EventType,
			&msg.Payload,
			&msg.CreatedAt,
			&msg.ProcessedAt,
			&msg.RetryCount,
			&msg.Error,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &msg)
	}

	if err := rows.Err(); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return messages, nil
}

func (r *Repo) UpdateOutboxBatch(ctx context.Context, items []*entity.Outbox) error {
	q, ok := pgxtx.GetTx(ctx)
	if !ok {
		return pgxtx.ErrNoTx
	}

	query := `
		UPDATE outbox_messages
		SET processed_at = $1
		WHERE id = $2
	`

	for _, item := range items {
		_, err := q.Exec(ctx, query, item.ProcessedAt, item.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repo) UpdateOutboxSingle(ctx context.Context, item *entity.Outbox) error {
	q, ok := pgxtx.GetTx(ctx)
	if !ok {
		return pgxtx.ErrNoTx
	}

	query := `
		UPDATE outbox_messages
		SET aggregate_id = $2, aggregate_type = $3, event_type = $4, payload = $5, created_at = $6, processed_at = $7, retry_count = $8, error = $9
		WHERE id = $1
	`

	_, err := q.Exec(ctx, query, item.ID, item.AggregateID, item.AggregateType, item.EventType, item.Payload, item.CreatedAt, item.ProcessedAt, item.RetryCount, item.Error)
	if err != nil {
		return err
	}

	return nil
}
