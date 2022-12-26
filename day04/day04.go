package day04

import (
	"fmt"
	"strings"
)

type Range struct {
	s int
	e int
}

func (r1 Range) Contains(r2 Range) bool {
	return r1.s <= r2.s && r1.e >= r2.e
}

func (r1 Range) Overlap(r2 Range) bool {
	return r1.s <= r2.s && r1.e >= r2.s || r2.s <= r1.s && r2.e >= r1.s
}

type Command struct {
}

func (c *Command) Execute(input string) (any, any) {
	ranges := c.parse(input)

	return c.part1(ranges), c.part2(ranges)
}

func (c *Command) part1(ranges [][]Range) int {
	var countOverlap int
	for _, r := range ranges {
		if r[0].Contains(r[1]) || r[1].Contains(r[0]) {
			countOverlap++
		}
	}

	return countOverlap
}

func (c *Command) part2(ranges [][]Range) int {
	var countOverlap int
	for _, r := range ranges {
		if r[0].Overlap(r[1]) {
			countOverlap++
		}
	}
	return countOverlap
}

func (c *Command) parse(input string) [][]Range {
	ranges := [][]Range{}

	for _, line := range strings.Split(input, "\n") {
		var r1, r2 Range

		if _, err := fmt.Sscanf(line, "%d-%d,%d-%d", &r1.s, &r1.e, &r2.s, &r2.e); err != nil {
			continue
		}

		ranges = append(ranges, []Range{r1, r2})
	}

	return ranges
}
