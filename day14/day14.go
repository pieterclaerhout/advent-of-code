package day14

import (
	"bufio"
	_ "embed"
	"strings"

	"golang.org/x/exp/slog"
)

//go:embed input.txt
var input string

type Command struct {
	RegisterX   int
	CycleNumber int
	FinalValue  int
}

func (c *Command) Execute() {
	c.part1()
	c.part2()
}

func (c *Command) part1() {
	caveMap := c.parse()
	slog.Info("Part 1", slog.Any("result", caveMap.simulateFallingSand()))
}

func (c *Command) part2() {
	caveMap := c.parse()
	slog.Info("Part 2", slog.Any("result", caveMap.simulateBlockSource()))
}

func (c *Command) parse() CaveMap {
	caveMap := make(CaveMap, 1000)
	for i := range caveMap {
		caveMap[i] = make([]bool, 1000)
	}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		points := strings.Split(sc.Text(), " -> ")
		pStart := strings.Split(points[0], ",")
		for i := range points[1:] {
			pNext := strings.Split(points[i+1], ",")
			caveMap.draw(pStart, pNext)
			pStart = pNext
		}
	}

	return caveMap
}
