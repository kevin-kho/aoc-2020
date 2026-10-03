package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

func GetCards(data []byte) ([]int, error) {
	var res []int

	for card := range bytes.SplitSeq(data, []byte{'\n'}) {
		cardInt, err := strconv.Atoi(string(card))
		if err != nil {
			return res, err
		}
		res = append(res, cardInt)
	}

	return res, nil
}

func CalculateScore(cards []int) int {
	var res int
	deckSize := len(cards)

	for _, card := range cards {
		res += (deckSize * card)
		deckSize--
	}

	return res
}

func SolvePartOne(p1Cards, p2Cards []int) int {

	for len(p1Cards) > 0 && len(p2Cards) > 0 {
		p1 := p1Cards[0]
		p1Cards = p1Cards[1:]

		p2 := p2Cards[0]
		p2Cards = p2Cards[1:]

		if p1 > p2 {
			p1Cards = append(p1Cards, p1, p2)
		} else if p1 < p2 {
			p2Cards = append(p2Cards, p2, p1)

		} else {
			p1Cards = append(p1Cards, p1)
			p2Cards = append(p2Cards, p2)
		}
	}

	var score int
	if len(p1Cards) > 0 {
		score = CalculateScore(p1Cards)
	}

	if len(p2Cards) > 0 {
		score = CalculateScore(p2Cards)
	}

	return score

}

func main() {
	// p1Data, err := common.ReadInput("player1Example.txt")
	p1Data, err := common.ReadInput("player1.txt")
	if err != nil {
		log.Fatal(err)
	}
	p1Data = common.TrimNewLineSuffix(p1Data)
	p1Cards, err := GetCards(p1Data)
	if err != nil {
		log.Fatal(err)
	}

	// p2Data, err := common.ReadInput("player2Example.txt")
	p2Data, err := common.ReadInput("player2.txt")
	if err != nil {
		log.Fatal(err)
	}
	p2Data = common.TrimNewLineSuffix(p2Data)
	p2Cards, err := GetCards(p2Data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(p1Cards, p2Cards)
	fmt.Println(res)

}
