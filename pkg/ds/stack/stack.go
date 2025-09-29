package stack

import "fmt"

type Stack struct {
	items []string
}

func New() *Stack {
	return &Stack{items: make([]string, 0)}
}

func (s *Stack) Push(a string) {
	s.items = append(s.items, a)
}

func (s *Stack) Pop() (string, error) {
	if len(s.items) == 0 {
		return "", fmt.Errorf("stack is empty")
	}
	i := len(s.items) - 1
	v := s.items[i]
	s.items[i] = ""
	s.items = s.items[:i]
	return v, nil

}

func (s *Stack) Peek() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool { return len(s.items) == 0 }
func (s *Stack) Len() int      { return len(s.items) }
