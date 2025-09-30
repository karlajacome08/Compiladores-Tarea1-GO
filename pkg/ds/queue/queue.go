package queue

import "fmt"

type Queue struct {
	items []string
}

func New() *Queue {
	return &Queue{items: make([]string, 0)}
}

func (q *Queue) Enqueue(i string) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() (string, error) {
	if len(q.items) == 0 {
		return "", fmt.Errorf("queue is empty")
	}
	v := q.items[0]
	q.items[0] = ""
	q.items = q.items[1:]
	return v, nil
}

func (q *Queue) Peek() (string, error) {
	if len(q.items) == 0 {
		return "", fmt.Errorf("queue is empty")
	}
	return q.items[0], nil
}
func (q *Queue) Len() int      { return len(q.items) }
func (q *Queue) IsEmpty() bool { return len(q.items) == 0 }
