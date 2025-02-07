package id

func NewInMemory(start int64) Generator {
	return &inMemory{
		counter: start,
	}
}

type inMemory struct {
	counter int64
}

func (m *inMemory) Generate() (int64, error) {
	m.counter++
	return m.counter, nil
}
