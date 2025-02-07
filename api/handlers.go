package api

import "net/http"

func handleCreate() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.WriteHeader(http.StatusCreated)
		responseWriter.Write([]byte(`{"url": "/your-short-url"}`))
	}
}
