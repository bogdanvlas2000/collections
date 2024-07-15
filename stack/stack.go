package collections

type Stack[T any] struct {
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

	if s.head == nil {
		s.head = newNode
		s.tail = newNode
		return
	}

	newNode.prev = s.tail
	s.tail.next = newNode
	s.tail = newNode
}

func (s *Stack[T]) Pop() (value T, ok bool) {
	if s.tail == nil {
		var emptyVal T
		return emptyVal, false
	}

	value = s.tail.value

	if s.tail == s.head {
		s.head = nil
		s.tail = nil
		return value, true
	}

	prevNode := s.tail.prev
	prevNode.next = nil
	s.tail = prevNode

	return value, true
}

func (s *Stack[T]) Peek() (value T, ok bool) {
	if s.tail == nil {
		var emptyVal T
		return emptyVal, false
	}

	return s.tail.value, true
}

func (s *Stack[T]) Size() int {
	if s.head == nil {
		return 0
	}
	if s.head == s.tail {
		return 1
	}

	walker := s.head
	counter := 0
	for walker != nil {
		walker = walker.next
		counter++
	}
	return counter
}
