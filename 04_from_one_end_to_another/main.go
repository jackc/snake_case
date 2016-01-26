// Week 04*​: From One End to the Other: Find the smallest possible (positive)
// integer that ends in a six such that if that six is removed and placed in
// front of the remaining digits of the number, the resulting number will be
// four times as large as the original.

package main

import (
	"fmt"
	"os"
)

func main() {
	mid := 0
	mult := 10
	for {
		for d := 0; d < mult; d++ {
			original := mid*10 + 6
			rotated := 6*mult + mid

			if rotated%original == 0 && rotated/original == 4 {
				fmt.Println(original, rotated)
				os.Exit(0)
			}
			mid++
		}
		mult *= 10
	}
}
