package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Passport map[string]string

func (p Passport) IsValid() (bool, error) {

	fields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	// Check all fields present
	for _, f := range fields {
		if _, ok := p[f]; !ok {
			return false, nil
		}
	}

	// check byr
	byr, err := strconv.Atoi(p["byr"])
	if err != nil {
		return false, err
	}
	if !(1920 <= byr && byr <= 2002) {
		fmt.Println("invalid byr:", byr)
		return false, nil
	}

	// check iyr
	iyr, err := strconv.Atoi(p["iyr"])
	if err != nil {
		return false, err
	}
	if !(2010 <= iyr && iyr <= 2020) {
		fmt.Println("invalid iyr:", iyr)
		return false, err
	}

	// check eyr
	eyr, err := strconv.Atoi(p["eyr"])
	if err != nil {
		return false, err
	}
	if !(2020 <= eyr && eyr <= 2030) {
		return false, err
	}

	// check hgt
	var sfx string
	hgt := p["hgt"]
	if strings.HasSuffix(hgt, "cm") {
		sfx = "cm"
	} else {
		sfx = "in"
	}

	h, _, _ := strings.Cut(hgt, sfx)
	hInt, err := strconv.Atoi(h)
	if err != nil {
		return false, err
	}
	if sfx == "cm" && !(150 <= hInt && hInt <= 193) {
		fmt.Println("invalid height", p["hgt"])
		return false, err
	}
	if sfx == "in" && !(59 <= hInt && hInt <= 76) {
		fmt.Println("invalid height", p["hgt"])
		return false, err
	}

	// check hcl
	_, after, found := strings.Cut(p["hcl"], "#")
	if !found {
		fmt.Println("invalid hcl:", p["hcl"])
		return false, nil
	}
	for _, c := range after {
		if (48 <= c && c <= 57) || (97 <= c && c <= 102) {
			continue
		}
		fmt.Println("invalid hcl:", p["hcl"], after, c)
		return false, nil
	}

	// check ecl
	validEcl := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	if !validEcl[p["ecl"]] {
		fmt.Println("invalid ecl", p["ecl"])
		return false, nil
	}

	// check pid
	pid := p["pid"]
	if len(pid) != 9 {
		fmt.Println("invalid pid", p["pid"])
		return false, nil
	}
	for _, c := range pid {
		if !(48 <= c && c <= 57) {
			fmt.Println("invalid pid", p["pid"], c)
			return false, nil
		}
	}

	fmt.Println("valid passport")
	return true, nil

}

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

func SolvePartTwo(passports []Passport) int {
	var count int
	for _, p := range passports {
		v, err := p.IsValid()
		if err != nil {
			log.Fatal(err)
		}
		if v {
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

	res2 := SolvePartTwo(passports)
	fmt.Println(res2)
}
