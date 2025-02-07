package main

import (
	"log/slog"
	"net/http"

	"github.com/bpoetzschke/go-url-shortner/api"
	"github.com/bpoetzschke/go-url-shortner/businesslogic"
	"github.com/gorilla/mux"
)

func main() {
	// TODO init storage
	shortener := businesslogic.NewShortener()

	router := mux.NewRouter()
	api.AddRoutes(router, shortener)

	slog.Info("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}
