package hash

import "testing"

func TestHash_Init(t *testing.T) {
	// Validamos que se inicializa correctamente
	h := Init()
	if h == nil {
		t.Fatal("Init() devolvió nil")
	}
	if len(h.array) != ArraySize {
		t.Fatalf("len(array) = %d; want %d", len(h.array), ArraySize)
	}
	// Validamos que todos los buckets están inicializados (no nil)
	for i, b := range h.array {
		if b == nil {
			t.Fatalf("bucket %d es nil; Init debe inicializar todos", i)
		}
	}
}

// Insert y Search
func TestHash_InsertSearch(t *testing.T) {
	h := Init()
	// Insertamos dos claves y validamos que Search las encuentra.
	h.Insert("uno")
	h.Insert("dos")

	if !h.Search("uno") {
		t.Fatal("Search('uno') = false; want true")
	}
	if !h.Search("dos") {
		t.Fatal("Search('dos') = false; want true")
	}
	// Buscamos una clave que no existe
	if h.Search("x") {
		t.Fatal("Search('x') = true; want false (no insertado)")
	}
}

// Valida que se elimina correctamente una clave
func TestHash_Delete(t *testing.T) {
	h := Init()

	// Insertamos y luego borramos una clave
	h.Insert("borrar")
	if !h.Search("borrar") {
		t.Fatal("debe existir antes de borrar")
	}

	h.Delete("borrar")
	if h.Search("borrar") {
		t.Fatal("tras Delete('borrar'), aún se encuentra la clave")
	}
}

// Valida que se manejan colisiones simples
func TestHash_Colision(t *testing.T) {
	h := Init()

	// Inserttamos dos claves que sabemos colisionan
	h.Insert("a")
	h.Insert("h")

	if !h.Search("a") {
		t.Fatal("debe encontrar 'a' (colisión)")
	}
	if !h.Search("h") {
		t.Fatal("debe encontrar 'h' (colisión)")
	}
}
