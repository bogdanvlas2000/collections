package set

type Set[T comparable] struct {
	entries map[T]struct{}
}

func New[T comparable](keys ...T) *Set[T] {
	entries := make(map[T]struct{})

	for _, key := range keys {
		entries[key] = struct{}{}
	}

	return &Set[T]{
		entries: entries,
	}
}

func (s *Set[T]) Size() int {
	return len(s.entries)
}

func (s *Set[T]) Put(key T) bool {
	_, ok := s.entries[key]
	if ok {
		return false
	}
	s.entries[key] = struct{}{}
	return true
}

func (s *Set[T]) Remove(key T) bool {
	_, ok := s.entries[key]
	if !ok {
		return false
	}
	delete(s.entries, key)
	return true
}

func (s *Set[T]) Contains(key T) bool {
	_, ok := s.entries[key]
	return ok
}

func (s *Set[T]) Equals(set *Set[T]) bool {
	if len(s.entries) != set.Size() {
		return false
	}

	for key, _ := range s.entries {
		if !set.Contains(key) {
			return false
		}
	}

	return true
}
