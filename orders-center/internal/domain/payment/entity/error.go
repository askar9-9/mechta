package entity

import "errors"

var (
	ErrOrderIDIsEmpty         = errors.New("order ID is empty")
	ErrSumIsZero              = errors.New("payment sum must be greater than zero")
	ErrPaymentTypeIsEmpty     = errors.New("payment type is empty")
	ErrContractNumberIsEmpty  = errors.New("contract number is empty")
	ErrExternalIDIsEmpty      = errors.New("external ID is empty")
	ErrCreditDataInvalid      = errors.New("credit data is invalid")
	ErrCardPaymentDataInvalid = errors.New("card payment data is invalid")
)
