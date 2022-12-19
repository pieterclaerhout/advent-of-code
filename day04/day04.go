package day04

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

func (c *Command) Execute() {
	c.part1()
	c.part2()
}

func (c *Command) part1() {
	var countOverlap int
	for _, pair := range c.parse() {
		if pair.OverlapsCompletely() {
			countOverlap++
		}
	}

	slog.Info("Part 1", slog.Any("countOverlap", countOverlap))
}

func (c *Command) part2() {
	var countOverlap int
	for _, pair := range c.parse() {
		if pair.Overlaps() {
			countOverlap++
		}
	}

	slog.Info("Part 2", slog.Any("countOverlap", countOverlap))
}

func (c *Command) parse() []Pair {
	result := []Pair{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		result = append(result, NewPair(sc.Text()))
	}

	return result
}
