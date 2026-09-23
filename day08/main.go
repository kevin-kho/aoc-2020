package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Command struct {
	Action string
	Value  int
}

func GetCommands(data []byte) ([]Command, error) {
	var res []Command

	for entry := range bytes.Lines(data) {
		entry = bytes.TrimSpace(entry)

		action, value, _ := strings.Cut(string(entry), " ")
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return res, err
		}

		res = append(res, Command{
			Action: action,
			Value:  valueInt,
		})

	}

	return res, nil

}

func SolvePartOne(cmds []Command) int {
	var res int
	prevI := 0
	i := 0
	seen := make(map[int]bool)
	var broke bool

	for i < len(cmds) {

		if seen[i] {
			fmt.Println("cycle found at index: ", i, prevI, cmds[i])
			broke = true
			break
		}

		seen[i] = true

		prevI = i
		switch cmds[i].Action {
		case "nop":
			i++
		case "acc":
			res += cmds[i].Value
			i++
		case "jmp":
			i += cmds[i].Value
		}

	}

	if !broke {
		fmt.Println("exited normally")
	}

	return res

}

func FindCandidates(cmds []Command) {
	// In input.txt, you need to find a way to the last 5 commands
	dest := make(map[int]bool)
	for i := len(cmds) - 5; i < len(cmds); i++ {
		dest[i] = true
		fmt.Println(i)
	}

}

func SolvePartTwo(cmds []Command) {

	var idxes []int // indexes of all nop and jmp cmds
	for i, cmd := range cmds {
		if cmd.Action == "nop" || cmd.Action == "jmp" {
			idxes = append(idxes, i)
		}
	}

	for _, i := range idxes {
		cmds := slices.Clone(cmds)
		switch cmds[i].Action {
		case "nop":
			cmds[i].Action = "jmp"
		case "jmp":
			cmds[i].Action = "nop"
		}
		res := SolvePartOne(cmds)
		fmt.Println(res)
	}

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

	SolvePartTwo(cmds)

}
