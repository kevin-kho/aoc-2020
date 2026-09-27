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

type Grid struct {
	Active   map[Pos]bool
	Inactive map[Pos]bool
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
		}
	}

	return Grid{
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

func main() {
	data, err := common.ReadInput("inputExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	grid := CreateGrid(data)
	fmt.Println(grid)

}
