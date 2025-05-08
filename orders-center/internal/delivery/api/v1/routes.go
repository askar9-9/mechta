package v1

import "net/http"

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	h := NewHandler()
	mux.HandleFunc("POST /api/v1/orders", h.CreateOrder)

	return mux
}
