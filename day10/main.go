package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

func GetAdapters(data []byte) ([]int, error) {

	var res []int
	for entry := range bytes.Lines(data) {
		entry = bytes.TrimSpace(entry)

		i, err := strconv.Atoi(string(entry))
		if err != nil {
			return res, err
		}
		res = append(res, i)
	}

	return res, nil

}

func SolvePartOne(adapters []int) int {

	diffs := make(map[int]int)
	curr := 0

	slices.Sort(adapters)
	for _, jolt := range adapters {
		d := jolt - curr
		diffs[d]++
		curr = jolt
	}

	diffs[3]++

	return diffs[1] * diffs[3]

}

// Basically climbing stairs
func SolvePartTwo(adapters []int) int {
	slices.Sort(adapters)

	adapters = slices.Concat([]int{0}, adapters)
	dp := make([]int, len(adapters))
	dp[0] = 1

	for i := 1; i < len(adapters); i++ {
		currVal := adapters[i]
		i1 := i - 1
		i2 := i - 2
		i3 := i - 3

		var val int

		if i1 >= 0 && currVal-adapters[i1] <= 3 {
			// val += 1
			val += dp[i1]
		}

		if i2 >= 0 && currVal-adapters[i2] <= 3 {
			// val += 1
			val += dp[i2]
		}

		if i3 >= 0 && currVal-adapters[i3] <= 3 {
			// val += 1
			val += dp[i3]
		}

		dp[i] = val

	}

	return dp[len(dp)-1]

}

func main() {

	// data, err := common.ReadInput("inputExample.txt")
	// data, err := common.ReadInput("inputExample2.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	adapters, err := GetAdapters(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(adapters)
	fmt.Println(res)

	res = SolvePartTwo(adapters)
	fmt.Println(res)

}
