package collections

type Stack[T any] interface {
	Size() int
	Push(value T)
	Pop() (T, bool)
	Peek() (T, bool)
}
