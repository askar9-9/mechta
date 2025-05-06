package entity

import (
	entity2 "orders-center/internal/domain/cart/entity"
	entity4 "orders-center/internal/domain/history/entity"
	"orders-center/internal/domain/order/entity"
	entity3 "orders-center/internal/domain/payment/entity"
)

type OrderFull struct {
	Order    entity.Order           `json:"order"`
	Items    []entity2.OrderItem    `json:"items"`
	Payments []entity3.OrderPayment `json:"payments"`
	History  []entity4.History      `json:"history"`
}
