package collections

type Dequeue[T any] interface {
	Stack[T]
	PullFirst() (T, bool)
	First() (T, bool)
}
