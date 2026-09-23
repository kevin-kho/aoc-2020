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

}
