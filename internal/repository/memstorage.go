package repository

type MemStorage[T any] struct {
	storage map[string]T
}

func NewMemStorage[T any]() *MemStorage[T] {
	return &MemStorage[T]{
		storage: make(map[string]T),
	}
}

func (s *MemStorage[T]) Set(key string, value T) {
	s.storage[key] = value
}

func (s *MemStorage[T]) Get(key string) T {
	return s.storage[key]
}
