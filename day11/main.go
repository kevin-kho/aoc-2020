package main

import (
	"bytes"
	"fmt"
	"log"
	"maps"

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
	XLimit int
	YLimit int
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

	yLimit := len(grd)
	xLimit := len(grd[0])

	return Grid{
		Empty:  empty,
		Filled: filled,
		XLimit: xLimit,
		YLimit: yLimit,
	}
}

func SolvePartOne(grid Grid) int {

	for {
		empty := make(map[Pos]bool)
		filled := make(map[Pos]bool)

		// assess empty
		for s := range grid.Empty {
			var count int
			adj := s.GetAdjacentPos()
			for _, a := range adj {
				if grid.Filled[a] {
					count++
				}
			}
			if count == 0 {
				filled[s] = true
			} else {
				empty[s] = true
			}
		}

		// assess filled
		for s := range grid.Filled {
			var count int
			adj := s.GetAdjacentPos()
			for _, a := range adj {
				if grid.Filled[a] {
					count++
				}
			}
			if count >= 4 {
				empty[s] = true
			} else {
				filled[s] = true
			}
		}

		if maps.Equal(empty, grid.Empty) && maps.Equal(filled, grid.Filled) {
			break
		}

		grid.Empty = empty
		grid.Filled = filled
	}

	return len(grid.Filled)

}

func SolvePartTwo(grid Grid) int {

	X := grid.XLimit
	Y := grid.YLimit

	var dfs func(curr Pos, delta Pos) *Pos
	dfs = func(curr, delta Pos) *Pos {

		// case: out of bounds
		if !(0 <= curr.X && curr.X < X) || !(0 <= curr.Y && curr.Y < Y) {
			return nil
		}

		// case: found a chair
		if grid.Empty[curr] || grid.Filled[curr] {
			return new(curr)
		}

		// Keep travelling
		return dfs(Pos{
			X: curr.X + delta.X,
			Y: curr.Y + delta.Y,
		}, delta)

	}

	for {
		empty := make(map[Pos]bool)
		filled := make(map[Pos]bool)

		// assess empty
		for s := range grid.Empty {
			var count int
			for _, d := range CreateDeltas() {
				res := dfs(Pos{
					X: s.X + d.X,
					Y: s.Y + d.Y,
				}, d)
				if res != nil && grid.Filled[*res] {
					count++
				}
			}
			if count == 0 {
				filled[s] = true
			} else {
				empty[s] = true
			}
		}

		// assess filled
		for s := range grid.Filled {
			var count int
			for _, d := range CreateDeltas() {
				res := dfs(Pos{
					X: s.X + d.X,
					Y: s.Y + d.Y,
				}, d)
				if res != nil && grid.Filled[*res] {
					count++
				}
			}
			if count >= 5 {
				empty[s] = true
			} else {
				filled[s] = true
			}

		}

		if maps.Equal(empty, grid.Empty) && maps.Equal(filled, grid.Filled) {
			break
		}
		grid.Empty = empty
		grid.Filled = filled

	}

	return len(grid.Filled)

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	grid := CreateGrid(data)

	res := SolvePartOne(grid)
	fmt.Println(res)

	res2 := SolvePartTwo(grid)
	fmt.Println(res2)

}
