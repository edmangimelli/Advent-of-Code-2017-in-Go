package main

import (
	"fmt"
	"math"
	"os"
)

var data = 325489

var x, y = 0, 0

func main() {
	steps := 1
	for {
		right(steps)
		up(steps)
		steps++
		left(steps)
		down(steps)
		steps++
	}
}

func right(steps int) { move(steps, &x, +1) }
func  left(steps int) { move(steps, &x, -1) }
func    up(steps int) { move(steps, &y, +1) }
func  down(steps int) { move(steps, &y, -1) }

func move(steps int, axis *int, inc int) {
	for i := 0; i < steps; i++ {
		if data == 1 {
			fmt.Printf("x = %v, y = %v\n", x, y)
			fmt.Println("therefore:")
			fmt.Printf("%v steps\n", math.Abs(float64(x)) + math.Abs(float64(y)))
			os.Exit(0)
		}
		*axis += inc
		data--
	}
}

