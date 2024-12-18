package priorityqueue

// Define an interface that Cell implements,
// so that we can extend Cell with more fields if needed.
type Cellable interface {
	GetCost() int
}

type Cell struct {
	Row, Col, Cost int
}

func (c Cell) GetCost() int {
	return c.Cost
}

type PriorityQueue []Cellable

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].GetCost() < pq[j].GetCost()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Cellable)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
