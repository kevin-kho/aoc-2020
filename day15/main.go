package main

import (
	"fmt"

	"github.com/kevin-kho/aoc-utilities/common"
)

func SolvePartOne(arr []int, turnCount int) {

	// key: int
	// values: the last two occurances; treat it like a queue
	lastSeen := make(map[int][]int)

	// initialize values
	for i, val := range arr {
		lastSeen[val] = append(lastSeen[val], i+1)
		fmt.Printf("turn: %v; speak %v\n", i+1, val)
	}

	turn := len(arr) + 1
	prev := arr[len(arr)-1]
	for turn < turnCount+1 {

		// case: prev not seen, speak 0
		if len(lastSeen[prev]) == 1 {
			if len(lastSeen[0]) == 2 {
				lastSeen[0] = lastSeen[0][1:]
			}
			lastSeen[0] = append(lastSeen[0], turn)
			prev = 0
		} else {
			num := common.IntAbs(lastSeen[prev][0] - lastSeen[prev][1])
			if len(lastSeen[num]) == 2 {
				lastSeen[num] = lastSeen[num][1:]
			}
			lastSeen[num] = append(lastSeen[num], turn)
			prev = num
		}

		fmt.Printf("turn: %v; speak %v\n", turn, prev)

		turn++
	}

}

func main() {
	fmt.Println("hi")
	// SolvePartOne([]int{0, 3, 6})
	SolvePartOne([]int{6, 4, 12, 1, 20, 0, 16}, 2000)
	SolvePartOne([]int{6, 4, 12, 1, 20, 0, 16}, 30000000)
}
