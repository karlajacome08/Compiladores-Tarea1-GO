package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()
	if !s.IsEmpty() {
		t.Fatal("debe iniciar vacía")
	}

	s.Push("a")
	s.Push("b")
	if top, _ := s.Peek(); top != "b" {
		t.Fatalf("peek=b got %q", top)
	}

	if v, _ := s.Pop(); v != "b" {
		t.Fatalf("pop1=b got %q", v)
	}
	if v, _ := s.Pop(); v != "a" {
		t.Fatalf("pop2=a got %q", v)
	}
	if _, err := s.Pop(); err == nil {
		t.Fatal("esperaba error con stack vacía")
	}
}
