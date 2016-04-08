# Problem

You've just arrived in sunny St. Augustine, and find yourself amazed by the visionary civic planning that would result in the area in which you now stand: a street grid exactly 10 blocks square.

You're in the northwest corner of this 10 by 10 block area, and would like to take a scenic walk to the southeast corner, while only ever moving south or east.

As you begin walking, you wonder to yourself, "how many different paths could I take from this northwest corner to the southeast corner?"

You quickly note that if the downtown area were only a 2 block by 2 block grid, there would be 6 distinct paths from one corner to the other:

So, how many distinct paths are there through the 10 by 10 downtown area?

# Compiling and Running

Ruby

```
ruby main.rb
```

Go

```
cd go
go build -o snake_case
time ./snake_case
```

C

```
cc -o snake_case main.c
chmod +x snake_case
time ./snake_case
```

C assembly output

```
cc -S -mllvm --x86-asm-syntax=intel main.c
```

C with popcnt instruction

```
cc -mpopcnt -o snake_case_popcnt main.c
chmod +x snake_case_popcnt
time ./snake_case_popcnt
```

C with popcnt instruction assembly output

```
cc -S -mllvm --x86-asm-syntax=intel -mpopcnt main.c && mv main.s main-popcnt.s
```
