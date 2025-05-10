package pgxtx

import (
	"context"
	"github.com/jackc/pgx/v5"
	"orders-center/internal/pkg/tx"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxTransactionManager struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) tx.TransactionManager {
	return &PgxTransactionManager{pool: pool}
}

func (m *PgxTransactionManager) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := m.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		}
	}()

	ctxWithTx := InjectTx(ctx, tx)

	if err := fn(ctxWithTx); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}
