package businesslogic

import (
	"strings"

	"github.com/bpoetzschke/go-url-shortner/id"
	"github.com/bpoetzschke/go-url-shortner/storage"
)

var base62Map = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Shortener interface {
	Create(longURL string) (string, error)
}

func NewShortener(idGenerator id.Generator, storage storage.Storage) Shortener {
	return &shortener{
		idGenerator: idGenerator,
		storage:     storage,
	}
}

type shortener struct {
	idGenerator id.Generator
	storage     storage.Storage
}

func (s *shortener) Create(longURL string) (string, error) {
	nextID, err := s.idGenerator.Generate()
	if err != nil {
		return "", err
	}

	stringBuilder := &strings.Builder{}
	for nextID > 0 {
		stringBuilder.WriteRune(base62Map[nextID%62])
		nextID /= 62
	}

	shortURL := stringBuilder.String()
	err = s.storage.Save(longURL, shortURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}
