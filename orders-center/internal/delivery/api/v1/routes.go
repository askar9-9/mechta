package v1

import (
	"net/http"
)

func Routes(svc OrderService) *http.ServeMux {
	mux := http.NewServeMux()

	h := NewHandler(svc)
	mux.HandleFunc("POST /api/v1/orders", h.CreateOrder)

	return mux
}
