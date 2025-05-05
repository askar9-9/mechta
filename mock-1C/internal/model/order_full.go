package model

type OrderFull struct {
	Order    Order          `json:"order"`
	Items    []OrderItem    `json:"items"`
	Payments []OrderPayment `json:"payments"`
}
