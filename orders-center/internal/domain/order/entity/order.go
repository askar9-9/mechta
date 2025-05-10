package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          string
	Type        string
	Status      string
	City        string
	Subdivision string
	Price       float64
	Platform    string
	GeneralID   uuid.UUID
	OrderNumber string
	Executor    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (o *Order) Validate() error {
	if o.ID == "" {
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
