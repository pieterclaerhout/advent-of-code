package day05

import (
	"fmt"
	"strings"
)

type Command struct{}

func (cmd *Command) Execute(input string) (any, any) {
	return cmd.part1(input), cmd.part2(input)
}

func (cmd *Command) part1(input string) string {
	stacks := cmd.parseStacks(input)

	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "move") {
			var toMove int
			var from int
			var to int
			fmt.Sscanf(line, "move %d from %d to %d", &toMove, &from, &to)

			for move := 0; move < toMove; move++ {
				stacks[to-1].Push(stacks[from-1].Pop())
			}
		}
	}

	result := ""
	for _, s := range stacks {
		result += string(s.Pop())
	}

	return result
}

func (cmd *Command) part2(input string) string {
	stacks := cmd.parseStacks(input)

	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "move") {
			var toMove int
			var from int
			var to int
			fmt.Sscanf(line, "move %d from %d to %d", &toMove, &from, &to)

			stacks[to-1].Push(stacks[from-1].PopN(toMove)...)
		}
	}

	result := ""
	for _, s := range stacks {
		result += string(s.Pop())
	}

	return result
}

func (cmd *Command) parseStacks(input string) []Stack {
	stacks := []Stack{}

	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")

	// Needed as we need to know the number of stacks
	size := len(strings.Fields(lines[len(lines)-1]))
	for i := 0; i < size; i++ {
		stacks = append(stacks, Stack{})
	}

	for _, line := range lines[:len(lines)-1] {
		for i, r := range line {
			if r != ' ' && r != '[' && r != ']' {
				stacks[i/4].AddToBottom(r)
			}
		}
	}

	return stacks
}
