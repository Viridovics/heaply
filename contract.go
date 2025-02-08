package heaply

type Heap[T any] interface {
	Top() (top T, ok bool)
	Pop() (top T, ok bool)
	Push(elem T)
	Size() int
}
