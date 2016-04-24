package solver_test

import (
	"testing"

	"github.com/jackc/snake_case/12_8_queens/solver"
)

func BenchmarkSolve8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solver := solver.New(8, 8, 8)
		if len(solver.Solutions()) != 92 {
			b.Fatalf("Expected 92 solutions, got: %d", len(solver.Solutions()))
		}
	}
}

func BenchmarkSolve10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solver := solver.New(10, 10, 10)
		if len(solver.Solutions()) != 724 {
			b.Fatalf("Expected 724 solutions, got: %d", len(solver.Solutions()))
		}
	}
}
