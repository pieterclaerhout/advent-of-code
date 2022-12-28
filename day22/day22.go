package day22

import (
	_ "embed"
	"strings"

	"github.com/pieterclaerhout/advent-of-code/day22/part1"
	"github.com/pieterclaerhout/advent-of-code/day22/part2"
)

// //go:embed input.txt
// var input string

type Command struct {
}

func (cmd *Command) Execute(input string) (any, any) {

	parsed := strings.Split(input, "\n")

	return part1.Solve(parsed), part2.Solve(parsed)
}

// func (c *Command) part1() {
// 	parsed := c.parse()
// 	result := part1.Solve(parsed)
// 	slog.Info("Part 1", slog.Any("result", result))
// }

// func (c *Command) part2() {
// 	parsed := c.parse()
// 	result := part2.Solve(parsed)
// 	slog.Info("Part 2", slog.Any("result", result))
// }

// func (c *Command) parse(input string) []string {
// 	lines := []string{}

// 	for _, line := range strings.Split(input, "\n") {
// 		lines = append(lines, line)
// 	}

// 	return lines

// }
