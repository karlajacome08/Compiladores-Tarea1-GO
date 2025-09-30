package main

import (
	"fmt"

	"compiladores-tarea1-go/pkg/ds/queue"
)

func main() {
	q := queue.New()
	q.Enqueue("A")
	q.Enqueue("B")
	q.Enqueue("C")

	if x, err := q.Dequeue(); err == nil {
		fmt.Println("dequeue:", x)
	}
	if front, err := q.Peek(); err == nil {
		fmt.Println("peek:", front)
	}
	fmt.Println("len:", q.Len(), "empty?:", q.IsEmpty())
}
