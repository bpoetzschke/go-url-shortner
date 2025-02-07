package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/bpoetzschke/go-url-shortner/businesslogic"
	"github.com/bpoetzschke/go-url-shortner/errors"
	"github.com/gorilla/mux"
)

func handleCreate(shortener businesslogic.Shortener) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var createRequestPayload CreateRequestPayload
		err := json.NewDecoder(request.Body).Decode(&createRequestPayload)
		if err != nil {
			slog.Error("Error decoding request payload", "err", err)
			write500Response(responseWriter, "error decoding request payload")
			return
		}

		if createRequestPayload.URL == "" {
			write400Response(responseWriter, "empty URL")
			return
		}

		shortURL, err := shortener.Create(createRequestPayload.URL)
		if err != nil {
			slog.Error("Error creating short URL", "err", err)
			write500Response(responseWriter, "error creating short URL")
			return
		}
		slog.Info("Created short URL", "shortURL", shortURL)
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusCreated)
		responseWriter.Write([]byte(fmt.Sprintf(`{"url": "/%s"}`, shortURL)))
	}
}

func handleRedirect(shortener businesslogic.Shortener) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		shortURL := vars["shortURL"]

		slog.Info("Received redirect request for short URL", "shortURL", shortURL)
		longURL, err := shortener.Get(shortURL)
		if err == errors.ErrorNotFound {
			responseWriter.WriteHeader(http.StatusNotFound)
			return
		}

		http.Redirect(responseWriter, request, longURL, http.StatusMovedPermanently)
	}
}

func write500Response(responseWriter http.ResponseWriter, errorMsg string) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusInternalServerError)
	responseWriter.Write([]byte(`{"error": "` + errorMsg + `"}`))
}

func write400Response(responseWriter http.ResponseWriter, msg string) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusBadRequest)
	responseWriter.Write([]byte(`{"msg": "` + msg + `"}`))
}
