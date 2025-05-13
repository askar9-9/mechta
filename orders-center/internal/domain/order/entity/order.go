package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID `json:"id"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	City        string    `json:"city"`
	Subdivision string    `json:"subdivision"`
	Price       float64   `json:"price"`
	Platform    string    `json:"platform"`
	GeneralID   uuid.UUID `json:"general_id"`
	OrderNumber string    `json:"order_number"`
	Executor    string    `json:"executor"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (o *Order) Validate() error {
	if o.ID == uuid.Nil {
		return ErrIDRequired
	}
	if o.Type == "" {
		return ErrTypeRequired
	}
	if o.Status == "" {
		return ErrStatusRequired
	}
	if o.City == "" {
		return ErrCityRequired
	}
	if o.Subdivision == "" {
		return ErrSubdivisionRequired
	}
	if o.Price < 0 {
		return ErrPriceInvalid
	}
	if o.Platform == "" {
		return ErrPlatformRequired
	}
	if o.GeneralID == uuid.Nil {
		return ErrGeneralIDInvalid
	}
	if o.OrderNumber == "" {
		return ErrOrderNumberRequired
	}
	if o.Executor == "" {
		return ErrExecutorRequired
	}
	if o.CreatedAt.IsZero() {
		return ErrCreatedAtRequired
	}
	if o.UpdatedAt.IsZero() {
		return ErrUpdatedAtRequired
	}
	return nil
}
