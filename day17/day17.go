package day17

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
	solver := NewSolver(c.parse(), Rocks)
	slog.Info("Part 1", slog.Any("result", solver.Solve(2022)))
}

func (c *Command) part2() {
	solver := NewSolver(c.parse(), Rocks)
	slog.Info("Part 2", slog.Any("result", solver.Solve(1_000_000_000_000)))
}

func (c *Command) parse() []int8 {
	var output []int8

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	parseByte := func(r byte) int8 {
		switch r {
		case '<':
			return -1
		case '>':
			return 1
		}
		return 0
	}

	if sc.Scan() {
		for _, r := range []byte(sc.Text()) {
			output = append(output, parseByte(r))
		}
	}

	return output
}
