package main

import (
	"fmt"
)

const BoardWidth = 8
const BoardHeight = 8
const QueenCount = 8

type Queen struct {
	X int8
	Y int8
}

type BoardState struct {
	Queens []Queen
	XUsed  []bool
	YUsed  []bool
}

func Solve(width, height, queenCount int8) chan []Queen {
	solChan := make(chan []Queen)

	emptyState := BoardState{
		Queens: make([]Queen, 0, queenCount),
		XUsed:  make([]bool, width),
		YUsed:  make([]bool, width),
	}

	return solChan
}

func NewBoardState(width, height, queenCount int8) *BoardState {
	return &BoardState{
		Queens: make([]Queen, 0, queenCount),
		XUsed:  make([]bool, width),
		YUsed:  make([]bool, width),
	}
}

func main() {
	bs := NewBoardState(BoardWidth, BoardHeight, QueenCount)

	for x := 0; x < BoardWidth; x++ {
		for y := 0; y < BoardHeight; y++ {

		}
	}
	gameCount := gameCountByTeamCount(64)
	fmt.Println("Games:", gameCount)
	combinations := big.NewInt(1)
	big2 := big.NewInt(2)
	for i := 0; i < gameCount; i++ {
		combinations.Mul(combinations, big2)
	}

	fmt.Println("Combinations:", combinations)
}
