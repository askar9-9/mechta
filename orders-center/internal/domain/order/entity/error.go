package entity

import "errors"

var (
	ErrIDRequired          = errors.New("ID is required")
	ErrTypeRequired        = errors.New("type is required")
	ErrStatusRequired      = errors.New("status is required")
	ErrCityRequired        = errors.New("city is required")
	ErrSubdivisionRequired = errors.New("subdivision is required")
	ErrPriceInvalid        = errors.New("price cannot be negative")
	ErrPlatformRequired    = errors.New("platform is required")
	ErrGeneralIDInvalid    = errors.New("general ID is invalid")
	ErrOrderNumberRequired = errors.New("order number is required")
	ErrExecutorRequired    = errors.New("executor is required")
	ErrCreatedAtRequired   = errors.New("created at is required")
	ErrUpdatedAtRequired   = errors.New("updated at is required")
)

// Service errors
var (
	ErrOrderRequired   = errors.New("order is required")
	ErrOrderIDRequired = errors.New("order ID is required")
)

// Repository errors
var (
	ErrOrderNotFound = errors.New("order not found")
)
