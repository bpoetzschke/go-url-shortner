package businesslogic

import (
	"strings"

	"github.com/bpoetzschke/go-url-shortner/id"
)

var base62Map = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Shortener interface {
	Create(longURL string) (string, error)
}

func NewShortener(idGenerator id.Generator) Shortener {
	return &shortener{
		idGenerator: idGenerator,
	}
}

type shortener struct {
	idGenerator id.Generator
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

	return stringBuilder.String(), nil
}
