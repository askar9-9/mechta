package pgxtx

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type txKeyType struct{}

var txKey = txKeyType{}

func InjectTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

func GetTx(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(txKey).(pgx.Tx)
	return tx, ok
}
