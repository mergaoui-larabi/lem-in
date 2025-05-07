package dsa

type Node struct {
	v     int
	index int
	room  string
}

type Queue []*Node

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i].v < q[j].v
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = j
	q[j].index = i
}

func (q *Queue) Push(x interface{}) {
	n := len(*q)
	item := x.(*Node) // cast interface to *Node
	item.index = n
	*q = append(*q, item)

}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[:n-1]
	return item
}
