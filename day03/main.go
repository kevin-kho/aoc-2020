package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

func GetGrid(data []byte) [][]byte {
	return bytes.Split(data, []byte{'\n'})
}

func SolvePartOne(grid [][]byte) int {
	// tree = '#'
	var count int
	X := len(grid[0])
	x := 0
	y := 0
	for y < len(grid) {

		xMod := x % X

		if grid[y][xMod] == '#' {
			count++
		}

		x += 3
		y += 1
	}

	return count

}

func main() {
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	grid := GetGrid(data)

	res := SolvePartOne(grid)
	fmt.Println(res)

}
