package entity

import (
	cart "orders-center/internal/domain/cart/entity"
	history "orders-center/internal/domain/history/entity"
	order "orders-center/internal/domain/order/entity"
	payment "orders-center/internal/domain/payment/entity"
)

type OrderFull struct {
	Order    order.Order            `json:"order"`
	Items    []cart.OrderItem       `json:"items"`
	Payments []payment.OrderPayment `json:"payments"`
	History  []history.History      `json:"history"`
}
