package day21

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
	monkeys := c.parse()
	result := monkeys.Solve("root")
	slog.Info("Part 1", slog.Any("result", result))
}

func (c *Command) part2() {
	monkeys := c.parse()
	result := monkeys.RootEquality()

	slog.Info("Part 2", slog.Any("result", result))
}

func (c *Command) parse() Monkeys {
	monkeys := Monkeys{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		lineParts := strings.Split(sc.Text(), ": ")
		monkeys[lineParts[0]] = lineParts[1]
	}

	return monkeys
}
