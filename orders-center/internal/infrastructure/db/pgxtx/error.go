package pgxtx

import "errors"

var (
	ErrNoTx = errors.New("no transaction found in context")
)
