package entity

import "time"

type History struct {
	Type     string    `json:"type"`
	TypeId   int       `json:"type_id"`
	OldValue []byte    `json:"old_value"`
	Value    []byte    `json:"value"`
	Date     time.Time `json:"date"`
	UserID   string    `json:"user_id"`
	OrderID  string    `json:"order_id"`
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
	if h.UserID == "" {
		return ErrUserIDRequired
	}
	if h.OrderID == "" {
		return ErrOrderIDRequired
	}
	return nil
}
