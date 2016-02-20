// Week 7: Thirty Four: A 4 by 4 grid can be filled with the numbers from 1 to
// 16 such that each row, column, and both of the diagonals all add up to thirty
// four. What is the total number of ways that a 4 by 4 grid can be filled in
// this way?

package main

import (
	"fmt"
)

type row struct {
	values [4]byte
	bits   int32
}

func (r row) String() string {
	return fmt.Sprintf("%d %d %d %d (%b)", r.values[0], r.values[1], r.values[2], r.values[3], r.bits)
}

func newRow(i, j, k, l byte) row {
	var r row
	r.values = [4]byte{i, j, k, l}
	r.bits = (1 << i) | (1 << j) | (1 << k) | (1 << l)
	return r
}

type square [4]row

func (s square) String() string {
	return fmt.Sprintf("%d\t%d\t%d\t%d\n%d\t%d\t%d\t%d\n%d\t%d\t%d\t%d\n%d\t%d\t%d\t%d", s[0].values[0], s[0].values[1], s[0].values[2], s[0].values[3],
		s[1].values[0], s[1].values[1], s[1].values[2], s[1].values[3],
		s[2].values[0], s[2].values[1], s[2].values[2], s[2].values[3],
		s[3].values[0], s[3].values[1], s[3].values[2], s[3].values[3])
}

func main() {
	validRows := make([]row, 0)

	for i := byte(1); i <= 16; i++ {
		for j := byte(1); j <= 16; j++ {
			if i == j {
				continue
			}

			for k := byte(1); k <= 16; k++ {
				if k == i || k == j {
					continue
				}

				for l := byte(1); l <= 16; l++ {
					if l != i && l != j && l != k && i+j+k+l == 34 {
						validRows = append(validRows, newRow(i, j, k, l))
					}
				}
			}

		}
	}

	// fmt.Println(validRows)

	// fmt.Println(len(validRows))

	var validSquares []square

	var l1, l2, l3, l4 int64

	for _, r1 := range validRows {
		l1++
		usedBits := r1.bits
		for _, r2 := range validRows {
			l2++
			if r2.bits&usedBits != 0 {
				continue
			}

			usedBits := usedBits | r2.bits
			for _, r3 := range validRows {
				l3++
				if r3.bits&usedBits != 0 {
					continue
				}

				usedBits := usedBits | r3.bits
				for _, r4 := range validRows {
					l4++
					if r4.bits&usedBits == 0 &&
						r1.values[0]+r2.values[0]+r3.values[0]+r4.values[0] == 34 &&
						r1.values[1]+r2.values[1]+r3.values[1]+r4.values[1] == 34 &&
						r1.values[2]+r2.values[2]+r3.values[2]+r4.values[2] == 34 &&
						r1.values[3]+r2.values[3]+r3.values[3]+r4.values[3] == 34 &&
						r1.values[0]+r2.values[1]+r3.values[2]+r4.values[3] == 34 &&
						r1.values[3]+r2.values[2]+r3.values[1]+r4.values[0] == 34 {
						validSquares = append(validSquares, square{r1, r2, r3, r4})
						fmt.Printf("%v\n\n", square{r1, r2, r3, r4})
					}
				}
			}
		}
	}

	fmt.Println("l1:", l1)
	fmt.Println("l2:", l2)
	fmt.Println("l3:", l3)
	fmt.Println("l4:", l4)
	fmt.Println("Magic squares:", len(validSquares))
}

// l1: 2064
// l2: 4260096
// l3: 2556057600
// l4: 268435980288
// Magic squares: 7040

// real	15m29.509s
// user	15m29.850s
// sys	0m0.370s
