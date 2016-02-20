// Week 7: Thirty Four: A 4 by 4 grid can be filled with the numbers from 1 to
// 16 such that each row, column, and both of the diagonals all add up to thirty
// four. What is the total number of ways that a 4 by 4 grid can be filled in
// this way?

package main

import (
	"fmt"
	"sort"
)

type row struct {
	values [4]byte
	bits   int32
}

func (r row) String() string {
	return fmt.Sprintf("%d %d %d %d (%b)", r.values[0], r.values[1], r.values[2], r.values[3], r.bits)
}

func newRow(x, y, z, w byte) row {
	var r row
	r.values = [4]byte{x, y, z, w}
	r.bits = (1 << x) | (1 << y) | (1 << z) | (1 << w)
	return r
}

type square [4]row

func (s square) String() string {
	return fmt.Sprintf("%d\t%d\t%d\t%d\n%d\t%d\t%d\t%d\n%d\t%d\t%d\t%d\n%d\t%d\t%d\t%d", s[0].values[0], s[0].values[1], s[0].values[2], s[0].values[3],
		s[1].values[0], s[1].values[1], s[1].values[2], s[1].values[3],
		s[2].values[0], s[2].values[1], s[2].values[2], s[2].values[3],
		s[3].values[0], s[3].values[1], s[3].values[2], s[3].values[3])
}

const allBitsUsed = 131070 // Bit pattern 11111111111111110

type orderedValidRows []row

func (rows orderedValidRows) filterPrefix1(x byte) orderedValidRows {
	startIdx := sort.Search(len(rows), func(i int) bool {
		values := rows[i].values
		return x <= values[0]
	})

	rows = rows[startIdx:]
	for i, r := range rows {
		if r.values[0] != x {
			return rows[:i]
		}
	}

	return rows
}

func (rows orderedValidRows) filterPrefix2(x, y byte) orderedValidRows {
	startIdx := sort.Search(len(rows), func(i int) bool {
		values := rows[i].values
		return x < values[0] || (x == values[0] && y <= values[1])
	})

	rows = rows[startIdx:]
	for i, r := range rows {
		if r.values[0] != x || r.values[1] != y {
			return rows[:i]
		}
	}

	return rows
}

func computeValidRows() orderedValidRows {
	validRows := make(orderedValidRows, 0)

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

	return validRows
}

func main() {
	validRows := computeValidRows()

	// for _, r := range validRows {
	//  fmt.Println(r)
	// }
	// return

	// fmt.Println(len(validRows))

	var validSquares []square

	var l0, l1, l2 int64

	for _, r0 := range validRows {
		l0++
		usedBits := r0.bits

		for _, r1 := range validRows {
			l1++

			if r1.bits&usedBits != 0 {
				continue
			}

			usedBits := usedBits | r1.bits

			c0ValidRows := validRows.filterPrefix2(r0.values[0], r1.values[0])
			c0r2Min := c0ValidRows[0].values[2]
			c0r2Max := c0ValidRows[len(c0ValidRows)-1].values[2]

			for i := c0r2Min; i <= c0r2Max; i++ {
				for _, r2 := range validRows.filterPrefix1(i) {
					l2++
					if r2.bits&usedBits != 0 {
						continue
					}

					usedBits := usedBits | r2.bits

					// Compute final row, and ensure that all values are between 1 and 16
					x := 34 - r0.values[0] - r1.values[0] - r2.values[0]
					if x < 1 || 16 < x {
						continue
					}
					y := 34 - r0.values[1] - r1.values[1] - r2.values[1]
					if y < 1 || 16 < y {
						continue
					}
					z := 34 - r0.values[2] - r1.values[2] - r2.values[2]
					if z < 1 || 16 < z {
						continue
					}
					w := 34 - r0.values[3] - r1.values[3] - r2.values[3]
					if w < 1 || 16 < w {
						continue
					}

					r3 := newRow(x, y, z, w)

					if r3.bits|usedBits == allBitsUsed && // check that all bits are used -- this will find collisions in last calculated row with itself as well as any of the other rows
						// rows 0, 1, and 2 are guaranteed to == 34 because we pulled them from the list
						// of known good rows
						r3.values[0]+r3.values[1]+r3.values[2]+r3.values[3] == 34 && // row 4 was computed so must check
						// All columns are correct because it was computed
						r0.values[0]+r1.values[1]+r2.values[2]+r3.values[3] == 34 && // diagonal 1
						r0.values[3]+r1.values[2]+r2.values[1]+r3.values[0] == 34 { // diagonal 2
						validSquares = append(validSquares, square{r0, r1, r2, r3})
						// fmt.Printf("%v\n\n", square{r0, r1, r2, r3})
					}
				}
			}

		}
	}

	fmt.Println("l0:", l0)
	fmt.Println("l1:", l1)
	fmt.Println("l2:", l2)
	fmt.Println("Magic squares:", len(validSquares))
}

// l0: 2064
// l1: 4260096
// l2: 1684010304
// Magic squares: 7040

// real	0m10.817s
// user	0m10.868s
// sys	0m0.043s
