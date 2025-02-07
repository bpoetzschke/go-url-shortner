package businesslogic

type Shortener interface {
	Create(longURL string) (string, error)
}

func NewShortener() Shortener {
	return &shortener{}
}

type shortener struct {
}

func (s *shortener) Create(longURL string) (string, error) {
	return "abcd", nil
}
