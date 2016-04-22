package solver

type Queen struct {
	X int8
	Y int8
}

type Solver struct {
	solChan     chan []Queen
	boardWidth  int8
	boardHeight int8
	queenCount  int8
}

type boardState struct {
	queens []Queen
	xUsed  []bool
	yUsed  []bool
}

func NewSolver(boardWidth, boardHeight, queenCount int8) *Solver {
	solver := &Solver{
		solChan:     make(chan []Queen),
		boardWidth:  boardWidth,
		boardHeight: boardHeight,
		queenCount:  queenCount,
	}

	for x := 0; x < solver.boardWidth; x++ {
		for y := 0; y < solver.boardHeight; y++ {

		}
	}

	return solver
}

func (solver *Solver) solve()

func (solver *Solver) SolChan() <-chan []Queen {
	return solver.solChan
}
