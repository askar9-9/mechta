package entity

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Outbox struct {
	ID            uuid.UUID
	AggregateID   uuid.UUID
	AggregateType AggregateType
	EventType     EventType
	Payload       json.RawMessage
	CreatedAt     time.Time
	SyncAt        time.Time
	ProcessedAt   time.Time
	RetryCount    int
	Error         string
}

type AggregateType string

var (
	AggregateTypeOrder AggregateType = "order"
)

type EventType string

var (
	EventTypeOrderCreated EventType = "order_created"
)

func (o *Outbox) IsProcessed() bool {
	return o.ProcessedAt != time.Time{}
}

func (o *Outbox) SetError(err error) {
	if err != nil {
		o.Error = err.Error()
	}
}

func (o *Outbox) SetProcessedAt() {
	o.ProcessedAt = time.Now()
}

func (o *Outbox) SetSyncAt() {
	o.SyncAt = time.Now()
}

func (o *Outbox) SetRetryCount(count int) {
	o.RetryCount = count
}
