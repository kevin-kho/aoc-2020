package main

import (
	"bytes"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Pos struct {
	X int
	Y int
}

func (p Pos) GetAdjacentPos() []Pos {
	var res []Pos
	for _, d := range CreateDeltas() {
		res = append(res, Pos{
			X: p.X + d.X,
			Y: p.Y + d.Y,
		})
	}
	return res
}

func CreateDeltas() []Pos {
	var res []Pos
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 {
				continue
			}
			res = append(res, Pos{
				X: x,
				Y: y,
			})
		}
	}
	return res
}

type Grid struct {
	Empty  map[Pos]bool
	Filled map[Pos]bool
}

func CreateGrid(data []byte) Grid {
	grd := bytes.Split(data, []byte{'\n'})

	empty := make(map[Pos]bool)
	filled := make(map[Pos]bool)

	for y, row := range grd {
		for x, val := range row {
			if val == 'L' {
				empty[Pos{
					X: x,
					Y: y,
				}] = true
			}

			if val == '#' {
				filled[Pos{
					X: x,
					Y: y,
				}] = true
			}
		}
	}

	return Grid{
		Empty:  empty,
		Filled: filled,
	}
}

func SolvePartOne(grid Grid) {

}

func main() {
	data, err := common.ReadInput("inputExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	grid := CreateGrid(data)

	SolvePartOne(grid)

}
