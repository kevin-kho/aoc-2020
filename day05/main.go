package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Instruction struct {
	Seat   string
	Column string
}

type Seat struct {
	Row    int
	Column int
}

func (s Seat) GetSeatId() int {
	return s.Row*8 + s.Column
}

func (i Instruction) GetSeat() Seat {

	var row int
	var col int

	// find row
	l := 0
	r := 127
	for _, c := range i.Seat {
		mid := (l + r) / 2
		switch c {
		case 'F':
			r = mid
		case 'B':
			l = mid + 1
		}
	}
	if l != r {
		log.Fatal()
	}
	row = l

	// find col
	l = 0
	r = 7
	for _, c := range i.Column {
		mid := (l + r) / 2
		switch c {
		case 'R':
			l = mid + 1
		case 'L':
			r = mid
		}
	}
	if l != r {
		log.Fatal()
	}
	col = l

	return Seat{
		Row:    row,
		Column: col,
	}

}

func CreateInstructions(data []byte) []Instruction {
	var res []Instruction

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {

		entryStr := string(entry)
		s := entryStr[:7]
		c := entryStr[7:]

		res = append(res, Instruction{
			Seat:   s,
			Column: c,
		})
	}

	return res
}

func SolvePartOne(instructions []Instruction) int {
	var res int
	for _, i := range instructions {
		id := i.GetSeat().GetSeatId()
		res = max(res, id)
	}

	return res
}

func SolvePartTwo(instructions []Instruction) int {
	var res int
	var ids []int
	for _, i := range instructions {
		id := i.GetSeat().GetSeatId()
		ids = append(ids, id)
	}
	slices.Sort(ids)

	for i := 1; i < len(ids); i++ {
		l := ids[i-1]
		r := ids[i]

		if l+2 == r {
			res = l + 1
			break
		}

	}

	return res

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	ins := CreateInstructions(data)
	res := SolvePartOne(ins)
	fmt.Println(res)

	res2 := SolvePartTwo(ins)
	fmt.Println(res2)

}
