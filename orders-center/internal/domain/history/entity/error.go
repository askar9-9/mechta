package entity

import "errors"

// Validation errors
var (
	ErrTypeRequired     = errors.New("type is required")
	ErrTypeIdRequired   = errors.New("type ID is required")
	ErrOldValueRequired = errors.New("old value is required")
	ErrValueRequired    = errors.New("value is required")
	ErrDateRequired     = errors.New("date is required")
	ErrUserIDRequired   = errors.New("user ID is required")
	ErrOrderIDRequired  = errors.New("order ID is required")
)

// Service errors
var (
	ErrHistoryItemsRequired = errors.New("history items is required")
)

// Repository errors
var (
	ErrHistoryNotFound = errors.New("history not found")
)
