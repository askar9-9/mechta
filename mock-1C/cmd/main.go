package main

import (
	"log"
	"mock-1C/internal/handler"
	"mock-1C/internal/storage"
	"net/http"
	"time"
)

func main() {
	store := storage.NewStorage()

	router := handler.NewRouter(store)

	server := &http.Server{
		Addr:         ":9900",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Listening on port%s", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
