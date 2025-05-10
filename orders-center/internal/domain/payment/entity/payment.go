package entity

import "github.com/google/uuid"

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

type OrderPayment struct {
	ID              uuid.UUID        `json:"id"`
	OrderID         uuid.UUID        `json:"order_id"`
	Type            PaymentType      `json:"type"`
	Sum             float64          `json:"sum"`
	Payed           bool             `json:"payed"`
	Info            string           `json:"info"`
	CreditData      *CreditData      `json:"credit_data"`
	ContractNumber  string           `json:"contract_number"`
	CardPaymentData *CardPaymentData `json:"card_payment_data"`
	ExternalID      string           `json:"external_id"`
}

type CreditData struct {
	Bank           string  `json:"bank"`
	Type           string  `json:"type"`
	NumberOfMonths int16   `json:"number_of_months"`
	PaySumPerMonth float64 `json:"pay_sum_per_month"`
	BrokerID       int32   `json:"broker_id"`
	IIN            string  `json:"iin"`
}

type CardPaymentData struct {
	Provider      string `json:"provider"`
	TransactionId string `json:"transaction_id"`
}

func (o *OrderPayment) Validate() error {
	if o.OrderID == uuid.Nil {
		return ErrOrderIDIsEmpty
	}
	if o.Sum <= 0 {
		return ErrSumIsZero
	}
	if o.Type == "" {
		return ErrPaymentTypeIsEmpty
	}
	if o.ContractNumber == "" {
		return ErrContractNumberIsEmpty
	}
	if o.ExternalID == "" {
		return ErrExternalIDIsEmpty
	}

	if o.Type == PaymentTypeCredit && o.CreditData != nil {
		if err := o.CreditData.Validate(); err != nil {
			return err
		}
	}

	if (o.Type == PaymentTypeCard || o.Type == PaymentTypeCardOnline) && o.CardPaymentData != nil {
		if err := o.CardPaymentData.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (cd *CreditData) Validate() error {
	if cd.Bank == "" || cd.Type == "" || cd.NumberOfMonths <= 0 || cd.PaySumPerMonth <= 0 || cd.BrokerID <= 0 || cd.IIN == "" {
		return ErrCreditDataInvalid
	}
	return nil
}

func (cp *CardPaymentData) Validate() error {
	if cp.Provider == "" || cp.TransactionId == "" {
		return ErrCardPaymentDataInvalid
	}
	return nil
}
