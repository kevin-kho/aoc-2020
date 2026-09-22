package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

func GetIntSlc(data []byte) ([]int, error) {
	var res []int
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		i, err := strconv.Atoi(string(entry))
		if err != nil {
			return res, err
		}
		res = append(res, i)
	}
	return res, nil
}

// Two Sum
func SolvePartOne(intSlc []int) int {
	seen := make(map[int]bool)

	for _, i := range intSlc {
		comp := 2020 - i
		if seen[comp] {
			return comp * i
		}
		seen[i] = true
	}

	return -1
}

// Three Sum
func SolvePartTwo(intSlc []int) int {
	nums := slices.Clone(intSlc)
	slices.Sort(nums)

	for i, num := range nums {

		l := i + 1
		r := len(nums) - 1
		for l < r {
			summed := nums[l] + num + nums[r]
			if summed == 2020 {
				return nums[l] * num * nums[r]
			}

			if summed > 2020 {
				r -= 1
			} else {
				l += 1
			}

		}

	}

	return -1
}

func main() {
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	intSlc, err := GetIntSlc(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(intSlc)
	fmt.Println(res)

	res2 := SolvePartTwo(intSlc)
	fmt.Println(res2)

}
