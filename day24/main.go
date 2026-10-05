package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Floor struct {
	Black map[Pos]bool
}

type Directions []Pos

type Pos struct {
	X int
	Y int
}

func (p *Pos) Move(d Pos) {
	p.X += d.X
	p.Y += d.Y
}

func DirToPos(dir string) Pos {
	mp := map[string]Pos{
		"ne": {1, 1},
		"e":  {2, 0},
		"se": {1, -1},
		"sw": {-1, -1},
		"w":  {-2, 0},
		"nw": {-1, 1},
	}
	return mp[dir]
}

func GetAdjs() []Pos {
	var res []Pos
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {

			if y == 0 && (x == 0 || common.IntAbs(x) == 1) {
				continue
			}
			if x == 0 {
				continue
			}

			res = append(res, Pos{
				X: x,
				Y: y,
			})
		}
	}

	res = append(res, Pos{0, 2})
	res = append(res, Pos{0, -2})

	return res
}

func ParseDirections(dirs []byte) Directions {
	var prefix byte
	var directions []string

	for _, b := range dirs {
		switch b {
		case 'n', 's':
			prefix = b
		default:
			// Trim off null byte
			dir := bytes.TrimLeft([]byte{prefix, b}, "\x00")
			directions = append(directions, string(dir))
			prefix = 0
		}
	}

	var res Directions
	for _, d := range directions {
		res = append(res, DirToPos(d))
	}

	return res

}

func GetDirections(data []byte) []Directions {
	var res []Directions

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		res = append(res, ParseDirections(entry))
	}

	return res

}

func GetInitialFloor(dirs []Directions) Floor {
	mp := make(map[Pos]bool) // false: white, true: black
	for _, dir := range dirs {
		tile := Pos{0, 0}
		for _, d := range dir {
			tile.Move(d)
		}
		mp[tile] = !mp[tile]
	}

	black := make(map[Pos]bool)

	for pos, isBlack := range mp {
		if isBlack {
			black[pos] = true
		}
	}

	return Floor{
		Black: black,
	}

}

func SolvePartTwo(floor Floor) {

	fmt.Println(floor)
	adj := GetAdjs()
	fmt.Println(adj)

	for range 10 {

		newBlack := make(map[Pos]bool)
		white := make(map[Pos]int)
		for pos := range floor.Black {
			var blackNei int
			for _, d := range adj {
				adjPos := Pos{
					X: pos.X + d.X,
					Y: pos.Y + d.Y,
				}

				if floor.Black[adjPos] {
					blackNei++
				} else {
					white[adjPos]++
				}
			}
			if blackNei == 0 || blackNei > 2 {
				continue
			} else {
				newBlack[pos] = true
			}

		}

		for pos, ct := range white {
			if ct == 2 {
				newBlack[pos] = true
			}
		}

		floor.Black = newBlack

		fmt.Println(len(floor.Black))
	}

}

func main() {

	data, err := common.ReadInput("inputExample.txt")
	// data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	dirs := GetDirections(data)
	floor := GetInitialFloor(dirs)
	fmt.Println(len(floor.Black))

	SolvePartTwo(floor)

}
