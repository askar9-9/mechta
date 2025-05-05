package model

import "github.com/google/uuid"

type OrderPayment struct {
	ID              uuid.UUID        `json:"id"`
	OrderID         uuid.UUID        `json:"orderID"`
	Type            PaymentType      `json:"type"`
	Sum             float64          `json:"sum"`
	Payed           bool             `json:"payed"`
	Info            string           `json:"info"`
	CreditData      *CreditData      `json:"creditData"`
	ContractNumber  string           `json:"contractNumber"`
	CardPaymentData *CardPaymentData `json:"cardPaymentData"`
	ExternalID      string           `json:"externalId"`
}

type PaymentType string

const (
	PaymentTypeCashAtShop    PaymentType = "cash_at_shop"
	PaymentTypeCashToCourier PaymentType = "cash_to_courier"
	PaymentTypeCard          PaymentType = "card"
	PaymentTypeCardOnline    PaymentType = "card_online"
	PaymentTypeCredit        PaymentType = "credit"
	PaymentTypeBonuses       PaymentType = "bonuses"
	PaymentTypeCashless      PaymentType = "cashless"
	PaymentTypePrepayment    PaymentType = "prepayment"
)

type CreditData struct {
	Bank           string  `json:"bank"`
	Type           string  `json:"type"`
	NumberOfMonths int16   `json:"numberOfMonths"`
	PaySumPerMonth float64 `json:"paySumPerMonth"`
	BrokerID       int32   `json:"brokerId"`
	IIN            string  `json:"iin"`
}

type CardPaymentData struct {
	Provider      string `json:"provider"`
	TransactionId string `json:"transactionId"`
}
