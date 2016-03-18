package main

import (
	"fmt"
	"math/big"
)

func main() {
	gameCount := gameCountByTeamCount(64)
	fmt.Println("Games:", gameCount)
	combinations := big.NewInt(1)
	big2 := big.NewInt(2)
	for i := 0; i < gameCount; i++ {
		combinations.Mul(combinations, big2)
	}

	fmt.Println("Combinations:", combinations)
}

func gameCountByTeamCount(teamCount int) int {
	return teamCount - 1
}
