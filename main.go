package main

import (
	"log/slog"
	"net/http"

	"github.com/bpoetzschke/go-url-shortner/api"
	"github.com/bpoetzschke/go-url-shortner/businesslogic"
	"github.com/bpoetzschke/go-url-shortner/id"
	"github.com/bpoetzschke/go-url-shortner/storage"
	"github.com/gorilla/mux"
)

func main() {
	// TODO init storage
	idGenerator := id.NewInMemory(1_000_000)
	storage := storage.NewInMemory()
	shortener := businesslogic.NewShortener(idGenerator, storage)

	router := mux.NewRouter()
	api.AddRoutes(router, shortener)

	slog.Info("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}
