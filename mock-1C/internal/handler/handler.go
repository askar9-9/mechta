package handler

import (
	"errors"
	"io"
	"mock-1C/internal/model"
	"mock-1C/internal/storage"
	"mock-1C/pkg/xjson"
	"net/http"
)

type Handler struct {
	store storage.StorageModule
}

func NewHandler(store storage.StorageModule) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orders := h.store.GetAll()

	data, err := xjson.MarshalJson[[]model.OrderFull](orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	order, err := h.store.GetByID(id)
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := xjson.MarshalJson[model.OrderFull](order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid content type", http.StatusUnsupportedMediaType)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newOrder, err := xjson.UnmarshalJson[model.OrderFull](data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.store.Save(newOrder)

	w.WriteHeader(http.StatusCreated)
}
