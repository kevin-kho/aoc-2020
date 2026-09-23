package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

func GetIntArr(data []byte) ([]int, error) {
	var res []int

	for entry := range bytes.Lines(data) {
		entry = bytes.TrimSpace(entry)
		val, err := strconv.Atoi(string(entry))
		if err != nil {
			return res, err
		}
		res = append(res, val)

	}

	return res, nil
}

func SolvePartOne(intArr []int, preamble int) int {
	// pointers for last X digits
	l := 0
	r := preamble - 1

	// pointer for current digit
	i := preamble

	for i < len(intArr) {

		// even := intArr[i]%2 == 0
		found := false

		for x := l; x < r+1; x++ {
			for y := x + 1; y < r+1; y++ {
				a := intArr[x]
				b := intArr[y]
				if a+b == intArr[i] {
					found = true
				}
			}
		}

		if !found {
			return intArr[i]
		}

		l++
		r++
		i++
	}

	return -1

}

func SolvePartTwo(intArr []int, target int) int {

	// sliding window
	l := 0
	var end int
	curr := 0

	for r := range len(intArr) {
		curr += intArr[r]

		if curr >= target {

			// slide window for as long as curr > target
			for curr > target {
				curr -= intArr[l]
				l++
			}

			// found item
			if curr == target && r-l+1 >= 2 {
				end = r
				break
			}
		}
	}

	mn := intArr[l]
	mx := intArr[l]
	for _, val := range intArr[l : end+1] {
		mn = min(mn, val)
		mx = max(mx, val)
	}

	return mn + mx

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	// intArr is not guaranteed to have unique values
	intArr, err := GetIntArr(data)
	if err != nil {
		log.Fatal(err)
	}
	// res := SolvePartOne(intArr, 5)
	res := SolvePartOne(intArr, 25)
	fmt.Println(res)

	res2 := SolvePartTwo(intArr, res)
	fmt.Println(res2)

}
