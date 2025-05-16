package model

type OrderFull struct {
	Order    Order          `json:"order"`
	Items    []OrderItem    `json:"items"`
	Payments []OrderPayment `json:"payment"`
	History  []History      `json:"history"`
}
