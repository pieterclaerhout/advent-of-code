package day02

import (
	"strings"
)

type Command struct {
}

func (c *Command) Execute(input string) (any, any) {
	return c.calculateScore(input, c.scoresPart1()), c.calculateScore(input, c.scoresPart2())
}

func (c *Command) scoresPart1() map[string]int {
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

func (c *Command) scoresPart2() map[string]int {
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

func (c *Command) calculateScore(input string, scores map[string]int) int {
	var totalScore int
	for _, round := range strings.Split(input, "\n") {
		if score, ok := scores[round]; ok {
			totalScore += score
		}
	}
	return totalScore
}
