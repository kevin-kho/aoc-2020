package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type BagMap map[string][]string

func GetBagMap(data []byte) (BagMap, error) {
	res := make(BagMap)
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		key, values, _ := strings.Cut(string(entry), "contain")
		key = strings.TrimSpace(key)
		key = strings.TrimSuffix(key, "s")

		var v []string
		for value := range strings.SplitSeq(values, ",") {
			value = strings.TrimSpace(value)
			value = strings.TrimSuffix(value, ".")
			value = strings.TrimSuffix(value, "s")

			if value == "no other bag" {
				continue
			}

			qty, bg, _ := strings.Cut(value, " ")
			qtyInt, err := strconv.Atoi(qty)
			if err != nil {
				return res, err
			}
			for range qtyInt {
				v = append(v, bg)
			}
		}

		res[key] = v
	}

	return res, nil

}

func main() {
	data, err := common.ReadInput("inputExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	mp, err := GetBagMap(data)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range mp {
		fmt.Println(k, v)
	}

}
