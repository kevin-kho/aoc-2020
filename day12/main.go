package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Command struct {
	Action Direction
	Value  int
}

type Pos struct {
	X int
	Y int
}

type Direction int

const (
	N Direction = iota
	E
	S
	W
	L
	R
	F
)

type Ship struct {
	Pos
	Direction Direction
}

func (s *Ship) TurnLeft(val int) {
	val = val / 90
	newDir := Direction((int(s.Direction) - val) % 4)
	if newDir < 0 {
		newDir += 4
	}

	s.Direction = newDir
}

func (s *Ship) TurnRight(val int) {
	val = val / 90
	newDir := Direction((int(s.Direction) + val) % 4)

	s.Direction = newDir

}

func (s *Ship) Move(delta Pos) {
	s.X += delta.X
	s.Y += delta.Y
}

func (s *Ship) HandleMove(d Direction, val int) {

	if d == F {
		d = s.Direction
	}

	switch d {
	case N:
		s.Move(Pos{0, val})
	case E:
		s.Move(Pos{val, 0})
	case S:
		s.Move(Pos{0, -val})
	case W:
		s.Move(Pos{-val, 0})
	case L:
		s.TurnLeft(val)
	case R:
		s.TurnRight(val)
	}

}

func GetCommands(data []byte) ([]Command, error) {
	var res []Command
	for entry := range bytes.Lines(data) {
		entry = bytes.TrimSpace(entry)

		var dir Direction
		switch string(entry)[0] {
		case 'N':
			dir = 0
		case 'E':
			dir = 1
		case 'S':
			dir = 2
		case 'W':
			dir = 3
		case 'L':
			dir = 4
		case 'R':
			dir = 5
		case 'F':
			dir = 6
		}
		val, err := strconv.Atoi(string(entry)[1:])
		if err != nil {
			return res, err
		}

		res = append(res, Command{
			Action: dir,
			Value:  val,
		})

	}

	return res, nil

}

func SolvePartOne(cmds []Command) int {

	// Ship starts facing east
	ship := Ship{
		Pos:       Pos{0, 0},
		Direction: E,
	}
	for _, cmd := range cmds {
		ship.HandleMove(cmd.Action, cmd.Value)
	}

	return common.IntAbs(ship.X) + common.IntAbs(ship.Y)

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	cmds, err := GetCommands(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(cmds)
	fmt.Println(res)

}
