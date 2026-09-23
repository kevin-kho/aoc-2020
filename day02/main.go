package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Policy struct {
	Character rune
	Min       int
	Max       int
}

type Password struct {
	Value  string
	Policy Policy
}

func (p Password) IsValid() bool {
	var count int
	for _, r := range p.Value {
		if r == p.Policy.Character {
			count++
		}
	}
	return p.Policy.Min <= count && count <= p.Policy.Max
}

func GetPasswords(data []byte) ([]Password, error) {
	var res []Password

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		strArr := strings.Split(string(entry), ":")

		policyStr := strArr[0]
		r := strings.Split(policyStr, " ")[0]
		minVal := strings.Split(r, "-")[0]
		minValInt, err := strconv.Atoi(minVal)
		if err != nil {
			return res, err
		}
		maxVal := strings.Split(r, "-")[1]
		maxValInt, err := strconv.Atoi(maxVal)
		if err != nil {
			return res, err
		}
		char := rune(strings.Split(policyStr, " ")[1][0])

		pwdStr := strArr[1]
		pwdStr = strings.TrimSpace(pwdStr)

		res = append(res, Password{
			Value: pwdStr,
			Policy: Policy{
				Character: char,
				Min:       minValInt,
				Max:       maxValInt,
			},
		})

	}

	return res, nil
}

func SolvePartOne(pwds []Password) int {
	var count int
	for _, p := range pwds {
		if p.IsValid() {
			count++
		}
	}

	return count
}

func main() {

	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	pwds, err := GetPasswords(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(pwds)
	fmt.Println(res)

}
