package hash

import "testing"

func TestInit_CreaBuckets(t *testing.T) {
	h := Init()
	if h == nil {
		t.Fatal("Init() devolvió nil")
	}
	if len(h.array) != ArraySize {
		t.Fatalf("ArraySize=%d, got len=%d", ArraySize, len(h.array))
	}
	for i, b := range h.array {
		if b == nil {
			t.Fatalf("bucket %d es nil; Init debe inicializar todos", i)
		}
	}
}

func TestInsertYSearch_Basico(t *testing.T) {
	h := Init()
	h.Insert("uno")
	h.Insert("dos")
	h.Insert("tres")

	if !h.Search("uno") {
		t.Fatal("debe encontrar 'uno'")
	}
	if !h.Search("dos") {
		t.Fatal("debe encontrar 'dos'")
	}
	if !h.Search("tres") {
		t.Fatal("debe encontrar 'tres'")
	}
	if h.Search("x") {
		t.Fatal("no debe encontrar 'x'")
	}
}

func TestColisiones_SeparateChaining(t *testing.T) {
	// con ArraySize=7, letras separadas por +7 colisionan: 'a'(97), 'h'(104), 'o'(111)
	h := Init()
	keys := []string{"a", "h", "o", "v"}
	for _, k := range keys {
		h.Insert(k)
	}
	for _, k := range keys {
		if !h.Search(k) {
			t.Fatalf("debe encontrar clave en colisión: %q", k)
		}
	}
}

func TestDelete_HeadMiddleTail(t *testing.T) {
	// Preparamos varias claves que caen en el mismo bucket para cubrir:
	// - borrar head
	// - borrar middle
	// - borrar tail
	h := Init()
	// mismas colisiones que antes
	keys := []string{"a", "h", "o", "v"} // Insert en orden: v -> o -> h -> a (se insertan al frente)
	for _, k := range keys {
		h.Insert(k)
	}

	// 1) borrar head actual (última insertada "v" queda a la cabeza)
	h.Delete("v")
	if h.Search("v") {
		t.Fatal("v no debe existir tras Delete(head)")
	}
	// 2) borrar middle (por ejemplo "h")
	h.Delete("h")
	if h.Search("h") {
		t.Fatal("h no debe existir tras Delete(middle)")
	}
	// 3) borrar tail (probablemente "a" o "o" según inserción)
	h.Delete("a")
	if h.Search("a") {
		t.Fatal("a no debe existir tras Delete(tail)")
	}
	// todavía debe quedar "o"
	if !h.Search("o") {
		t.Fatal("o debería seguir existiendo")
	}
}

func TestDelete_Inexistente_NoPanic(t *testing.T) {
	h := Init()

	// No debe paniquear borrar en tabla vacía
	h.Delete("x") // si tu bucket.delete no valida b.head==nil, esto podría paniquear

	// Inserta y borra inexistente en mismo bucket
	h.Insert("a")
	h.Delete("zzz") // no debe paniquear ni afectar a "a"
	if !h.Search("a") {
		t.Fatal("a no debe borrarse cuando se elimina inexistente")
	}
}

func TestInsert_Duplicate_NoDup(t *testing.T) {
	h := Init()
	h.Insert("dup")
	h.Insert("dup") // no debe duplicar

	// Si hubo duplicado, tras un único Delete, Search("dup") seguiría true.
	h.Delete("dup")
	if h.Search("dup") {
		t.Fatal("Insert debe evitar duplicados; tras un Delete no debería existir")
	}
}
