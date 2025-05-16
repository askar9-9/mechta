package entity

import (
	"context"
	"time"
)

type BaseJob struct {
	id          string
	processFunc func(ctx context.Context) error
	onComplete  func(ctx context.Context) error
	onError     func(ctx context.Context, err error) error
	maxRetries  int
	retryDelay  time.Duration
	timeout     time.Duration
}

func NewBaseJob(
	id string,
	processFunc func(ctx context.Context) error,
	onComplete func(ctx context.Context) error,
	onError func(ctx context.Context, err error) error,
) *BaseJob {
	job := &BaseJob{
		id:          id,
		processFunc: processFunc,
		onComplete:  onComplete,
		onError:     onError,
		maxRetries:  3,
		retryDelay:  500 * time.Millisecond,
		timeout:     5 * time.Second,
	}

	return job
}

func (j *BaseJob) ID() string {
	return j.id
}

func (j *BaseJob) Process(ctx context.Context) error {
	return j.processFunc(ctx)
}

func (j *BaseJob) OnComplete(ctx context.Context) error {
	if j.onComplete == nil {
		return nil
	}
	return j.onComplete(ctx)
}

func (j *BaseJob) OnError(ctx context.Context, err error) error {
	if j.onError == nil {
		return err
	}
	return j.onError(ctx, err)
}

func (j *BaseJob) MaxRetries() int {
	return j.maxRetries
}

func (j *BaseJob) RetryDelay() time.Duration {
	return j.retryDelay
}

func (j *BaseJob) Timeout() time.Duration {
	return j.timeout
}
