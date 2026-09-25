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

func main() {
	data, err := common.ReadInput("inputExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	cmds, err := GetCommands(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cmds)

}
