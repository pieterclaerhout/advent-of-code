package day05

import (
	"bufio"
	_ "embed"
	"fmt"
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
	stacks := [9]stack{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), "[") {
			for i, r := range sc.Text() {
				if r != ' ' && r != '[' && r != ']' {
					stacks[i/4].addToBottom(r)
				}
			}
		}

		if strings.HasPrefix(sc.Text(), "move") {
			var toMove int
			var from int
			var to int
			fmt.Sscanf(sc.Text(), "move %d from %d to %d", &toMove, &from, &to)

			for move := 0; move < toMove; move++ {
				stacks[to-1].push(stacks[from-1].pop())
			}
		}
	}

	result := ""
	for _, s := range stacks {
		result += string(s.pop())
	}

	slog.Info("Part 1", slog.Any("result", result))
}

func (c *Command) part2() {
	stacks := [9]stack{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), "[") {
			for i, r := range sc.Text() {
				if r != ' ' && r != '[' && r != ']' {
					stacks[i/4].addToBottom(r)
				}
			}
		}

		if strings.HasPrefix(sc.Text(), "move") {
			var toMove int
			var from int
			var to int
			fmt.Sscanf(sc.Text(), "move %d from %d to %d", &toMove, &from, &to)

			stacks[to-1].push(stacks[from-1].popN(toMove)...)
		}
	}

	result := ""
	for _, s := range stacks {
		result += string(s.pop())
	}

	slog.Info("Part 2", slog.Any("result", result))
}
