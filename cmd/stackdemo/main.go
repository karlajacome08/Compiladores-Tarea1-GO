package main

import (
	"fmt"

	"compiladores-tarea1-go/pkg/ds/stack"
)

func main() {
	s := stack.New()
	s.Push("Dog")
	s.Push("Cat")
	s.Push("Fish")

	fmt.Println("len:", s.Len())
	if x, err := s.Pop(); err == nil {
		fmt.Println("pop:", x)
	}
	if top, err := s.Peek(); err == nil {
		fmt.Println("peek:", top)
	}
	fmt.Println("empty?:", s.IsEmpty())
}
