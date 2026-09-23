package main

import (
	"bytes"
	"fmt"
	"log"
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
	i := 0
	seen := make(map[int]bool)

	for i < len(cmds) {

		if seen[i] {
			break
		}

		seen[i] = true

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

	return res

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
