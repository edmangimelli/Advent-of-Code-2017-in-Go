package main

import (
	"fmt"
	"os"
	"time"
)


var puzzleInput = 325489

var memory = []int{1, 1, 2, 4, 5, 10, 11, 23, 25, 26, 54, 57}
var c = 12
var previousSumList = []int{1, 2}
var sumList []int
var patternWidth = 3

var start, stop time.Time

func main() {
	start = time.Now()

	for {
		for i := 0; i < 2; i++ {
			sumList = []int{previousSumList[1]}
			test(sumList)

			for i := 0; i < patternWidth; i++ {
				switch i {
				case 0:
					sumList = []int{previousSumList[0], previousSumList[0]+1, c-2}
					test(sumList)
				case 1:
					sumList = previousSumList
					sumList[2] = sumList[1]+1
					test(sumList)
				default:
					sumList = previousSumList
					sumList[0]++
					sumList[1]++
					sumList[2]++
					test(sumList)
				case patternWidth-1:
					sumList = []int{previousSumList[1], previousSumList[2]}
					test(sumList)
				}
			}
		}
		patternWidth++
	}
}


func test(list []int) {
	sum := memory[c-1]

	for _, v := range list {
		sum += memory[v]
	}

	memory = append(memory, sum)

	if memory[c] > puzzleInput {
		stop = time.Now()
		fmt.Printf("The value at memory location [%v] is %v\n", c, memory[c])
		fmt.Printf("Which is the first value larger than %v (the puzzle input)\n", puzzleInput)
		fmt.Printf("(%v elapsed)\n", stop.Sub(start))
		os.Exit(0)
	}

	previousSumList = sumList
	c++
}
