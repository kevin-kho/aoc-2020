package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Pos struct {
	X int
	Y int
	Z int
}

type Pos4d struct {
	Pos
	W int
}

type Grid struct {
	Active   map[Pos]bool
	Inactive map[Pos]bool
}

type Grid4d struct {
	Active   map[Pos4d]bool
	Inactive map[Pos4d]bool
}

func CreateGrid(data []byte) Grid {

	active := make(map[Pos]bool)
	inactive := make(map[Pos]bool)

	for y, row := range bytes.Split(data, []byte{'\n'}) {
		for x, val := range row {
			switch val {
			case '#':
				active[Pos{
					X: x,
					Y: y,
					Z: 0,
				}] = true
			case '.':
				inactive[Pos{
					X: x,
					Y: y,
					Z: 0,
				}] = true
			}

			// Set all x,y along z = -1 and z = +1 as inactive
			inactive[Pos{
				X: x,
				Y: y,
				Z: -1,
			}] = true
			inactive[Pos{
				X: x,
				Y: y,
				Z: 1,
			}] = true
		}
	}

	return Grid{
		Active:   active,
		Inactive: inactive,
	}

}

func CreateGrid4d(data []byte) Grid4d {
	active := make(map[Pos4d]bool)
	inactive := make(map[Pos4d]bool)

	for y, row := range bytes.Split(data, []byte{'\n'}) {
		for x, val := range row {
			switch val {
			case '#':
				active[Pos4d{
					X: x,
					Y: y,
					Z: 0,
					W: 0,
				}] = true
			case '.':
				inactive[Pos4d{
					X: x,
					Y: y,
					Z: 0,
					W: 0,
				}] = true
			}

			// Set all x,y along z, w = -1 and z, w = +1 as inactive
			for z := -1; z < 2; z++ {
				for w := -1; w < 2; w++ {
					if z == 0 && w == 0 {
						continue
					}
					inactive[Pos4d{
						X: x,
						Y: y,
						Z: z,
						W: w,
					}] = true

				}
			}

		}
	}

	return Grid4d{
		Active:   active,
		Inactive: inactive,
	}

}

func GetDeltas() []Pos {
	var res []Pos

	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			for dz := -1; dz < 2; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				res = append(res, Pos{
					X: dx,
					Y: dy,
					Z: dz,
				})
			}
		}
	}

	return res

}

func GetDeltas4d() []Pos4d {
	var res []Pos4d

	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			for dz := -1; dz < 2; dz++ {
				for dw := -1; dw < 2; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					res = append(res, Pos4d{
						X: dx,
						Y: dy,
						Z: dz,
						W: dw,
					})
				}
			}
		}
	}

	return res
}

func SolvePartOne(grid Grid) int {
	deltas := GetDeltas()

	for range 6 {
		active := make(map[Pos]bool)
		inactive := make(map[Pos]bool)

		// assess active
		for p := range grid.Active {
			var nei int
			for _, d := range deltas {
				n := Pos{
					X: p.X + d.X,
					Y: p.Y + d.Y,
					Z: p.Z + d.Z,
				}
				if grid.Active[n] {
					nei++
				} else {
					grid.Inactive[n] = true
				}
			}
			if nei == 2 || nei == 3 {
				active[p] = true
			} else {
				inactive[p] = true
			}
		}

		// assess inactive
		for p := range grid.Inactive {
			var nei int
			for _, d := range deltas {
				n := Pos{
					X: p.X + d.X,
					Y: p.Y + d.Y,
					Z: p.Z + d.Z,
				}
				if grid.Active[n] {
					nei++
				}
			}
			if nei == 3 {
				active[p] = true
			} else {
				inactive[p] = true
			}
		}

		grid.Active = active
		grid.Inactive = inactive
	}

	return len(grid.Active)

}

func SolvePartTwo(grid Grid4d) {
	deltas := GetDeltas4d()

	for range 6 {
		active := make(map[Pos4d]bool)
		inactive := make(map[Pos4d]bool)

		// assess active
		for p := range grid.Active {
			var nei int
			for _, d := range deltas {
				n := Pos4d{
					X: p.X + d.X,
					Y: p.Y + d.Y,
					Z: p.Z + d.Z,
					W: p.W + d.W,
				}
				if grid.Active[n] {
					nei++
				} else {
					grid.Inactive[n] = true
				}
			}
			if nei == 2 || nei == 3 {
				active[p] = true
			} else {
				inactive[p] = true
			}
		}

		// assess inactive
		for p := range grid.Inactive {
			var nei int
			for _, d := range deltas {
				n := Pos4d{
					X: p.X + d.X,
					Y: p.Y + d.Y,
					Z: p.Z + d.Z,
					W: p.W + d.W,
				}
				if grid.Active[n] {
					nei++
				}
			}
			if nei == 3 {
				active[p] = true
			} else {
				inactive[p] = true
			}
		}

		grid.Active = active
		grid.Inactive = inactive
	}

	fmt.Println(len(grid.Active))
	// return len(grid.Active)
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

	grid4d := CreateGrid4d(data)
	SolvePartTwo(grid4d)

}
