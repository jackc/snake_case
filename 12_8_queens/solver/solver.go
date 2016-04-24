package solver

const MaxQueens = 16
const MaxBoardWidth = 16
const MaxBoardHeight = 16

type bit32 uint32

func (bits bit32) get(n uint8) bool {
	return ((bits >> n) & 1) == 1
}

func (bits bit32) set(n uint8) bit32 {
	return bits | (1 << n)
}

func (bits bit32) unset(n uint8) bit32 {
	return bits &^ (1 << n)
}

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
	xUsed        bit32
	yUsed        bit32
	dPosUsed     bit32
	dNegUsed     bit32
	queensPlaced int
}

func (bs *boardState) validQueen(solver *Solver, queen Queen) bool {
	x := queen.X
	y := queen.Y

	dPos := y + x
	dNeg := solver.boardWidth + y - x
	return !(bs.xUsed.get(uint8(x)) || bs.yUsed.get(uint8(y)) || bs.dPosUsed.get(uint8(dPos)) || bs.dNegUsed.get(uint8(dNeg)))
}

func (bs *boardState) pushQueen(solver *Solver, queen Queen) {
	x := queen.X
	y := queen.Y

	dPos := y + x
	dNeg := solver.boardWidth + y - x

	bs.queens[bs.queensPlaced] = queen
	bs.queensPlaced++
	bs.xUsed = bs.xUsed.set(uint8(x))
	bs.yUsed = bs.yUsed.set(uint8(y))
	bs.dPosUsed = bs.dPosUsed.set(uint8(dPos))
	bs.dNegUsed = bs.dNegUsed.set(uint8(dNeg))
}

func (bs *boardState) popQueen(solver *Solver) {
	queen := bs.queens[bs.queensPlaced-1]
	bs.queensPlaced--
	x := queen.X
	y := queen.Y

	dPos := y + x
	dNeg := solver.boardWidth + y - x

	bs.xUsed = bs.xUsed.unset(uint8(x))
	bs.yUsed = bs.yUsed.unset(uint8(y))
	bs.dPosUsed = bs.dPosUsed.unset(uint8(dPos))
	bs.dNegUsed = bs.dNegUsed.unset(uint8(dNeg))
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

			bs.pushQueen(solver, Queen{X: x, Y: y})

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

			bs.pushQueen(solver, Queen{X: x, Y: y})

			if bs.queensPlaced == int(solver.queenCount) {
				sol := make([]Queen, 0, bs.queensPlaced)
				sol = append(sol, bs.queens[:bs.queensPlaced]...)
				solver.solChan <- sol
				bs.popQueen(solver)
				continue
			}

			solver.solve(bs)

			bs.popQueen(solver)
		}
		startX = 0
	}
}

func (solver *Solver) SolChan() <-chan []Queen {
	return solver.solChan
}
