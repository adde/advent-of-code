package queue

type Queue[T any] []T

func New[T any](item T) Queue[T] {
	q := make(Queue[T], 0)
	q.Append(item)
	return q
}

func (q *Queue[T]) Append(item T) {
	*q = append(*q, item)
}

func (q *Queue[T]) Pop() T {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q Queue[T]) Size() int {
	return len(q)
}

func (q Queue[T]) IsEmpty() bool {
	return len(q) == 0
}
