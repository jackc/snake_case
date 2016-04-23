# Results

Intel i7 4770 - Ubuntu 14.04 - Go 1.6.2

```
commit 43c242e65d2656835ba48f44e8b88217f177646b
Author: Jack Christensen <jack@jackchristensen.com>
Date:   Sat Apr 23 10:02:48 2016 -0500

    Add benchmark to queens solver

jack@edi:~/dev/go/src/github.com/jackc/snake_case/12_8_queens/solver$ go test -bench=. -benchmem
testing: warning: no tests to run
PASS
BenchmarkSolve8-8        100    15144887 ns/op  22844238 B/op   713870 allocs/op
BenchmarkSolve10-8         2   991666435 ns/op  1929117680 B/op 45217839 allocs/op



commit 6a0945a7e635c84eecf3f0e9378c33715bc37cc0
Author: Jack Christensen <jack@jackchristensen.com>
Date:   Sat Apr 23 10:03:08 2016 -0500

    Prune search space at start

jack@edi:~/dev/go/src/github.com/jackc/snake_case/12_8_queens/solver$ go test -bench=. -benchmem
testing: warning: no tests to run
PASS
BenchmarkSolve8-8        100    10168858 ns/op  14624461 B/op   457010 allocs/op
BenchmarkSolve10-8         2   660099829 ns/op  1282691200 B/op 30065465 allocs/op



commit 7f679d55b7480f597833409497b19a700f23e148
Author: Jack Christensen <jack@jackchristensen.com>
Date:   Sat Apr 23 10:26:53 2016 -0500

    Reduce memory allocations by inlining board state

    Note that this also limits board size to 16x16 with 16 queens.

jack@edi:~/dev/go/src/github.com/jackc/snake_case/12_8_queens/solver$ go test -bench=. -benchmem
testing: warning: no tests to run
PASS
BenchmarkSolve8-8        300     4329858 ns/op  10968031 B/op    76185 allocs/op
BenchmarkSolve10-8         5   225693069 ns/op  721509200 B/op   5011566 allocs/op
ok    github.com/jackc/snake_case/12_8_queens/solver  3.090s



commit ba7265ce4d1a00b841eb2fda4bc7a8bdf230a24e
Author: Jack Christensen <jack@jackchristensen.com>
Date:   Sat Apr 23 10:47:42 2016 -0500

    Use bit operations for solving

jack@edi:~/dev/go/src/github.com/jackc/snake_case/12_8_queens/solver$ go test -bench=. -benchmem
testing: warning: no tests to run
PASS
BenchmarkSolve8-8        500     3619480 ns/op   4874855 B/op    76174 allocs/op
BenchmarkSolve10-8        10   180722660 ns/op  320672425 B/op   5010949 allocs/op
ok    github.com/jackc/snake_case/12_8_queens/solver  4.166s
```
