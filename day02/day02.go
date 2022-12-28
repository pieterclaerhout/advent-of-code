package day02

import (
	"strings"
)

type Command struct{}

func (cmd *Command) Execute(input string) (any, any) {
	return cmd.calculateScore(input, cmd.scoresPart1()), cmd.calculateScore(input, cmd.scoresPart2())
}

func (cmd *Command) scoresPart1() map[string]int {
	return map[string]int{
		"B X": 1,
		"C Y": 2,
		"A Z": 3,
		"A X": 4,
		"B Y": 5,
		"C Z": 6,
		"C X": 7,
		"A Y": 8,
		"B Z": 9,
	}
}

func (cmd *Command) scoresPart2() map[string]int {
	return map[string]int{
		"B X": 1,
		"C X": 2,
		"A X": 3,
		"A Y": 4,
		"B Y": 5,
		"C Y": 6,
		"C Z": 7,
		"A Z": 8,
		"B Z": 9,
	}
}

func (cmd *Command) calculateScore(input string, scores map[string]int) int {
	var totalScore int
	for _, round := range strings.Split(input, "\n") {
		if score, ok := scores[round]; ok {
			totalScore += score
		}
	}
	return totalScore
}
