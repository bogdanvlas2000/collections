package linkedlist

type LinkedList[T any] struct {
	size int
	head *node[T]
	tail *node[T]
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

func (l *LinkedList[T]) Push(value T) {
	newNode := &node[T]{
		value: value,
	}

	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}

	l.size++
}

func (l *LinkedList[T]) Pop() (value T, ok bool) {
	if l.size == 0 {
		var emptyVal T
		return emptyVal, false
	}

	value = l.tail.value

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size--
		return value, true
	}

	prevNode := l.tail.prev
	prevNode.next = nil
	l.tail = prevNode
	l.size--

	return value, true
}

func (l *LinkedList[T]) Peek() (value T, ok bool) {
	if l.size == 0 {
		var emptyVal T
		return emptyVal, false
	}

	return l.tail.value, true
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) PullFirst() (value T, ok bool) {
	if l.size == 0 {
		var emptyVal T
		return emptyVal, false
	}

	value = l.head.value

	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		nextNode := l.head.next
		nextNode.prev = nil
		l.head = nextNode
	}

	l.size--
	return value, true
}

func (l *LinkedList[T]) First() (value T, ok bool) {
	if l.size == 0 {
		var emptyVal T
		return emptyVal, false
	}

	return l.head.value, true
}
