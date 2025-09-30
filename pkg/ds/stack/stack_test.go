package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()
	//Validamos que se inicialice vacia
	if !s.IsEmpty() {
		t.Fatal("debe iniciar vacía")
	}

	s.Push("a")
	s.Push("b")
	//Validamos que se agreguen los elementos y que el tope sea el correcto
	if top, _ := s.Peek(); top != "b" {
		t.Fatalf("peek=b got %q", top)
	}

	//Validamos que se eliminen los elementos y que el tope sea el correcto
	if v, _ := s.Pop(); v != "b" {
		t.Fatalf("pop1=b got %q", v)
	}
	//Validamos que se elimine el ultimo elemento y que el stack quede vacio
	if v, _ := s.Pop(); v != "a" {
		t.Fatalf("pop2=a got %q", v)
	}
	if _, err := s.Pop(); err == nil {
		t.Fatal("esperaba error con stack vacía")
	}
}
