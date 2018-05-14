package main

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

var stateComparisons = 0

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
			fmt.Printf("%v state comparisons performed\n", stateComparisons)
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
	len := len(states)
	lastState := states[len-1]
	next:
	for i := 0; i < len-1; i++ {
		for b := range states[i] {
			stateComparisons++
			if states[i][b] != lastState[b] {
				continue next
			}
		}
		fmt.Printf("state %v matches state %v\n%v cycles comprise the loop\n\n", i, len-1, len-1-i) // the only difference between part1/solution2/main.go and this file is the addition of this line
		return true
	}
	return false
}

