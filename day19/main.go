package main

import (
	"bytes"
	"log"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

func GetRules(data []byte) (map[string][][]string, error) {
	res := make(map[string][][]string)

	for entry := range bytes.Lines(data) {
		entry = bytes.TrimSpace(entry)

		keyStr, valueStr, _ := strings.Cut(string(entry), ":")
		keyStr = strings.TrimSpace(keyStr)

		valueStr = strings.TrimSpace(valueStr)
		var valueStrs [][]string
		for v := range strings.SplitSeq(valueStr, "|") {
			v = strings.TrimSpace(v)

			var vs []string
			for i := range strings.SplitSeq(v, " ") {
				i = strings.TrimPrefix(i, "\"")
				i = strings.TrimSuffix(i, "\"")
				vs = append(vs, i)
			}
			valueStrs = append(valueStrs, vs)
		}
		res[keyStr] = valueStrs

	}

	return res, nil

}

func GetMessages(data []byte) []string {
	return strings.Split(string(data), "\n")
}

func SolvePartOne(msgs []string, rules map[string][][]string) {
	// 0 is root

}

func main() {
	ruleData, err := common.ReadInput("exampleRules.txt")
	if err != nil {
		log.Fatal(err)
	}
	ruleData = common.TrimNewLineSuffix(ruleData)

	rules, err := GetRules(ruleData)
	if err != nil {
		log.Fatal(err)
	}

	msgData, err := common.ReadInput("exampleMsgs.txt")
	if err != nil {
		log.Fatal(err)
	}
	msgData = common.TrimNewLineSuffix(msgData)
	msgs := GetMessages(msgData)

	// fmt.Println(rules)
	// fmt.Println(msgs)

	SolvePartOne(msgs, rules)

}
