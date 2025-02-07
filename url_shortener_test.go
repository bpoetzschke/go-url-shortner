package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bpoetzschke/go-url-shortner/api"
	"github.com/bpoetzschke/go-url-shortner/businesslogic"
	"github.com/bpoetzschke/go-url-shortner/id"
	"github.com/bpoetzschke/go-url-shortner/storage"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestUrlShortenerCreate(t *testing.T) {
	idGenerator := id.NewInMemory(1_000_000)
	storage := storage.NewInMemory()
	shortener := businesslogic.NewShortener(idGenerator, storage)
	router := mux.NewRouter()
	api.AddRoutes(router, shortener)
	body := bytes.NewBufferString(`{"url": "https://www.google.com"}`)
	request, err := http.NewRequest("POST", "/create", body)
	require.NoError(t, err)
	requestRecorder := httptest.NewRecorder()

	router.ServeHTTP(requestRecorder, request)

	require.Equal(t, http.StatusCreated, requestRecorder.Code)
	require.JSONEq(t, `{"url": "/39c4"}`, requestRecorder.Body.String())
}

func TestURLShortenerReturns400ForEmptyURL(t *testing.T) {
	router := mux.NewRouter()
	api.AddRoutes(router, nil)

	body := bytes.NewBufferString(`{}`)
	request, err := http.NewRequest("POST", "/create", body)
	require.NoError(t, err)
	requestRecorder := httptest.NewRecorder()

	router.ServeHTTP(requestRecorder, request)
	require.Equal(t, http.StatusBadRequest, requestRecorder.Code)
}

func TestURLShortenerReturns500ForInvalidJSONPayload(t *testing.T) {
	router := mux.NewRouter()
	api.AddRoutes(router, nil)

	body := bytes.NewBufferString(`{`)
	request, err := http.NewRequest("POST", "/create", body)
	require.NoError(t, err)
	requestRecorder := httptest.NewRecorder()

	router.ServeHTTP(requestRecorder, request)
	require.Equal(t, http.StatusInternalServerError, requestRecorder.Code)
}

func TestURLShortenerRedirectsForExistingShortURL(t *testing.T) {
	idGenerator := id.NewInMemory(1_000_000)
	storage := storage.NewInMemory()
	shortener := businesslogic.NewShortener(idGenerator, storage)
	router := mux.NewRouter()
	api.AddRoutes(router, shortener)
	body := bytes.NewBufferString(`{"url": "https://www.google.com"}`)
	request, err := http.NewRequest("POST", "/create", body)
	require.NoError(t, err)
	requestRecorder := httptest.NewRecorder()
	router.ServeHTTP(requestRecorder, request)
	var responseBody map[string]string
	json.Unmarshal(requestRecorder.Body.Bytes(), &responseBody)

	body = bytes.NewBufferString("")
	request, err = http.NewRequest("GET", responseBody["url"], body)
	requestRecorder = httptest.NewRecorder()
	router.ServeHTTP(requestRecorder, request)

	require.Equal(t, http.StatusMovedPermanently, requestRecorder.Code)
	require.Equal(t, "https://www.google.com", requestRecorder.Header().Get("Location"))
}
