package main

import (
	"bytes"
	"fmt"

	"github.com/jackc/snake_case/12_8_queens/solver"
)

const BoardWidth = 8
const BoardHeight = 8
const QueenCount = 8

func main() {
	solver := solver.New(BoardWidth, BoardHeight, QueenCount)

	solCount := 0
	rb := newRasterizedBoard(BoardWidth, BoardHeight)

	for queens := range solver.SolChan() {
		solCount++

		rb.clear()
		for _, q := range queens {
			rb.set(q.X, q.Y)
		}

		fmt.Println(rasterizedBoardToString(rb, 'Q', '-'))
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
		squares: make([]bool, width*height),
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

	for y := 0; y < BoardHeight; y++ {
		for x := 0; x < BoardWidth; x++ {
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
