package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Slope struct {
	X int
	Y int
}

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

func SolvePartTwo(grid [][]byte) int {
	var trees []int
	slopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	X := len(grid[0])
	for _, d := range slopes {
		x := 0
		y := 0
		var count int
		for y < len(grid) {
			xMod := x % X

			if grid[y][xMod] == '#' {
				count++
			}

			x += d.X
			y += d.Y
		}
		trees = append(trees, count)

	}

	res := 1
	for _, t := range trees {
		res *= t
	}

	return res

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

	res2 := SolvePartTwo(grid)
	fmt.Println(res2)

}
