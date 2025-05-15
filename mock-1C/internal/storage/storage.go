package storage

import (
	"mock-1C/internal/model"
	"sync"
)

type MemoryStorage struct {
	mu     sync.RWMutex
	orders []model.OrderFull
}

func NewStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Save(order model.OrderFull) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.orders = append(s.orders, order)
}

func (s *MemoryStorage) GetAll() []model.OrderFull {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.orders
}

func (s *MemoryStorage) GetByID(id string) (model.OrderFull, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, order := range s.orders {
		if order.Order.ID == id {
			return order, nil
		}
	}
	return model.OrderFull{}, model.ErrOrderNotFound
}
