package queue

import "testing"

func TestQueue_EnqueueDequeue(t *testing.T) {
	var q Queue

	// Validamos que inicie vacía.
	if !q.IsEmpty() {
		t.Fatal("la cola debe iniciar vacía")
	}

	// Encolamos dos elementos: "a", luego "b".
	q.Enqueue("a")
	q.Enqueue("b")

	// El primer Dequeue debe devolver "a" y quitarlo
	v, err := q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() inesperó error: %v", err)
	}
	if v != "a" {
		t.Fatalf("Dequeue() => %q; want 'a'", v)
	}

	// Debe quedar 1 elemento
	if q.Len() != 1 {
		t.Fatalf("Len() => %d; want 1", q.Len())
	}
}

//Peek no debe quitar el elemento.
func TestQueue_Peek_NoRemueve(t *testing.T) {
	var q Queue
	q.Enqueue("x")

	// Peek mira el frente sin quitarlo.
	v, err := q.Peek()
	if err != nil {
		t.Fatalf("Peek() inesperó error: %v", err)
	}
	if v != "x" {
		t.Fatalf("Peek() => %q; want 'x'", v)
	}

	// La longitud no cambia después de Peek.
	if q.Len() != 1 {
		t.Fatalf("Len() tras Peek => %d; want 1", q.Len())
	}

	// Valida que el siguiente Dequeue devuelve el mismo valor.
	v2, err := q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() inesperó error: %v", err)
	}
	if v2 != "x" {
		t.Fatalf("Dequeue() => %q; want 'x'", v2)
	}
}

// Valida que si se hace un dequeue vacío debe devolver error
func TestQueue_Dequeue_EnVacio(t *testing.T) {
	var q Queue
	if _, err := q.Dequeue(); err == nil {
		t.Fatal("Dequeue() en vacío: se esperaba error, llegó nil")
	}
}

//Peek en vacío debe dar error
func TestQueue_Peek_EnVacio(t *testing.T) {
	var q Queue
	if _, err := q.Peek(); err == nil {
		t.Fatal("Peek() en vacío: se esperaba error, llegó nil")
	}
}
