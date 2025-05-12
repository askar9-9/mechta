package entity

import (
	"github.com/google/uuid"
	"time"
)

type History struct {
	Type     string    `json:"type"`
	TypeId   int       `json:"type_id"`
	OldValue []byte    `json:"old_value"`
	Value    []byte    `json:"value"`
	Date     time.Time `json:"date"`
	UserID   uuid.UUID `json:"user_id"`
	OrderID  uuid.UUID `json:"order_id"`
}

func (h *History) Validate() error {
	if h.Type == "" {
		return ErrTypeRequired
	}
	if h.TypeId == 0 {
		return ErrTypeIdRequired
	}
	if h.OldValue == nil {
		return ErrOldValueRequired
	}
	if h.Value == nil {
		return ErrValueRequired
	}
	if h.Date.IsZero() {
		return ErrDateRequired
	}
	if h.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if h.OrderID == uuid.Nil {
		return ErrOrderIDRequired
	}
	return nil
}
