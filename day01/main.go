package main

import (
	"bytes"
	"fmt"
	"log"
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

}
