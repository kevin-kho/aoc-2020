package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

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
		"e":  {1, 0},
		"se": {1, -1},
		"sw": {-1, -1},
		"w":  {-1, 0},
		"nw": {-1, 1},
	}
	return mp[dir]
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

func SolvePartOne(dirs []Directions) int {
	mp := make(map[Pos]bool) // false: white, true: black
	for _, dir := range dirs {
		tile := Pos{0, 0}
		for _, d := range dir {
			tile.Move(d)
		}
		mp[tile] = !mp[tile]
	}

	var count int
	for _, isBlack := range mp {
		if isBlack {
			count++
		}
	}

	return count

}

func main() {

	data, err := common.ReadInput("inputExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	dirs := GetDirections(data)
	res := SolvePartOne(dirs)
	fmt.Println(res)

}
