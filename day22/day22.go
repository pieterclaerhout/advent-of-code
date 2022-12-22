package day22

import (
	"bufio"
	_ "embed"
	"strings"

	"github.com/pieterclaerhout/advent-of-code/day22/part1"
	"github.com/pieterclaerhout/advent-of-code/day22/part2"
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
	parsed := c.parse()
	result := part1.Solve(parsed)
	slog.Info("Part 1", slog.Any("result", result))
}

func (c *Command) part2() {
	parsed := c.parse()
	result := part2.Solve(parsed)
	slog.Info("Part 2", slog.Any("result", result))
}

func (c *Command) parse() []string {
	lines := []string{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines

}
