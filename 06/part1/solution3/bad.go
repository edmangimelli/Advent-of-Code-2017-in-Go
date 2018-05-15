package main

import (
	"fmt"
	"strings"
	"strconv"
	"time"
	"os"
)

const puzzleInput = "11	11	13	7	0	15	5	5	4	4	1	1	7	1	15	11"

func main() {
	start := time.Now()

	states := make([][]int, 1) // bank states
	states[0] = func() []int { // initialize with puzzle input (state 0)
		banks := make([]int, 0)
		for _, block := range strings.Split(puzzleInput, "\t") {
			num, _ := strconv.Atoi(block)
			banks = append(banks, num)
		}
		return banks
	}()

	newStateIndex := make(chan int)

	go balanceBanks(newStateIndex, states)

	for {
		if yes, i := identicalToAnOldState(<-newStateIndex, states); yes {
			stop := time.Now()
			fmt.Printf("%v cycles\n", i)
			fmt.Printf("%v elapsed\n", stop.Sub(start))
			os.Exit(0)
		}
	}
}

func balanceBanks(conduit chan int, states [][]int) {
	for {
		banks := func() []int { // make a copy of current state
			c := len(states)-1
			temp := make([]int, len(states[c]))
			copy(temp, states[c])
			return temp
		}()

		bankWithMostBlocks := 0 // find bank with most blocks
		for i, len := 1, len(banks); i < len; i++ {
			if banks[i] > banks[bankWithMostBlocks] {
				bankWithMostBlocks = i
			}
		}

		store := banks[bankWithMostBlocks] // blocks to redistribute
		banks[bankWithMostBlocks] = 0

		b := bankWithMostBlocks // redistribute
		for store > 0 {
			b++
			if b >= len(banks) { b = 0 }
			banks[b]++
			store--
		}

		states = append(states, banks) // add new state to states

		conduit <-(len(states)-1) //send index of new state
	}
}

func identicalToAnOldState(newStateIndex int, states [][]int) (bool, int) {
	next:
	for i := 0; i < newStateIndex; i++ {
		for b := range states[i] {
			if states[newStateIndex][b] != states[i][b] {
				continue next
			}
		}
		return true, newStateIndex
	}
	return false, -1
}

