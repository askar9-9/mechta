package storage

import "mock-1C/internal/model"

type StorageModule interface {
	Save(order model.OrderFull)
	GetAll() []model.OrderFull
	GetByID(id string) (model.OrderFull, error)
}
