package solver

const MaxQueens = 16
const MaxBoardWidth = 16
const MaxBoardHeight = 16

type Queen struct {
	X int8
	Y int8
}

type Solver struct {
	solChan     chan []Queen
	doneChan    chan struct{}
	workerCount int
	boardWidth  int8
	boardHeight int8
	queenCount  int8
}

type boardState struct {
	queens       [MaxQueens]Queen
	xUsed        [MaxBoardWidth]bool
	yUsed        [MaxBoardHeight]bool
	dPosUsed     [MaxBoardHeight + MaxBoardWidth]bool
	dNegUsed     [MaxBoardHeight + MaxBoardWidth]bool
	queensPlaced int
}

func (bs *boardState) validQueen(solver *Solver, queen Queen) bool {
	x := queen.X
	y := queen.Y

	dPos := y + x
	dNeg := solver.boardWidth + y - x
	return !(bs.xUsed[x] || bs.yUsed[y] || bs.dPosUsed[dPos] || bs.dNegUsed[dNeg])
}

func (bs *boardState) addQueen(solver *Solver, queen Queen) {
	x := queen.X
	y := queen.Y

	dPos := y + x
	dNeg := solver.boardWidth + y - x

	bs.queens[bs.queensPlaced] = queen
	bs.queensPlaced++
	bs.xUsed[x] = true
	bs.yUsed[y] = true
	bs.dPosUsed[dPos] = true
	bs.dNegUsed[dNeg] = true
}

func (bs *boardState) deepCopy() *boardState {
	newBS := *bs
	return &newBS
}

func New(boardWidth, boardHeight, queenCount int8) *Solver {
	solver := &Solver{
		solChan:     make(chan []Queen),
		doneChan:    make(chan struct{}),
		boardWidth:  boardWidth,
		boardHeight: boardHeight,
		queenCount:  queenCount,
	}

	// Since n queens use at least n rows, don't try searches that leave too may
	// rows empty such that a solution is impossible.
	for y := int8(0); y < (solver.boardHeight - solver.queenCount + 1); y++ {
		for x := int8(0); x < solver.boardWidth; x++ {
			bs := &boardState{}

			bs.addQueen(solver, Queen{X: x, Y: y})

			solver.workerCount++
			go func(bs *boardState) {
				solver.solve(bs)
				solver.doneChan <- struct{}{}
			}(bs)
		}
	}

	go solver.doneWatcher()

	return solver
}

func (solver *Solver) doneWatcher() {
	for i := 0; i < solver.workerCount; i++ {
		<-solver.doneChan
	}

	close(solver.solChan)
}

func (solver *Solver) solve(bs *boardState) {
	lastQueen := bs.queens[bs.queensPlaced-1]
	startY := lastQueen.Y
	startX := lastQueen.X + 1
	for y := startY; y < solver.boardHeight; y++ {
		for x := startX; x < solver.boardWidth; x++ {
			queen := Queen{X: x, Y: y}
			if !bs.validQueen(solver, queen) {
				continue
			}

			bs := bs.deepCopy()
			bs.addQueen(solver, Queen{X: x, Y: y})

			if bs.queensPlaced == int(solver.queenCount) {
				solver.solChan <- bs.queens[:bs.queensPlaced]
				continue
			}

			solver.solve(bs)
		}
		startX = 0
	}
}

func (solver *Solver) SolChan() <-chan []Queen {
	return solver.solChan
}
