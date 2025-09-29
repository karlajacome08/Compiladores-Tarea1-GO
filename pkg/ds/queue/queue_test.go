package queue

import "testing"

func TestQueue_TableDriven(t *testing.T) {
	type step struct {
		op    string // "enq", "deq", "peek"
		value string // para enq
	}

	tests := []struct {
		name      string
		steps     []step
		wantFront string // valor al frente al FINAL (si no está vacía)
		wantLen   int
	}{
		{
			name: "solo_enqueue",
			steps: []step{
				{"enq", "a"},
				{"enq", "b"},
			},
			wantFront: "a", // <- CORRECTO: el frente es "a"
			wantLen:   2,
		},
		{
			name: "enq_deq_intercalado",
			steps: []step{
				{"enq", "a"},
				{"enq", "b"},
				{"deq", ""},  // sale "a"; queda ["b"]
				{"enq", "c"}, // ["b","c"]
				{"deq", ""},  // sale "b"; queda ["c"]
			},
			wantFront: "c",
			wantLen:   1,
		},
		{
			name: "peek_entre_operaciones",
			steps: []step{
				{"enq", "a"},
				{"peek", ""}, // observa "a"
				{"enq", "b"},
				{"deq", ""}, // sale "a"; queda ["b"]
			},
			wantFront: "b",
			wantLen:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var q Queue
			for _, st := range tt.steps {
				switch st.op {
				case "enq":
					q.Enqueue(st.value)
				case "deq":
					if _, err := q.Dequeue(); err != nil {
						t.Fatalf("Dequeue inesperado: %v", err)
					}
				case "peek":
					if _, err := q.Peek(); err != nil {
						t.Fatalf("Peek inesperado: %v", err)
					}
				default:
					t.Fatalf("op desconocida: %s", st.op)
				}
			}

			if q.Len() != tt.wantLen {
				t.Fatalf("Len esperado=%d, got=%d", tt.wantLen, q.Len())
			}
			if tt.wantLen > 0 {
				got, err := q.Peek()
				if err != nil {
					t.Fatalf("Peek final error: %v", err)
				}
				if got != tt.wantFront {
					t.Fatalf("frente esperado=%q, got=%q", tt.wantFront, got)
				}
			}
		})
	}
}
