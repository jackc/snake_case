package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/jackc/snake_case/12_8_queens/solver"
)

var options struct {
	boardWidth  int
	boardHeight int
	queenCount  int
	draw        bool
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage:  %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.IntVar(&options.boardWidth, "width", 8, "board width")
	flag.IntVar(&options.boardHeight, "height", 8, "board height")
	flag.IntVar(&options.queenCount, "queens", 8, "number of queens to place")
	flag.BoolVar(&options.draw, "draw", true, "draw board solutions")
	flag.Parse()

	if options.boardWidth < 1 || solver.MaxBoardWidth < options.boardWidth {
		fmt.Fprintf(os.Stderr, "width must be between 1 and %d", solver.MaxBoardWidth)
		os.Exit(1)
	}
	if options.boardHeight < 1 || solver.MaxBoardHeight < options.boardHeight {
		fmt.Fprintf(os.Stderr, "height must be between 1 and %d", solver.MaxBoardHeight)
		os.Exit(1)
	}
	if options.queenCount < 1 || solver.MaxQueens < options.queenCount {
		fmt.Fprintf(os.Stderr, "queens must be between 1 and %d", solver.MaxQueens)
		os.Exit(1)
	}

	solver := solver.New(int8(options.boardWidth), int8(options.boardHeight), int8(options.queenCount))

	solCount := 0
	rb := newRasterizedBoard(int8(options.boardWidth), int8(options.boardHeight))

	for queens := range solver.SolChan() {
		solCount++

		if options.draw {
			rb.clear()
			for _, q := range queens {
				rb.set(q.X, q.Y)
			}

			fmt.Println(rasterizedBoardToString(rb, 'Q', '-'))
		}
	}

	fmt.Println("Solutions:", solCount)
}

type rasterizedBoard struct {
	squares []bool
	width   int8
	height  int8
}

func newRasterizedBoard(width, height int8) *rasterizedBoard {
	return &rasterizedBoard{
		width:   width,
		height:  height,
		squares: make([]bool, int(width)*int(height)),
	}
}

func (rb *rasterizedBoard) clear() {
	for i := range rb.squares {
		rb.squares[i] = false
	}
}

func (rb *rasterizedBoard) set(x, y int8) {
	rb.squares[rb.coordToIdx(x, y)] = true
}

func (rb *rasterizedBoard) get(x, y int8) bool {
	return rb.squares[rb.coordToIdx(x, y)]
}

func (rb *rasterizedBoard) coordToIdx(x, y int8) int {
	return int(y)*int(rb.width) + int(x)
}

func rasterizedBoardToString(rb *rasterizedBoard, queen, empty rune) string {
	var buf bytes.Buffer

	for y := int8(0); y < rb.width; y++ {
		for x := int8(0); x < rb.height; x++ {
			if rb.get(int8(x), int8(y)) {
				buf.WriteRune(queen)
			} else {
				buf.WriteRune(empty)
			}
		}
		buf.WriteRune('\n')
	}

	return buf.String()
}
