package api

import "github.com/gorilla/mux"

func AddRoutes(router *mux.Router) {
	router.HandleFunc("/create", handleCreate()).Methods("POST")
}
