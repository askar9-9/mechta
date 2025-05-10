package entity

import "errors"

// Validation errors
var (
	ErrProductIDRequired  = errors.New("product ID is required")
	ErrExternalIDRequired = errors.New("external ID is required")
	ErrInvalidStatus      = errors.New("invalid order item status")
	ErrBasePriceInvalid   = errors.New("base price cannot be negative")
	ErrPriceInvalid       = errors.New("price cannot be negative")
	ErrBonusesInvalid     = errors.New("bonuses cannot be negative")
	ErrOwnerIDRequired    = errors.New("owner ID is required")
	ErrWarehouseRequired  = errors.New("warehouse is required")
	ErrOrderIDInvalid     = errors.New("order ID is invalid")
)

// Repository errors
var (
	ErrItemNotFound = errors.New("order item not found")
)
