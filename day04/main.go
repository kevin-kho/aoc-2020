package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Passport map[string]string

func GetPassports(data []byte) []Passport {

	var res []Passport

	for passport := range bytes.SplitSeq(data, []byte{'\n', '\n'}) {

		p := make(Passport)

		for row := range strings.SplitSeq(string(passport), "\n") {
			for entry := range strings.SplitSeq(row, " ") {
				key, value, _ := strings.Cut(entry, ":")
				p[key] = value

			}
		}

		res = append(res, p)

	}

	return res

}

func SolvePartOne(passports []Passport) int {

	var count int
	for _, p := range passports {
		if len(p) == 8 {
			count++
		}

		if _, ok := p["cid"]; !ok && len(p) == 7 {
			count++
		}

	}

	return count

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	passports := GetPassports(data)

	res := SolvePartOne(passports)
	fmt.Println(res)
}
