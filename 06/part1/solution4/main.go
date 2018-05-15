package main

import (
	"fmt"
	"strings"
	"strconv"
	"time"
	"os"
	"sync"
)

const puzzleInput = "11	11	13	7	0	15	5	5	4	4	1	1	7	1	15	11"

var states = make([][]int, 1) // bank states

var mutex sync.Mutex

func init() { // initialize with puzzle input (state 0)
	states[0] = func() []int {
		banks := make([]int, 0)
		for _, block := range strings.Split(puzzleInput, "\t") {
			num, _ := strconv.Atoi(block)
			banks = append(banks, num)
		}
		return banks
	}()
}

func main() {
	start := time.Now()
	newStateIndex := make(chan int)

	go balanceBanks(newStateIndex)

	for {
		if yes, i := identicalToAnOldState(<-newStateIndex); yes {
			stop := time.Now()
			fmt.Printf("%v cycles\n", i)
			fmt.Printf("%v elapsed\n", stop.Sub(start))
			os.Exit(0)
		}
	}
}

func balanceBanks(conduit chan int) {
	for {
		mutex.Lock()
		banks := func() []int { // make a copy of current state
			c := len(states)-1
			temp := make([]int, len(states[c]))
			copy(temp, states[c])
			return temp
		}()
		mutex.Unlock()

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

		mutex.Lock()
		states = append(states, banks) // add new state to states
		last := len(states)-1
		mutex.Unlock()

		conduit <-last //send index of new state
	}
}

func identicalToAnOldState(newStateIndex int) (bool, int) {
	next:
	for i := 0; i < newStateIndex; i++ {
		mutex.Lock()
		len := len(states[i])
		mutex.Unlock()
		for b := 0; b < len; b++  {
			mutex.Lock()
			misMatch := states[newStateIndex][b] != states[i][b]
			mutex.Unlock()
			if misMatch {
				continue next
			}
		}
		return true, newStateIndex
	}
	return false, -1
}

