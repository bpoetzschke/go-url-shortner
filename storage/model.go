package storage

type Storage interface {
	Save(longURL string, shortURL string) error
	Get(shortURL string) (string, error)
}
