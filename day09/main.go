package main

import (
	"bytes"
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

func SolvePartOne(intArr []int) {

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	intArr, err := GetIntArr(data)
	if err != nil {
		log.Fatal(err)
	}
	SolvePartOne(intArr)

}
