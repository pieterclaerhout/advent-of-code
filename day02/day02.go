package day02

import (
	"bufio"
	_ "embed"
	"strings"

	"golang.org/x/exp/slog"
)

//go:embed input.txt
var input string

type Command struct {
}

func (c Command) Execute() {
	c.part1()
	c.part2()
}

func (c Command) part1() {
	scores := map[string]int{
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

	var totalScore int
	for _, round := range c.parse() {
		if score, ok := scores[round]; ok {
			totalScore += score
		}
	}

	slog.Info("Total score 1", slog.Any("score", totalScore))
}

func (c Command) part2() {
	scores := map[string]int{
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

	var totalScore int
	for _, round := range c.parse() {
		if score, ok := scores[round]; ok {
			totalScore += score
		}
	}

	slog.Info("Total score 2", slog.Any("score", totalScore))
}

func (c Command) parse() []string {
	result := []string{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		result = append(result, sc.Text())
	}

	return result
}
