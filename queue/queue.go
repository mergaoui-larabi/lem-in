package queue

import "fmt"

type Node struct {
	value string
	next  *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (q *Queue) IsEmpty() bool {
	return q.Head == nil
}

func (q *Queue) Enqueue(v string) {
	n := &Node{
		value: v,
		next:  nil,
	}
	if q.Tail == nil {
		q.Head = n
		q.Tail = n
	} else {
		q.Tail.next = n
		q.Tail = n
	}
	q.Length++
}

func (q *Queue) Dequeue() string {
	if q.Head == nil {
		return ""
	}
	deq := q.Head.value
	q.Head = q.Head.next
	q.Length--

	if q.Head == nil {
		q.Tail = nil
	}

	return deq
}

func (q *Queue) Print() {
	
	if q.Head == nil {
		fmt.Println("Queue is empty")
		return
	}
	travel := q.Head
	for travel != nil {
		fmt.Print(travel.value, " -> ")
		travel = travel.next
	}
	fmt.Println("nil")
}
