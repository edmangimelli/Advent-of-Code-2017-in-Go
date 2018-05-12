package main

/*
  Whoa! I don't think I've ever made my computer work so hard!
  It took so long to get the answer, I thought for sure my code
  was broken. So, I started printing the state of the current
  bank and capped the loop to 1,000,000 cycles, just to see
  what it was doing. To my surprise, it found the answer!
  The answer being over 4000 cycles makes the lengthiness of
  the computation time make sense. When you get past 4000
  cycles, you're getting into 8,000,000 (4001*(4000/2)) state
  comparisons per cycle. 
  Print the current bank state at each cycle for a neat visual
  effect!
*/

import (
	"fmt"
	"strings"
	"strconv"
	"time"
)

const puzzleInput = "11	11	13	7	0	15	5	5	4	4	1	1	7	1	15	11"
var banks = make([]int, 0)

func init() {
	for _, block := range strings.Split(puzzleInput, "\t") {
		num, _ := strconv.Atoi(block)
		banks = append(banks, num)
	}
}

func main() {
	start := time.Now()
	cycles := 0
	states := make([][]int, 0)

	for {
		temp := make([]int, len(banks))
		copy(temp, banks)
		states = append(states, temp)
		if matchingStates(states) {
			stop := time.Now()
			fmt.Printf("%v cycles\n", cycles)
			fmt.Printf("%v elapsed\n", stop.Sub(start))
			break
		}

		chosenBank := len(banks)-1
		for i := len(banks)-2; i >= 0; i-- {
			if banks[i] >= banks[chosenBank] { chosenBank = i }
		}

		store := banks[chosenBank]
		banks[chosenBank] = 0

		c := chosenBank
		for store > 0 {
			c++
			if c >= len(banks) { c = 0 }
			banks[c]++
			store--
		}

		cycles++
	}
}

func matchingStates(states [][]int) bool {
	for i := 0; i < len(states)-1; i++ {
		next:
		for j := i+1; j < len(states); j++ {
			for b := range states[i] {
				if states[i][b] != states[j][b] {
					continue next
				}
			}
			return true
		}
	}
	return false
}

/*
cycles = 0
loop
	save state
	compare to other states
		if one states equals another, done
	find largest number, ties go to the earliest location (simply go backwards)
	empty into register
	distribute until empty
*/
