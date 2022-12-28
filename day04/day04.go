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

type Command struct{}

func (cmd *Command) Execute(input string) (any, any) {
	ranges := cmd.parse(input)

	return cmd.part1(ranges), cmd.part2(ranges)
}

func (cmd *Command) part1(ranges [][]Range) int {
	var countOverlap int
	for _, r := range ranges {
		if r[0].Contains(r[1]) || r[1].Contains(r[0]) {
			countOverlap++
		}
	}

	return countOverlap
}

func (cmd *Command) part2(ranges [][]Range) int {
	var countOverlap int
	for _, r := range ranges {
		if r[0].Overlap(r[1]) {
			countOverlap++
		}
	}
	return countOverlap
}

func (cmd *Command) parse(input string) [][]Range {
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
