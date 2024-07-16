package collections

type Stack[T any] struct {
	size int
	head *node[T]
	tail *node[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

func (s *Stack[T]) Push(value T) {
	newNode := &node[T]{
		value: value,
	}

	if s.size == 0 {
		s.head = newNode
		s.tail = newNode
	} else {
		newNode.prev = s.tail
		s.tail.next = newNode
		s.tail = newNode
	}

	s.size++
}

func (s *Stack[T]) Pop() (value T, ok bool) {
	if s.size == 0 {
		var emptyVal T
		return emptyVal, false
	}

	value = s.tail.value

	if s.size == 1 {
		s.head = nil
		s.tail = nil
		s.size--
		return value, true
	}

	prevNode := s.tail.prev
	prevNode.next = nil
	s.tail = prevNode
	s.size--

	return value, true
}

func (s *Stack[T]) Peek() (value T, ok bool) {
	if s.size == 0 {
		var emptyVal T
		return emptyVal, false
	}

	return s.tail.value, true
}

func (s *Stack[T]) Size() int {
	return s.size
}
