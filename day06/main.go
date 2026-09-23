package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Group struct {
	Answers []string
}

func (g Group) GetQuestions() map[rune]int {

	var sb strings.Builder
	for _, a := range g.Answers {
		sb.WriteString(a)
	}

	q := make(map[rune]int)

	for _, c := range sb.String() {
		q[c]++
	}

	return q

}

func GetGroups(data []byte) []Group {
	var res []Group

	for group := range bytes.SplitSeq(data, []byte{'\n', '\n'}) {

		ans := strings.Split(string(group), "\n")
		res = append(res, Group{
			Answers: ans,
		})
	}

	return res
}

func SolvePartOne(grps []Group) int {
	var res int

	for _, g := range grps {
		res += len(g.GetQuestions())
	}

	return res

}

func SolvePartTwo(grps []Group) int {
	var res int
	for _, g := range grps {
		for _, ct := range g.GetQuestions() {
			if ct == len(g.Answers) {
				res++
			}
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

	grps := GetGroups(data)

	res := SolvePartOne(grps)
	fmt.Println(res)

	res2 := SolvePartTwo(grps)
	fmt.Println(res2)
}
