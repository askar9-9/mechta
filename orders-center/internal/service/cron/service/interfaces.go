package service

import (
	"context"
	"time"
)

type Job interface {
	ID() string
	Process(ctx context.Context) error
	OnComplete(ctx context.Context) error
	OnError(ctx context.Context, err error) error
	MaxRetries() int
	RetryDelay() time.Duration
	Timeout() time.Duration
}
