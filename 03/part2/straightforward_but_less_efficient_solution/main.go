package main

import (
	"fmt"
	"os"
	"time"
)


var puzzleInput = 325489

const (
	n  = 1 << iota
	ne = 1 << iota
	e  = 1 << iota
	se = 1 << iota
	s  = 1 << iota
	sw = 1 << iota
	w  = 1 << iota
	nw = 1 << iota
	right = w|nw|n|ne
	left  = sw|s|se|e
	up    = nw|w|sw|s
	down  = n|ne|e|se
)

type coord struct {
	x, y int
}


var c = coord{0,0} // current location
var memory = map[coord]int{c: 1}

var start, stop time.Time

func main() {
	start = time.Now()

	steps := 1
	for {
		move( right, steps, &c.x, +1)
		move( up,    steps, &c.y, +1)
		steps++
		move( left,  steps, &c.x, -1)
		move( down,  steps, &c.y, -1)
		steps++
	}
}

func move(locations uint8, steps int, axis *int, inc int) {
	for i := 0; i < steps; i++ {
		*axis += inc
		sumLocations(locations)
		if memory[c] > puzzleInput {
			stop = time.Now()
			fmt.Printf("The value at location (%v, %v)\n", c.x, c.y)
			fmt.Printf("is %v\n", memory[c])
			fmt.Printf("which is larger than the puzzle input (%v)\n", puzzleInput)
			fmt.Printf("(%v elapsed)\n", stop.Sub(start))
			os.Exit(0)
		}
	}
}

func sumLocations(bits uint8) {
	if bits & n  != 0 { addValueFromCoord(coord{c.x  , c.y+1}) }
	if bits & ne != 0 { addValueFromCoord(coord{c.x+1, c.y+1}) }
	if bits & e  != 0 { addValueFromCoord(coord{c.x+1, c.y  }) }
	if bits & se != 0 { addValueFromCoord(coord{c.x+1, c.y-1}) }
	if bits & s  != 0 { addValueFromCoord(coord{c.x  , c.y-1}) }
	if bits & sw != 0 { addValueFromCoord(coord{c.x-1, c.y-1}) }
	if bits & w  != 0 { addValueFromCoord(coord{c.x-1, c.y  }) }
	if bits & nw != 0 { addValueFromCoord(coord{c.x-1, c.y+1}) }
}

func addValueFromCoord(somewhere coord) {
	if v, ok := memory[somewhere]; ok {
		memory[c] += v
	}
}

