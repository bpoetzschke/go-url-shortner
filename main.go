package main

import (
	"net/http"

	"github.com/bpoetzschke/go-url-shortner/api"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
)

func main() {
	// TODO init storage
	router := mux.NewRouter()
	api.AddRoutes(router)

	slog.Info("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}
