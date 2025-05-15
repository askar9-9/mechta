package cron

import "context"

type Job interface {
	ID() string
	Process(ctx context.Context) error
	OnComplete(ctx context.Context) error
}
