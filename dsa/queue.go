package main

import "fmt"

type Queue struct {
	items []int
}

// A queue works in FIFO order
func CreateQueue() *Queue {
	return &Queue{
		items: make([]int, 0),
	}
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Dequeue() int {
	item := q.items[0]

	q.items = q.items[1:]
	fmt.Println("Dequeued", item)
	return item
}
