package entity

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Outbox struct {
	ID            uuid.UUID
	AggregateID   uuid.UUID
	AggregateType string
	EventType     string
	Payload       json.RawMessage
	CreatedAt     time.Time
	ProcessedAt   time.Time
	RetryCount    int
	Error         string
}
