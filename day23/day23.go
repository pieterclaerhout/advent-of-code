package day23

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
	a := c.parse()

	for i := 0; i < 10; i++ {
		a.DoRound()
	}
	slog.Info("Part 1", slog.Any("result", a.countEmptyCells()))
}

func (c *Command) part2() {
	a := c.parse()

	somebodyMoved := true
	rounds := 0
	for somebodyMoved {
		somebodyMoved = a.DoRound()
		rounds++
	}

	slog.Info("Part 2", slog.Any("result", rounds))
}

func (c *Command) parse() *Arena {
	lines := []string{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return NewArena(lines)
}
