package main

import (
	"fmt"

	"github.com/cznic/mathutil"
)

func main() {
	var paths int
	for i := uint32(0); i < uint32(1<<20); i++ {
		bitCount := mathutil.PopCountUint32(i)
		if bitCount == 10 {
			paths++
		}
	}

	fmt.Println(paths)
}
