package main

import (
	"bytes"
	"cmp"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Criteria struct {
	Name      string
	Intervals []Interval
}

type Interval struct {
	Start int
	End   int
}

type Ticket []int

func ReadFile(filePath string) ([]byte, error) {
	data, err := common.ReadInput(filePath)
	if err != nil {
		return data, err
	}
	common.TrimNewLineSuffix(data)

	return data, nil

}

func GetCriteria(filePath string) ([]Criteria, error) {
	var res []Criteria
	data, err := ReadFile(filePath)
	if err != nil {
		return res, err
	}

	for entry := range bytes.Lines(data) {
		entryStr := strings.TrimSpace(string(entry))

		name, rangeStr, _ := strings.Cut(entryStr, ": ")
		name = strings.TrimSpace(name)
		ranges := strings.Split(strings.TrimSpace(rangeStr), " or ")

		var intervals []Interval
		for _, r := range ranges {

			st, ed, _ := strings.Cut(r, "-")
			stInt, err := strconv.Atoi(st)
			if err != nil {
				return res, err
			}
			edInt, err := strconv.Atoi(ed)
			if err != nil {
				return res, err
			}

			intervals = append(intervals, Interval{stInt, edInt})

		}

		res = append(res, Criteria{
			Name:      name,
			Intervals: intervals,
		})
	}

	return res, nil

}

func GetNearbyTickets(filePath string) ([]Ticket, error) {

	var res []Ticket
	data, err := ReadFile(filePath)
	if err != nil {
		return res, err
	}

	for entry := range bytes.Lines(data) {
		entry = bytes.TrimSpace(entry)

		var ticket Ticket

		for val := range strings.SplitSeq(string(entry), ",") {
			valInt, err := strconv.Atoi(val)
			if err != nil {
				return res, err
			}
			ticket = append(ticket, valInt)
		}
		res = append(res, ticket)
	}

	return res, nil

}

func GetMergedIntervals(criteria []Criteria) []Interval {
	var res []Interval

	var intervals []Interval
	for _, c := range criteria {
		intervals = append(intervals, c.Intervals...)
	}

	// Sort by start time then end time.
	// Ensures later end times come later
	slices.SortFunc(intervals, func(a, b Interval) int {
		return cmp.Or(
			cmp.Compare(a.Start, b.Start),
			cmp.Compare(a.End, b.End),
		)
	})

	res = append(res, intervals[0])
	i := 1
	for i < len(intervals) {
		prevI := len(res) - 1
		prev := res[prevI]

		curr := intervals[i]

		if prev.Start <= curr.Start {
			st := min(prev.Start, curr.Start)
			ed := max(prev.End, curr.End)
			res[prevI].Start = st
			res[prevI].End = ed
		} else {
			res = append(res, curr)
		}

		i++
	}

	return res

}

func SolvePartOne(criteria []Criteria, tickets []Ticket) int {
	intervals := GetMergedIntervals(criteria)
	var invalidValues []int
	var res int
	for _, t := range tickets {
		for _, val := range t {
			valid := true
			for _, i := range intervals {
				if !(i.Start <= val && val <= i.End) {
					valid = false
				}
			}
			if !valid {
				invalidValues = append(invalidValues, val)
			}
		}
	}

	for _, v := range invalidValues {
		res += v
	}

	return res

}

func main() {
	// First parse valid-criteria
	criteria, err := GetCriteria("valid-criteria.txt")
	if err != nil {
		log.Fatal(err)
	}
	tickets, err := GetNearbyTickets("nearby-tickets.txt")
	if err != nil {
		log.Fatal(err)
	}
	res := SolvePartOne(criteria, tickets)
	fmt.Println(res)

}
