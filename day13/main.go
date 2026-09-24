package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type State struct {
	Time  int
	Buses []int
}

func GetCurrentState(data []byte) (State, error) {

	var res State

	timeStr, busStr, _ := strings.Cut(string(data), "\n")

	time, err := strconv.Atoi(timeStr)
	if err != nil {
		return res, err
	}

	var buses []int
	for val := range strings.SplitSeq(busStr, ",") {
		if val == "x" {
			continue
		}

		id, err := strconv.Atoi(val)
		if err != nil {
			return res, err
		}
		buses = append(buses, id)
	}

	res.Time = time
	res.Buses = buses

	return res, nil

}

func SolvePartOne(state State) int {

	mp := make(map[int]int) // key: busId, value: minutes

	for _, busId := range state.Buses {
		rem := state.Time % busId

		// Bus is available from the start
		if rem == 0 {
			return busId * rem
		}

		mp[busId] = busId - rem

	}

	var busId int
	var minutes int

	// Choose a random value in the map for starting value
	for k, v := range mp {
		busId = k
		minutes = v
		break
	}

	for id, mn := range mp {
		if mn < minutes {
			busId = id
			minutes = mn
		}
	}

	return busId * minutes

}

func main() {

	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	state, err := GetCurrentState(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(state)
	fmt.Println(res)

}
