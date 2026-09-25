package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type CommandType int

type Memory map[string]string

func (m Memory) Totalize() (int64, error) {
	var res int64
	for _, val := range m {
		i, err := strconv.ParseInt(val, 2, 0)
		if err != nil {
			return res, err
		}
		res += i
	}
	return res, nil
}

const (
	CommandTypeMask CommandType = iota
	CommandTypeWrite
)

type Command struct {
	Type    CommandType
	Address int
	Value   string // 36-bit binary
}

func (c Command) GetValue() (int64, error) {
	i, err := strconv.ParseInt(c.Value, 2, 0)
	if err != nil {
		return i, err
	}

	return i, nil
}

func CreateMaskCmd(data []byte) Command {

	_, val, _ := strings.Cut(string(data), "mask = ")

	return Command{
		Type:    CommandTypeMask,
		Address: 0,
		Value:   val,
	}

}

func CreateWriteCmd(data []byte) (Command, error) {

	var res Command

	addr, val, _ := strings.Cut(string(data), " = ")

	addr = strings.TrimPrefix(addr, "mem[")
	addr = strings.TrimSuffix(addr, "]")
	addrInt, err := strconv.Atoi(addr)
	if err != nil {
		return res, err
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return res, err
	}
	val = fmt.Sprintf("%036b", valInt)

	res.Type = CommandTypeWrite
	res.Address = addrInt
	res.Value = val
	return res, nil

}

func GetCommands(data []byte) ([]Command, error) {

	var res []Command

	for entry := range bytes.Lines(data) {
		entry = bytes.TrimSpace(entry)

		var cmd Command
		var err error
		if bytes.HasPrefix(entry, []byte("mask = ")) {
			cmd = CreateMaskCmd(entry)
		} else {
			cmd, err = CreateWriteCmd(entry)
			if err != nil {
				return res, err
			}
		}

		res = append(res, cmd)

	}

	return res, nil

}

func MaskValue(mask string, val string) string {
	var sb strings.Builder
	for i := range len(mask) {
		m := mask[i]
		v := val[i]

		if m == 'X' {
			sb.WriteByte(v)
		} else {
			sb.WriteByte(m)
		}
	}

	return sb.String()
}

func MaskAddress(mask string, address string) string {
	var sb strings.Builder

	for i := range len(mask) {
		m := mask[i]
		a := address[i]

		switch m {
		case '0':
			sb.WriteByte(a)
		case 'X', '1':
			sb.WriteByte(m)
		}

	}

	return sb.String()
}

func SolvePartOne(cmds []Command) (int64, error) {
	mp := make(Memory)
	var mask string
	for _, cmd := range cmds {
		switch cmd.Type {
		case CommandTypeMask:
			mask = cmd.Value
		case CommandTypeWrite:
			val := MaskValue(mask, cmd.Value)
			mp[fmt.Sprintf("%036b", cmd.Address)] = val
		}
	}

	res, err := mp.Totalize()
	if err != nil {
		return res, err
	}

	return res, nil

}

func UnfloatAddress(addrFloat string) []string {
	var res []string

	var dfs func(i int, curr string)
	dfs = func(i int, curr string) {

		// exit condition
		if i == len(addrFloat) {
			res = append(res, curr)
			return
		}

		char := addrFloat[i]
		switch char {
		case 'X':
			dfs(i+1, curr+"0")
			dfs(i+1, curr+"1")
		default:
			dfs(i+1, curr+string(char))
		}

	}

	dfs(0, "")

	return res
}

func SolvePartTwo(cmds []Command) (int64, error) {

	mp := make(Memory)
	var mask string
	for _, cmd := range cmds {
		switch cmd.Type {
		case CommandTypeMask:
			mask = cmd.Value
		case CommandTypeWrite:
			addrFloat := MaskAddress(mask, fmt.Sprintf("%036b", cmd.Address))
			for _, addr := range UnfloatAddress(addrFloat) {
				mp[addr] = cmd.Value
			}
		}
	}

	res, err := mp.Totalize()
	if err != nil {
		return res, err
	}
	return res, nil

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

	res, err := SolvePartOne(cmds)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	res2, err := SolvePartTwo(cmds)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res2)

}
