package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Hand []int

func (h Hand) GetString() string {
	var sb strings.Builder

	for _, val := range h {
		sb.WriteString(strconv.Itoa(val))
		sb.WriteByte(',')
	}

	return sb.String()

}

func GetCards(data []byte) (Hand, error) {
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

func CaclulateScore(cards []int) int {
	var res int
	deckSize := len(cards)

	for _, card := range cards {
		res += (deckSize * card)
		deckSize--
	}

	return res
}

func SolvePartOne(p1Cards, p2Cards Hand) int {

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
		score = CaclulateScore(p1Cards)
	}

	if len(p2Cards) > 0 {
		score = CaclulateScore(p2Cards)
	}

	return score

}

func SolvePartTwo(p1Cards, p2Cards Hand) int {

	gameCount := 1
	// bool represents if p1 wins
	var recurse func(p1Hand, p2Hand Hand, p1Seen, p2Seen map[string]bool, game int, round int) (bool, Hand, Hand)
	recurse = func(p1Hand, p2Hand Hand, p1Seen, p2Seen map[string]bool, game int, round int) (bool, Hand, Hand) {

		// fmt.Println(p1Hand, p2Hand, game, round)

		// exit condition: player one lost
		if len(p1Hand) == 0 {
			return false, p1Hand, p2Hand
		}

		// exit condition: player one wins
		if len(p2Hand) == 0 {
			return true, p1Hand, p2Hand
		}

		if p1Seen[p1Hand.GetString()] && p2Seen[p2Hand.GetString()] {
			return true, p1Hand, p2Hand
		}

		p1Seen[p1Hand.GetString()] = true
		p2Seen[p2Hand.GetString()] = true
		p1Card, p2Card := p1Hand[0], p2Hand[0]
		p1Hand, p2Hand = p1Hand[1:], p2Hand[1:]

		var p1Wins bool
		// Subgame condition
		if p1Card <= len(p1Hand) && p2Card <= len(p2Hand) {
			gameCount++
			res, _, _ := recurse(slices.Clone(p1Hand[:p1Card]), slices.Clone(p2Hand[:p2Card]), map[string]bool{}, map[string]bool{}, gameCount, 1)
			p1Wins = p1Wins || res
		} else if p1Card > p2Card {
			p1Wins = true
		} else if p1Card < p2Card {
			p1Wins = false
		} else {
			// Tie
			p1Hand = append(p1Hand, p1Card)
			p2Hand = append(p2Hand, p2Card)
			return recurse(p1Hand, p2Hand, p1Seen, p2Seen, game, round+1)
		}

		// Non-tie situations
		if p1Wins {
			p1Hand = append(p1Hand, p1Card, p2Card)
		} else {
			p2Hand = append(p2Hand, p2Card, p1Card)
		}

		return recurse(p1Hand, p2Hand, p1Seen, p2Seen, game, round+1)

	}

	p1Wins, p1HandFinal, p2HandFinal := recurse(p1Cards, p2Cards, map[string]bool{}, map[string]bool{}, gameCount, 1)

	var score int
	switch p1Wins {
	case true:
		score = CaclulateScore(p1HandFinal)
	case false:
		score = CaclulateScore(p2HandFinal)
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

	res2 := SolvePartTwo(p1Cards, p2Cards)
	fmt.Println(res2)

}
