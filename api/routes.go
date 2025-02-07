package api

import (
	"github.com/bpoetzschke/go-url-shortner/businesslogic"
	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router, shortener businesslogic.Shortener) {
	router.HandleFunc("/create", handleCreate(shortener)).Methods("POST")
	router.HandleFunc("/{shortURL}", handleRedirect(shortener)).Methods("GET")
}
