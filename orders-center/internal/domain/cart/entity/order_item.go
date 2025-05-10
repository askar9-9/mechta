package entity

import (
	"github.com/google/uuid"
)

type OrderItem struct {
	ProductID     string    `json:"product_id"`
	ExternalID    string    `json:"external_id"`
	Status        string    `json:"status"`
	BasePrice     float64   `json:"base_price"`
	Price         float64   `json:"price"`
	EarnedBonuses float64   `json:"earned_bonuses"`
	SpentBonuses  float64   `json:"spent_bonuses"`
	Gift          bool      `json:"gift"`
	OwnerID       string    `json:"owner_id"`
	DeliveryID    string    `json:"delivery_id"`
	ShopAssistant string    `json:"shop_assistant"`
	Warehouse     string    `json:"warehouse"`
	OrderID       uuid.UUID `json:"order_id"`
}

func (o *OrderItem) Validate() error {
	if o.ProductID == "" {
		return ErrProductIDRequired
	}
	if o.ExternalID == "" {
		return ErrExternalIDRequired
	}
	if o.Status == "" {
		return ErrInvalidStatus
	}
	if o.BasePrice < 0 {
		return ErrBasePriceInvalid
	}
	if o.Price < 0 {
		return ErrPriceInvalid
	}
	if o.EarnedBonuses < 0 || o.SpentBonuses < 0 {
		return ErrBonusesInvalid
	}
	if o.OwnerID == "" {
		return ErrOwnerIDRequired
	}
	if o.Warehouse == "" {
		return ErrWarehouseRequired
	}
	if o.OrderID == uuid.Nil {
		return ErrOrderIDInvalid
	}
	return nil
}
