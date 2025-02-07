package storage

import "github.com/bpoetzschke/go-url-shortner/errors"

func NewInMemory() Storage {
	return &inMemory{
		data: make(map[string]string),
	}
}

type inMemory struct {
	data map[string]string
}

func (m *inMemory) Save(longURL string, shortURL string) error {
	m.data[shortURL] = longURL
	return nil
}

func (m *inMemory) Get(shortURL string) (string, error) {
	longURL, found := m.data[shortURL]
	if !found {
		return "", errors.ErrorNotFound
	}

	return longURL, nil
}
