package solver_test

import (
	"testing"

	"github.com/jackc/snake_case/12_8_queens/solver"
)

func BenchmarkSolve8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solver := solver.New(8, 8, 8)
		for range solver.SolChan() {
		}
	}
}

func BenchmarkSolve10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solver := solver.New(10, 10, 10)
		for range solver.SolChan() {
		}
	}
}
