package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bpoetzschke/go-url-shortner/api"
	"github.com/bpoetzschke/go-url-shortner/businesslogic"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestUrlShortenerCreate(t *testing.T) {
	shortener := businesslogic.NewShortener()
	router := mux.NewRouter()
	api.AddRoutes(router, shortener)

	body := bytes.NewBufferString(`{"url": "https://www.google.com"}`)
	request, err := http.NewRequest("POST", "/create", body)
	require.NoError(t, err)
	requestRecorder := httptest.NewRecorder()
	router.ServeHTTP(requestRecorder, request)

	require.Equal(t, http.StatusCreated, requestRecorder.Code)
	require.JSONEq(t, `{"url": "/abcd"}`, requestRecorder.Body.String())
}
