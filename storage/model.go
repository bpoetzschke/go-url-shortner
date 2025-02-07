package storage

import "errors"

type Storage interface {
	Save(longURL string, shortURL string) error
	Get(shortURL string) (string, error)
}

var (
	ErrorNotFound = errors.New("not_found")
)
