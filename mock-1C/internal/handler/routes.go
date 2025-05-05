package handler

import (
	"mock-1C/internal/storage"
	"net/http"
)

func NewRouter(serv storage.StorageModule) *http.ServeMux {
	router := http.NewServeMux()

	h := NewHandler(serv)

	router.HandleFunc("GET /api/v1/orders", h.GetOrders)
	router.HandleFunc("GET /api/v1/order", h.CreateOrder)
	router.HandleFunc("POST /api/v1/order/{id}", h.GetOrderByID)

	return router
}
