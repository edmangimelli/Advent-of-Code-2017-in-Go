package main

import (
	"fmt"
	"math"
	"os"
)

/*

using their example:

17  16  15  14  13
18   5   4   3  12
19   6   1   2  11
20   7   8   9  10
21  22  23---> ...

starting from 1, you move like so:
r u l l d d r r r u u u l l l l d d d d

here's the pattern:
r 1
u 1
l 2
d 2
r 3
u 3
l 4
d 4
...and so on

my code below simply steps through a graph
(assuming 1's spot is [0,0])
moving in the spiral pattern, all the while
counting (actually, decrementing) and testing
to see if we reached our number.
If we have, we simply add the x and y coordinates
(absolute value) and we have our # of steps.

*/

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

