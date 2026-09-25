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

type Memory map[int]string

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

func ApplyMask(mask string, val string) string {
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

func SolvePartOne(cmds []Command) (int64, error) {
	mp := make(Memory)
	var mask string
	for _, cmd := range cmds {
		switch cmd.Type {
		case CommandTypeMask:
			mask = cmd.Value
		case CommandTypeWrite:
			val := ApplyMask(mask, cmd.Value)
			mp[cmd.Address] = val
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

}
