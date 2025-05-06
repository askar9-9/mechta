package service

import (
	"orders-center/internal/domain/order/repository"
)

type Order struct {
	repo repository.OrderRepository
}
