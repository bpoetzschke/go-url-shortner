package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/bpoetzschke/go-url-shortner/businesslogic"
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
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusCreated)
		responseWriter.Write([]byte(fmt.Sprintf(`{"url": "/%s"}`, shortURL)))
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
