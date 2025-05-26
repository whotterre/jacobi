package utils

// Queue is a simple queue implementation
type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
	
}

func (q *Queue) PopLeft() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item

}
