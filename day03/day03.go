package day03

import (
	"strings"
)

type Command struct{}

func (cmd *Command) Execute(input string) (any, any) {
	lines := strings.Split(input, "\n")
	return cmd.part1(lines), cmd.part2(lines)
}

func (cmd *Command) part1(lines []string) int {
	var prioritySum int

	for _, line := range lines {
		bagSize := len(line)
		comp1, comp2 := line[:bagSize/2], line[bagSize/2:]
		common := intersect(comp1, comp2)
		prioritySum += itemPriority(common)
	}

	return prioritySum
}

func (cmd *Command) part2(lines []string) int {
	var prioritySum int

	for i := 0; i < len(lines); i = i + 3 {
		common := intersect(lines[i+2], intersect(lines[i], lines[i+1]))
		prioritySum += itemPriority(common)
	}

	return prioritySum
}

func intersect(a, b string) string {
	intersection := strings.Builder{}
	for _, c := range a {
		if strings.ContainsRune(b, c) && !strings.ContainsRune(intersection.String(), c) {
			intersection.WriteRune(c)
		}
	}
	return intersection.String()
}

func itemPriority(item string) int {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, char := range item {
		for i, c := range alphabet {
			if c == char {
				return i + 1
			}
		}
	}
	return 0
}
