package entity

import (
	"context"
	"orders-center/internal/delivery/client"
	"orders-center/internal/domain/outbox/entity"
)

type Task struct {
	client     *client.OneCClient
	msg        *entity.Outbox
	onComplete func(ctx context.Context) error
}

func NewTask(client *client.OneCClient, msg *entity.Outbox, onComplete func(ctx context.Context) error) *Task {
	return &Task{
		client:     client,
		msg:        msg,
		onComplete: onComplete,
	}
}

func (t *Task) ID() string {
	return t.msg.ID.String()
}

func (t *Task) Process(ctx context.Context) error {
	if err := t.client.SendMessage(ctx, t.msg); err != nil {
		return err
	}

	t.msg.SetProcessedAt()
	return nil
}

func (t *Task) OnComplete(ctx context.Context) error {
	return t.onComplete(ctx)
}
