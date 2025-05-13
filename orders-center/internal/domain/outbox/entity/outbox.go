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
